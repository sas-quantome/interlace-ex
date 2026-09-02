/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// createGeneProduct creates the Gene entity Product asset bin files.
func (app *Application) createGeneProduct() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.geneProductWithEnum(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-u -n")
}

// geneProductWithEnum parses the Entrez Gene DB files, creates the Products asset, and writes the enumeration bin files.
func (app *Application) geneProductWithEnum() error {
	if err := app.refiner.Enum.Create(productsEnum()); err != nil {
		return err
	}
	if err := app.ensemblIdMap(); err != nil {
		return err
	}
	if err := app.refSeqIdMap(); err != nil {
		return err
	}
	if err := app.geneInfo(); err != nil {
		return err
	}
	if err := app.uniProtIdMap(); err != nil {
		return err
	}
	return app.refiner.Enum.Save(app.refiner.EnumFp(enumName))
}

// parseGeneID replaces geneID. It returns the parsed int64 and a boolean indicating success.
func (app *Application) parseGeneID(s string) (int64, bool) {
	gid, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		app.refiner.Mgr.War(fmt.Sprintf("failed parsing GeneID from s=[%s]", s))
		return 0, false
	}
	return gid, true
}

// decodeURL example: 2%27,3%27-Cyclic-nucleotide_3%27-phosphodiesterase -> 2',3'-Cyclic-nucleotide_3'-phosphodiesterase
func decodeURL(s string) string {
	if strings.ContainsRune(s, '%') {
		if r, e := url.QueryUnescape(s); e == nil {
			return r
		}
	}
	return s
}
