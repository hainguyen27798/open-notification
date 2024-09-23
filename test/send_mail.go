package main

import (
	"context"
	"fmt"
	"github.com/hainguyen27798/open-notification/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	c := proto.NewMailerClient(conn)
	ctx := context.Background()
	r, err := c.SendSMTP(ctx, &proto.SendSMTPRequest{
		TemplateName: "otp-email",
		Subject:      "OTP Verification",
		To:           []string{"hainguyen.freelancer@gmail.com"},
		Data:         map[string]string{"otp": "test1234"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("mailer res: %v", r.GetMessage())
}
