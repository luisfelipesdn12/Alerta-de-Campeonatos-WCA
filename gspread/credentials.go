package gspread

import (
	"gopkg.in/Iwark/spreadsheet.v2"
)

// CredentialStruct is the struct based in the Sheet
// rows format.
type CredentialStruct struct {
	Email    string
	Password string
}

const (
	spreadsheetID string = "1CVO5kb-4Rjga1scWIQHSOzA9b7JgC2RZtUe1_bjAMKM"
)

// GetCredentialsData fetch the data stored in the
// spreadsheet and tranform it in a `CredentialStruct`
// as described above.
func GetCredentialsData() (CredentialStruct, error) {

	credentials := CredentialStruct{}

	// Make the connection to the gspread API using
	// the `client_secret.json` file in the project
	// root. If an error happen, the function returns
	// a empty `CredentialStruct` and the error.
	service, err := spreadsheet.NewService()
	if err != nil {
		return credentials, err
	}

	// Fetch the spreadsheet with the `spreadsheetID` value.
	// If an error happen, the function returns a empty
	// `CredentialStruct` and the error.
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return credentials, err
	}

	// Fetch the specific credentials sheet with the
	// `spreadsheet` value. If an error happen, the
	// function returns a empty `CredentialStruct`
	// and the error.
	credentialSheet, err := spreadsheet.SheetByTitle("Credentials")
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
