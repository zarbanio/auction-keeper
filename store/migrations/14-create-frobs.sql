CREATE TABLE if not exists vat_frobs
(
    id     SERIAL PRIMARY KEY,
    i      TEXT,
    u      TEXT,
    v      TEXT,
    w      TEXT,
    dink   NUMERIC(50, 0),
    dart   NUMERIC(50, 0),
    log_id INTEGER UNIQUE REFERENCES logs (id)
);

CREATE INDEX if not exists vat_frobs_log_id_index on vat_frobs(log_id);
