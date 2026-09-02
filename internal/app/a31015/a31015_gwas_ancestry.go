/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31015

import (
	"path"
	"strconv"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
	"github.com/sas-quantome/interlace-go/pkg/storekeeper"
)

/*
 0  STUDY ACCESSION
 1  PUBMEDID
 2  FIRST AUTHOR
 3  DATE
 4  INITIAL SAMPLE DESCRIPTION
 5  REPLICATION SAMPLE DESCRIPTION
 6  STAGE
 7  NUMBER OF INDIVDUALS
 8  BROAD ANCESTRAL CATEGORY
 9  COUNTRY OF ORIGIN
 10 COUNTRY OF RECRUITMENT
 11 ADDITONAL ANCESTRY DESCRIPTION

*/

// gwasAncestryParser parses the GWAS Ancestry file.
func (app *Application) gwasAncestryParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row := 0, 0
	var cols [12]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array instead of using strings.Split
		if runner.FillCols(ska.Text(), cols[:]) >= 12 && row > 1 {
			records++

			// MICRO-OPTIMIZATION: Inline checks, handling both empty strings and dashes
			ancestry := cols[8]
			if ancestry == "" || ancestry == "-" {
				continue
			}

			pid, err := strconv.ParseUint(cols[1], 10, 64)
			if err != nil {
				app.refiner.Mgr.War("int parsing error from string: " + cols[1])
				continue
			}

			if pid != 0 {
				app.parseAncestry(ancestry, uint32(pid))
			}
		}
	}

	err = app.saveGeneAncestry()
	app.refiner.ReportCounts(row, records, "GWAS Ancestry")
	return runner.ValidateErrors(err, ska.Close())
}

// saveGeneAncestry writes the GWAS ancestry bit flags to the Gene entity Genomic Variation asset.
func (app *Application) saveGeneAncestry() error {
	bit := make([]uint32, app.pubMedIds.Len)
	for i, pid := range app.pubMedIds.UArr {
		bit[i] = app.bitFlags[pid]
		app.refiner.CountInc()
	}
	fp := app.refiner.Base.EntityPropertyPath(app.refiner.EntityNS)
	co := storekeeper.NewContainer("udk_bit_flags_u.bin")
	return co.SaveNumArr(path.Join(fp, "udk_bit_flags_u.bin"), bit)
}

// parseAncestry parses a comma-separated list of ancestries and links the PubMed ID to the corresponding ancestry flags.
func (app *Application) parseAncestry(text string, pubMedId uint32) {
	buf := text
	for {
		anc, rest, found := strings.Cut(buf, ", ")
		if anc != "" {
			// CRITICAL FIX: Purely functional, stateless flag fetching
			flag := getAncestryFlag(strings.TrimSpace(anc), app)
			if flag != 0 {
				app.bitFlags[pubMedId] |= flag
			}
		}
		if !found {
			break
		}
		buf = rest
	}
}
