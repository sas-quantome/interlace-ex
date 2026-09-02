/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31003

import (
	"strconv"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file: gene_info.gz from NCBI Gene.

  0 #tax_id
  1 GeneID
  2 Symbol
  4 Synonyms
  5 dbXrefs
  6 chromosome
  8 description
  9 type_of_gene
 13 Other_designations

*/

// geneInfo parses the Gene Info and writes the temporary metadata file.
func (app *Application) geneInfo() error {
	var err error
	var gmd runner.LgBitMask
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	records, row := 0, 0
	var cols [16]string
	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split (Zero-allocation)
		if runner.FillCols(ska.Text(), cols[:]) >= 16 && cols[0] == defaultTaxonId {
			records++
			tax := runner.DataOrNothing(cols[0], "-")
			gid := runner.DataOrNothing(cols[1], "-")
			sym := runner.DataOrNothing(cols[2], "-")
			syn := runner.DataOrNothing(cols[4], "-")
			dbx := runner.DataOrNothing(cols[5], "-")
			chr := runner.DataOrNothing(cols[6], "-")
			dsc := runner.DataOrNothing(cols[8], "-")
			aka := runner.DataOrNothing(cols[13], "-")

			if i, e := strconv.ParseInt(gid, 10, 32); e == nil {
				geneID32 := int32(i)
				hashID := app.index[geneID32]
				setGeneFlags(&gmd, runner.DataOrNothing(cols[9], "-"))

				// OPTIMIZATION: Bypassing reflection overhead for numeric types
				err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
					gid,
					tax,
					chr,
					sym,
					dsc,
					syn,
					aka,
					dbx,
					strconv.FormatUint(uint64(gmd.Flag), 10),
					strconv.FormatUint(hashID, 10),
				)

				if err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "gene info")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
