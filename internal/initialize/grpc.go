package initialize

import (
	"github.com/hainguyen27798/open-notification/internal/wires"
	"github.com/hainguyen27798/open-notification/proto"
	"google.golang.org/grpc"
)

func RegisterGrpc(s *grpc.Server) {
	mailerController, _ := wires.InitMailerController()
	proto.RegisterMailerServer(s, mailerController)
}
