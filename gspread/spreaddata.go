package gspread

import (
	"log"

	"gopkg.in/Iwark/spreadsheet.v2"
)

const (
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
		return spreadData, err
	}

	// Fetch the spreadsheet with the `spreadsheetID` value.
	// If an error happen, the function returns a empty
	// `spreadsheet.Spreadsheet` and the error.
	log.Println("Fetching the spreadsheet in account")
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return spreadData, err
	}

	return spreadsheet, nil
}
