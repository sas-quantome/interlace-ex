/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31012

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// reactomePathwayParser parses the Reactome pathways and creates the Pathways asset.
func (app *Application) reactomePathwayParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [6]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 6 {
			records++

			// MICRO-OPTIMIZATION: Inline checks to avoid function call overhead
			gid := cols[0]
			txt := cols[3]

			// Ensure valid data and that gid starts with a digit
			if gid != "" && gid != "-" && txt != "" && txt != "-" && unicode.IsDigit(rune(gid[0])) {

				rid := reactomeIdToInt(cols[1])

				if rid > 0 {
					if hid := app.refiner.Cat.HashText(txt); hid > 0 {
						if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {

							// MICRO-OPTIMIZATION: Format integers to strings to bypass reflection overhead
							ridStr := strconv.Itoa(rid)
							hashStr := strconv.FormatUint(hid, 10)

							if err = app.refiner.Writer("%s\t%s\t%s\n", gid, ridStr, hashStr); err != nil {
								return runner.ValidateErrors(err, ska.Close())
							}
						}
					}
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Reactome pathways")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}

// reactomeIdToInt extracts the integer ID from a Reactome string.
// Note: taxonID parameter removed as it was statically compared to a global constant.
func reactomeIdToInt(reactomeID string) int {
	if strings.HasPrefix(reactomeID, "R-HSA-") {
		// MICRO-OPTIMIZATION: Byte slice [6:] is faster than calling strings.TrimPrefix
		if n, e := strconv.Atoi(reactomeID[6:]); e == nil {
			return n
		}
	}
	return 0
}
