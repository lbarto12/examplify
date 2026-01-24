package serviceaccess

import "database/sql"

type Access struct {
	Postgres *sql.DB
}
