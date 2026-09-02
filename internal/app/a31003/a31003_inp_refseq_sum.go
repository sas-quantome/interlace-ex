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

Input file: gene_refseq_summary.gz from NCBI Gene.

  0 #tax_id
  1 GeneID
  2 Source
  3 Summary

*/

func (app *Application) indexGenSum() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(summaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	// pre-sized to a reasonable guess: Human Entrez Gene unique IDs ~193,797
	app.index = make(map[int32]uint64, 1024*256)
	records, row := 0, 0
	var cols [4]string // Stack-allocated array for 4 columns
	for row = 1; ska.Scan(); row++ {
		line := ska.Text()
		if runner.FillCols(line, cols[:]) >= 4 && cols[0] == defaultTaxonId {
			records++
			gid := runner.DataOrNothing(cols[1], "-")
			txt := runner.DataOrNothing(cols[3], "-")
			if gid != "" && txt != "" {
				if hid := app.refiner.Cat.HashText(txt); hid > 0 {
					if geneID, e := strconv.ParseInt(gid, 10, 32); e == nil {
						if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {
							app.index[int32(geneID)] = hid
							app.refiner.CountInc()
						}
					} else {
						err = e
					}
				}
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Entrez Gene RefSeq summary")
	return ska.Close()
}
