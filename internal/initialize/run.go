package initialize

import (
	"flag"
	"fmt"
	"github.com/hainguyen27798/open-notification/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func Run() {
	LoadConfig()
	InitLogger()
	SetupFirestore()

	s := grpc.NewServer()
	RegisterGrpc(s)

	port := flag.Int("port", global.Config.Server.Port, "grpc server port")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		global.Logger.Error("failed to listen", zap.Error(err))
		panic(err)
	}
	global.Logger.Info(
		fmt.Sprintf("server listening on port %v", lis.Addr()),
	)
	if err := s.Serve(lis); err != nil {
		global.Logger.Error("failed to serve", zap.Error(err))
		panic(err)
	}
}
