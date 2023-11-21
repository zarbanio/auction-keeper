CREATE TABLE if not exists vow_flogs
(
    id     SERIAL PRIMARY KEY,
	era    NUMERIC(50, 0),
    log_id INTEGER UNIQUE REFERENCES evm_logs (id)
);

CREATE INDEX if not exists flogs_log_id_index on vow_flogs(log_id);
