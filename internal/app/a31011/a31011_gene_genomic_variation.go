/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31011

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneGenomicVariation creates the Gene entity Genomic Variation asset bin files.
func (app *Application) createGeneGenomicVariation() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.oboParser(); err != nil {
		return err
	}
	if err := app.gwasCatalogParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-n")
}
