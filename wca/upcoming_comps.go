package wca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	turnLogsOn bool   = false
	layoutISO  string = "2006-01-02"
	apiBaseURI string = "https://www.worldcubeassociation.org/api/v0"
)

// WCAAPIResponse is the struct based in the JSON format
// that will be requested bellow. That's not all the data
// provided by the API, but that´s the data witch will be
// probally usefull later
type WCAAPIResponse []struct {
	URL             string   `json:"url"`
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Website         string   `json:"website"`
	City            string   `json:"city"`
	StartDate       string   `json:"start_date"`
	CompetitorLimit int      `json:"competitor_limit"`
	EventIds        []string `json:"event_ids"`
}

func init() {
	if !(turnLogsOn) {
		log.SetOutput(ioutil.Discard)
	}
}

// UpcomingCopetitions given a city name, do a request to the
// WCA API and compare the start date of each competition with
// the current date. So, it returns the number of upcoming copetitions
func UpcomingCopetitions(cityName string) (int, error) {
	upcomingCompetitions := 0

	currentDate := time.Now().String()[:10] // Current date in format: "yyyy-MM-dd"

	// Format the complete URI with the request params
	// defined in the `cityName` param and the currentDate
	// to find just the upcoming competitions
	queryParam := fmt.Sprintf("?start=%v&q=%v", currentDate, url.QueryEscape(cityName))
	URI := fmt.Sprintf(
		"%v/competitions/%v",
		apiBaseURI, queryParam,
	)

	// Do the request, if an error the function returns 0 and the error
	log.Printf("Doing a resquest to: %v\n", URI)
	resp, err := http.Get(URI)
	if err != nil {
		return upcomingCompetitions, err
	}

	// Get the response JSON string, if an error
	// happen, the function returns 0 and the error
	log.Println("Geting the JSON string from request")
	defer resp.Body.Close()
	responseAsJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return upcomingCompetitions, err
	}

	// Transform the response JSON string in the
	// structure WCAAPIResponse. If an error
	// happen, the function returns 0 and the error
	log.Println("Converting the JSON string to response structure")
	responseAsStruct := WCAAPIResponse{}
	err = json.Unmarshal(responseAsJSON, &responseAsStruct)
	if err != nil {
		return upcomingCompetitions, err
	}

	// The upcomingCompetitions is defined by the
	// length of the results
	upcomingCompetitions = len(responseAsStruct)

	return upcomingCompetitions, nil
}
