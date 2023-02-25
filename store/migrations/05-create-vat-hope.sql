CREATE TABLE if not exists vat_hopes
(
    id              SERIAL PRIMARY KEY,
    usr             TEXT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists vat_hope_usr_index on vat_hopes(usr);
