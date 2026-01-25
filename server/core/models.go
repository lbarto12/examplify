package core

import "github.com/google/uuid"

type Collection struct {
	ID   uuid.UUID
	Type string
}

type Document struct {
	ID           uuid.UUID
	CollectionID uuid.UUID
	MimeType     string
	S3Location   string
}
