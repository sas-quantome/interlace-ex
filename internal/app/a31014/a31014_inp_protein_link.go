/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31014

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/sas-quantome/interlace-go/pkg/cowriter"
	"github.com/sas-quantome/interlace-go/pkg/runner"
)

// proteinInteractionParser parses the STRING DB FILE and creates the protein interaction asset.
func (app *Application) proteinInteractionParser() error {
	cw := cowriter.NewCowriter()
	fp := app.refiner.Base.EventArtifactPath(app.refiner.EventNS)
	if err := cw.Open(fp, "string_db_*.tsv.gz", true, true); err != nil {
		return runner.ValidateErrors(err, cw.Close())
	}
	if err := app.parseProteinInteraction(cw); err != nil {
		return runner.ValidateErrors(err, cw.Close())
	}
	if err := cw.Close(); err != nil {
		return err
	}
	if err := cw.Sort("-n"); err != nil {
		return err
	}
	if err := app.createProteinInteraction(cw.File.Name()); err != nil {
		return err
	}
	return cw.Remove()
}

// parseProteinInteraction adds the GeneID to the STRING DB protein 1 and writes to temp file 1.
// protein1 protein2 combined_score
// 9606.ENSP00000000233 9606.ENSP00000356607 173
func (app *Application) parseProteinInteraction(cw *cowriter.Cowriter) error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	found, missing, row := 0, 0, 0
	var lid, nid string
	var gidStr string // Cache the stringified Gene ID
	idx := -1

	for row = 1; ska.Scan(); row++ {
		line := ska.Text()
		if len(line) == 0 || !unicode.IsDigit(rune(line[0])) {
			continue
		}

		col1, rest, ok1 := strings.Cut(line, " ")
		col2, score, ok2 := strings.Cut(rest, " ")

		if ok1 && ok2 {
			if len(col1) > 5 && col1[:5] == "9606." {
				col1 = col1[5:]
			}
			if len(col2) > 5 && col2[:5] == "9606." {
				col2 = col2[5:]
			}

			if lid != col1 {
				lid = col1
				nid = ""
				idx = runner.BinarySearchItem(app.synonyms.SArr, col1)
				if idx >= 0 {
					found++
					nid = runner.ExtractNumber(col1, "", "ENSP")
					// MICRO-OPTIMIZATION: Cache the string conversion
					gidStr = strconv.FormatUint(uint64(app.eGeneIDs.UArr[idx]), 10)
				} else {
					missing++
					app.refiner.Mgr.Say(fmt.Sprintf("not found: %s", col1))
				}
			}

			if nid != "" {
				// MICRO-OPTIMIZATION: Bypass Fprintf reflection
				if err = cw.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", col2, nid, gidStr, score)); err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	app.refiner.Mgr.Say(fmt.Sprintf("total STRING DB rows counted: %d", row))
	app.refiner.Mgr.Say(fmt.Sprintf("total STRING DB records with GeneID: %d", found))
	app.refiner.Mgr.Say(fmt.Sprintf("total STRING DB records with no GeneID: %d", missing))
	return ska.Close()
}

// createProteinInteraction adds the GeneID to the STRING DB protein 2 and writes to temp file 2.
func (app *Application) createProteinInteraction(fp string) error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(fp); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	var nid, lid string
	var gidStr string // Cache the stringified Gene ID
	missing, found, row := 0, 0, 0
	idx := -1

	for row = 1; ska.Scan(); row++ {
		line := ska.Text()
		col1, rest, ok1 := strings.Cut(line, "\t")
		col2, rest, ok2 := strings.Cut(rest, "\t")
		col3, score, ok3 := strings.Cut(rest, "\t")

		if ok1 && ok2 && ok3 {
			if lid != col1 {
				lid = col1
				nid = ""
				idx = runner.BinarySearchItem(app.synonyms.SArr, col1)
				if idx >= 0 {
					found++
					nid = runner.ExtractNumber(col1, "", "ENSP")
					if nid == "" {
						continue
					}
					// MICRO-OPTIMIZATION: Cache the string conversion
					gidStr = strconv.FormatUint(uint64(app.eGeneIDs.UArr[idx]), 10)
				} else {
					missing++
					app.refiner.Mgr.Say(fmt.Sprintf("not found: %s", col1))
					continue
				}
			}

			if nid != "" {
				// MICRO-OPTIMIZATION: Bypass Fprintf reflection
				if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n", col3, gidStr, col2, nid, score); err != nil {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}

	app.refiner.Mgr.Say(fmt.Sprintf("total protein interaction rows counted: %d", row))
	app.refiner.Mgr.Say(fmt.Sprintf("total protein interaction records with GeneID: %d", found))
	app.refiner.Mgr.Say(fmt.Sprintf("total protein interaction records with no GeneID: %d", missing))
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}
