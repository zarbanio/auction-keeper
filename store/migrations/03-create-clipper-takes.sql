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

CREATE INDEX if not exists clipper_take_who_index on clipper_takes(who);
