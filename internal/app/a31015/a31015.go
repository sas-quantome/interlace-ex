/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31015

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
	"github.com/sas-quantome/interlace-go/pkg/runner"
	"github.com/sas-quantome/interlace-go/pkg/storekeeper"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "10"                  // Increment this number for each release.
	primaryInpFile = "GWAS_ANCESTRY"       // GWAS ancestry genomic variants.
	assetName      = "genomic_variation"   // The entity's asset name.
	udkPubMedIds   = "udk_pubmed_id_u.bin" // Link to the GWAS ancestry data.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner   *contractor.Refiner    // refiner configures and handles the data refinement.
	pubMedIds *storekeeper.Container // pubMedIds from Genomic Variation asset.
	bitFlags  map[uint32]uint32      // bitFlags maps the GWAS ancestry data with the PubMed IDs.
}

// Start configures the Application, verifies parameters, begins logging, and returns true unless the help page is requested.
// It exits safely with an explicit error message whenever the configuration is invalid.
func (app *Application) Start() bool {
	var helpPageState bool
	app.pubMedIds = storekeeper.NewContainer(udkPubMedIds)
	app.refiner, helpPageState = contractor.NewRefiner(
		true,
		false,
		arranger.EntityConcept,
		arranger.GeneWorkflow,
		arranger.ProjectName,
		"a31015",
		"inserts the GWAS Ancestry data as bit flags to the Gene entity Genomic Variation asset.",
		programVersion,
		[]string{primaryInpFile, udkPubMedIds})
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
	app.refiner.AddBin("udk_bit_flags_u.bin")
	app.refiner.Mgr.Say("integrating the gene genomic variation GWAS ancestry data ...")
	return app.createGeneGenomicVariationAncestry()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(runner.ValidateErrors(err, app.pubMedIds.CloseMappedArr()))
}
