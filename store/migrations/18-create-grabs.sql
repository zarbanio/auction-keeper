CREATE TABLE if not exists grabs
(
    id     SERIAL PRIMARY KEY,
    i      TEXT,
    u      TEXT,
    v      TEXT,
    w      TEXT,
    dink   NUMERIC(50, 0),
    dart   NUMERIC(50, 0),
    log_id INTEGER UNIQUE REFERENCES evm_logs (id)
);

CREATE INDEX if not exists grabs_log_id_index on grabs(log_id);
