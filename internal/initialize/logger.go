package initialize

import (
	"github.com/hainguyen27798/open-notification/global"
	"github.com/hainguyen27798/open-notification/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
