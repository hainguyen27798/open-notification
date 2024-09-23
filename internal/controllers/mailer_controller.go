package controllers

import (
	"context"
	"github.com/hainguyen27798/open-notification/global"
	"github.com/hainguyen27798/open-notification/internal/services"
	"github.com/hainguyen27798/open-notification/proto"
	"go.uber.org/zap"
)

type MailerController struct {
	proto.UnimplementedMailerServer
	MailerService services.IMailService
}

func NewMailerController(mailerService services.IMailService) *MailerController {
	return &MailerController{
		MailerService: mailerService,
	}
}

func (s *MailerController) SendSMTP(ctx context.Context, req *proto.SendSMTPRequest) (*proto.SendSMTPResponse, error) {
	err := s.MailerService.SendSTMP(
		req.GetTemplateName(),
		req.GetSubject(),
		req.GetTo(),
		req.GetData(),
	)
	if err != nil {
		return nil, err
	}
	global.Logger.Info("send mail success", zap.Any("req", req))
	return &proto.SendSMTPResponse{
		Code:    200,
		Message: "Send Email Success",
	}, nil
}
