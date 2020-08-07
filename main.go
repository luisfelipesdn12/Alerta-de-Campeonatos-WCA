package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"./email"
	"./gspread"
	"./wca"
)

func main() {
	recipients, _ := gspread.GetRecipientsData()

	for _, recipient := range recipients {
		fmt.Println(recipient.CurrentVerificationDate)
		recipient.CurrentVerificationDate = time.Now().String()[:19]
		result, err := wca.UpcomingCopetitions(recipient.City.Value)
		recipient.CurrentUpcomingCompetitions = result
		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Printf(
			"\n%v de %v tinha %v e agora tem %v copetições futuras\n",
			recipient.Name.Value,
			recipient.City.Value,
			recipient.UpcomingCompetitions.Value,
			recipient.CurrentUpcomingCompetitions,
		)

		upcomingCompetitionsInInteger, err := strconv.Atoi(recipient.UpcomingCompetitions.Value)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if upcomingCompetitionsInInteger == recipient.CurrentUpcomingCompetitions {
			continue
		}

		email.SendEmail(recipient)
	}
}
