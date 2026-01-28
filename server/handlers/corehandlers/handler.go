package corehandlers

import (
	"server/api/serviceaccess"
	"server/core"
	"server/sqlc/sqlgen"
)

type Handler struct {
	Services *serviceaccess.Access
	Core     *core.Core
	Queries  *sqlgen.Queries
}
