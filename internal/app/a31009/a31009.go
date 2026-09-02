/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31009

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "13"                    // Increment this number for each release.
	defaultTaxonId = "9606"                  // The Homo sapiens NCBI Taxonomy ID is required for human data aggregation.
	primaryInpFile = "interactions.gz"       // Entrez Gene RIF interactions.
	assetName      = "molecular_interaction" // The entity's asset name.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner *contractor.Refiner // refiner configures and handles the data refinement.
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
		"a31009",
		"creates the Homo sapiens Gene entity Molecular Interaction asset from the NCBI Gene RIF.",
		programVersion,
		[]string{primaryInpFile})
	return helpPageState
}

// Commit executes the application's work.
func (app *Application) Commit() error {
	app.refiner.Mgr.Say("setting up the refiner ...")
	if commit, err := app.refiner.Commit(assetName); err != nil {
		return err
	} else if commit == false {
		return nil
	}
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin")             // CW tsv column 0
	app.refiner.AddBin("udk_taxonomy_id_u.bin")                // CW tsv column 1
	app.refiner.AddBin("udk_interactant_taxonomy_id_u.bin")    // CW tsv column 2
	app.refiner.AddBin("udk_interactant_entrez_gene_id_u.bin") // CW tsv column 3
	app.refiner.AddBin("udk_pubmed_id_set_s.bin")              // CW tsv column 4
	app.refiner.Mgr.Say("creating the gene molecular interaction asset ...")
	return app.createGeneMolecularInteraction()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
