/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31003

import "github.com/sas-quantome/interlace-go/pkg/runner"

const (
	GeneBiologicalRegion uint32 = 1 << iota
	GeneMiscRNA
	GeneNcRNA
	GeneOther
	GeneProteinCoding
	GenePseudo
	GeneRRNA
	GeneScRNA
	GeneSnRNA
	GeneSnoRNA
	GeneTRNA
	GeneUnknown
)

// bitFlagsEnum returns the flag names aligned to the masking order.
func bitFlagsEnum() []string {
	a := []string{"biological-region",
		"miscRNA",
		"ncRNA",
		"other-RNA",
		"protein-coding",
		"pseudogene",
		"rRNA",
		"scRNA",
		"snRNA",
		"snoRNA",
		"tRNA"}
	return a
}

// setGeneFlags captures the gene type and sets the corresponding flag.
func setGeneFlags(flag *runner.LgBitMask, gType string) {
	flag.Flag = 0
	if gType != "" {
		switch gType {
		case "biological-region":
			flag.Add(GeneBiologicalRegion)
		case "miscRNA":
			flag.Add(GeneMiscRNA)
		case "ncRNA":
			flag.Add(GeneNcRNA)
		case "other":
			flag.Add(GeneOther)
		case "protein-coding":
			flag.Add(GeneProteinCoding)
		case "pseudo":
			flag.Add(GenePseudo)
		case "rRNA":
			flag.Add(GeneRRNA)
		case "scRNA":
			flag.Add(GeneScRNA)
		case "snRNA":
			flag.Add(GeneSnRNA)
		case "snoRNA":
			flag.Add(GeneSnoRNA)
		case "tRNA":
			flag.Add(GeneTRNA)
		default:
			flag.Add(GeneUnknown)
		}
	}
}
