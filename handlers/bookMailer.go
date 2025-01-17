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
	configs.LoadEnv()

	sender := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	server := os.Getenv("SERVER")

	receiverCal := os.Getenv("RECEIVER")
	receiverCustomer := data.Email

	// Internal email body
	businessBody := fmt.Sprintf(`
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

	// External email body
	customerBody := fmt.Sprintf(`
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
				<h1>Vielen Dank für Ihre Buchungsanfrage</h1>
			</div>
			<div class="section">
				<h3>Wir haben folgende Details erhalten:</h3>
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
			<div class="signature">
		  		<br></br>
				<p>Wir werden uns in Kürze bei Ihnen melden.</p>
				<p>Liebe Grüsse,</p>
        		<p>Café Au Lait Team</p>
		  </div>
			<div class="footer">
				<p>Dies ist eine automatisch generierte. Bestätigung Ihrer Anfrage.</p>
			</div>
			</div>
		</body>
		</html>`,
		data.Name, data.Phone, data.Email, data.Organization, data.Location, data.Duration,
		data.NumberOfGuests, data.EventDate, data.EventTime, data.Budget, data.Comment)

	// Send to business
	m1 := gomail.NewMessage()
	m1.SetHeader("From", sender)
	m1.SetHeader("To", receiverCal)
	m1.SetHeader("Subject", "Neue Buchungsanfrage eingegangen")
	m1.SetBody("text/html", businessBody)

	// Send to customer
	m2 := gomail.NewMessage()
	m2.SetHeader("From", sender)
	m2.SetHeader("To", receiverCustomer)
	m2.SetHeader("Subject", "Bestätigung Ihrer Buchungsanfrage")
	m2.SetBody("text/html", customerBody)

	// Settings for SMTP server
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d := gomail.NewDialer(server, 587, sender, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email to business
	if err := d.DialAndSend(m1); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Send email to customer
	if err := d.DialAndSend(m2); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// return
}
