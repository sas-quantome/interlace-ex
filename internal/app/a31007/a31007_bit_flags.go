/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31007

import (
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

const (
	OmimTypeIsGene uint32 = 1 << iota
	OmimTypeIsPhenotype
	OmimCommentHasQTL1
	OmimCommentHasQTL2
	OmimCommentHasModifier
	OmimCommentHasNonDisease
	OmimCommentHasQuestion
	OmimCommentHasSomatic
	OmimCommentHasSusceptibility
)

// bitFlagsEnum returns the flag names aligned to the masking order.
func bitFlagsEnum() []string {
	a := []string{"gene",
		"phenotype",
		"QTL1",
		"QTL2",
		"modifier",
		"nondisease",
		"question",
		"somatic",
		"susceptibility"}
	return a
}

// setOmimFlags flags the OMIM's type and comment categories.
func setOmimFlags(lg *runner.LgBitMask, oType, oComment string) {
	lg.Flag = 0
	if oType != "" {
		switch strings.TrimSpace(oType) {
		case "gene":
			lg.Add(OmimTypeIsGene)
		case "phenotype":
			lg.Add(OmimTypeIsPhenotype)
		}
	}
	if oComment != "" {
		buf := oComment
		for {
			txt, rest, found := strings.Cut(buf, ";")
			if txt != "" {
				switch txt {
				case "QTL 1":
					lg.Add(OmimCommentHasQTL1)
				case "QTL 2":
					lg.Add(OmimCommentHasQTL2)
				case "modifier":
					lg.Add(OmimCommentHasModifier)
				case "nondisease":
					lg.Add(OmimCommentHasNonDisease)
				case "question":
					lg.Add(OmimCommentHasQuestion)
				case "somatic":
					lg.Add(OmimCommentHasSomatic)
				case "susceptibility":
					lg.Add(OmimCommentHasSusceptibility)
				}
			}
			if !found {
				break
			}
			buf = rest
		}
	}
}
