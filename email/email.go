package email

import (
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

type EmailMessage struct {
	Subject string
	Body    string
}

func SendEmail(message *EmailMessage) error {
	// compose the message
	m := email.NewMessage(message.Subject, message.Body)
	m.From = mail.Address{Name: "From", Address: "from@example.com"}
	m.To = []string{"to@example.com"}

	// add attachments
	//if err := m.Attach("email.go"); err != nil {
	//	log.Fatal(err)
	//}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "from@example.com", "pwd", "smtp.zoho.com")
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		return err
	}
	return nil
}
