CREATE TABLE if not exists vow_fesses
(
    id     SERIAL PRIMARY KEY,
	tab    NUMERIC(50, 0),
    log_id INTEGER UNIQUE REFERENCES evm_logs (id)
);

CREATE INDEX if not exists fesses_log_id_index on vow_fesses(log_id);
