//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/open-notification/internal/consumers"
	"github.com/hainguyen27798/open-notification/internal/services"
)

func InitMailerConsumer() (*consumers.MailerConsumer, error) {
	wire.Build(
		services.NewMailService,
		consumers.NewMailerConsumer,
	)
	return new(consumers.MailerConsumer), nil
}
