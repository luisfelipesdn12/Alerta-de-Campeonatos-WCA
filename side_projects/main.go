package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/side_projects/sendmelog"
	"github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/side_projects/updategists"
)

// secretEnvStruct stores useful and confidential variables as
// telegram tokens and IDs, secret GitHub gists etc.
type secretEnvStruct struct {
	// The token provided by Telegram to control the bot.
	// Know more in: https://core.telegram.org/bots
	TelegramBotToken string
	// The Telegram ID of the user who will receive the message.
	// Know more in: https://core.telegram.org/constructor/user
	TelegramRecipientID int
	// The ID of the secret GitHub gist witch stores the main.log.
	GitHubMainLogGistID string
	// The ID of the secret GitHub gist witch stores the resume.json.
	GitHubResumeJSONGistID string
}

var secretEnv secretEnvStruct

// Will fetch the token from the `secret.json` file
// and assign to the `secretEnv` structure so it
// can be used by the other packages.
func setSecretEnv() error {
	secretEnv.TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")

	recipientID, err := strconv.Atoi(os.Getenv("TELEGRAM_RECIPIENT_ID"))
	secretEnv.TelegramRecipientID = recipientID

	secretEnv.GitHubMainLogGistID = os.Getenv("GH_MAIN_LOG_GIST_ID")

	secretEnv.GitHubResumeJSONGistID = os.Getenv("GH_RESUME_GIST_ID")

	return err
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Was not founded an .env file")
	}
}

func main() {
	// Set the secretEnv variable with the property
	// `GitHubMainLogGistLastCommitHash` already updated.
	log.Println("Setting necessary environment variables")
	err := setSecretEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Update the `resume.json` and the `main.log` file to
	// their GitHub gists.
	log.Println("Updating gists")
	err = updategists.UpdateGists(secretEnv.GitHubMainLogGistID, secretEnv.GitHubResumeJSONGistID)
	if err != nil {
		log.Fatal(err)
	}

	// Sends a Telegram message with the main.log file.
	log.Println("Sending results to Telegram")
	sendmelog.Send(
		secretEnv.TelegramBotToken,
		secretEnv.TelegramRecipientID,
		secretEnv.GitHubMainLogGistID,
	)
}
