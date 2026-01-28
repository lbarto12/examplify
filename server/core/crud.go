package core

import (
	"context"
	"net/url"
	"server/sqlc/sqlgen"

	"github.com/google/uuid"
)

func (core Core) CreateCollection(ctx context.Context, userID uuid.UUID, params Collection) (*Collection, error) {
	collection, err := core.Queries.CreateCollection(ctx, sqlgen.CreateCollectionParams{
		UserID: userID,
		Title:  params.Title,
		Course: params.Course,
		Type:   params.Type,
	})
	if err != nil {
		return nil, err
	}

	return &Collection{
		ID:     collection.ID,
		Title:  collection.Title,
		Course: collection.Course,
		Type:   collection.Type,
	}, nil
}

func (core Core) GetCollection(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*Collection, error) {
	collection, err := core.Queries.GetCollection(ctx, sqlgen.GetCollectionParams{
		UserID: userID,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	return &Collection{
		ID:     collection.ID,
		Title:  collection.Title,
		Course: collection.Course,
		Type:   collection.Type,
	}, nil
}

func (core Core) CreateDocument(ctx context.Context, userID uuid.UUID, doc Document) (*url.URL, error) { // awful code, presigned upload URL returned
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
		UploadBucket,
		fileID.String(),
		PresignedExpiry,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

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

func (core Core) PresignedGetDocument(ctx context.Context, userID uuid.UUID, id uuid.UUID) (*url.URL, error) {
	document, err := core.Queries.GetDocument(ctx, sqlgen.GetDocumentParams{
		UserID: userID,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	result, err := core.Services.Minio.PresignedGetObject(ctx, UploadBucket, document.S3Location, PresignedExpiry, make(url.Values))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (core Core) GetCollectionDocuments(ctx context.Context, userID uuid.UUID, collectionID uuid.UUID) ([]Document, error) {
	documents, err := core.Queries.GetCollectionDocuments(ctx, sqlgen.GetCollectionDocumentsParams{
		UserID:       userID,
		CollectionID: collectionID,
	})
	if err != nil {
		return nil, err
	}

	result := []Document{}

	for _, doc := range documents {
		result = append(result, Document{
			ID:           doc.ID,
			CollectionID: doc.CollectionID,
			Title:        doc.Title,
			MimeType:     doc.MimeType,
			S3Location:   doc.S3Location,
		})
	}

	return result, nil
}
