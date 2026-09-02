/*
 * Copyright 2026 Quantome SAS
 * SPDX-License-Identifier: Apache-2.0
 */

package a31011

// oboParser loads the HP Mappings.
func (app *Application) oboParser() error {
	app.obo = make(map[string]string, 128*1024)
	if err := app.gwasEboParser(); err != nil {
		return err
	}
	if err := app.refiner.OboParser(hpoInpFile, "HPO", "id: HP:", app.obo); err != nil {
		return err
	}
	if err := app.refiner.OboParser(obaInpFile, "OBA", "id: ", app.obo); err != nil {
		return err
	}
	if err := app.refiner.OboParser(gssInpFile, "GSSO", "id: ", app.obo); err != nil {
		return err
	}
	if err := app.refiner.OboParser(monInpFile, "MONDO", "id: ", app.obo); err != nil {
		return err
	}
	return nil
}
