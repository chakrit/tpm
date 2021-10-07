CREATE TABLE IF NOT EXISTS problems (
    id          SERIAL PRIMARY KEY,
    summary     TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);