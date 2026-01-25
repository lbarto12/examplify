CREATE INDEX IF NOT EXISTS idx_documents_collection_id
ON documents (collection_id);

CREATE INDEX IF NOT EXISTS idx_extractions_document_id
ON document_extractions (document_id);

CREATE INDEX IF NOT EXISTS idx_snapshots_collection_id
ON collection_snapshots (collection_id);

CREATE INDEX IF NOT EXISTS idx_analyses_snapshot_id
ON collection_analyses (snapshot_id);
