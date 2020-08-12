/*
	Alerta-de-Campeonatos-WCA - A script witch send an e-mail when there's a new WCA competition.
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

package gspread

import (
	"log"

	"gopkg.in/Iwark/spreadsheet.v2"
)

// CredentialStruct is the struct based in the Sheet
// rows format.
type CredentialStruct struct {
	Email    string
	Password string
}

// GetCredentialsData fetch the data stored in the
// spreadsheet and transform it in a `CredentialStruct`
// as described above.
func GetCredentialsData(spreadData spreadsheet.Spreadsheet) (CredentialStruct, error) {

	credentials := CredentialStruct{}

	// Fetch the specific credentials sheet with the
	// `spreadData` value. If an error happen, the
	// function returns a empty `CredentialStruct`
	// and the error.
	log.Println(`Fetching the specific sheet "Credentials"`)
	credentialSheet, err := spreadData.SheetByTitle("Credentials")
	if err != nil {
		return credentials, err
	}

	// Select the row where email and password is allocated.
	rowCells := credentialSheet.Rows[1]
	credentials = CredentialStruct{
		Email:    rowCells[0].Value,
		Password: rowCells[1].Value,
	}

	return credentials, nil
}
