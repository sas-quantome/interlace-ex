/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31010

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneGeneticVariation creates the Gene entity Genetic Variation asset bin files.
func (app *Application) createGeneGeneticVariation() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.bitFlagsToEnum(); err != nil {
		return err
	}
	if err := app.refiner.Enum.Create(app.terms); err != nil {
		return err
	}
	if err := app.clinVarSummaryParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	if err := app.refiner.Enum.Save(app.refiner.EnumFp("bit_flags_cv")); err != nil {
		return err
	}
	return app.refiner.CloseCowriterWithCatalog("-n")
}
