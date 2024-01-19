package notification

import (
	"bytes"
	"log"

	"html/template"

	"gopkg.in/gomail.v2"
)

type EmailParams struct {
	ToEmail   string
	Subject   string
	HTMLFile  string
	Name      string
	Date      string
	StartHour string
	EndHour   string
}

func SendEmail(params EmailParams) error {
	from := "go.haircut2024@gmail.com"
	password := "dfzu stxn lqwz wdpi"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", params.ToEmail)
	m.SetHeader("Subject", params.Subject)

	tmpl, err := template.ParseFiles(params.HTMLFile)
	if err != nil {
		log.Println("Error parsing template:", err)
		return err
	}

	data := struct {
		Name      string
		Date      string
		StartHour string
		EndHour   string
	}{
		Name:      params.Name,
		Date:      params.Date,
		StartHour: params.StartHour,
		EndHour:   params.EndHour,
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Println("Error executing template:", err)
		return err
	}

	m.SetBody("text/html", tpl.String())
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	return d.DialAndSend(m)
}
