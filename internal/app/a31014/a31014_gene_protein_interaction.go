/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31014

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneProteinInteraction creates the Gene entity Protein Interaction asset bin files.
func (app *Application) createGeneProteinInteraction() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.proteinInteractionParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-n")
}
