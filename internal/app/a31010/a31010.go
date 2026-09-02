/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31010

import (
	"strings"

	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "14"                     // Increment this number for each release.
	primaryInpFile = "variant_summary.txt.gz" // NCBI ClinVar.
	assetName      = "genetic_variation"      // The entity's asset name.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner *contractor.Refiner // refiner configures and handles the data refinement.
	buf     strings.Builder
	cache   [16]string
	terms   []string
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
		"a31010",
		"creates the Homo sapiens Gene entity Genetic Variation asset from the NCBI ClinVar DB.",
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
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin")    // CW tsv column 0
	app.refiner.AddBin("udk_variation_id_u.bin")      // CW tsv column 1
	app.refiner.AddBin("udk_refsnp_id_u.bin")         // CW tsv column 2
	app.refiner.AddBin("udk_omim_id_u.bin")           // CW tsv column 3
	app.refiner.AddBin("udk_bit_flags_m.bin")         // CW tsv column 4
	app.refiner.AddBin("udk_conical_allele_id_u.bin") // CW tsv column 5
	app.refiner.AddBin("udk_variation_hash_id_m.bin") // CW tsv column 6
	app.refiner.Mgr.Say("creating the gene genetic variation asset ...")
	return app.createGeneGeneticVariation()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
