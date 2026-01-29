package core

import (
	"context"
	"encoding/json"
	"net/url"
	"server/api/serviceaccess"
	"server/environment"
	"server/sqlc/sqlgen"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/openai/openai-go/v3"
)

type core_interface interface {
	// Collection operations
	CreateCollection(ctx context.Context, userID uuid.UUID, params Collection) (*Collection, error)
	GetCollection(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*Collection, error)
	GetCollectionDocuments(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID) ([]Document, error)

	// Document operations
	CreateDocument(ctx context.Context, userID uuid.UUID, doc Document) (*url.URL, error)
	GetDocument(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*Document, error)
	PresignedGetDocument(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*url.URL, error)

	// Analysis operations
	AnalyzeCollection(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID, kind sqlgen.AnalysisType) (*CollectionAnalysis, error)
	GetCollectionAnalyses(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID) ([]CollectionAnalysis, error)

	// Internal
	extractDocumentContent(ctx context.Context, doc sqlgen.Document) (*DocumentTextExtraction, error)
	createSnapshot(ctx context.Context, collectionID uuid.UUID) (*CollectionSnapshot, error)
	runAnalysis(ctx context.Context, content string, kind sqlgen.AnalysisType) (json.RawMessage, error)
}

type Core struct {
	Services        *serviceaccess.Access
	Queries         *sqlgen.Queries
	UploadBucket    string
	PresignedExpiry time.Duration
}

func NewCore(services *serviceaccess.Access, env *environment.Vars) (*Core, error) {
	bucketName := env.UploadBucketName
	presignedExpiry := time.Minute * time.Duration(env.PresignedExpiryMins)

	// Check / Create bucket for temp assets
	bucketExists, err := services.Minio.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, err
	}

	if !bucketExists {
		if err := services.Minio.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}

		// If I were a kinder man I would enable auto-deleting, but I have no respect for server space rn
	}

	var intf core_interface = &Core{
		Services:        services,
		Queries:         sqlgen.New(services.Postgres),
		UploadBucket:    bucketName,
		PresignedExpiry: presignedExpiry,
	}

	return intf.(*Core), nil
}

const ( // Consts
	ChatModel = openai.ChatModelGPT5ChatLatest
)
