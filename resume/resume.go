/*
	Alerta-de-Campeonatos-WCA - A script which send an e-mail when there's a new WCA competition.
	Copyright (C) 2020  Luis Felipe Santos do Nascimento

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Package resume is a local package and implements
// objects and functions to export a resume.json when
// some verification is complete.
package resume

import (
	"encoding/json"
	"io/ioutil"
)

// Information is the data witch will be exported
// to the json file.
type Information struct {
	EmailsSended    int
	RequestsSended  int
	UsersChecked    int
	StartIn         string
	RuntimeDuration string
}

// ExportResume transform the `Information` to a json file.
func (info Information) ExportResume(dir string) error {
	jsonString, err := json.MarshalIndent(info, "", "\t")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dir, jsonString, 0644)

	if err != nil {
		return err
	}

	return nil
}
