/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31008

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneOrtholog creates the Gene entity Ortholog asset bin files.
func (app *Application) createGeneOrtholog() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.geneOrthologParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-n")
}
