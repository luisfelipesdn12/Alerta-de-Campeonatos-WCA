package gspread

import (
	"log"
	"strconv"

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
	Language                    spreadsheet.Cell
	UpcomingCompetitions        spreadsheet.Cell
	LastVerification            spreadsheet.Cell
	Sheet                       *spreadsheet.Sheet
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
	log.Println("Connecting with Google SpreadSheets API")
	service, err := spreadsheet.NewService()
	if err != nil {
		return recipients, err
	}

	// Fetch the spreadsheet with the `spreadsheetID` value.
	// If an error happen, the function returns a empty
	// slice of `RecipientStruct` and the error
	log.Println("Fetching the spreadsheet in account")
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	if err != nil {
		return recipients, err
	}

	// Fetch the specific recipients sheet with the
	// `spreadsheet` value. If an error happen, the
	// function returns a empty slice of `RecipientStruct`
	// and the error
	log.Println(`Fetching the specific sheet "Recipients"`)
	recipientsSheet, err := spreadsheet.SheetByTitle("Recipients")
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
			Language:             rowCells[4],
			UpcomingCompetitions: rowCells[5],
			LastVerification:     rowCells[6],
			Sheet:                recipientsSheet,
		}
		recipients = append(recipients, recipientData)
	}

	return recipients, nil
}

func (recipient RecipientStruct) UpdateUpcomingCompetitions() error {
	recipient.Sheet.Update(
		int(recipient.UpcomingCompetitions.Row),
		int(recipient.UpcomingCompetitions.Column),
		strconv.Itoa(recipient.CurrentUpcomingCompetitions),
	)

	recipient.Sheet.Update(
		int(recipient.LastVerification.Row),
		int(recipient.LastVerification.Column),
		string(recipient.CurrentVerificationDate),
	)

	log.Printf("Updating and Syncronyzing the upcoming competitions from %v\n", recipient.Name.Value)
	err := recipient.Sheet.Synchronize()
	if err != nil {
		return err
	}

	return nil
}
