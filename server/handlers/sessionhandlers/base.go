package sessionhandlers

import (
	"server/api/serviceaccess"
	"server/sqlc/sqlgen"
)

type Handler struct {
	Services *serviceaccess.Access
	Queries  *sqlgen.Queries
}
