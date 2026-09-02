/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31010

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

/*

Input file:

   0 #AlleleID
   3 GeneID
   6 ClinicalSignificance : bit
   9 RS# (dbSNP)
  11 RCVaccession
  13 PhenotypeList : cat
  14 Origin : bit
  24 ReviewStatus : bit
  28 OtherIDs : https://reg.clinicalgenome.org/redmine/projects/registry/genboree_registry/by_caid?caid=CA251392
  30 VariationID
  34 SomaticClinicalImpact : bit
  37 Oncogenicity : bit
  40 SCVsForAggregateGermlineClassification
  41 SCVsForAggregateSomaticClinicalImpact
  42 SCVsForAggregateOncogenicityClassification

*/

// clinVarSummaryParser parses the ClinVar summary file and creates the Genetic Variation asset.
func (app *Application) clinVarSummaryParser() error {
	var err error
	var xl runner.XlBitMask
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}

	zeros, records, row := 0, 0, 0
	var cols [43]string

	for row = 1; ska.Scan(); row++ {
		// Fill the pre-allocated array (Zero-allocation string splitting)
		if runner.FillCols(ska.Text(), cols[:]) >= 43 && cols[0][0] != '#' {
			records++

			gid := cols[3]
			vid := cols[30]

			if gid != "" && vid != "" {

				// MICRO-OPTIMIZATION: Hoist calculations OUTSIDE the phenotype loop
				xl.Flag = 0
				xl.Add(app.setClinVarFlags(cols[6]))
				xl.Add(app.setClinVarFlags(cols[14]))
				xl.Add(app.setClinVarFlags(cols[24]))
				xl.Add(app.setClinVarFlags(cols[34]))
				xl.Add(app.setClinVarFlags(cols[37]))

				// Pre-format flag string to bypass Fprintf reflection
				flagStr := strconv.FormatUint(xl.Flag, 10)

				rid := cols[9]
				if rid == "-1" || rid == "" {
					rid = "0"
				}

				mim, cai := omimAndCanonicalAlleleId(cols[28])

				buf := cols[13]
				for {
					txt, rest, found := strings.Cut(buf, "|")
					if txt != "" && txt != "-" {
						if hid := app.refiner.Cat.HashText(txt); hid > 0 {
							if err = app.refiner.Cat.WriteTuple(hid, txt); err == nil {

								// Bypass reflection for the Hash ID
								hashStr := strconv.FormatUint(hid, 10)
								if gid == "-1" {
									gid = "0"
									zeros++
								}
								if err = app.refiner.Writer("%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
									gid, vid, rid, mim, flagStr, cai, hashStr); err != nil {
									return runner.ValidateErrors(err, ska.Close())
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
	}
	app.refiner.Mgr.Say(fmt.Sprintf("total ClinVar summary records without GeneID: %d", zeros))
	app.refiner.ReportCounts(row, records, "ClinVar summary")
	return runner.ValidateErrors(app.refiner.FlushAll(), ska.Close())
}

func omimAndCanonicalAlleleId(s string) (string, string) {
	if s == "" || s == "-" {
		return "0", "0"
	}

	o := "0"
	c := "0"

	// ClinGen:CA114071,UniProtKB:P00387#VAR_004620,OMIM:613213.0004
	if k := strings.Index(s, "OMIM:"); k != -1 {
		var buf strings.Builder
		for i := k + 5; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				buf.WriteByte(s[i])
			} else {
				break
			}
		}
		if buf.Len() > 0 {
			o = buf.String()
		}
	}

	if k := strings.Index(s, "ClinGen:CA"); k != -1 {
		var buf strings.Builder
		for i := k + 10; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				buf.WriteByte(s[i])
			} else {
				break
			}
		}
		if buf.Len() > 0 {
			c = buf.String()
		}
	}
	return o, c
}
