CREATE TABLE if not exists dog_barks
(
    id              SERIAL PRIMARY KEY,
    ilk             TEXT,
	urn             TEXT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);
