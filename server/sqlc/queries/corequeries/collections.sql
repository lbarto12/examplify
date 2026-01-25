-- name: CreateCollection :one
INSERT INTO collections (type)
VALUES (@type)
RETURNING *;

-- name: CreateDocument :one
INSERT INTO documents (id, collection_id, mime_type, s3_location)
VALUES (@id, @collection_id, @mime_type, @s3_location)
RETURNING *;

-- name: GetCollection :one
SELECT * FROM collections
WHERE id = @id;

-- name: GetDocument :one
SELECT * FROM documents
WHERE id = @id;

-- name: GetCollectionDocuments :many
SELECT * FROM documents
WHERE collection_id = @collection_id;
