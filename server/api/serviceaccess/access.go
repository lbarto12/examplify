package serviceaccess

import (
	"database/sql"

	"github.com/minio/minio-go/v7"
)

type Access struct {
	Postgres *sql.DB
	Minio    *minio.Client
}
