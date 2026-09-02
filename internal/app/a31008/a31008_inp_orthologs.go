/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31008

import "github.com/sas-quantome/interlace-go/pkg/runner"

/*

Input file:

  0 #tax_id
  1 GeneID
  2 relationship
  3 Other_tax_id
  4 Other_GeneID

*/

// geneOrthologParser parses the GeneDB ortholog file and creates the Gene Ortholog asset.
func (app *Application) geneOrthologParser() error {
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
			otherGid := cols[4]

			if gid != "" && gid != "-" && otherGid != "" && otherGid != "-" {
				otherTax := cols[3]
				if otherTax == "-" {
					otherTax = ""
				}

				// cols[0] are already guaranteed to be defaultTaxonId
				if err = app.refiner.Writer("%s\t%s\t%s\t%s\n", gid, cols[0], otherTax, otherGid); err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene Orthologs")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
