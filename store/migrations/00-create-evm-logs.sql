create table if not exists evm_logs
(
    id           SERIAL PRIMARY KEY,
    timestamp    BIGINT,
    address      TEXT,
    block_number BIGINT,
    tx_hash      TEXT,
    tx_index     INTEGER,
    block_hash   TEXT,
    index        INTEGER,
    data         BYTEA,
    topics       TEXT[]
);

CREATE UNIQUE INDEX if not exists evm_logs_unique_index ON evm_logs (tx_hash, index);
CREATE INDEX if not exists evm_logs_block_number_index ON evm_logs (block_number);
CREATE INDEX if not exists evm_logs_id_index on evm_logs (id);