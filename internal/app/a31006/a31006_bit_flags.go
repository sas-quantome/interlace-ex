/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31006

import "github.com/sas-quantome/interlace-go/pkg/runner"

const (
	GoComponent uint64 = 1 << iota
	GoFunction
	GoProcess
	GoEXP
	GoHDA
	GoHEP
	GoHGI
	GoHMP
	GoHTP
	GoIBA
	GoIC
	GoIDA
	GoIEA
	GoIEP
	GoIGC
	GoIGI
	GoIKR
	GoIMP
	GoIPI
	GoISA
	GoISM
	GoISO
	GoISS
	GoNAS
	GoND
	GoRCA
	GoTAS
	GoNotActsUpstreamOf
	GoNotActsUpstreamOfOrWithin
	GoNotActsUpstreamOfOrWithinNegativeEffect
	GoNotColocalizesWith
	GoNotContributesTo
	GoNotEnables
	GoNotInvolvedIn
	GoNotIsActiveIn
	GoNotLocatedIn
	GoNotPartOf
	GoActsUpstreamOf
	GoActsUpstreamOfNegativeEffect
	GoActsUpstreamOfOrWithin
	GoActsUpstreamOfOrWithinNegativeEffect
	GoActsUpstreamOfOrWithinPositiveEffect
	GoActsUpstreamOfPositiveEffect
	GoColocalizesWith
	GoContributesTo
	GoEnables
	GoInvolvedIn
	GoIsActiveIn
	GoLocatedIn
	GoPartOf
)

// bitFlagsEnum returns the flag names aligned to the masking order.
func bitFlagsEnum() []string {
	a := []string{"component",
		"function",
		"process",
		"EXP",
		"HDA",
		"HEP",
		"HGI",
		"GHMP",
		"HTP",
		"IBA",
		"IC",
		"IDA",
		"IEA",
		"IEP",
		"IGC",
		"IGI",
		"IKR",
		"IMP",
		"IPI",
		"ISA",
		"ISM",
		"ISO",
		"ISS",
		"NAS",
		"ND",
		"RCA",
		"TAS",
		"not_acts_upstream_of",
		"not_acts_upstream_of_or_within",
		"not_acts_upstream_of_or_within_negative_effect",
		"not_colocalizes_with",
		"not_contributes_to",
		"not_enables",
		"not_involved_in",
		"not_is_active_in",
		"not_located_in",
		"not_part_of",
		"acts_upstream_of",
		"acts_upstream_of_negative_effect",
		"acts_upstream_of_or_within",
		"acts_upstream_of_or_within_negative_effect",
		"acts_upstream_of_or_within_positive_effect",
		"acts_upstream_of_positive_effect",
		"colocalizes_with",
		"contributes_to",
		"enables",
		"involved_in",
		"is_active_in",
		"located_in",
		"part_of"}
	return a
}

func setGoFlags(xl *runner.XlBitMask, evidence, category, qualifier string) {
	// If you are using the generic BitMask we refactored earlier, it has a Clear method.
	// We can manually reset it to 0 here to be safe before adding flags.
	xl.Flag = 0

	if category != "" {
		switch category {
		case "Component":
			xl.Add(GoComponent)
		case "Function":
			xl.Add(GoFunction)
		case "Process":
			xl.Add(GoProcess)
		}
	}

	if evidence != "" {
		switch evidence {
		case "EXP":
			xl.Add(GoEXP)
		case "HDA":
			xl.Add(GoHDA)
		case "HEP":
			xl.Add(GoHEP)
		case "HGI":
			xl.Add(GoHGI)
		case "HMP":
			xl.Add(GoHMP)
		case "HTP":
			xl.Add(GoHTP)
		case "IBA":
			xl.Add(GoIBA)
		case "IC":
			xl.Add(GoIC)
		case "IDA":
			xl.Add(GoIDA)
		case "IEA":
			xl.Add(GoIEA)
		case "IEP":
			xl.Add(GoIEP)
		case "IGC":
			xl.Add(GoIGC)
		case "IGI":
			xl.Add(GoIGI)
		case "IKR":
			xl.Add(GoIKR)
		case "IMP":
			xl.Add(GoIMP)
		case "IPI":
			xl.Add(GoIPI)
		case "ISA":
			xl.Add(GoISA)
		case "ISM":
			xl.Add(GoISM)
		case "ISO":
			xl.Add(GoISO)
		case "ISS":
			xl.Add(GoISS)
		case "NAS":
			xl.Add(GoNAS)
		case "RCA":
			xl.Add(GoRCA)
		case "TAS":
			xl.Add(GoTAS)
		default:
			xl.Add(GoND)
		}
	}

	if qualifier != "" {
		switch qualifier {
		case "NOT acts_upstream_of":
			xl.Add(GoNotActsUpstreamOf)
		case "NOT acts_upstream_of_or_within":
			xl.Add(GoNotActsUpstreamOfOrWithin)
		case "NOT acts_upstream_of_or_within_negative_effect":
			xl.Add(GoNotActsUpstreamOfOrWithinNegativeEffect)
		case "NOT colocalizes_with":
			xl.Add(GoNotColocalizesWith)
		case "NOT contributes_to":
			xl.Add(GoNotContributesTo)
		case "NOT enables":
			xl.Add(GoNotEnables)
		case "NOT involved_in":
			xl.Add(GoNotInvolvedIn)
		case "NOT is_active_in":
			xl.Add(GoNotIsActiveIn)
		case "NOT located_in":
			xl.Add(GoNotLocatedIn)
		case "NOT part_of":
			xl.Add(GoNotPartOf)
		case "acts_upstream_of":
			xl.Add(GoActsUpstreamOf)
		case "acts_upstream_of_negative_effect":
			xl.Add(GoActsUpstreamOfNegativeEffect)
		case "acts_upstream_of_or_within":
			xl.Add(GoActsUpstreamOfOrWithin)
		case "acts_upstream_of_or_within_negative_effect":
			xl.Add(GoActsUpstreamOfOrWithinNegativeEffect)
		case "acts_upstream_of_or_within_positive_effect":
			xl.Add(GoActsUpstreamOfOrWithinPositiveEffect)
		case "acts_upstream_of_positive_effect":
			xl.Add(GoActsUpstreamOfPositiveEffect)
		case "colocalizes_with":
			xl.Add(GoColocalizesWith)
		case "contributes_to":
			xl.Add(GoContributesTo)
		case "enables":
			xl.Add(GoEnables)
		case "involved_in":
			xl.Add(GoInvolvedIn)
		case "is_active_in":
			xl.Add(GoIsActiveIn)
		case "located_in":
			xl.Add(GoLocatedIn)
		case "part_of":
			xl.Add(GoPartOf)
		}
	}
}
