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

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"./email"
	"./gspread"
	"./resume"
	"./wca"
)

const (
	turnLogsOn bool = true
)

var (
	startIn = time.Now()

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

	// Create a struct to allocate information about
	// the runtime. It will be exported to a json file,
	// but just when the runtime is complete successfully.
	resumeInformation := resume.Information{}

	// Using the `gspred` local package to do the
	// connection with the Google SpreadSheet API,
	// fetch the specific spreadsheet of the project
	// and return the spreadsheet data.
	spreadData, err := gspread.GetSpreadData()
	checkError(err)

	// Using the `gspred` local package to get the data
	// of each recipient from Google SpreadSheets.
	recipients, err := gspread.GetRecipientsData(spreadData)
	checkError(err)
	resumeInformation.UsersChecked = len(recipients)

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
			resumeInformation.RequestsSended++
			if err != nil {
				log.Println(err)
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

		// Convert the `Upcoming Competitions` of the sheet (a string)
		// to an integer value, so it can be an operating. If an error
		// happen because the `recipient.UpcomingCompetitions.Value`
		// is a non-convertable string (for example, when this is the
		// first verification of the recipient and the value is ""),
		// the value will be defined as the current upcoming
		// competitions number and no emails will be sended. If an
		// non-predicate error happen, it is logged and the loop goes
		// to the next recipient in the slice.
		upcomingCompetitionsInInteger, err := strconv.Atoi(recipient.UpcomingCompetitions.Value)
		if err != nil {
			if err.Error() == `strconv.Atoi: parsing "`+recipient.UpcomingCompetitions.Value+`": invalid syntax` {
				upcomingCompetitionsInInteger = recipient.CurrentUpcomingCompetitions
			} else {
				log.Println(err)
				continue
			}
		}

		// Update the recipient cell in the spreadsheet.
		// If an error happen, it is logged and the loop
		// goes to the next recipient in the slice.
		err = recipient.UpdateUpcomingCompetitions()
		if err != nil {
			log.Println(err)
			continue
		}

		// Test if the obsolete and the current value are the
		// same, if are it finishes the loop and goes to the
		// next recipient in the slice, if not it goes through
		// and send an email notifying the recipient.
		if upcomingCompetitionsInInteger == recipient.CurrentUpcomingCompetitions {
			continue
		}

		// Send the notification email, if an error occur it may
		// turn back the upcoming competitions value in the sheet
		// because if not the user will lost the notification.
		err = email.SendEmail(recipient, credentials)
		if err != nil {
			log.Printf(
				"An error occurred while send an notification email, the upcoming competitions of %v will be turned back to %v\n",
				recipient.Name.Value, upcomingCompetitionsInInteger,
			)

			recipient.CurrentUpcomingCompetitions = upcomingCompetitionsInInteger

			updateErr := recipient.UpdateUpcomingCompetitions()
			if updateErr != nil {
				log.Println("An error occurred while turn turn back in the sheet YOU SHOULD do it manually for now :/")
				log.Fatal("The error was:", updateErr)
			}

			log.Fatal("The error while sending email was:", err)
		}

		// If the email was successfully sended, incrase in the resume.
		resumeInformation.EmailsSended++
	}

	log.Printf("The cache was this: %v\n", cityUpcomingCompetitionsCache)

	resumeInformation.StartIn = startIn.String()
	resumeInformation.RuntimeDuration = time.Since(startIn).String()
	resumeInformation.ExportResume("resume.json")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
