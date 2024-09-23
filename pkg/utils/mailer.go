package utils

import (
	"bytes"
	"fmt"
	"github.com/hainguyen27798/open-notification/global"
	"go.uber.org/zap"
	"html/template"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) string {
	msg := "MINE-version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Name)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\n", mail.Body)
	return msg
}

func GetEmailTemplate(name string, data map[string]string) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.ParseFiles(fmt.Sprintf("templates/%s.html", name)))
	err := t.Execute(htmlTemplate, data)
	if err != nil {
		global.Logger.Error("Get email template error", zap.Error(err))
		return "", err
	}
	return htmlTemplate.String(), nil
}
