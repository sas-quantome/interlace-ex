/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31015

const (
	AncNR uint32 = 1 << iota
	AncAboAus
	AncAfrAme
	AncAfrUns
	AncAsiUns
	AncCenAsi
	AncEasAsi
	AncEur
	AncGrMiEa
	AncHisLat
	AncMiEaPe
	AncNatAme
	AncOce
	AncOth
	AncOthAdm
	AncSouAsi
	AncSoEaAs
	AncSubSah
)

// bitFlagsEnum returns the flag names aligned to the masking order.
func bitFlagsEnum() []string {
	a := []string{"not_reported",
		"Aboriginal_Australian",
		"African_American_or_Afro-Caribbean",
		"African_unspecified",
		"Asian_unspecified",
		"Central_Asian",
		"East_Asian",
		"European",
		"Greater_Middle_Eastern",
		"Hispanic_or_Latin_American",
		"Greater_Middle_Eastern_Persian",
		"Native_American",
		"Oceanian",
		"other_ancestry",
		"other_admixed_ancestry",
		"South_Asian",
		"South_East_Asian",
		"Sub-Saharan_African"}
	return a
}

// getAncestryFlag returns the ancestry bit flag or warns if the ancestry is unknown.
// This is now purely functional and stateless.
func getAncestryFlag(ancestry string, app *Application) uint32 {
	switch ancestry {
	case "", "Greater Middle Eastern (Middle Eastern":
		return 0
	case "NR":
		return AncNR
	case "Aboriginal Australian":
		return AncAboAus
	case "African American or Afro-Caribbean":
		return AncAfrAme
	case "African unspecified":
		return AncAfrUns
	case "Asian unspecified":
		return AncAsiUns
	case "Central Asian":
		return AncCenAsi
	case "East Asian":
		return AncEasAsi
	case "European":
		return AncEur
	case "Greater Middle Eastern":
		return AncGrMiEa
	case "Hispanic or Latin American":
		return AncHisLat
	case "North African or Persian)":
		return AncMiEaPe
	case "Native American":
		return AncNatAme
	case "Oceanian":
		return AncOce
	case "Other":
		return AncOth
	case "Other admixed ancestry":
		return AncOthAdm
	case "South Asian":
		return AncSouAsi
	case "South East Asian":
		return AncSoEaAs
	case "Sub-Saharan African":
		return AncSubSah
	default:
		app.refiner.Mgr.War("unknown term: " + ancestry)
		return 0
	}
}
