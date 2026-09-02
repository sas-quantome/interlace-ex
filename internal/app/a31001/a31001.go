/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 interlace-go version string.
const (
	programVersion = "12"                          // Increment this number for each release.
	defaultTaxonId = "9606"                        // The Homo sapiens NCBI Taxonomy ID is required for human data aggregation.
	primaryInpFile = "gene_info.gz"                // Entrez Gene Info contains the gene metadata.
	ensInpFile     = "gene2ensembl.gz"             // Entrez Gene to Ensembl.
	accInpFile     = "gene2accession.gz"           // Entrez Gene to RefSeq accession numbers.
	idsInpFile     = "HUMAN_9606_idmapping.dat.gz" // NCBI Entrez Gene to external database ID maps.
	assetName      = "product"                     // The entity's asset name.
	enumName       = "synonym_category_cv"         // The enumerator and property Go files require the same name.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner *contractor.Refiner // refiner configures and handles the data refinement.
	ids     idTrio              // ids link primary identifiers to their products.
}

// Start configures the Application, verifies parameters, begins logging, and returns true unless the help page is requested.
// It exits safely with an explicit error message whenever the configuration is invalid.
func (app *Application) Start() bool {
	var helpPageState bool
	app.refiner, helpPageState = contractor.NewRefiner(
		true,
		false,
		arranger.EntityConcept,
		arranger.GeneWorkflow,
		arranger.ProjectName,
		"a31001",
		"creates the Homo sapiens Gene entity Product asset from NCBI Gene DB and UniProt.",
		programVersion,
		[]string{primaryInpFile, ensInpFile, accInpFile, idsInpFile})
	return helpPageState
}

// Commit executes the work if the refiner is ready, unless the commit flag is false or an input error occurs.
func (app *Application) Commit() error {
	app.refiner.Mgr.Say("setting up the refiner ...")
	if commit, err := app.refiner.Commit(assetName); err != nil {
		return err
	} else if commit == false {
		return nil
	}
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin")             // CW tsv column 0
	app.refiner.AddBin("udk_protein_id_s.bin")                 // CW tsv column 1
	app.refiner.AddBin("udk_synonym_s.bin")                    // CW tsv column 2
	app.refiner.AddBin(app.refiner.PropertyEnumName(enumName)) // CW tsv column 3
	app.refiner.Mgr.Say("creating the gene product asset ...")
	return app.createGeneProduct()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
