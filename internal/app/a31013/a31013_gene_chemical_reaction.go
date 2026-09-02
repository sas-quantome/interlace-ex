/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31013

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneChemicalReaction creates the Gene entity Chemical Reaction asset bi files.
func (app *Application) createGeneChemicalReaction() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.reactomeReactionsWithEnum(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-u -n")
}

// reactomeReactionsWithEnum creates the Chemical Reactions asset and writes the Reaction enum bin files.
func (app *Application) reactomeReactionsWithEnum() error {
	if err := app.refiner.Enum.Create(reactionsEnum()); err != nil {
		return err
	}
	if err := app.indexReactomePubMedIds(); err != nil {
		return err
	}
	app.refiner.ResetCounter()
	if err := app.geneToReactome(); err != nil {
		return err
	}
	return app.refiner.Enum.Save(app.refiner.EnumFp(enumName))
}
