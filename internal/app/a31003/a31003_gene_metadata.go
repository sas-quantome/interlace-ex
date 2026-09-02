/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31003

import (
	"errors"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// createGeneMetadata creates the Gene entity Metadata and Catalog bin files.
//
//goland:noinspection DuplicatedCode
func (app *Application) createGeneMetadata() error {
	if err := app.refiner.OpenCowriterWithCatalog(assetName); err != nil {
		return err
	}
	if err := app.refiner.MakeEnum("bit_flags_cv", bitFlagsEnum()); err != nil {
		return err
	}
	if err := app.geneMetadata(); err != nil {
		return runner.ValidateErrors(err, app.refiner.CW.Close())
	}
	return app.refiner.CloseCowriterWithCatalog("-n")
}

// geneMetadata extracts data from Gene Info and RefSeq summary.
func (app *Application) geneMetadata() error {
	var err error
	if err = app.indexGenSum(); err != nil {
		return err
	}
	if app.refiner.Count() == 0 {
		return errors.New("failed indexing the gene summary data")
	}
	app.refiner.ResetCounter()
	return app.geneInfo()
}
