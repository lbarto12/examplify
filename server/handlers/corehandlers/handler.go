package corehandlers

import (
	"server/api/serviceaccess"
	"server/core"
)

type Handler struct {
	Services *serviceaccess.Access
	Core     *core.Core
}
