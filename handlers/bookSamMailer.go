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

func SendBookSamEmail(data data.BookSamData) error{
	configs.LoadEnv()

	sender := os.Getenv("SENDER_SAM")
	password := os.Getenv("PASSWORD_SAM")
	server := os.Getenv("SERVER_SAM")

	receiverSam := os.Getenv("RECEIVER_SAM")
	// receiverCustomer := data.Email

	// Internal email body - plain text
	businessBody := fmt.Sprintf(`
Neue Buchungsanfrage eingegangen

Hier sind die Details:

Name: %s %s
Telefonnummer: %s
Emailadresse: %s
Organisation: %s
Veranstaltungsort: %s
Gewünschte Dauer der Darbietung: %s
Gästeanzahl: %s
Veranstaltungsdatum: %s
Zusätzliche Infos oder Fragen: %s

---
Diese Nachricht wurde automatisch generiert.`,
		data.Name, data.LastName, data.Phone, data.Email, data.EventDate,
		data.Location, data.NumberOfGuests, data.EventDate,
		data.Comment)

	// External email body - COMMENTED OUT
	/*
	customerBody := fmt.Sprintf(`
Vielen Dank für Ihre Buchungsanfrage

Wir haben folgende Details erhalten:

Name: %s %s
Telefonnummer: %s
Emailadresse: %s
Organisation: %s
Veranstaltungsort: %s
Gewünschte Dauer der Darbietung: %s
Gästeanzahl: %s
Veranstaltungsdatum: %s
Veranstaltungszeit: %s
Budget: %s
Zusätzliche Infos oder Fragen: %s

Wir werden uns in Kürze bei Ihnen melden.

Liebe Grüsse,
Café Au Lait Team

---
Dies ist eine automatisch generierte Bestätigung Ihrer Anfrage.`,
		data.Name, data.LastName, data.Phone, data.Email, data.Organization,
		data.Location, data.Duration, data.NumberOfGuests, data.EventDate,
		data.EventTime, data.Budget, data.Comment)
	*/

	// Send to business
	m1 := gomail.NewMessage()
	m1.SetHeader("From", sender)
	m1.SetHeader("To", receiverSam)
	m1.SetHeader("Subject", "Neue Buchungsanfrage eingegangen")
	m1.SetBody("text/plain", businessBody)

	// Send to customer - COMMENTED OUT
	/*
	m2 := gomail.NewMessage()
	m2.SetHeader("From", sender)
	m2.SetHeader("To", receiverCustomer)
	m2.SetHeader("Subject", "Bestätigung Ihrer Buchungsanfrage")
	m2.SetBody("text/plain", customerBody)
	*/

	// Settings for SMTP server
	d := gomail.NewDialer(server, 587, sender, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email to business
	if err := d.DialAndSend(m1); err != nil {
		fmt.Println(err)
		log.Printf("Failed to send email: %v", err)
		return err
	}

	// Send email to customer - COMMENTED OUT
	/*
	if err := d.DialAndSend(m2); err != nil {
		fmt.Println(err)
		panic(err)
	}
	*/

	log.Println("Email sent successfully")
	return nil
}
