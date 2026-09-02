/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31005

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneFunction creates the Gene Function asset bin files.
func (app *Application) createGeneFunction() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.geneRifBasic(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-n")
}
