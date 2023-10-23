CREATE TABLE IF NOT exists logs_cache_keys (
    id SERIAL PRIMARY KEY,
    from_block BIGINT NOT NULL,
    to_block BIGINT NOT NULL,
    addresses TEXT[] NOT NULL
);

CREATE INDEX IF NOT exists idx_cache_keys_blocks ON logs_cache_keys (from_block, to_block);
CREATE INDEX IF NOT exists idx_cache_keys_addresses ON logs_cache_keys USING GIN (addresses);
