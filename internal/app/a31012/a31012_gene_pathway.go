/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31012

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGenePathway creates the Gene entity Pathway asset bin files.
func (app *Application) createGenePathway() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.reactomePathwayParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-u -n")
}
