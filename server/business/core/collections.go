package core

import (
	"context"
	"server/sqlc/sqlgen"

	"github.com/google/uuid"
)

// CreateCollection creates a new collection for a user
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

// GetCollection retrieves a collection by ID for a user
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

// GetCollectionDocuments retrieves all documents in a collection
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
