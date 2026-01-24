package sessionhandlers

import "server/api/serviceaccess"

type Handler struct {
	Services *serviceaccess.Access
}
