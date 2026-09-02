/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31010

import (
	"errors"
	"math"
	"slices"
	"strings"

	"github.com/sas-quantome/interlace-go/pkg/runner"
)

func (app *Application) bitFlagsToEnum() error {
	var err error
	ska := runner.NewScanner(0)
	if err = ska.Start(app.refiner.SrcInpFile(primaryInpFile)); err != nil {
		return runner.ValidateErrors(err, ska.Close())
	}
	var cols [43]string
	for ska.Scan() {
		if runner.FillCols(ska.Text(), cols[:]) >= 43 && cols[0][0] != '#' {
			gid := cols[3]
			vid := cols[30]
			if gid != "" && gid != "-" && vid != "" && vid != "-" {
				app.splitTerms(cols[6])
				app.splitTerms(cols[14])
				app.splitTerms(cols[24])
				app.splitTerms(cols[34])
				app.splitTerms(cols[37])
			}
		}
	}
	if len(app.terms) == 0 {
		return runner.ValidateErrors(errors.New("zero enum terms collected"), ska.Close())
	}
	slices.Sort(app.terms)
	return ska.Close()
}

func (app *Application) splitTerms(s string) {
	app.buf.Reset()
	if s == "" || s == "-" {
		return
	}
	for i := range s {
		if s[i] == ';' || s[i] == '/' {
			app.buf.WriteRune('\t')
		} else {
			app.buf.WriteByte(s[i])
		}
	}
	if runner.FillCols(app.buf.String(), app.cache[:]) > 0 {
		for _, v := range app.cache {
			if v != "" {
				app.terms = runner.InsertOnce(app.terms, strings.TrimSpace(v))
			}
		}
	}
}

func (app *Application) setClinVarFlags(s string) uint64 {
	if s == "" {
		return 0
	}
	var xl runner.XlBitMask
	app.terms = app.terms[:0]
	app.splitTerms(s)
	for _, t := range app.terms {
		e := app.refiner.Enum.GetCode(t)
		if e >= 0 && e < 64 {
			xl.Add(uint64(math.Pow(2, float64(e))))
		}
	}
	return xl.Flag
}
