CREATE TABLE IF NOT EXISTS logs (
    id              SERIAL PRIMARY KEY,
    message         TEXT,
    fields          JSONB,
    created_at      TIMESTAMP default current_timestamp
)