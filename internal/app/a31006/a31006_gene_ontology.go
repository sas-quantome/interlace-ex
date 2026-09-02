/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31006

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneOntology creates the Gene entity Metadata asset bin files.
//
//goland:noinspection DuplicatedCode
func (app *Application) createGeneOntology() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.refiner.MakeEnum("bit_flags_cv", bitFlagsEnum()); err != nil {
		return err
	}
	if err := app.geneOntologyParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-n")
}
