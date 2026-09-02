/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31001

import "strings"

// Recall that order must align with the columns in the CW tsv file.
// app.refiner.AddBin("sdk_entrez_gene_id_u.bin")          // CW tsv column 0
// app.refiner.AddBin("udk_protein_id_s.bin")              // CW tsv column 1
// app.refiner.AddBin("udk_synonym_s.bin")                 // CW tsv column 2
// app.refiner.AddBin(app.refiner.PropertyEnumName(enumName)) // CW tsv column 3

// saveGeneEnsembl caches the IDs.
func (app *Application) saveGeneEnsembl(gid int64, alias, protein string, enumCode int) error {
	if gid < 1 || alias == "" || alias[0] != 'E' {
		return nil
	}
	return app.refiner.Writer("%d\t%s\t%s\t%d\n", gid, protein, decodeURL(alias), enumCode)
}

// saveXdbRef captures prefixed ids.
func (app *Application) saveXdbRef(gid int64, s string, hgn, img, mim, mir int) error {
	if s == "" || s == "-" {
		return nil
	}
	var err error
	t := s
	for {
		x, rest, found := strings.Cut(t, "|")
		if x != "" {
			if after, ok := strings.CutPrefix(x, "HGNC:HGNC:"); ok {
				err = app.refiner.Writer("%d\t\t%s\t%d\n", gid, after, hgn)
			} else if after, ok = strings.CutPrefix(x, "IMGT/GENE-DB:"); ok {
				err = app.refiner.Writer("%d\t\t%s\t%d\n", gid, after, img)
			} else if after, ok = strings.CutPrefix(x, "MIM:"); ok {
				err = app.refiner.Writer("%d\t\t%s\t%d\n", gid, after, mim)
			} else if after, ok = strings.CutPrefix(x, "miRBase:"); ok {
				err = app.refiner.Writer("%d\t\t%s\t%d\n", gid, after, mir)
			}
			if err != nil {
				return err
			}
		}
		if !found {
			break
		}
		t = rest
	}
	return nil
}

// saveSynonym writes the aliases and symbols.
func (app *Application) saveSynonym(gid int64, s string, c int) error {
	if s == "-" || s == "" {
		return nil
	}
	t := s
	for {
		x, rest, found := strings.Cut(t, "|")
		if x != "" {
			if err := app.refiner.Writer("%d\t\t%s\t%d\n", gid, decodeURL(x), c); err != nil {
				return err
			}
		}
		if !found {
			break
		}
		t = rest
	}
	return nil
}

// saveRefIds from RefSeq.
func (app *Application) saveRefIds(gid int64, alias, protein string, enumCode int) error {
	if alias == "" || alias == "-" {
		return nil
	}
	return app.refiner.Writer("%d\t%s\t%s\t%d\n", gid, protein, decodeURL(alias), enumCode)
}

// saveProtein IDs
func (app *Application) saveProtein(noGeneID *int) error {
	// CRITICAL FIX: Check if index is >= 0 (since 0 is a valid index)
	if app.ids.products >= 0 {
		for i := 0; i < app.ids.lastAliasIndex; i++ {
			if i != app.ids.products {
				if err := app.refiner.Writer("%s\t%s\t%s\t%d\n",
					app.ids.alias[app.ids.products], // The GeneID
					app.ids.protein[i],              // The UniProt ID
					decodeURL(app.ids.alias[i]),     // The external alias
					app.ids.cat[i]); err != nil {    // Category Code
					return err
				}
			}
		}
		return nil
	}
	*noGeneID = *noGeneID + 1
	return nil
}
