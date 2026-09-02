/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31007

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGeneMedicalGenetics creates the Gene entity Medical Genetics asset bin files.
//
//goland:noinspection DuplicatedCode
func (app *Application) createGeneMedicalGenetics() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.refiner.MakeEnum("bit_flags_cv", bitFlagsEnum()); err != nil {
		return err
	}
	if err := app.geneMedGenParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-n")
}
