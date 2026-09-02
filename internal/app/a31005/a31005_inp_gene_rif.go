/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31005

import (
	"strconv"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

  0 #Tax ID
  1 Gene ID
  2 PubMed ID (PMID) list
  3 last update timestamp
  4 GeneRIF text

*/

// geneRifBasic parses the Gene RIF file to create the Gene Function asset.
func (app *Application) geneRifBasic() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [5]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 5 && cols[0] == defaultTaxonId {
			records++

			// MICRO-OPTIMIZATION: Inline checks avoid function call overhead
			gid := cols[1]
			txt := cols[4]

			if gid != "" && gid != "-" && txt != "" && txt != "-" {
				if hid := app.refiner.Cat.HashText(txt); hid > 0 {
					if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {

						// MICRO-OPTIMIZATION: Convert uint64 to string once to bypass reflection overhead
						hashStr := strconv.FormatUint(hid, 10)

						buf := cols[2]
						for {
							pid, rest, found := strings.Cut(buf, ",")

							// Inline check for pid
							if pid != "" && pid != "-" {
								if err = app.refiner.Writer("%s\t%s\t%s\t%s\n", gid, pid, cols[0], hashStr); err != nil {
									return runner.ValidateErrors(err, ska.Close())
								}
							}

							if !found {
								break
							}
							buf = rest
						}
					}
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene RIF")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
