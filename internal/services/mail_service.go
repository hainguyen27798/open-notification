package services

import (
	"fmt"
	"github.com/hainguyen27798/open-notification/global"
	"github.com/hainguyen27798/open-notification/pkg/utils"
	"go.uber.org/zap"
	"net/smtp"
)

type IMailService interface {
	SendSTMP(templateName string, subject string, to []string, data map[string]string) error
}

type mailService struct{}

func NewMailService() IMailService {
	return &mailService{}
}

func (ms *mailService) SendSTMP(templateName string, subject string, to []string, data map[string]string) error {
	body, err := utils.GetEmailTemplate(templateName, data)
	from := global.Config.SMTP.From
	SMTPHost := global.Config.SMTP.Host
	SMTPPort := global.Config.SMTP.Port
	SMTPUsername := global.Config.SMTP.Username
	SMTPPassword := global.Config.SMTP.Password

	if err != nil {
		return err
	}

	content := utils.Mail{
		From:    utils.EmailAddress{Address: from, Name: from},
		To:      to,
		Subject: subject,
		Body:    body,
	}

	message := utils.BuildMessage(content)

	// smtp auth
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)

	host := fmt.Sprintf("%s:%d", SMTPHost, SMTPPort)
	if err := smtp.SendMail(host, auth, from, to, []byte(message)); err != nil {
		global.Logger.Error("Send mail error:", zap.Error(err))
		return err
	}

	return nil
}
