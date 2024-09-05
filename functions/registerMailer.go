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
	to := []string {
		"info@cafeaulait.ch",
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// message
	message := []byte("This is a test email message")

	// authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// sending email
	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully!")
}