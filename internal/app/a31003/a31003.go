/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31003

import (
	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "13"              // Increment this number for each release.
	defaultTaxonId = "9606"            // The Homo sapiens NCBI Taxonomy ID is required for human data aggregation.
	primaryInpFile = "gene_info.gz"    // Entrez Gene Info contains the gene metadata.
	summaryInpFile = "gene_summary.gz" // Entrez Gene summary contains the RefSeq description.
	assetName      = "metadata"        // The entity's asset name.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner *contractor.Refiner // refiner configures and handles the data refinement.
	index   map[int32]uint64    // index links asset properties to their text hashes.
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
		"a31003",
		"creates the Homo sapiens Gene entity Metadata asset from the NCBI Gene DB.",
		programVersion,
		[]string{primaryInpFile, summaryInpFile})
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
	app.refiner.AddBin("suk_entrez_gene_id_u.bin")      // CW tsv column 0
	app.refiner.AddBin("udk_taxonomy_id_u.bin")         // CW tsv column 1
	app.refiner.AddBin("udk_chromosome_s.bin")          // CW tsv column 2
	app.refiner.AddBin("udk_gene_symbol_s.bin")         // CW tsv column 3
	app.refiner.AddBin("udk_gene_name_s.bin")           // CW tsv column 4
	app.refiner.AddBin("udk_symbol_synonyms_set_s.bin") // CW tsv column 5
	app.refiner.AddBin("udk_name_synonyms_set_s.bin")   // CW tsv column 6
	app.refiner.AddBin("udk_xdb_id_set_s.bin")          // CW tsv column 7
	app.refiner.AddBin("udk_bit_flags_u.bin")           // CW tsv column 8
	app.refiner.AddBin("udk_metadata_hash_id_m.bin")    // CW tsv column 9
	app.refiner.Mgr.Say("creating the gene metadata asset ...")
	return app.createGeneMetadata()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(err)
}
