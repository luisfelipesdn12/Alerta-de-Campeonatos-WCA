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
	"strconv"
	"strings"

	"gopkg.in/Iwark/spreadsheet.v2"
)

// RecipientStruct is the struct based in the Sheet rows
// format. Every recipient has a name, email, city, upcoming
// competitions number and when was the last time that this
// number was checked. All this properties are not literal
// values, are `spreadsheet.Cell` objects witch contains useful
// information about the cell where the value is stored in
// addition to the value itself.
type RecipientStruct struct {
	// Basic data, for send emails and do the verifications.
	Name  spreadsheet.Cell
	Email spreadsheet.Cell
	City  spreadsheet.Cell

	// Language, for emails be responsive to it.
	Language spreadsheet.Cell

	// The upcoming competitions number in the last verification.
	UpcomingCompetitions spreadsheet.Cell
	// The date/time information in string of the last verification.
	LastVerification spreadsheet.Cell

	// The pointer to spreadsheet object itself, it helps
	// to synchronize changes using just the `RecipientStruct`,
	// by example.
	Sheet *spreadsheet.Sheet

	// The upcoming competitions number that was/will be
	// checked in this exactly runtime. Can be a zero value
	// if it was not checked yet.
	CurrentUpcomingCompetitions int
	// The date/time information in string of the
	// verification in this exactly runtime. Can be
	// a zero value if it was not checked yet.
	CurrentVerificationDate string
}

// GetRecipientsData fetch the data stored in the
// spreadsheet, transform it in a slice of `RecipientStruct`
// as described above.
func GetRecipientsData(spreadData spreadsheet.Spreadsheet) ([]RecipientStruct, error) {

	recipients := []RecipientStruct{}

	// Fetch the specific recipients sheet with the
	// `spreadData` value. If an error happen, the
	// function returns a empty slice of `RecipientStruct`
	// and the error.
	log.Println(`Fetching the specific sheet "Recipients"`)
	recipientsSheet, err := spreadData.SheetByTitle("Recipients")
	if err != nil {
		return recipients, err
	}

	// Uses the exact recipients number to make a slice
	// where the final `RecipientStruct` values will be
	// stored. The "-1" is because the columns title is
	// also considerated a row, but it will not be returned.
	recipients = make(
		[]RecipientStruct, 0, len(recipientsSheet.Rows)-1,
	)

	// Transform the each row in the `RecipientStruct`
	// and append it to the `recipients` slice.
	for row := 1; row < len(recipientsSheet.Rows); row++ {

		// returns a slice of `spreadsheet.Cell` objects
		// that cells are equivalent to the data of a
		// specific recipient.
		rowCells := recipientsSheet.Rows[row]

		// Using the position in the spreadsheet,
		// defina a `RecipientStruct`.
		recipientData := RecipientStruct{
			Name:                 rowCells[1],
			Email:                rowCells[2],
			City:                 rowCells[3],
			Language:             rowCells[4],
			UpcomingCompetitions: rowCells[5],
			LastVerification:     rowCells[6],
			Sheet:                recipientsSheet,
		}

		// The properties `Name`, `Email` and `City`
		// are free user inputs, so it can be with
		// spaces in extremities like " My Name" or
		// "   My City " . It is very bad when the
		// program sends an email or verify if a city
		// already exists in `main.cityUpcomingCompetitionsCache`.
		// So the lines below strip this strings using
		// the function `stripIfNecessary`.
		stripIfNecessary(&recipientData.Name.Value)
		stripIfNecessary(&recipientData.Email.Value)
		stripIfNecessary(&recipientData.City.Value)

		// Add the `RecipientStruct` in the slice.
		recipients = append(recipients, recipientData)
	}

	return recipients, nil
}

// UpdateUpcomingCompetitions update the recipient date
// in the spreadsheet with the new upcoming competitions
// number and last verifications date.
func (recipient RecipientStruct) UpdateUpcomingCompetitions() error {

	log.Printf("Updating and Syncronyzing the upcoming competitions from %v\n", recipient.Name.Value)

	// Update the `UpcomingCompetitions` propertie
	// in the spreadsheet.
	recipient.Sheet.Update(
		int(recipient.UpcomingCompetitions.Row),
		int(recipient.UpcomingCompetitions.Column),
		strconv.Itoa(recipient.CurrentUpcomingCompetitions),
	)

	// Update the `LastVerification` propertie
	// in the spreadsheet.
	recipient.Sheet.Update(
		int(recipient.LastVerification.Row),
		int(recipient.LastVerification.Column),
		string(recipient.CurrentVerificationDate),
	)

	// Synchronize the updates with the original
	// spreadsheet in the Google account.
	err := recipient.Sheet.Synchronize()
	if err != nil {
		return err
	}

	return nil
}

// stripIfNecessary delete all spaces in the prefix
// and suffix string, given it pointer.
func stripIfNecessary(s *string) {
	if strings.HasPrefix(*s, "") || strings.HasSuffix(*s, " ") {
		*s = strings.TrimSpace(*s)
	}
}
