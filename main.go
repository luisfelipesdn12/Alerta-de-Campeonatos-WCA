package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"./email"
	"./gspread"
	"./wca"
)

const (
	turnLogsOn bool = true
)

var (
	// This log file variable is defined globally
	// because it need to be visible for init() and
	// main().
	logFile, err = os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
)

// init basically check the option turnLogsOn above,
// if this options is true, it set the `logFile` as
// output of logs; if is false, if discard the logs.
func init() {
	if !(turnLogsOn) {
		log.SetOutput(ioutil.Discard)
	} else {
		// this check the error of the var declaration
		// with `os.OpenFile()`
		checkError(err)

		// this clear the `logFile` before starting
		// to write new logs on it
		err = os.Truncate("main.log", 0)
		checkError(err)

		log.SetOutput(logFile)
	}
}

func main() {

	defer logFile.Close()

	// Using the `gspred` local package to get the data
	// of each recipient from Google SpreadSheets.
	recipients, err := gspread.GetRecipientsData()
	checkError(err)

	// Using the `gspred` local package to get the
	// credentials data from Google SpreadSheets. It
	// will be used to send the emails below.
	credentials, err := gspread.GetCredentialsData()
	checkError(err)

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

		email.SendEmail(recipient, credentials)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
