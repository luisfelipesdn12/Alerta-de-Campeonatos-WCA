// Package sendmelog package implements a function to
// send me via an Telegram message the `main.log` file.
package sendmelog

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const notifyWhenDone = false
const notifyWhenNotDone = true

// Send actually send to the Telegram receiver the main.log file.
func Send(TelegramBotToken string, TelegramRecipientID int, GitHubMainLogGistURL, GitHubMainLogGistLastCommitHash string) {
	logFile, err := ioutil.ReadFile("../main.log")
	if err != nil {
		log.Fatalln("Error while opening main.log: " + err.Error())
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  TelegramBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln(err)
	}

	recipient := tb.ChatID(TelegramRecipientID)

	switch wasDone(string(logFile)) {
	// If the execution was done.
	case true:
		if notifyWhenDone {
			bot.Send(recipient,
				fmt.Sprintf(
					"The execution of WCA-Alert WAS done\n\n[See more here!](%v)",
					(GitHubMainLogGistURL+"/"+GitHubMainLogGistLastCommitHash),
				),
				&tb.SendOptions{
					ParseMode: "Markdown",
				},
			)
		}

	// If the execution was not done.
	case false:
		if notifyWhenNotDone {
			bot.Send(recipient,
				fmt.Sprintf(
					"The execution of WCA-Alert WAS NOT done\n\n[See more here!](%v)",
					(GitHubMainLogGistURL+"/"+GitHubMainLogGistLastCommitHash),
				),
				&tb.SendOptions{
					ParseMode: "Markdown",
				},
			)
		}
	}
}

func wasDone(logFileInString string) bool {
	return strings.Contains(logFileInString, "THE EXECUTION WAS DONE")
}
