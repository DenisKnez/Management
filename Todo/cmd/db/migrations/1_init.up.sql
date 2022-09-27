CREATE TABLE IF NOT EXISTS todos (
    id uuid PRIMARY KEY,
    "text" text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted boolean NOT NULL
);