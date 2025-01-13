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
	// load the env file
	configs.LoadEnv()

	// sender data
	sender := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	receiver := os.Getenv("RECEIVER")

	// email body
	body := fmt.Sprintf(`
	<html>
		<body>
			<h1>Ein weiteres Kind möchte bei uns mittanzen</h1>
			<h2>Hier sind die Details:</h2>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Geburtsdatum:</strong> %s</li>
				<li><strong>Herkunft:</strong> %s</li>
				<li><strong>Motivationsgrund:</strong> %s</li>
			</ul>
			<h2>Erziehungsberechtigte:r 1</h2>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Phone:</strong> %s</li>
				<li><strong>Strasse:</strong> %s</li>
				<li><strong>Nr:</strong> %s</li>
				<li><strong>PLZ:</strong> %s</li>
				<li><strong>Ort:</strong> %s</li>
			</ul>
			<h2>Erziehungsberechtigte:r 2</h2>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Phone:</strong> %s</li>
				<li><strong>Strasse:</strong> %s</li>
				<li><strong>Nr:</strong> %s</li>
				<li><strong>PLZ:</strong> %s</li>
				<li><strong>Ort:</strong> %s</li>
			</ul>
		</body>
	</html>`, data.Name, data.Birthday, data.Origin, data.Motivation,
		data.ParentOneName, data.ParentOneEmail, data.ParentOnePhone, data.ParentOneStreet,
		data.ParentOneHouseNumber, data.ParentOnePostcode, data.ParentOneLocation,
		data.ParentTwoName, data.ParentTwoEmail, data.ParentTwoPhone, data.ParentTwoStreet,
		data.ParentTwoHouseNumber, data.ParentTwoPostcode, data.ParentTwoLocation,
	)

	// gomail magic starts here
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender)

	// Set E-Mail receivers
	m.SetHeader("To", receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", "Ein weiteres Kind möchte bei uns mittanzen")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer("asmtp.mail.hostpoint.ch", 587, sender, password)

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
