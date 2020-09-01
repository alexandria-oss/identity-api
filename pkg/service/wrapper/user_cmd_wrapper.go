package wrapper

import (
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/alexandria-oss/identity-api/pkg/service/middleware"
	log "github.com/sirupsen/logrus"
)

// NewUserCommandHandler Chain-of-responsibility service wrapper (logging)
func NewUserCommandHandler(svc service.UserCommandHandler, logger *log.Logger) service.UserCommandHandler {
	var wrapSvc service.UserCommandHandler
	wrapSvc = middleware.UserCommandHandlerLog{
		Logger: logger,
		Next:   svc,
	}

	return wrapSvc
}
