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

// Package email is a local package and implements
// functions to send a notification email given the
// `gspread.RecipientStruct` structure.
package email

import (
	"fmt"
	"log"
	"strings"
	"time"

	"../gspread"
	"gopkg.in/gomail.v2"
)

// ReturnATwoWordName uses the strings package to transform
// a string value in a []string and then join it again with
// just the first and the second string.
//
// Example:
// 		n := ReturnATwoWordName("Gabriel Toshio Omiya")
// 		fmt.Println(n)
// 		// Output: "Gabriel Toshio"
//
// If the string value already have 1 or 2 words, it returns
// the same string.
//
// Example:
// 		n := ReturnATwoWordName("Gabriel Toshio")
// 		fmt.Println(n)
// 		// Output: "Gabriel Toshio"
//
// Example:
// 		n := ReturnATwoWordName("Gabriel")
// 		fmt.Println(n)
// 		// Output: "Gabriel"
func ReturnATwoWordName(s string) string {
	separatedNames := strings.Split(s, " ")
	if len(separatedNames) == 1 || len(separatedNames) == 2 {
		return s
	}
	return strings.Join(strings.Split(s, " ")[:2], " ")
}

// SendEmail send an email to the recipient notifying
// the difference between the current number of upcoming
// competitions and the obsolete number.
func SendEmail(r gspread.RecipientStruct, credentials gspread.CredentialStruct) {

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

	var emailLiteralTemplate string
	var headerImageLiteral string
	var emailSubject string

	// Check the recipient language to set the email body
	// template, the subject and the header image. By default
	// it is English.
	switch r.Language.Value {
	case "Português":
		emailLiteralTemplate = `
			<img src="https://raw.githubusercontent.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/master/images/%v" style="%v">
			<h1>Olá, %v</h1>

			<i>Este email é enviado automaticamente e tem informações sobre competições oficiais da WCA na cidade %v.</i>

			<p>Há uma alteração de <b>%v</b> competições futuras, para <b>%v</b>;</p>

			<h2>Informações Gerais:</h2>
			<p><b>Copetições futuras: </b>%v</p>
			<p><b>Verificação obsoleta: </b>%v</p>
			<br>
			<p><b>Copetições futuras: </b>%v</p>
			<p><b>Última verificação: </b>%v</p>

			<h3>Veja <a href="https://www.worldcubeassociation.org/competitions?utf8=%v&search=%v">aqui</a>.</h3>

			<p>Teve alguma dúvida? Problema? Sugestão? Contate o e-mail <a href="mailto:apisbyluisfelipesdn12@gmail.com">apisbyluisfelipesdn12@gmail.com</a> ou abra uma "Issue" no <a href="https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA">GitHub</a>. Obrigado!</p>
			`

		headerImageLiteral = `Email%20Header%20Portuguese.png`
		emailSubject = ("Olá, " + recipientName + "! Atualizações nas competições da WCA em " + r.City.Value + " - " + time.Now().String()[:16])

	default:
		emailLiteralTemplate = `
			<img src="https://raw.githubusercontent.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/master/images/%v" style="%v">
			<h1>Hello, %v</h1>

			<i>This message is automatic sended and have information about official WCA competitions in the city of %v.</i>

			<p>There is a change from <b>%v</b> upcoming competitions, to <b>%v</b>;</p>

			<h2>Information:</h2>
			<p><b>Upcoming competitions: </b>%v</p>
			<p><b>Obsolete verification: </b>%v</p>
			<br>
			<p><b>Upcoming competitions: </b>%v</p>
			<p><b>Last verification: </b>%v</p>

			<h3>See more <a href="https://www.worldcubeassociation.org/competitions?utf8=%v&search=%v">here</a>.</h3>

			<p>Some doubt? Issue? Suggestion? Contact the email <a href="mailto:apisbyluisfelipesdn12@gmail.com">apisbyluisfelipesdn12@gmail.com</a> or open an "Issue" in <a href="https://github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA">GitHub</a>. Thank you!</p>
			`

		headerImageLiteral = `Email%20Header%20English.png`
		emailSubject = ("Hello, " + recipientName + "! News about WCA competitions in " + r.City.Value + " - " + time.Now().String()[:16])
	}

	m.SetHeader("Subject", emailSubject)
	m.SetBody(
		"text/html",
		fmt.Sprintf(
			emailLiteralTemplate,

			headerImageLiteral,
			`max-width: 100%; max-height: 100%;`,
			recipientName, r.City.Value,
			r.UpcomingCompetitions.Value,
			r.CurrentUpcomingCompetitions,
			r.UpcomingCompetitions.Value,
			r.LastVerification.Value,
			r.CurrentUpcomingCompetitions,
			r.CurrentVerificationDate,
			`%E2%9C%93`, r.City.Value,
		),
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
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
