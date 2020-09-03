package wrapper

import (
	"github.com/alexandria-oss/identity-api/pkg/service"
	"github.com/alexandria-oss/identity-api/pkg/service/middleware"
	log "github.com/sirupsen/logrus"
)

// NewUserQuery Chain-of-responsibility service wrapper (logging)
func NewUserQuery(svc service.UserQuery, logger *log.Logger) service.UserQuery {
	var wrapSvc service.UserQuery
	wrapSvc = middleware.UserQueryLog{
		Logger: logger,
		Next:   svc,
	}

	return wrapSvc
}
