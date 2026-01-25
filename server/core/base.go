package core

import (
	"context"
	"net/url"
	"server/api/serviceaccess"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type core_interface interface {
	CreateCollection(ctx context.Context, collectionType string) (*Collection, error)
	GetCollection(ctx context.Context, id uuid.UUID) (*Collection, error)

	CreateDocument(ctx context.Context, doc Document) (*url.URL, error) // awful code, presigned upload URL returned
	GetDocument(ctx context.Context, id uuid.UUID) (*Document, error)
	PresignedGetDocument(ctx context.Context, id uuid.UUID) (*url.URL, error)

	GetCollectionDocuments(ctx context.Context, collectionID uuid.UUID) ([]Document, error)
}

type Core struct {
	Services *serviceaccess.Access
}

const (
	UploadBucket    string        = "image-analysis-images"
	PresignedExpiry time.Duration = time.Minute * 5
)

func NewCore(services *serviceaccess.Access) (*Core, error) {

	// Check / Create bucket for temp assets
	bucketExists, err := services.Minio.BucketExists(context.Background(), UploadBucket)
	if err != nil {
		return nil, err
	}

	if !bucketExists {
		if err := services.Minio.MakeBucket(context.Background(), UploadBucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}

		// If I were a kinder man I would enable auto-deleting, but I have no respect for server space rn
	}

	var intf core_interface = &Core{
		Services: services,
	}

	return intf.(*Core), nil
}
