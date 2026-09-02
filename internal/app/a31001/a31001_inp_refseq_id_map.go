/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import "github.com/sas-quantome/interlace-go/pkg/runner"

// refSeqIdMap get the RefSeq DNA/RNA/Protein accessions the Entrez gene2accession.gz and writes to the tmp file.
func (app *Application) refSeqIdMap() error {
	var err error
	dna := app.refiner.Enum.Assure("RefSeq_NT")
	pro := app.refiner.Enum.Assure("RefSeq_PRO")
	rna := app.refiner.Enum.Assure("RefSeq_TRS")
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(accInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	records, row := 0, 0
	var cols [16]string
	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split
		if runner.FillCols(ska.Text(), cols[:]) >= 16 && cols[0] == defaultTaxonId {
			if gid, ok := app.parseGeneID(cols[1]); ok {
				records++
				pid := cols[5]
				if pid == "-" {
					pid = ""
				} else {
					err = app.saveRefIds(gid, cols[5], pid, pro)
				}
				if err == nil {
					if err = app.saveRefIds(gid, cols[3], pid, rna); err == nil {
						if err = app.saveRefIds(gid, cols[9], pid, dna); err == nil {
							err = app.saveRefIds(gid, cols[7], pid, dna)
						}
					}
				}
				if err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene RefSeq")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
