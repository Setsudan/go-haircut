package notification

import (
	"log"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, subject string, htmlFilePath string, name string, date string, startHour string, endHour string) {
	from := "go.haircut2024@gmail.com"
	password := "dfzu stxn lqwz wdpi"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)

	body, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier HTML: %s", err)
	}
	htmlContent := string(body)

	htmlContent = strings.Replace(htmlContent, "{{NAME}}", name, -1)
	htmlContent = strings.Replace(htmlContent, "{{DATE}}", date, -1)
	htmlContent = strings.Replace(htmlContent, "{{START_HOUR}}", startHour, -1)
	htmlContent = strings.Replace(htmlContent, "{{END_HOUR}}", endHour, -1)

	m.SetBody("text/html", htmlContent)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Erreur lors de l'envoi de l'email: %s", err)
	}
}
