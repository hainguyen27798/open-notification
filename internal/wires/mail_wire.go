//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/open-notification/internal/controllers"
	"github.com/hainguyen27798/open-notification/internal/services"
)

func InitMailerController() (*controllers.MailerController, error) {
	wire.Build(
		services.NewMailService,
		controllers.NewMailerController,
	)
	return new(controllers.MailerController), nil
}
