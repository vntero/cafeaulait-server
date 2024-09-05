// package functions

// import (
// 	"cafeaulait-server/models"
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/smtp"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func SendRegisterEmail(data models.RegisterInput) {
// 	// load the env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	// sender data
// 	from := os.Getenv("SENDER")
// 	password := os.Getenv("PASSWORD")
// 	receiver := os.Getenv("RECEIVER")

// 	// receiver data
// 	to := []string{ receiver }

// 	// smtp server configuration
// 	smtpHost := "smtp.mail.me.com"
// 	smtpPort := "587"

// 	// email body
// 	subject := "Subject: Child registration for dance classes\r\n"
// 	contentType := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body := fmt.Sprintf(`
// 	<html>
// 		<body>
// 			<h1>We have another child insterested in joining our classes!</h1>
// 			<p>Here are the registration details:</p>
// 			<ul>
// 				<li><strong>Name:</strong> %s</li>
// 				<li><strong>Birthday:</strong> %s</li>
// 				<li><strong>Origin:</strong> %s</li>
// 				<li><strong>Motivation:</strong> %s</li>
// 			</ul>
// 			<h2>Parent 1 Details</h2>
// 			<ul>
// 				<li><strong>Name:</strong> %s</li>
// 				<li><strong>Email:</strong> %s</li>
// 				<li><strong>Phone:</strong> %s</li>
// 				<li><strong>Street:</strong> %s</li>
// 				<li><strong>House Number:</strong> %s</li>
// 				<li><strong>Postcode:</strong> %s</li>
// 				<li><strong>Location:</strong> %s</li>
// 			</ul>
// 			<h2>Parent 2 Details</h2>
// 			<ul>
// 				<li><strong>Name:</strong> %s</li>
// 				<li><strong>Email:</strong> %s</li>
// 				<li><strong>Phone:</strong> %s</li>
// 				<li><strong>Street:</strong> %s</li>
// 				<li><strong>House Number:</strong> %s</li>
// 				<li><strong>Postcode:</strong> %s</li>
// 				<li><strong>Location:</strong> %s</li>
// 			</ul>
// 		</body>
// 	</html>`, data.Name, data.Birthday, data.Origin, data.Motivation,
// 		data.ParentOneName, data.ParentOneEmail, data.ParentOnePhone, data.ParentOneStreet,
// 		data.ParentOneHouseNumber, data.ParentOnePostcode, data.ParentOneLocation,
// 		data.ParentTwoName, data.ParentTwoEmail, data.ParentTwoPhone, data.ParentTwoStreet,
// 		data.ParentTwoHouseNumber, data.ParentTwoPostcode, data.ParentTwoLocation,
// 	)

// 	// message
// 	message := []byte(subject + contentType + body)

// 	// SMTP connection
//     tlsconfig := &tls.Config{
//         InsecureSkipVerify: true, // Set to false in production
//         ServerName: smtpHost,
//     }

//     conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsconfig)
//     if err != nil {
//         log.Fatalf("Failed to dial: %v", err)
//     }

//     client, err := smtp.NewClient(conn, smtpHost)
//     if err != nil {
//         log.Fatalf("Failed to create SMTP client: %v", err)
//     }

//     defer client.Close()

// 	// authentication
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	// sending email
// 	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if err != nil {
// 		log.Fatalf("Error sending email: %v", err)
// 	}
// 	fmt.Println("Email sent successfully!")
// }

package functions

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendRegisterEmail () {
  m := gomail.NewMessage()

  // Set E-Mail sender
  m.SetHeader("From", "vntero@icloud.com")

  // Set E-Mail receivers
  m.SetHeader("To", "shin.antero@gmail.com")

  // Set E-Mail subject
  m.SetHeader("Subject", "Gomail test subject")

  // Set E-Mail body. You can set plain text or html with text/html
  m.SetBody("text/plain", "This is Gomail test body")

  // Settings for SMTP server
  d := gomail.NewDialer("smtp.mail.me.com", 587, "vntero@icloud.com", "")

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