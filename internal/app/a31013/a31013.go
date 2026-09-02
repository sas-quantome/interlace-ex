/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31013

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "14"                             // Increment this number for each release.
	primaryInpFile = "NCBI2Reactome_PE_Reactions.txt" // Reactome reactions.
	chemPmcInpFile = "ReactionPMIDS.txt"              // Reactome reaction PMIDs.
	assetName      = "chemical_reaction"              // The entity's asset name.
	enumName       = "reaction_entity_cv"             // The enumerator and property Go files require the same name.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner *contractor.Refiner // refiner configures and handles the data refinement.
	kv      map[int]int         // kv lookup table for primary IDs.
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
		"a31013",
		"creates the Homo sapiens Gene entity Chemical Reaction asset from the Reactome DB.",
		programVersion,
		[]string{primaryInpFile, chemPmcInpFile})
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
	app.refiner.AddBin("udk_reactome_hsa_id_u.bin")            // CW tsv column 1
	app.refiner.AddBin("udk_pubmed_id_u.bin")                  // CW tsv column 2
	app.refiner.AddBin(app.refiner.PropertyEnumName(enumName)) // CW tsv column 3
	app.refiner.AddBin("udk_reaction_hash_id_m.bin")           // CW tsv column 4
	app.refiner.Mgr.Say("creating the gene chemical reaction asset ...")
	return app.createGeneChemicalReaction()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
