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

// Package wca is a local package and implements
// functions to access the data from the WCA's API.
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
	apiBaseURI string = "https://www.worldcubeassociation.org/api/v0"
)

// APIResponse is the struct based in the JSON format
// that will be requested bellow. That's not all the data
// provided by the API, but that´s the data witch will be
// probally useful later.
type APIResponse []struct {
	URL             string   `json:"url"`
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Website         string   `json:"website"`
	City            string   `json:"city"`
	StartDate       string   `json:"start_date"`
	CompetitorLimit int      `json:"competitor_limit"`
	EventIds        []string `json:"event_ids"`
}

// UpcomingCopetitions given a city name, do a request to the
// WCA API and compare the start date of each competition with
// the current date. So, it returns the number of upcoming copetitions.
func UpcomingCopetitions(cityName string) (int, error) {
	upcomingCompetitions := 0

	// Current date in format: "yyyy-MM-dd"
	currentDate := time.Now().String()[:10]

	// Format the complete URI with the request params
	// defined in the `cityName` param and the currentDate
	// to find just the upcoming competitions.
	queryParam := fmt.Sprintf("?start=%v&q=%v", currentDate, url.QueryEscape(cityName))
	URI := fmt.Sprintf(
		"%v/competitions/%v",
		apiBaseURI, queryParam,
	)

	// Do the request, if an error the function returns 0 and the error.
	log.Printf("Doing a resquest to: %v\n", URI)
	resp, err := http.Get(URI)
	if err != nil {
		return upcomingCompetitions, err
	}

	// Get the response JSON string, if an error
	// happen, the function returns 0 and the error.
	log.Println("Geting the JSON string from request")
	defer resp.Body.Close()
	responseAsJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return upcomingCompetitions, err
	}

	// Transform the response JSON string in the
	// structure APIResponse. If an error
	// happen, the function returns 0 and the error.
	log.Println("Converting the JSON string to response structure")
	responseAsStruct := APIResponse{}
	err = json.Unmarshal(responseAsJSON, &responseAsStruct)
	if err != nil {
		return upcomingCompetitions, err
	}

	// The upcomingCompetitions is defined by the
	// length of the results.
	upcomingCompetitions = len(responseAsStruct)

	return upcomingCompetitions, nil
}
