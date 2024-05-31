BEGIN;

CREATE TABLE IF NOT EXISTS todos (
    id uuid PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    pending_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
    );

COMMIT;