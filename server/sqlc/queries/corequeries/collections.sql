-- name: CreateCollection :one
-- The user creating a collection must be the creator themselves (creator_id = @user_id)
INSERT INTO collections (type, creator_id)
VALUES (@type, @user_id) -- force creator_id to match authenticated user
RETURNING *;

-- name: CreateDocument :one
-- Only allow creating a document in a collection the user owns
INSERT INTO documents (id, collection_id, mime_type, s3_location)
SELECT @id, @collection_id, @mime_type, @s3_location
FROM collections c
WHERE c.id = @collection_id
  AND c.creator_id = @user_id
RETURNING *;

-- name: GetCollection :one
SELECT *
FROM collections
WHERE id = @id
  AND creator_id = @user_id;

-- name: GetDocument :one
SELECT d.*
FROM documents d
JOIN collections c ON d.collection_id = c.id
WHERE d.id = @id
  AND c.creator_id = @user_id;

-- name: GetCollectionDocuments :many
SELECT d.*
FROM documents d
JOIN collections c ON d.collection_id = c.id
WHERE d.collection_id = @collection_id
  AND c.creator_id = @user_id;

