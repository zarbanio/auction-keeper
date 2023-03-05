CREATE TABLE if not exists vow_kisses
(
    id              SERIAL PRIMARY KEY,
    rad             BIGINT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists vow_kiss_rad_index on vow_kisses(rad);
