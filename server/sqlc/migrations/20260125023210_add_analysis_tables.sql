-- +goose Up
-- +goose StatementBegin
CREATE TABLE document_extractions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    document_id UUID NOT NULL REFERENCES documents(id),
    content TEXT NOT NULL
);

CREATE TABLE collection_snapshots (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    collection_id UUID NOT NULL REFERENCES collections(id),
    combined_content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TYPE analysis_type AS ENUM (
    'summary',
    'flashcards',
    'quiz',
    'deep_summary'
);

CREATE TABLE collection_analyses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    snapshot_id UUID NOT NULL REFERENCES collection_snapshots(id),
    type analysis_type NOT NULL,
    result JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS collection_analyses;
DROP TYPE IF EXISTS analysis_type;
DROP TABLE IF EXISTS collection_snapshots;
DROP TABLE IF EXISTS document_extractions;
-- +goose StatementEnd
