CREATE TABLE IF NOT EXISTS logs (
    id              SERIAL PRIMARY KEY,
    level           VARCHAR(20),
    service         VARCHAR(20),
    message         TEXT,
    fields          JSONB,
    created_at      TIMESTAMP default current_timestamp
)