CREATE TABLE if not exists transactions
(
    id      BIGINT PRIMARY KEY,
	time    BIGINT ,
	hash    TEXT,
	from    TEXT,
);

CREATE INDEX if not exists transactions_id_index ON transactions (id);
CREATE INDEX if not exists transactions_timestamp_index ON transactions (time);
CREATE INDEX if not exists transactions_hash_index ON transactions (hash);
