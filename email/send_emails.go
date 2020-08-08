package email

import (
	"fmt"
	"time"

	"../gspread"
	"gopkg.in/gomail.v2"
)

// SendEmail send an email to the recipient notifying
// the difference between the current number of upcoming
// competitions and the obsolete number.
func SendEmail(r gspread.RecipientStruct, credentials gspread.CredentialStruct) {

	// Compose the message object
	m := gomail.NewMessage()
	m.SetHeader("From", credentials.Email)
	m.SetHeader("To", r.Email.Value)
	m.SetHeader("Subject", ("Olá, " + r.Name.Value + "! Atualizações nas competições da WCA em " + r.City.Value + " - " + time.Now().String()[:16]))
	m.SetBody(
		"text/html",
		fmt.Sprintf(
			`
			<img src="https://camo.githubusercontent.com/64fa8a73b5e761b03cc07bd8a7602e3c4043f15e/68747470733a2f2f692e696d6775722e636f6d2f7047586d58524c2e706e67" style="%v">
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
			`,

			`max-width: 100%; max-height: 100%;`,
			r.Name.Value, r.City.Value,
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
	err := d.DialAndSend(m)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
