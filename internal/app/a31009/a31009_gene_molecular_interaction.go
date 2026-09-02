/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31009

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneMolecularInteraction creates the Gene entity Molecular Interaction asset bin files.
func (app *Application) createGeneMolecularInteraction() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.geneRifInteractionsParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-u -n")
}
