-- name: CreateDocumentExtraction :one
INSERT INTO document_extractions (document_id, content)
VALUES ($1, $2)
RETURNING id;

-- name: HasDocumentExtraction :one
SELECT EXISTS (
  SELECT 1
  FROM document_extractions
  WHERE document_id = $1
);

-- name: GetDocumentExtractionsByCollection :many
SELECT e.content
FROM document_extractions e
JOIN documents d ON d.id = e.document_id
WHERE d.collection_id = $1;