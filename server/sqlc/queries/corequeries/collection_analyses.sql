-- name: CreateCollectionAnalysis :one
INSERT INTO collection_analyses (snapshot_id, type, result)
VALUES ($1, $2, $3)
RETURNING id, type, result, created_at;

-- name: GetCollectionAnalysesByCollection :many
SELECT a.id,
       a.type,
       a.result,
       a.created_at
FROM collection_analyses a
JOIN collection_snapshots s ON s.id = a.snapshot_id
WHERE s.collection_id = $1
ORDER BY a.created_at DESC;


-- name: CreateCollectionSnapshot :one
INSERT INTO collection_snapshots (collection_id, combined_content)
VALUES ($1, $2)
RETURNING id, collection_id, combined_content, created_at;

-- name: GetCollectionSnapshots :many
SELECT id, collection_id, combined_content, created_at
FROM collection_snapshots
WHERE collection_id = $1
ORDER BY created_at DESC;
