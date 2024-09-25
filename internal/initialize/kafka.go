package initialize

import "github.com/hainguyen27798/open-notification/internal/wires"

func InitKafka() {
	mailerConsumer, _ := wires.InitMailerConsumer()
	go mailerConsumer.SendSMTP()
}
