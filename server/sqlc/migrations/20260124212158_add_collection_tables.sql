-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS collections (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    type VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS documents (
    id UUID PRIMARY KEY NOT NULL, -- must be generated server-side
    collection_id UUID NOT NULL REFERENCES collections(id),
    mime_type VARCHAR NOT NULL,
    s3_location VARCHAR NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS collections;
DROP TABLE IF EXISTS documents;
-- +goose StatementEnd
