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

	sender := os.Getenv("SENDER_ICLOUD")
	password := os.Getenv("PASSWORD_ICLOUD")
	server := os.Getenv("SERVER_ICLOUD")

	receiver := os.Getenv("RECEIVER_ICLOUD")

	// Email body
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

	// gomail magic starts here
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender)

	// Set E-Mail receiver
	m.SetHeader("To", receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", "Ein weiteres Kind möchte bei uns mittanzen")

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
