/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31002

import (
	"errors"
	"path"
	"strconv"

	"github.com/sas-quantome/interlace-go/pkg/runner"
	"github.com/sas-quantome/interlace-go/pkg/storekeeper"
)

// createGeneSynonym creates the Gene entity Synonym asset bin files.
func (app *Application) createGeneSynonym() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.geneProductToSynonym(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-u")
}

// geneProductToSynonym loads the Gene Product bin files to create the Gene Synonym asset, and copies the Products enumeration.
func (app *Application) geneProductToSynonym() error {
	var e1, e2, e3, e4 error
	productEntityNS, err := app.refiner.Mgr.EntityNS(queryAssetName)
	if err != nil {
		return err
	}

	// Load the Entrez Gene Product asset previously created.
	tt := app.refiner.Mgr.GetConfigString(runner.EventTimeTag)
	if err = app.refiner.Base.AddEntity("gene", queryAssetName, tt); err != nil {
		return err
	}

	fp := app.refiner.Base.EntityPropertyPath(productEntityNS)

	gid := storekeeper.NewContainer("sdk_entrez_gene_id_u.bin")
	if err = gid.OpenNumArr(path.Join(fp, "sdk_entrez_gene_id_u.bin"), false); err != nil {
		return err
	}
	defer func() {
		e1 = gid.CloseMappedArr()
	}()

	cat := storekeeper.NewContainer("udk_synonym_category_cv_d.bin")
	if err = cat.OpenNumArr(path.Join(fp, "udk_synonym_category_cv_d.bin"), false); err != nil {
		return runner.ValidateErrors(err, e1)
	}
	defer func() {
		e2 = cat.CloseMappedArr()
	}()

	pid := storekeeper.NewContainer("udk_protein_id_s.bin")
	if err = pid.OpenStrArr(path.Join(fp, "udk_protein_id_s.bin"), false); err != nil {
		return runner.ValidateErrors(err, e1, e2)
	}
	defer func() {
		e3 = pid.CloseMappedArr()
	}()

	syn := storekeeper.NewContainer("udk_synonym_s.bin")
	if err = syn.OpenStrArr(path.Join(fp, "udk_synonym_s.bin"), false); err != nil {
		return runner.ValidateErrors(err, e1, e2, e3)
	}
	defer func() {
		e4 = syn.CloseMappedArr()
	}()

	if !runner.ValidateSame(gid.Len, pid.Len, syn.Len, cat.Len) {
		err = errors.New("the Entrez Gene Index bin arrays are of different size")
		return runner.ValidateErrors(err, e1, e2, e3, e4)
	}

	for i := range gid.Len {
		ac, vn := runner.SplitAccessionVersion(syn.SArr[i])

		// OPTIMIZATION: Bypass fmt.Fprintf reflection overhead for fast integer-to-string conversion
		// Format: synonym \t geneID \t proteinID \t category \t version \n
		err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n",
			ac,
			strconv.FormatUint(uint64(gid.UArr[i]), 10),
			pid.SArr[i],
			strconv.Itoa(int(cat.DArr[i])),
			vn,
		)
		if err != nil {
			return runner.ValidateErrors(err, e1, e2, e3, e4)
		}
	}

	if err = app.refiner.FlushAll(); err != nil {
		return runner.ValidateErrors(err, e1, e2, e3, e4)
	}
	app.refiner.ReportCounts(gid.Len, gid.Len, "Entrez Gene RIF")

	// Copy the Gene Product Enumeration previously created to independently share the same vocabulary.
	src := path.Join(app.refiner.Base.EntityEnumerationPath(productEntityNS), "uuk_synonym_category_cv_s.bin")
	dst := path.Join(app.refiner.Base.EntityEnumerationPath(app.refiner.EntityNS), "uuk_synonym_category_cv_s.bin")

	err = runner.CopyFile(src, dst)
	return runner.ValidateErrors(err, e1, e2, e3, e4)
}
