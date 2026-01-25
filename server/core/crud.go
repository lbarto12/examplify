package core

import (
	"context"
	"net/url"
	"server/sqlc/sqlgen"

	"github.com/google/uuid"
)

func (core Core) CreateCollection(ctx context.Context, collectionType string) (*Collection, error) {
	qtx := sqlgen.New(core.Services.Postgres)

	collection, err := qtx.CreateCollection(ctx, collectionType)
	if err != nil {
		return nil, err
	}

	return &Collection{
		ID:   collection.ID,
		Type: collection.Type,
	}, nil
}

func (core Core) GetCollection(ctx context.Context, id uuid.UUID) (*Collection, error) {
	qtx := sqlgen.New(core.Services.Postgres)

	collection, err := qtx.GetCollection(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Collection{
		ID:   collection.ID,
		Type: collection.Type,
	}, nil
}

func (core Core) CreateDocument(ctx context.Context, doc Document) (*url.URL, error) { // awful code, presigned upload URL returned
	qtx := sqlgen.New(core.Services.Postgres)

	fileID := uuid.New()

	_, err := qtx.CreateDocument(ctx, sqlgen.CreateDocumentParams{
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
		UploadBucket,
		fileID.String(),
		PresignedExpiry,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (core Core) GetDocument(ctx context.Context, id uuid.UUID) (*Document, error) {
	qtx := sqlgen.New(core.Services.Postgres)

	document, err := qtx.GetDocument(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Document{
		ID:           document.ID,
		CollectionID: document.CollectionID,
		MimeType:     document.MimeType,
		S3Location:   document.S3Location,
	}, nil
}

func (core Core) PresignedGetDocument(ctx context.Context, id uuid.UUID) (*url.URL, error) {
	qtx := sqlgen.New(core.Services.Postgres)

	document, err := qtx.GetDocument(ctx, id)
	if err != nil {
		return nil, err
	}

	result, err := core.Services.Minio.PresignedGetObject(ctx, UploadBucket, document.S3Location, PresignedExpiry, make(url.Values))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (core Core) GetCollectionDocuments(ctx context.Context, collectionID uuid.UUID) ([]Document, error) {
	qtx := sqlgen.New(core.Services.Postgres)

	documents, err := qtx.GetCollectionDocuments(ctx, collectionID)
	if err != nil {
		return nil, err
	}

	result := []Document{}

	for _, doc := range documents {
		result = append(result, Document{
			ID:           doc.ID,
			CollectionID: doc.CollectionID,
			MimeType:     doc.MimeType,
			S3Location:   doc.S3Location,
		})
	}

	return result, nil
}
