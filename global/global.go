package global

import (
	"cloud.google.com/go/firestore"
	"github.com/hainguyen27798/open-notification/pkg/logger"
	"github.com/hainguyen27798/open-notification/pkg/setting"
)

var (
	Config          setting.Config
	Logger          *logger.Zap
	FirestoreClient *firestore.Client
)
