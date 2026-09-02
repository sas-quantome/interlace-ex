/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31007

import (
	"strconv"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

  0 #MIM number
  1 GeneID
  2 type
  3 Source
  4 MedGenCUI
  5 Comment

*/

// geneMedGenParser parses the OMIM GenMed ID map and creates the Medical Genetics asset.
func (app *Application) geneMedGenParser() error {
	var err error
	var lg runner.LgBitMask
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [6]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 6 && cols[0][0] != '#' {
			records++

			// MICRO-OPTIMIZATION: Inline checks avoid function call overhead
			gid := cols[1]
			if gid != "" && gid != "-" {

				mim := cols[0]
				if mim == "" || mim == "-" {
					mim = "0"
				}

				cat := cols[2]
				if cat == "-" {
					cat = ""
				}

				cui := cols[4]
				if cui == "-" {
					cui = ""
				}

				com := cols[5]
				if com == "-" {
					com = ""
				}

				setOmimFlags(&lg, cat, com)

				// MICRO-OPTIMIZATION: Pre-format to string to bypass Fprintf reflection
				flagStr := strconv.FormatUint(uint64(lg.Flag), 10)

				if err = app.refiner.Writer("%s\t%s\t%s\t%s\n", gid, mim, cui, flagStr); err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene OMIM")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
