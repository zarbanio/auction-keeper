CREATE TABLE if not exists clipper_takes
(
    id              SERIAL PRIMARY KEY,
    auction_id      INTEGER NOT NULL,
    amt             TEXT,
	max             TEXT,
	who             TEXT,
	data            TEXT,    
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists init_log_id_index on vat_inits(log_id);
