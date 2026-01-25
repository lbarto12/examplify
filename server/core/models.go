package core

import "github.com/google/uuid"

type Collection struct {
	ID     uuid.UUID
	Title  string
	Course string
	Type   string
}

type Document struct {
	ID           uuid.UUID
	CollectionID uuid.UUID
	Title        string
	MimeType     string
	S3Location   string
}
