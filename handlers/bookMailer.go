package handlers

import (
	"crypto/tls"
	"fmt"
	"os"

	"cafeaulait-server/configs"
	"cafeaulait-server/data"

	gomail "gopkg.in/mail.v2"
)

func SendBookEmail(data data.BookData) {
	// load the env file
	configs.LoadEnv()

	// sender data
	sender := os.Getenv("SENDER_ICLOUD")
	password := os.Getenv("PASSWORD_ICLOUD")
	receiver := os.Getenv("RECEIVER_ICLOUD")
	server := os.Getenv("SERVER_ICLOUD")

	// email body
	body := fmt.Sprintf(`
		<html>
		<head>
			<style>
			body {
				font-family: Arial, sans-serif;
				line-height: 1.6;
				color: #333333;
			}
			.container {
				max-width: 600px;
				margin: auto;
				border: 1px solid #dddddd;
				padding: 20px;
				border-radius: 10px;
				background-color: white;
			}
			.header {
				text-align: center;
				color: #EF4444; /* Matches Tailwind text-red-500 */
				padding: 10px 0;
				font-weight: 500; /* Matches Tailwind font-medium */
			}
			.section {
				margin-top: 20px;
			}
			h1, h3 {
				color: #EF4444; /* Matches Tailwind text-red-500 */
				font-weight: 500; /* Matches Tailwind font-medium */
			}
			ul {
				list-style: none;
				padding: 0;
			}
			ul li {
				margin-bottom: 8px;
			}
			.footer {
				text-align: center;
				margin-top: 20px;
				font-size: 12px;
				color: #aaaaaa;
			}
			</style>
		</head>
		<body>
			<div class="container">
			<div class="header">
				<h1>Neue Buchungsanfrage eingegangen</h1>
			</div>
			<div class="section">
				<h3>Hier sind die Details:</h3>
				<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Telefonnummer:</strong> %s</li>
				<li><strong>Emailadresse:</strong> %s</li>
				<li><strong>Organisation:</strong> %s</li>
				<li><strong>Veranstaltungsort:</strong> %s</li>
				<li><strong>Gewünschte Dauer der Darbietung:</strong> %s</li>
				<li><strong>Gästeanzahl:</strong> %s</li>
				<li><strong>Veranstaltungsdatum:</strong> %s</li>
				<li><strong>Veranstaltungszeit:</strong> %s</li>
				<li><strong>Budget:</strong> %s</li>
				<li><strong>Zusätzliche Infos oder Fragen:</strong> %s</li>
				</ul>
			</div>
			<div class="footer">
				<p>Diese Nachricht wurde automatisch generiert. Bitte nicht darauf antworten.</p>
			</div>
			</div>
		</body>
		</html>`,
		data.Name, data.Phone, data.Email, data.Organization, data.Location, data.Duration,
		data.NumberOfGuests, data.EventDate, data.EventTime, data.Budget, data.Comment)

	// gomail magic starts here
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender)

	// Set E-Mail receivers
	m.SetHeader("To", receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", "Neue Buchungsanfrage eingegangen")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer(server, 587, sender, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// return
}
