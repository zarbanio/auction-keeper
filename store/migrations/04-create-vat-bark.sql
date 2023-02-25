CREATE TABLE if not exists vat_barks
(
    id              SERIAL PRIMARY KEY,
    ilk             TEXT,
	urn             TEXT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists vat_bark_urn_index on vat_barks(urn);
