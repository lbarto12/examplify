package environment

import "github.com/caarlos0/env/v11"

type Vars struct {

	// SERVER
	ServerHost string `env:"SERVER_HOST,notEmpty"`
	ServerPort string `env:"SERVER_PORT,notEmpty"`

	// JWTs
	JWTSecretKey string `env:"JWT_SECRET_KEY,notEmpty"`

	// Postgres
	PostgresHost               string `env:"POSTGRES_HOST,notEmpty"`
	PostgresPort               string `env:"POSTGRES_PORT,notEmpty"`
	PostgresDatabase           string `env:"POSTGRES_DATABASE,notEmpty"`
	PostgresUser               string `env:"POSTGRES_USER,notEmpty"`
	PostgresPassword           string `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresMaxOpenConnections int64  `env:"POSTGRES_MAX_OPEN_CONNECTIONS,notEmpty"`
	PostgresSSLMode            string `env:"POSTGRES_SSL_MODE,notEmpty"`

	// Minio
	MinioEndpoint      string `env:"MINIO_ENDPOINT,notEmpty"`
	MinioUser          string `env:"MINIO_USER,notEmpty"`
	MinioPassword      string `env:"MINIO_PASSWORD,notEmpty"`
	MinioDefaultBucket string `env:"MINIO_DEFAULT_BUCKET,notEmpty"`
	MinioUseSSL        bool   `env:"MINIO_USE_SSL,notEmpty"`

	// Gemini
	GeminiAIKey string `env:"GEMINI_API_KEY,notEmpty"`

	// OpenAI
	OpenAIKey string `env:"OPENAI_API_KEY,notEmpty"`
}

func Get() (*Vars, error) {
	vars := Vars{}
	if err := env.Parse(&vars); err != nil {
		return nil, err
	}
	return &vars, nil
}
