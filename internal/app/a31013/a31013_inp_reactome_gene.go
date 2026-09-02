/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31013

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// geneToReactome parses the NCBI Reactome ID map and creates the chemical reaction asset.
func (app *Application) geneToReactome() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [8]string

	for row = 1; ska.Scan(); row++ {
		if runner.FillCols(ska.Text(), cols[:]) >= 8 {
			records++

			// MICRO-OPTIMIZATION: Inline checks
			gid := cols[0]
			txt := cols[5]

			if gid != "" && gid != "-" && txt != "" && txt != "-" && unicode.IsDigit(rune(gid[0])) {

				rid := reactomeID(cols[3])

				if rid > 0 {
					if hid := app.refiner.Cat.HashText(txt); hid > 0 {

						org := runner.PopOutString(cols[2], "[", "]")
						isa := app.refiner.Enum.Assure(org)

						if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {

							pid := 0
							if val, found := app.kv[rid]; found {
								pid = val
							}

							// MICRO-OPTIMIZATION: Pre-format to strings to bypass Fprintf reflection
							ridStr := strconv.Itoa(rid)
							pidStr := strconv.Itoa(pid)
							isaStr := strconv.Itoa(isa)
							hashStr := strconv.FormatUint(hid, 10)

							if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n", gid, ridStr, pidStr, isaStr, hashStr); err != nil {
								return runner.ValidateErrors(err, ska.Close())
							}
						}
					}
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Reactome reactions")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}

// reactomeID extracts the integer ID from a Reactome string.
// Note: taxonID parameter removed.
func reactomeID(reactomeID string) int {
	if strings.HasPrefix(reactomeID, "R-HSA-") {
		// MICRO-OPTIMIZATION: Byte slice [6:] is faster than calling strings.TrimPrefix
		if n, e := strconv.Atoi(reactomeID[6:]); e == nil {
			return n
		}
	}
	return 0
}
