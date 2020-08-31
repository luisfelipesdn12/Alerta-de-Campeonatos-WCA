package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"

	"./notification"
	"./sendmelog"
)

// SecretJSONStruct stores useful and confidential variables as
// telegram tokens adn IDs, secret GitHub gists etc.
type SecretJSONStruct struct {
	// The token provided by Telegram to control the bot.
	// Know more in: https://core.telegram.org/bots
	TelegramBotToken string
	// The Telegram ID of the user who will receive the message.
	// Know more in: https://core.telegram.org/constructor/user
	TelegramRecipientID int
	// The URL of the secret GitHub gist witch stores the main.log.
	GitHubMainLogGistURL string
}

var secretJSON SecretJSONStruct

// Will fetch the token from the `secret.json` file
// and assign to the `secretJSON` structure so it
// can be used by the other packages.
func init() {
	secretJSONFile, err := ioutil.ReadFile("./secret.json")
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(secretJSONFile, &secretJSON)
}

func main() {
	// Sends a Telegram message with the main.log file.
	sendmelog.Send(secretJSON.TelegramBotToken, secretJSON.TelegramRecipientID, secretJSON.GitHubMainLogGistURL)

	// Update the `resume.json` and the `main.log` file to
	// their GitHub gists.
	exec.Command("python.exe", "./update_gists/update_gists.py").Run()

	// Sends a toast notification with execution information.
	notification.Notify(secretJSON.GitHubMainLogGistURL)
}
