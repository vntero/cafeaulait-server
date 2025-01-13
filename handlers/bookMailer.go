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
	sender := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	receiver := os.Getenv("RECEIVER")

	// email body
	body := fmt.Sprintf(`
	<html>
		<body>
			<h1>Neue Buchungsanfrage eingegangen</h1>
			<h3>Hier sind die Details:</h3>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Telefonnummer:</strong> %s</li>
				<li><strong>Emailadresse:</strong> %s</li>
				<li><strong>Organisation:</strong> %s</li>
				<li><strong>Veranstaltungsort:</strong> %s</li>
				<li><strong>Gästeanzahl:</strong> %s</li>
				<li><strong>Number of guests:</strong> %s</li>
				<li><strong>Veranstaltungsdatum:</strong> %s</li>
				<li><strong>Veranstaltungszeit:</strong> %s</li>
				<li><strong>Budget:</strong> %s</li>
				<li><strong>Zusätzliche Infos oder Fragen:</strong> %s</li>
			</ul>
		</body>
	</html>`, data.Name, data.Phone, data.Email, data.Organization, data.Location, data.Duration, data.NumberOfGuests, data.EventDate, data.EventTime, data.Budget, data.Comment,
	)

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
