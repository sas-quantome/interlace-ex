/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import (
	"fmt"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

const maxIdMap = 1024 * 4

// idTrio holds the UniProt Map ID trios for a single GeneID.
type idTrio struct {
	lastAliasIndex, products int
	protein                  []string
	alias                    []string
	cat                      []int
}

// uniProtIdMap joins the id trio.
func (app *Application) uniProtIdMap() error {
	var err error
	app.ids.protein = make([]string, maxIdMap)
	app.ids.alias = make([]string, maxIdMap)
	app.ids.cat = make([]int, maxIdMap)
	app.ids.products = -1 // CRITICAL FIX: Initialize to -1 (not found)

	gidCode := app.refiner.Enum.Assure("GeneID")
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(idsInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	last := "nada"
	noGeneID, records, row := 0, 0, 0
	var cols [3]string

	for row = 1; ska.Scan(); row++ {
		if runner.FillCols(ska.Text(), cols[:]) >= 3 {
			if cols[1] == "GI" || cols[0] == "" || cols[2] == "" || cols[2] == "-" {
				continue
			}
			records++
			uniprotID, _, _ := strings.Cut(cols[0], "-")

			if last == uniprotID {
				app.ids.addTrio(last, cols[2], app.refiner.Enum.Assure(cols[1]), gidCode)
			} else {
				last = uniprotID
				if err = app.saveProtein(&noGeneID); err == nil {
					app.ids.lastAliasIndex = 0
					app.ids.products = -1 // CRITICAL FIX: Reset to -1 for the new cluster
					app.ids.addTrio(last, cols[2], app.refiner.Enum.Assure(cols[1]), gidCode)
				} else {
					return runner.ValidateErrors(err, ska.Close())
				}
			}
		}
	}
	err = runner.ValidateErrors(app.saveProtein(&noGeneID), app.refiner.FlushAll())
	app.refiner.ReportCounts(row, records, "UniProt ID Map")
	app.refiner.Mgr.Say(fmt.Sprintf("total UniProt ID Map clusters without GeneID: %d", noGeneID))
	return runner.ValidateErrors(err, ska.Close())
}

// addTrio appends identifiers and updates indexes.
func (aka *idTrio) addTrio(protein, alias string, cat, gid int) {
	if aka.lastAliasIndex < maxIdMap {
		aka.protein[aka.lastAliasIndex] = protein
		aka.alias[aka.lastAliasIndex] = alias
		aka.cat[aka.lastAliasIndex] = cat
	} else {
		aka.protein = append(aka.protein, protein)
		aka.alias = append(aka.alias, alias)
		aka.cat = append(aka.cat, cat)
	}
	if cat == gid {
		aka.products = aka.lastAliasIndex // Captures the index (which could be 0)
	}
	aka.lastAliasIndex++
}
