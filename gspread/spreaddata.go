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

// Package gspread is a local package and implements
// functions to access recipients and credentials
// data from the Google account with the Google Sheets API.
package gspread

import (
	"log"

	"gopkg.in/Iwark/spreadsheet.v2"
)

const (
	// The ID of the spreadsheet in my Google account
	// it is visible in the URI when you are with the
	// sheet opened.
	spreadsheetID string = "1CVO5kb-4Rjga1scWIQHSOzA9b7JgC2RZtUe1_bjAMKM"
)

// GetSpreadData fetch the data stored in the
// Google SpreadSheet account and return the
// specififc spreadsheet to this project.
func GetSpreadData() (spreadsheet.Spreadsheet, error) {

	spreadData := spreadsheet.Spreadsheet{}

	// Make the connection to the gspread API using
	// the `client_secret.json` file in the project
	// root. If an error happen, the function returns
	// a empty `spreadsheet.Spreadsheet` and the error.
	log.Println("Connecting with Google SpreadSheets API")
	service, err := spreadsheet.NewService()
	if err != nil {
		log.Fatal(err)
		return spreadData, err
	}

	// Fetch the spreadsheet with the `spreadsheetID` value.
	// If an error happen, the function returns a empty
	// `spreadsheet.Spreadsheet` and the error.
	log.Println("Fetching the spreadsheet in account")
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		log.Fatal(err)
		return spreadData, err
	}

	return spreadsheet, nil
}
