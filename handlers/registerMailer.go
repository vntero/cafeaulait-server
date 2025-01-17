package handlers

import (
	"crypto/tls"
	"fmt"
	"os"

	"cafeaulait-server/configs"
	"cafeaulait-server/data"

	gomail "gopkg.in/mail.v2"
)

func SendRegisterEmail(data data.ResgisterData) {
	configs.LoadEnv()

	sender := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	server := os.Getenv("SERVER")

	receiverCal := os.Getenv("RECEIVER")
	receiverCustomer := data.ParentOneEmail

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
			color: #EF4444;
			padding: 10px 0;
			font-weight: 500;
		  }
		  .section {
			margin-top: 20px;
		  }
		  h1, h2 {
			color: #EF4444;
			font-weight: 500;
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
			<h1>Ein weiteres Kind möchte bei uns mittanzen</h1>
		  </div>
		  <div class="section">
			<h2>Hier sind die Details:</h2>
			<ul>
			  <li><strong>Name:</strong> %s</li>
			  <li><strong>Geburtsdatum:</strong> %s</li>
			  <li><strong>Herkunft:</strong> %s</li>
			  <li><strong>Motivationsgrund:</strong> %s</li>
			</ul>
		  </div>
		  <div class="section">
			<strong>Erziehungsberechtigte:r 1</strong>
			<ul>
			  <li><strong>Name:</strong> %s</li>
			  <li><strong>Email:</strong> %s</li>
			  <li><strong>Telefonnummer:</strong> %s</li>
			  <li><strong>Strasse:</strong> %s</li>
			  <li><strong>Nr:</strong> %s</li>
			  <li><strong>PLZ:</strong> %s</li>
			  <li><strong>Ort:</strong> %s</li>
			</ul>
		  </div>
		  <div class="section">
			<strong>Erziehungsberechtigte:r 2</strong>
			<ul>
			  <li><strong>Name:</strong> %s</li>
			  <li><strong>Email:</strong> %s</li>
			  <li><strong>Telefonnummer:</strong> %s</li>
			  <li><strong>Strasse:</strong> %s</li>
			  <li><strong>Nr:</strong> %s</li>
			  <li><strong>PLZ:</strong> %s</li>
			  <li><strong>Ort:</strong> %s</li>
			</ul>
		  </div>
		  <div class="footer">
			<p>Diese Nachricht wurde automatisch generiert. Bitte nicht darauf antworten.</p>
		  </div>
		</div>
	  </body>
	</html>`,
		data.Name, data.Birthday, data.Origin, data.Motivation,
		data.ParentOneName, data.ParentOneEmail, data.ParentOnePhone, data.ParentOneStreet,
		data.ParentOneHouseNumber, data.ParentOnePostcode, data.ParentOneLocation,
		data.ParentTwoName, data.ParentTwoEmail, data.ParentTwoPhone, data.ParentTwoStreet,
		data.ParentTwoHouseNumber, data.ParentTwoPostcode, data.ParentTwoLocation)

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
				color: #EF4444;
				padding: 10px 0;
				font-weight: 500;
			}
			.section {
				margin-top: 20px;
			}
			h1, h2 {
				color: #EF4444;
				font-weight: 500;
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
				<h1>Vielen Dank für Ihr Interesse</h1>
			</div>
			<div class="section">
				<h2>Wir haben folgende Details erhalten:</h2>
				<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Geburtsdatum:</strong> %s</li>
				<li><strong>Herkunft:</strong> %s</li>
				<li><strong>Motivationsgrund:</strong> %s</li>
				</ul>
			</div>
			<div class="section">
				<strong>Erziehungsberechtigte:r 1</strong>
				<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Telefonnummer:</strong> %s</li>
				<li><strong>Strasse:</strong> %s</li>
				<li><strong>Nr:</strong> %s</li>
				<li><strong>PLZ:</strong> %s</li>
				<li><strong>Ort:</strong> %s</li>
				</ul>
			</div>
			<div class="section">
				<strong>Erziehungsberechtigte:r 2</strong>
				<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Telefonnummer:</strong> %s</li>
				<li><strong>Strasse:</strong> %s</li>
				<li><strong>Nr:</strong> %s</li>
				<li><strong>PLZ:</strong> %s</li>
				<li><strong>Ort:</strong> %s</li>
				</ul>
			</div>
			<div class="signature">
					<br></br>
					<p>Wir werden uns in Kürze bei Ihnen melden.</p>
					<p>Liebe Grüsse,</p>
					<p>Café Au Lait Team</p>
			</div>
			<div class="footer">
				<p>Diese Nachricht wurde automatisch generiert. Bitte nicht darauf antworten.</p>
			</div>
			</div>
		</body>
	</html>`,
	data.Name, data.Birthday, data.Origin, data.Motivation,
	data.ParentOneName, data.ParentOneEmail, data.ParentOnePhone, data.ParentOneStreet,
	data.ParentOneHouseNumber, data.ParentOnePostcode, data.ParentOneLocation,
	data.ParentTwoName, data.ParentTwoEmail, data.ParentTwoPhone, data.ParentTwoStreet,
	data.ParentTwoHouseNumber, data.ParentTwoPostcode, data.ParentTwoLocation)

	// Send to business
	m1 := gomail.NewMessage()
	m1.SetHeader("From", sender)
	m1.SetHeader("To", receiverCal)
	m1.SetHeader("Subject", "Ein weiteres Kind möchte bei uns mittanzen")
	m1.SetBody("text/html", businessBody)

	// Send to customer
	m2 := gomail.NewMessage()
	m2.SetHeader("From", sender)
	m2.SetHeader("To", receiverCustomer)
	m2.SetHeader("Subject", "Vielen Dank für Ihr Interesse")
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
