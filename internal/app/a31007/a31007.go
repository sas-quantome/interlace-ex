/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31007

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "14"               // Increment this number for each release.
	primaryInpFile = "mim2gene_medgen"  // NCBI MedGen.
	assetName      = "medical_genetics" // The entity's asset name.
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
		"a31007",
		"creates the Homo sapiens Gene entity Medical Genetics asset from the NCBI GenMed DB.",
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
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin") // CW tsv column 0
	app.refiner.AddBin("udk_omim_id_u.bin")        // CW tsv column 1
	app.refiner.AddBin("udk_medgen_cui_s.bin")     // CW tsv column 2
	app.refiner.AddBin("udk_bit_flags_u.bin")      // CW tsv column 3
	app.refiner.Mgr.Say("creating the gene medical genetics asset ...")
	return app.createGeneMedicalGenetics()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
