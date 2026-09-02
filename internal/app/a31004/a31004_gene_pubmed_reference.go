/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31004

import "github.com/sas-quantome/interlace-go/pkg/runner"

// createGenePubMedReference creates the Gene entity Metadata asset bin files.
func (app *Application) createGenePubMedReference() error {
	if err := app.refiner.OpenCowriter(assetName); err != nil {
		return err
	}
	if err := app.geneToPubMedParser(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriter("-n")
}
