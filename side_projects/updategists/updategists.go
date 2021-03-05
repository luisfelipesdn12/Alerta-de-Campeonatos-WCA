package updategists

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	resumeFile  string
	mainLogFile string
)

// GistFile is the file object on GitHub API
type GistFile struct {
	Content string `json:"content"`
}

// Gist is the gist object on GitHub API
type Gist struct {
	Description string              `json:"description"`
	Files       map[string]GistFile `json:"files"`
}

func init() {
	log.Println("Checking file:", "../main.log")
	mainLogFileBytes, err := ioutil.ReadFile("../main.log")
	if err != nil {
		log.Fatal("File error:", err)
	}
	mainLogFile = string(mainLogFileBytes)

	log.Println("Checking file:", "../resume.json")
	resumeFileBytes, err := ioutil.ReadFile("../resume.json")
	if err != nil {
		log.Fatal("File error:", err)
	}
	resumeFile = string(resumeFileBytes)
}

// UpdateGists update GitHub gist to controll
func UpdateGists(mainLogGistID, resumeFileGistID string) error {
	response, err := updateFileOnGist(
		mainLogGistID,
		"main.log",
		"Alerta-WCA main log file",
		mainLogFile,
	)
	if err != nil {
		return err
	}
	if !(response.StatusCode == http.StatusOK) {
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))

		return fmt.Errorf("main.log update not worked - status code %v", response.StatusCode)
	}

	response, err = updateFileOnGist(
		resumeFileGistID,
		"resume.json",
		"Information about Alerta-de-Campeonatos-WCA runtime",
		resumeFile,
	)
	if err != nil {
		return err
	}
	if !(response.StatusCode == http.StatusOK) {
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))
		return fmt.Errorf("resume.json update not worked - status code %v", response.StatusCode)
	}

	return nil
}

func updateFileOnGist(gistID, fileName, description, fileContent string) (*http.Response, error) {
	gist := Gist{
		Description: description,
		Files: map[string]GistFile{
			fileName: {
				Content: fileContent,
			},
		},
	}

	reqBody, err := json.Marshal(gist)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.github.com/gists/%v", gistID)
	log.Println("Requesting to:", url)

	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("token %v", os.Getenv("GH_TOKEN")))

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
