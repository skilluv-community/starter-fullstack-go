CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS notes_created_at_idx ON notes (created_at DESC);
