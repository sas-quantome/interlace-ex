/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31009

import (
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

  0 #tax_id
  1 gene_id
  2 accn.vers
  3 name
  4 keyphrase
  5 tax_id
  6 interactant_id
  7 interactant_id_type
  8 accn.vers
  9 name
  10 complex_id
  11 complex_id_type
  12 complex_name
  13 pubmed_id_list
  14 last_mod
  15 generif_text
  16 interaction_id
  17 interaction_id_type

*/

// geneRifInteractionsParser parses the GeneRIF interactions file and creates the Molecular Interaction asset.
func (app *Application) geneRifInteractionsParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [18]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 18 && cols[0] == defaultTaxonId {
			records++

			// MICRO-OPTIMIZATION: Inline checks avoid function call overhead
			gid := cols[1]
			otherGid := cols[6]
			pmids := cols[13]

			// Ensure required fields are valid and it's a GeneID interaction
			if gid != "" && gid != "-" && otherGid != "" && otherGid != "-" && pmids != "" && pmids != "-" && cols[7] == "GeneID" {

				otherTax := cols[5]
				if otherTax == "-" {
					otherTax = ""
				}

				pids := strings.ReplaceAll(pmids, ",", "|")
				if pids != "" {
					// cols[0] is already guaranteed to be defaultTaxonId
					if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n", gid, cols[0], otherTax, otherGid, pids); err != nil {
						return runner.ValidateErrors(err, ska.Close())
					}
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene RIF interactions")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
