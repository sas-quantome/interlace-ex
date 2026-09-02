/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import "github.com/sas-quantome/interlace-go/pkg/runner"

// geneInfo gets the Entrez symbol/HGNC/MIM/IMGT/miRBase accessions from the gene_info.gz info and writes to the file.
func (app *Application) geneInfo() error {
	var err error
	syn := app.refiner.Enum.Assure("Gene_Synonym")
	hgn := app.refiner.Enum.Assure("HGNC")
	mim := app.refiner.Enum.Assure("MIM")
	img := app.refiner.Enum.Assure("IMGT")
	mir := app.refiner.Enum.Assure("miRBase")
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	records, row := 0, 0
	var cols [11]string
	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split
		if runner.FillCols(ska.Text(), cols[:]) >= 11 && cols[0] == defaultTaxonId {
			if gid, ok := app.parseGeneID(cols[1]); ok {
				records++
				if err = app.saveRefIds(gid, cols[2], "", syn); err == nil {
					if err = app.saveSynonym(gid, cols[4], syn); err == nil {
						err = app.saveXdbRef(gid, cols[5], hgn, img, mim, mir)
					}
				}
				if err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene synonyms")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
