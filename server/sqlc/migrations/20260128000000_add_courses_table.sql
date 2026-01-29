-- +goose Up
CREATE TABLE IF NOT EXISTS courses (
    name TEXT NOT NULL,
    creator_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (name, creator_id)
);

-- +goose Down
DROP TABLE IF EXISTS courses;
