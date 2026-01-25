package serviceaccess

import (
	"database/sql"

	"github.com/minio/minio-go/v7"
	"github.com/openai/openai-go/v3"
	"google.golang.org/genai"
)

type Access struct {
	Postgres *sql.DB
	Minio    *minio.Client
	Gemini   *genai.Client
	OpenAI   *openai.Client
}
