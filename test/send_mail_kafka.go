package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	producer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9094"),
		Topic:    "send-mail-topic",
		Balancer: &kafka.LeastBytes{},
	}
	body := make(map[string]interface{})
	body["templateName"] = "otp-email"
	body["subject"] = "OTP Verification"
	body["to"] = []string{"hainguyen27798@gmail.com"}
	body["data"] = map[string]string{"otp": "test1234kafka"}
	message, _ := json.Marshal(body)
	err := producer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("opt-auth"),
		Value: []byte(message),
		Time:  time.Now(),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}
