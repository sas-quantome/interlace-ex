/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31015

import "path"

// createGeneGenomicVariationAncestry inserts the GWAS ancestry data to the Gene entity Genomic Variation asset.
func (app *Application) createGeneGenomicVariationAncestry() error {
	fp := app.refiner.Base.EntityPropertyPath(app.refiner.EntityNS)
	if err := app.pubMedIds.OpenNumArr(path.Join(fp, udkPubMedIds), false); err != nil {
		return err
	}
	app.bitFlags = make(map[uint32]uint32, app.pubMedIds.Len)
	if err := app.gwasAncestryParser(); err != nil {
		return err
	}
	return app.refiner.MakeEnum("bit_flags_cv", bitFlagsEnum())
}
