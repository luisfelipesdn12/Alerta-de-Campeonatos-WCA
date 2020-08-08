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

	// Close the log file at the end of the block
	defer logFile.Close()

	// Using the `gspred` local package to do the
	// connection with the Google SpreadSheet API,
	// fetch the specific spreadsheet of the project
	// and return the spreadsheet data.
	spreadData, err := gspread.GetSpreadData()

	// Using the `gspred` local package to get the data
	// of each recipient from Google SpreadSheets.
	recipients, err := gspread.GetRecipientsData(spreadData)
	checkError(err)

	// Using the `gspred` local package to get the
	// credentials data from Google SpreadSheets. It
	// will be used to send the emails below.
	credentials, err := gspread.GetCredentialsData(spreadData)
	checkError(err)

	// Create a map to allocate the cities and the
	// upcoming competitions number. It will be used
	// to check if the verification already exists and
	// improve the performance.
	cityUpcomingCompetitionsCache := make(map[string]int)

	// For each recipient object, get the current number
	// of upcoming competitions, compare with the obsolete
	// value, and send a an email notifying if the value
	// has changed since the last verification.
	for _, recipient := range recipients {

		// Current date in format: "yyyy-MM-dd hh-mm-ss"
		recipient.CurrentVerificationDate = time.Now().String()[:19]

		// If the value of upcoming competitions to the
		// recipients city already exists in the map
		// `cityUpcomingCompetitionsCache`, uses this value.
		// If not, uses the `wca` local package to access the
		// WCA's API and put the result in the above-mentioned
		// map, so it can be reused in the next times.
		if result, ok := cityUpcomingCompetitionsCache[recipient.City.Value]; ok {
			log.Printf(
				"The cache value for %v (%v) was reused to %v\n",
				recipient.City.Value,
				result,
				recipient.Name.Value,
			)
			recipient.CurrentUpcomingCompetitions = result
		} else {
			// Using the `wca` local package to get the number
			// of upcoming competitions in the city of the recipient.
			// If an error happen, it is logged and the loop goes
			// to the next recipient in the slice.
			result, err := wca.UpcomingCopetitions(recipient.City.Value)
			if err != nil {
				log.Fatal(err)
				continue
			}

			// If there is no errors, the result value is attributed
			// to the propriertie `CurrentUpcomingCompetitions` in the
			// `RecipientStruct` and goes to `cityUpcomingCompetitionsCache`
			// map, so it can be reused in the next times.
			recipient.CurrentUpcomingCompetitions = result
			cityUpcomingCompetitionsCache[recipient.City.Value] = result
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
		// operating. If an error happen because the
		// `recipient.UpcomingCompetitions.Value` is a
		// non-convertable string (for example, when this
		// is the first verification of the recipient and
		// the value is ""), the value will be defined as the
		// current upcoming competitions number and no emails
		// will be sended. If non-predicate an error happen,
		// it is logged and the loop goes to the next recipient
		// in the slice.
		upcomingCompetitionsInInteger, err := strconv.Atoi(recipient.UpcomingCompetitions.Value)
		if err != nil {
			if err.Error() == `strconv.Atoi: parsing "`+recipient.UpcomingCompetitions.Value+`": invalid syntax` {
				upcomingCompetitionsInInteger = recipient.CurrentUpcomingCompetitions
			} else {
				continue
			}
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

	log.Printf("The cache was this: %v\n", cityUpcomingCompetitionsCache)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
