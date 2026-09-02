/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import "github.com/sas-quantome/interlace-go/pkg/runner"

// ensemblIdMap gets the Ensembl Gene/RNA/Protein accessions from the Entrez gene2ensembl.gz and writes to the tmp file.
func (app *Application) ensemblIdMap() error {
	var err error
	gen := app.refiner.Enum.Assure("Ensembl")
	pro := app.refiner.Enum.Assure("Ensembl_PRO")
	rna := app.refiner.Enum.Assure("Ensembl_TRS")
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(ensInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	records, row := 0, 0
	var cols [7]string
	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split
		if runner.FillCols(ska.Text(), cols[:]) >= 7 && cols[0] == defaultTaxonId {
			if gid, ok := app.parseGeneID(cols[1]); ok {
				records++
				pid := cols[6]
				if pid == "-" {
					pid = ""
				} else {
					err = app.saveGeneEnsembl(gid, cols[6], pid, pro)
				}
				if err == nil {
					if err = app.saveGeneEnsembl(gid, cols[2], pid, gen); err == nil {
						err = app.saveGeneEnsembl(gid, cols[4], pid, rna)
					}
				}
				if err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene Ensembl")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
