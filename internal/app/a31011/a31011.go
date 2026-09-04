/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31011

import (
	"path"

	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
	"github.com/sas-quantome/interlace-go/pkg/runner"
	"github.com/sas-quantome/interlace-go/pkg/storekeeper"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "15"                       // Increment this number for each release.
	primaryInpFile = "GWAS_CATALOG"             // EBI GWAS Catalog.
	efoInpFile     = "GWAS_EFO"                 // EBI EFO trait Mappings.
	hpoInpFile     = "hp.obo"                   // Human Phenotype Ontology.
	obaInpFile     = "oba.obo"                  // Ontology of Biological Attributes.
	gssInpFile     = "gsso.obo"                 // Gender, Sex, and Sexual Orientation ontology.
	monInpFile     = "mondo.obo"                // Mondo Disease Ontology.
	assetName      = "genomic_variation"        // The entity's asset name.
	queryAssetName = "synonym"                  // The entity's query asset name.'
	sdkGenSyn      = "sdk_synonym_s.bin"        // Synonyms index.
	udkGenId       = "udk_entrez_gene_id_u.bin" // Associated Entrez Gene IDs.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner  *contractor.Refiner    // refiner configures and handles the data refinement.
	synonyms *storekeeper.Container // synonyms SDK array.
	eGeneIDs *storekeeper.Container // eGeneIDs UDK array.
	obo      map[string]string      // ebo maps EBO terms to their ontology IDs.
	unk      []string               // n is the number of OBO ids not found.
}

// Start configures the Application, verifies parameters, begins logging, and returns true unless the help page is requested.
// It exits safely with an explicit error message whenever the configuration is invalid.
func (app *Application) Start() bool {
	var helpPageState bool
	app.synonyms = storekeeper.NewContainer(sdkGenSyn)
	app.eGeneIDs = storekeeper.NewContainer(udkGenId)
	app.refiner, helpPageState = contractor.NewRefiner(
		true,
		false,
		arranger.EntityConcept,
		arranger.GeneWorkflow,
		arranger.ProjectName,
		"a31011",
		"creates the Homo sapiens Gene entity Genomic Variation asset from the EBI GWAS Catalog.",
		programVersion,
		[]string{primaryInpFile, efoInpFile, hpoInpFile, obaInpFile, gssInpFile, monInpFile, sdkGenSyn, udkGenId})
	return helpPageState
}

// Commit executes the application's work.
//
//goland:noinspection DuplicatedCode
func (app *Application) Commit() error {
	app.refiner.Mgr.Say("setting up the refiner ...")
	if commit, err := app.refiner.Commit(assetName); err != nil {
		return err
	} else if commit == false {
		return nil
	}
	var err error
	fp := app.refiner.Base.SwapAssetPath(app.refiner.EntityNS, assetName, queryAssetName)
	if err = app.synonyms.OpenStrArr(path.Join(fp, sdkGenSyn), false); err != nil {
		return err
	}
	if err = app.eGeneIDs.OpenNumArr(path.Join(fp, udkGenId), false); err != nil {
		return err
	}
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin") // CW tsv column 0
	app.refiner.AddBin("udk_refsnp_id_u.bin")      // CW tsv column 1
	app.refiner.AddBin("udk_pubmed_id_u.bin")      // CW tsv column 2
	app.refiner.AddBin("udk_ontology_id_s.bin")    // CW tsv column 3
	app.refiner.AddBin("udk_gwas_hash_id_m.bin")   // CW tsv column 4
	app.refiner.Mgr.Say("creating the gene genomic variation asset ...")
	return app.createGeneGenomicVariation()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(runner.ValidateErrors(err, app.synonyms.CloseMappedArr(), app.eGeneIDs.CloseMappedArr()))
}
