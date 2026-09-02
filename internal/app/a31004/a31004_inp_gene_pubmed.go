/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31004

import "github.com/sas-quantome/interlace-go/pkg/runner"

/*

Input file: gene2pubmed.gz

  0 #tax_id
  1 GeneID
  2 PubMed_ID

*/

// geneToPubMedParser parses the GeneDB PubMed ID map and creates the PubMed Reference asset.
func (app *Application) geneToPubMedParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [3]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 3 && cols[0] == defaultTaxonId {
			records++

			// MICRO-OPTIMIZATION: Inline checks avoid function call overhead in the hot loop
			if cols[1] != "" && cols[1] != "-" && cols[2] != "" && cols[2] != "-" {
				if err = app.refiner.Writer("%s\t%s\n", cols[1], cols[2]); err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}

	app.refiner.ReportCounts(row, records, "Entrez Gene PubMed")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
