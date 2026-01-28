package core

import (
	"context"
	"net/url"
	"server/sqlc/sqlgen"

	"github.com/google/uuid"
)

// CreateDocument creates a new document and returns a presigned upload URL
func (core Core) CreateDocument(ctx context.Context, userID uuid.UUID, doc Document) (*url.URL, error) {
	fileID := uuid.New()

	_, err := core.Queries.CreateDocument(ctx, sqlgen.CreateDocumentParams{
		UserID:       userID,
		ID:           fileID,
		CollectionID: doc.CollectionID,
		MimeType:     doc.MimeType,
		S3Location:   fileID.String(),
	})
	if err != nil {
		return nil, err
	}

	result, err := core.Services.Minio.PresignedPutObject(
		ctx,
		core.UploadBucket,
		fileID.String(),
		core.PresignedExpiry,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDocument retrieves a document by ID for a user
func (core Core) GetDocument(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*Document, error) {
	document, err := core.Queries.GetDocument(ctx, sqlgen.GetDocumentParams{
		UserID: userID,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	return &Document{
		ID:           document.ID,
		CollectionID: document.CollectionID,
		Title:        document.Title,
		MimeType:     document.MimeType,
		S3Location:   document.S3Location,
	}, nil
}

// PresignedGetDocument returns a presigned URL to download a document
func (core Core) PresignedGetDocument(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*url.URL, error) {
	document, err := core.Queries.GetDocument(ctx, sqlgen.GetDocumentParams{
		UserID: userID,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	result, err := core.Services.Minio.PresignedGetObject(ctx, core.UploadBucket, document.S3Location, core.PresignedExpiry, make(url.Values))
	if err != nil {
		return nil, err
	}

	return result, nil
}
