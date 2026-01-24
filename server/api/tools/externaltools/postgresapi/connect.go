package postgresapi

import (
	"database/sql"
	"fmt"
	"server/environment"

	_ "github.com/lib/pq"
)

func Connect(env *environment.Vars) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.PostgresHost,
		env.PostgresPort,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDatabase,
		env.PostgresSSLMode,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(int(env.PostgresMaxOpenConnections))

	return db, nil
}
