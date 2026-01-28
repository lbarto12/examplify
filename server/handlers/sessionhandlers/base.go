package sessionhandlers

import (
	"server/api/serviceaccess"
	"server/api/tools/features/sessions"
	"server/sqlc/sqlgen"
)

type Handler struct {
	Services       *serviceaccess.Access
	Queries        *sqlgen.Queries
	SessionManager *sessions.SessionManager
}
