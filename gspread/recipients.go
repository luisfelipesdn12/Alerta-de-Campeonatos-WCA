package gspread

import (
	"gopkg.in/Iwark/spreadsheet.v2"
)

// RecipientStruct is the struct based in the Sheet rows
// format. Every recipient has a name, email, city, upcoming
// competitions number and when was the last time that this
// number was checked. All this properties are not literal
// values, are `spreadsheet.Cell` objects witch contains useful
// information about the cell where the value is stored in
// addition to the value itself
type RecipientStruct struct {
	Name                        spreadsheet.Cell
	Email                       spreadsheet.Cell
	City                        spreadsheet.Cell
	UpcomingCompetitions        spreadsheet.Cell
	LastVerification            spreadsheet.Cell
	CurrentUpcomingCompetitions int
	CurrentVerificationDate     string
}

// GetRecipientsData fetch the data stored in the
// spreadsheet, tranform it in a slice of `RecipientStruct`
// as described above
func GetRecipientsData() ([]RecipientStruct, error) {

	recipients := []RecipientStruct{}

	// Make the connection to the gspread API using
	// the `client_secret.json` file in the project
	// root. If an error happen, the function returns
	// a empty slice of `RecipientStruct` and the error
	service, err := spreadsheet.NewService()
	if err != nil {
		return recipients, err
	}

	// Fetch the spreadsheet with the `spreadsheetID` value.
	// If an error happen, the function returns a empty
	// slice of `RecipientStruct` and the error
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return recipients, err
	}

	// Fetch the specific recipients sheet with the
	// `spreadsheet` value. If an error happen, the
	// function returns a empty slice of `RecipientStruct`
	// and the error
	recipientsSheet, err := spreadsheet.SheetByTitle("Betas")
	if err != nil {
		return recipients, err
	}

	// Uses the exact recipients number to make a slice
	// where the final `RecipientStruct` values will be
	// stored
	recipients = make(
		[]RecipientStruct, 0, len(recipientsSheet.Rows)-1,
	)

	// Tranform the each row in the `RecipientStruct`
	// and append it to the `recipients` slice
	for row := 1; row < len(recipientsSheet.Rows); row++ {

		rowCells := recipientsSheet.Rows[row]

		recipientData := RecipientStruct{
			Name:                 rowCells[1],
			Email:                rowCells[2],
			City:                 rowCells[3],
			UpcomingCompetitions: rowCells[4],
			LastVerification:     rowCells[5],
		}
		recipients = append(recipients, recipientData)
	}

	return []RecipientStruct{recipients[0]}, nil
	// return recipients, nil
}
