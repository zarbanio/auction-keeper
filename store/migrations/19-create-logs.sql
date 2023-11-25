CREATE TABLE if not exists logs (
    id              SERIAL PRIMARY KEY,
    message         TEXT,
    fields          JSONB,
    created_at      TIMESTAMP default current_timestamp
)