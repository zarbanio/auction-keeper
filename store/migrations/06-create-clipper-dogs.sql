CREATE TABLE if not exists clipper_dogs
(
    id              SERIAL PRIMARY KEY,
    ilk             TEXT,
	urn             TEXT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists clipper_dog_urn_index on clipper_dogs(urn);
