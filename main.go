package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"./email"
	"./gspread"
	"./wca"
)

const (
	turnLogsOn bool = true
)

func init() {
	if !(turnLogsOn) {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	// Using the `gspred` local package to get the data
	// of each recipient from Google SpreadSheets.
	recipients, err := gspread.GetRecipientsData()
	if err != nil {
		log.Fatal(err)
	}

	// For each recipient object, get the current number
	// of upcoming competitions, compare with the obsolete
	// value, and send a an email notifying if the value
	// has changed since the last verification.
	for _, recipient := range recipients {

		// Current date in format: "yyyy-MM-dd hh-mm-ss"
		recipient.CurrentVerificationDate = time.Now().String()[:19]

		// Using the `wca` local package to get the number
		// of upcoming competitions in the city of the recipient.
		// If an error happen, it is logged and the loop goes
		// to the next recipient in the slice.
		result, err := wca.UpcomingCopetitions(recipient.City.Value)
		recipient.CurrentUpcomingCompetitions = result
		if err != nil {
			log.Fatal(err)
			continue
		}

		log.Printf(
			"%v from %v has %v and now have %v upcoming competitions\n",
			recipient.Name.Value,
			recipient.City.Value,
			recipient.UpcomingCompetitions.Value,
			recipient.CurrentUpcomingCompetitions,
		)

		// Convert the `Upcoming Competitions` of the sheet
		// (a string) to an integer value, so it can be an
		// operating. If an error happen, it is logged and
		// the loop goes to the next recipient in the slice.
		upcomingCompetitionsInInteger, err := strconv.Atoi(recipient.UpcomingCompetitions.Value)
		if err != nil {
			log.Fatal(err)
			continue
		}

		// Update the recipient cell in the spreadsheet.
		// If an error happen, it is logged and the loop
		// goes to the next recipient in the slice.
		err = recipient.UpdateUpcomingCompetitions()
		if err != nil {
			log.Fatal(err)
			continue
		}

		// Test if the obsolete and the current value are the
		// same, if are it finishes the loop and goes to the
		// next recipient in the slice, if not it goes through
		// and send an email notifying the recipient.
		if upcomingCompetitionsInInteger == recipient.CurrentUpcomingCompetitions {
			continue
		}

		email.SendEmail(recipient)
	}
}
