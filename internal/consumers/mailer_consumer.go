package consumers

import (
	"encoding/json"
	"github.com/hainguyen27798/open-notification/global"
	"github.com/hainguyen27798/open-notification/internal/services"
	"github.com/hainguyen27798/open-notification/pkg/utils"
	"github.com/segmentio/kafka-go"
	"log"
)

type MailerConsumer struct {
	MailerService services.IMailService
}

type MailerParams struct {
	Subject      string            `json:"subject"`
	TemplateName string            `json:"templateName"`
	To           []string          `json:"to"`
	Data         map[string]string `json:"data"`
}

func NewMailerConsumer(mailerService services.IMailService) *MailerConsumer {
	return &MailerConsumer{
		MailerService: mailerService,
	}
}

func (mc *MailerConsumer) SendSMTP() {
	consumer := utils.NewKafkaConsumer(
		[]string{global.Config.Kafka.Broker},
		"send-mail-topic",
		"send-mail-group-1",
	)
	defer func(consumer *kafka.Reader) {
		err := consumer.Close()
		if err != nil {
		}
	}(consumer)

	for {
		message, err := consumer.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Error kafka message: %v", err)
		}
		params := MailerParams{}
		if err := json.Unmarshal(message.Value, &params); err != nil {
			log.Fatalf("Error unmarshal json: %v", err)
		} else {
			err := mc.MailerService.SendSTMP(
				params.TemplateName,
				params.Subject,
				params.To,
				params.Data,
			)
			if err != nil {
				log.Fatalf("Error sending mail: %v", err)
			}
		}
		log.Printf("Send mail successfully: %s", params.Subject)
	}
}
