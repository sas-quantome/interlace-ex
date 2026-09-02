/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31013

import (
	"strconv"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

func (app *Application) indexReactomePubMedIds() error {
	var err error

	// Initialize map with capacity hint (e.g., 60,000 reactions)
	app.kv = make(map[int]int, 60000)

	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(chemPmcInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [2]string // Just two columns: ReactomeID and PMID

	for row = 1; ska.Scan(); row++ {
		// CRITICAL FIX: Use FillCols instead of the buggy for/Cut loop
		if runner.FillCols(ska.Text(), cols[:]) >= 2 {
			records++

			rid := reactomeID(cols[0])
			pid := reactomePubMedID(cols[1])

			if pid > 0 && rid > 0 {
				app.kv[rid] = pid
				app.refiner.CountInc()
			}
		}
	}
	app.refiner.ReportCounts(row, records, "Reactome PubMed reference")
	return ska.Close()
}

func reactomePubMedID(s string) int {
	if n, e := strconv.Atoi(s); e == nil {
		return n
	}
	return 0
}
