package functions

import (
	"fmt"
	"net/smtp"
)

func RenderEmail() {
	// sender data
	from := "info@cafeaulait.ch"
	password := "superstrongpassword"

	// receiver data
	to := []string{
		"info@cafeaulait.ch",
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// email body
	subject := "Child registration for dance classes"
	contentType := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := `
	<html>
	<body>
		<h1>Thank you for Registering!</h1>
		<p>Here are your registration details:</p>
		<ul>
			<li>Name: John Doe</li>
			<li>Birthday: January 1, 1990</li>
			<li>Origin: Somewhere</li>
		</ul>
		<p>We hope to see you soon!</p>
	</body>
	</html>`

	// message
	message := []byte(subject + contentType + body)

	// authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully!")
}
