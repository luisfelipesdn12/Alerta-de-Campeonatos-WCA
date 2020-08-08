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
// spreadsheet and tranform it in a `CredentialStruct`
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
