/*
	Alerta-de-Campeonatos-WCA - A script which send an e-mail when there's a new WCA competition.
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

// Package email is a local package and implements
// functions to send a notification email given the
// `gspread.RecipientStruct` structure.
package email

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/gspread"
	"gopkg.in/gomail.v2"
)

var englishTemplate, portugueseTemplate string

// Load the template files in english and portuguese
func init() {
	HTMLEnglishBytes, err := ioutil.ReadFile("./email/template/email-en.html")
	if err != nil {
		log.Fatal(err)
	}

	HTMLPortugueseBytes, err := ioutil.ReadFile("./email/template/email-pt.html")
	if err != nil {
		log.Fatal(err)
	}

	englishTemplate = string(HTMLEnglishBytes)
	portugueseTemplate = string(HTMLPortugueseBytes)
}

// ReturnATwoWordName uses the strings package to transform
// a string value in a []string and then join it again with
// just the first and the second string.
func ReturnATwoWordName(s string) string {
	separatedNames := strings.Split(s, " ")
	if len(separatedNames) == 1 || len(separatedNames) == 2 {
		return s
	}
	return strings.Join(strings.Split(s, " ")[:2], " ")
}

// EmailData is the structure for the template can
// fill the variables
type EmailData struct {
	// The end propertie in the header image URI
	// witch is responsible by the language
	HeaderImageName string

	// The recipient structure
	Recipient gspread.RecipientStruct

	// The name witch returned from the
	// ReturnATwoWordName function
	RecipientName string
}

// SendEmail send an email to the recipient notifying
// the difference between the current number of upcoming
// competitions and the obsolete number.
func SendEmail(r gspread.RecipientStruct, credentials gspread.CredentialStruct) error {

	// In the email subject and body, it's not convenient to
	// use the intire name provided by the user in the form
	// so the function SendEmail will not use the `r.Name.Value`,
	// and in it place will use the `recipientName` value, witch
	// is defined by the return of the `ReturnATwoWordName` function
	// with the `r.Name.Value` as parameter.
	recipientName := ReturnATwoWordName(r.Name.Value)

	// Compose the message object
	m := gomail.NewMessage()
	m.SetHeader("From", credentials.Email)
	m.SetHeader("To", r.Email.Value)

	var emailBody string
	var emailSubject string

	emailData := EmailData{
		Recipient:     r,
		RecipientName: recipientName,
	}

	// Check the recipient language to set the email body
	// template, the subject and the header image. By default
	// it is English.
	switch r.Language.Value {
	case "Português":
		emailData.HeaderImageName = `Email%20Header%20Portuguese%20Compressed.png`

		// It will store the template result as a writer
		var templateHolder bytes.Buffer

		t, err := template.New("email").Parse(portugueseTemplate)
		if err != nil {
			return err
		}

		err = t.Execute(&templateHolder, emailData)
		if err != nil {
			return err
		}

		// Then, it is converted to string
		emailBody = templateHolder.String()
		emailSubject = ("Olá, " + recipientName + "! Atualizações nas competições da WCA em " + r.City.Value + " - " + time.Now().String()[:16])

	default:
		emailData.HeaderImageName = `Email%20Header%20English%20Compressed.png`

		var templateHolder bytes.Buffer

		t, err := template.New("email").Parse(englishTemplate)
		if err != nil {
			return err
		}

		err = t.Execute(&templateHolder, emailData)
		if err != nil {
			return err
		}

		emailBody = templateHolder.String()
		emailSubject = ("Hello, " + recipientName + "! News about WCA competitions in " + r.City.Value + " - " + time.Now().String()[:16])
	}

	m.SetHeader("Subject", emailSubject)
	m.SetBody(
		"text/html",
		emailBody,
	)

	// Do the "login" with the credentials to now
	// be able to send the email with the provided
	// email address above.
	d := gomail.NewDialer(
		"smtp.gmail.com", 587,
		credentials.Email,
		credentials.Password,
	)

	// Send the email.
	log.Printf("Sending a email to %v\n", r.Email.Value)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
