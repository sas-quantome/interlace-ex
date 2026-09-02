/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31011

import (
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// gwasEboParser loads the EBI EBO Mappings.
func (app *Application) gwasEboParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.Mgr.GetConfigString(efoInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	records, row := 0, 0
	var cols [5]string
	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split
		if runner.FillCols(ska.Text(), cols[:]) >= 5 && row > 1 {
			records++

			// MICRO-OPTIMIZATION: Fast ID extraction without path.Base
			url := strings.TrimSpace(cols[2])
			if idx := strings.LastIndexByte(url, '/'); idx >= 0 {
				oid := url[idx+1:]
				if oid != "" {
					// OPTIMIZATION: Faster byte replacement instead of strings.ReplaceAll
					obo := replaceUnderscore(oid)
					app.obo[obo] = strings.TrimSpace(cols[1])
					app.refiner.CountInc()
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "EFO Terms")
	app.refiner.ResetCounter()
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}

// replaceUnderscore replaces the first underscore with a colon efficiently.
// EFO IDs typically only have one underscore (e.g., EFO_0006335 -> EFO:0006335).
func replaceUnderscore(s string) string {
	idx := strings.IndexByte(s, '_')
	if idx < 0 {
		return s
	}
	b := []byte(s)
	b[idx] = ':'
	return string(b)
}
