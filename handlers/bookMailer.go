package handlers

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"cafeaulait-server/models"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func SendBookEmail(data models.BookInput) {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// sender data
	sender := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	receiver := os.Getenv("RECEIVER")

	// email body
	body := fmt.Sprintf(`
	<html>
		<body>
			<h1>Someone is interested in booking us!</h1>
			<h3>Here are the event details:</h3>
			<ul>
				<li><strong>Name:</strong> %s</li>
				<li><strong>Phone:</strong> %s</li>
				<li><strong>Email:</strong> %s</li>
				<li><strong>Location:</strong> %s</li>
				<li><strong>Duration:</strong> %s</li>
				<li><strong>Number of guests:</strong> %s</li>
				<li><strong>Event date:</strong> %s</li>
				<li><strong>Event time:</strong> %s</li>
				<li><strong>Budget:</strong> %s</li>
				<li><strong>Additional comments:</strong> %s</li>
			</ul>
		</body>
	</html>`, data.Name, data.Phone, data.Email, data.Location, data.Duration, data.NumberOfGuests, data.EventDate, data.EventTime, data.Budget, data.Comment,
	)

	// gomail magic starts here
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", sender)

	// Set E-Mail receivers
	m.SetHeader("To", receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", "Interest in booking a show")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.mail.me.com", 587, sender, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
