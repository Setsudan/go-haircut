package notification

import (
	"log"
	"net/smtp"
)

func SendEmail(toEmail string, subject string, body string) {
	from := "go.haircut2024@gmail.com"
	password := "Gohaircut2024*"
	to := []string{toEmail}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: " + subject + "\n\n" + body)

	auth := smtp.PlainAuth("Support Go Haircut", from, password, smtpHost)

	send := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if send != nil {
		log.Fatalf("Erreur lors de l'envoi de l'e-mail: %s", send)
	}
}
