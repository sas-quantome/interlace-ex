/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31006

import (
	"strconv"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

  0 #tax_id
  1 GeneID
  2 GO_ID
  3 Evidence
  4 Qualifier
  5 GO_term
  6 PubMed
  7 Category

*/

func (app *Application) geneOntologyParser() error {
	var err error
	var xl runner.XlBitMask
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [8]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 8 && cols[0] == defaultTaxonId {
			records++

			// MICRO-OPTIMIZATION: Inline checking avoiding function calls
			gid := cols[1]
			oid := cols[2]
			txt := cols[5]

			if gid != "" && gid != "-" && oid != "" && oid != "-" && txt != "" && txt != "-" {
				if hid := app.refiner.Cat.HashText(txt); hid > 0 {
					if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {

						setGoFlags(&xl, cols[3], cols[7], cols[4])

						// MICRO-OPTIMIZATION: Pre-format to strings to bypass Fprintf reflection in loop
						flagStr := strconv.FormatUint(xl.Flag, 10)
						hashStr := strconv.FormatUint(hid, 10)
						oidClean := runner.ExtractNumber(oid, "-", "GO:")

						buf := cols[6]
						for {
							pid, rest, found := strings.Cut(buf, "|")

							if pid != "" && pid != "-" {
								if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n", gid, pid, oidClean, flagStr, hashStr); err != nil {
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
	app.refiner.ReportCounts(row, records, "Entrez Gene GO")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
