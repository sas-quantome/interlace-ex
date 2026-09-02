/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31011

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

   0 DATE ADDED TO CATALOG
   1 PUBMEDID
  17 SNP_GENE_IDS
  23 SNP_ID_CURRENT
  34 MAPPED_TRAIT
  35 MAPPED_TRAIT_URI

*/

// gwasCatalogParser parses the EBI GWAS Catalog and creates the genomic variation asset.
func (app *Application) gwasCatalogParser() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	records, row, n := 0, 0, 0
	var cols [38]string
	var oids [128]string
	var terms [128]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 38 && row > 1 {
			records++

			// MICRO-OPTIMIZATION: Inline checks avoid function call overhead
			pid := cols[1]
			rid := cols[23]

			if pid == "" || pid == "-" || rid == "" || rid == "-" {
				continue
			}

			nOIDs := app.parseOIDs(cols[35], oids[:], terms[:], 128)
			if nOIDs >= 128 {
				app.refiner.Mgr.War("caught max number of ontology terms and associated ID pairs")
				continue
			}
			if nOIDs == 0 {
				oids[0] = ""
				terms[0] = cols[34]
				nOIDs = 1
				n++
			}

			buf := cols[17]
			for {
				uid, rest, found := strings.Cut(buf, ", ")
				if uid != "" {
					idx := runner.BinarySearchItem(app.synonyms.SArr, uid)
					if idx >= 0 {
						if gid := app.eGeneIDs.UArr[idx]; gid > 0 {

							// MICRO-OPTIMIZATION: Pre-format to string to bypass Fprintf reflection
							gidStr := strconv.FormatUint(uint64(gid), 10)

							for i := 0; i < nOIDs; i++ {
								if hid := app.refiner.Cat.HashText(terms[i]); hid > 0 {
									if err = app.refiner.Cat.WriteTuple(hid, terms[i]); err == nil {

										hashStr := strconv.FormatUint(hid, 10)

										if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\n",
											gidStr, rid, pid, oids[i], hashStr); err != nil {
											return runner.ValidateErrors(err, ska.Close())
										}
									}
								}
							}
						}
					}
				}
				if !found {
					break
				}
				buf = rest
			}
		}
	}
	if app.refiner.Mgr.GetBitFlag(runner.MgrFlgLoud) && len(app.unk) > 0 {
		for _, obo := range app.unk {
			app.refiner.Mgr.Say("ontology ID not found: " + obo)
		}
	}
	app.refiner.ReportCounts(row, records, "EBI GWAS Catalog")
	app.refiner.Mgr.Say(fmt.Sprintf("unique ontology IDs not found: %d", len(app.unk)))
	app.refiner.Mgr.Say(fmt.Sprintf("tuples without ontology ID: %d", n))
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}

// parseOIDs parses a comma-separated string of URLs, extracts ontology ID, and reformats the ID to the OBO-ID standard.
// Searches the EBO term from the ontology ID.
func (app *Application) parseOIDs(s string, oids, terms []string, size int) int {
	buf := s
	i := 0
	for i < size {
		urlStr, rest, found := strings.Cut(buf, ", ")
		urlStr = strings.TrimSpace(urlStr)

		if urlStr != "" {
			// Fast extraction without path.Base
			if idx := strings.LastIndexByte(urlStr, '/'); idx >= 0 {
				oid := urlStr[idx+1:]
				if oid != "" {
					obo := replaceUnderscore(oid)
					if term, ok := app.obo[obo]; ok {
						oids[i] = obo
						terms[i] = term
						i++
					} else {
						app.unk = runner.InsertOnce(app.unk, obo)
					}
				}
			}
		}
		if !found {
			break
		}
		buf = rest
	}
	return i
}
