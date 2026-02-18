package handlers

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"cafeaulait-server/configs"
	"cafeaulait-server/data"

	gomail "gopkg.in/mail.v2"
)

func SendBookSamEmail(data data.BookSamData) error {
	configs.LoadEnv()

	sender := os.Getenv("SENDER_SAM")
	smtpUser := os.Getenv("SENDER_ICLOUD") // Using main iCloud email only for SMTP authentication
	password := os.Getenv("PASSWORD_SAM")
	server := os.Getenv("SERVER_SAM")
	receiverSam := os.Getenv("RECEIVER_SAM")
	log.Printf("receiverSam: %v\n", receiverSam)
	receiverCustomer := data.Email

// Internal email body - plain text
	businessBody := fmt.Sprintf(`
You have received a new booking request with the following details:

Name: %s
Email: %s
Phone: %s
Preferred Event Date: %s
Location: %s
Number of guests: %s
Occasion: %s
Message: %s

---
This message was automatically generated.`,
		data.Name, data.Email, data.Phone, data.EventDate,
		data.Location, data.NumberOfGuests, data.Occasion,
		data.Message)

// External email body
	customerBody := fmt.Sprintf(`
Thank you for your booking request! I have received the following details:

Name: %s
Email: %s
Phone: %s
Preferred Event Date: %s
Location: %s
Number of guests: %s
Occasion: %s
Message: %s

I will get back to you shortly.

Kind regards
Sam In The Kitchen

---
This message was automatically generated.`,
		data.Name, data.Email, data.Phone, data.EventDate,
		data.Location, data.NumberOfGuests, data.Occasion,
		data.Message)

	// Send to business
	m1 := gomail.NewMessage()
	m1.SetHeader("From", sender)
	m1.SetHeader("To", receiverSam)
	m1.SetHeader("Subject", "New Booking Request Received")
	m1.SetBody("text/plain", businessBody)

	// Send to customer
	m2 := gomail.NewMessage()
	m2.SetHeader("From", sender)
	m2.SetHeader("To", receiverCustomer)
	m2.SetHeader("Subject", "Confirmation of Your Booking Request")
	m2.SetBody("text/plain", customerBody)
	

	// Settings for SMTP server
	d := gomail.NewDialer(server, 587, smtpUser, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email to business
	if err := d.DialAndSend(m1); err != nil {
		fmt.Println(err)
		log.Printf("Failed to send email: %v", err)
		return err
	}

	// Send email to customer	
	if err := d.DialAndSend(m2); err != nil {
		fmt.Println(err)
		log.Printf("Failed to send email to customer: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
