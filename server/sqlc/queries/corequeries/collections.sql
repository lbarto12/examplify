-- name: CreateCollection :one
-- The user creating a collection must be the creator themselves (creator_id = @user_id)
INSERT INTO collections (type, creator_id, title, course)
VALUES (@type, @user_id, @title, @course) -- force creator_id to match authenticated user
RETURNING *;

-- name: CreateDocument :one
-- Only allow creating a document in a collection the user owns
INSERT INTO documents (id, collection_id, mime_type, s3_location, title)
SELECT @id, @collection_id, @mime_type, @s3_location, @title
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


-- name: FilterCollections :many
SELECT * FROM collections
WHERE course = @course
AND type = @type
AND creator_id = @user_id;
