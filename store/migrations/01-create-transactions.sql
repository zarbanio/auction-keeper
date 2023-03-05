CREATE TABLE if not exists transactions
(
    id                  SERIAL PRIMARY KEY,
    hash                TEXT NOT NULL,
    from_address        TEXT,
    cost                BIGINT,
    data                TEXT,
    value               BIGINT,
    chain_id            BIGINT,
    gas_price           BIGINT,
    gas                 INT,
    nonce               INT,
    to_address          TEXT,
    v                   TEXT,
    r                   TEXT,
    s                   TEXT,
    block_timestamp     INT,
    block_number        BIGINT,
    cumulative_gas_used INT,
    block_hash          TEXT,
    status              INT
);


CREATE INDEX if not exists transactions_block_timestamp_index ON transactions (block_timestamp);
CREATE INDEX if not exists transactions_block_hash_index ON transactions (block_hash);
CREATE UNIQUE INDEX if not exists transactions_hash_index ON transactions (hash);
