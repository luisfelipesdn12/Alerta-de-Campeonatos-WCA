package notification

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/toast.v1"
)

const notifyWhenDone = false
const notifyWhenNotDone = true

var gitHubMainLogGistURL, gitHubMainLogGistLastCommitHash string

// Notify throws an toast notification with information
// about the execution.
func Notify(GitHubMainLogGistURL, GitHubMainLogGistLastCommitHash string) {
	gitHubMainLogGistURL = GitHubMainLogGistURL
	gitHubMainLogGistLastCommitHash = GitHubMainLogGistLastCommitHash

	logFile, err := ioutil.ReadFile("../main.log")
	if err != nil {
		fmt.Println(err)
	}

	switch wasDone(string(logFile)) {
	case true:
		if notifyWhenDone {
			sendTheNotification("WCA-Alert - Verification done!")
		}
	case false:
		if notifyWhenNotDone {
			sendTheNotification("WCA-Alert - Verification WAS NOT done!")
		}
	}
}

func sendTheNotification(title string) {
	notification := toast.Notification{
		AppID:   "Alerta-de-Campeonatos-WCA",
		Title:   title,
		Message: "See more information in main.log",
		Actions: []toast.Action{
			{"protocol", "Open resume", `https://luisfelipesdn12.github.io/Runtime-Information-WCA-Alert/`},
			{"protocol", "Open main.log", (gitHubMainLogGistURL + "/" + gitHubMainLogGistLastCommitHash)},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func wasDone(logFileInString string) bool {
	return strings.Contains(logFileInString, "THE EXECUTION WAS DONE")
}
