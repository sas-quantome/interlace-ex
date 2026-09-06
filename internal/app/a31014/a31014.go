/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31014

import (
	"path"

	"github.com/sas-quantome/interlace-ex/internal/pkg/arranger"
	"github.com/sas-quantome/interlace-go/pkg/contractor"
	"github.com/sas-quantome/interlace-go/pkg/runner"
	"github.com/sas-quantome/interlace-go/pkg/storekeeper"
)

// constant program version number: to be appended to the ISO 8601 go-interlace version string.
const (
	programVersion = "16"                       // Increment this number for each release.
	primaryInpFile = "STRING_DB"                // STRING DB protein interactions.
	assetName      = "protein_interaction"      // The entity's asset name.
	queryAssetName = "synonym"                  // The entity's query asset name.'
	sdkGenSyn      = "sdk_synonym_s.bin"        // Synonyms index.
	udkGenId       = "udk_entrez_gene_id_u.bin" // Associated Entrez Gene IDs.
)

// Application represents the main struct for managing resources, computing, and coordinating deterministic data flows.
type Application struct {
	refiner  *contractor.Refiner    // refiner configures and handles the data refinement.
	synonyms *storekeeper.Container // synonyms SDK array.
	eGeneIDs *storekeeper.Container // eGeneIDs UDK array.
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
		"a31014",
		"creates the Homo sapiens Gene entity Protein Interaction asset from the STRING DB.",
		programVersion,
		[]string{primaryInpFile, sdkGenSyn, udkGenId})
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
	app.refiner.AddBin("sdk_entrez_gene_id_u.bin")             // CW tsv column 0
	app.refiner.AddBin("udk_association_entrez_gene_id_u.bin") // CW tsv column 1
	app.refiner.AddBin("udk_ensp_id_i.bin")                    // CW tsv column 2
	app.refiner.AddBin("udk_association_ensp_id_i.bin")        // CW tsv column 3
	app.refiner.AddBin("udk_likelihood_score_n.bin")           // CW tsv column 4
	app.refiner.Mgr.Say("creating the gene protein interaction asset ...")
	return app.createGeneProteinInteraction()
}

// End safely exits the Application.
func (app *Application) End(err error) {
	app.refiner.End(runner.ValidateErrors(err, app.synonyms.CloseMappedArr(), app.eGeneIDs.CloseMappedArr()))
}
