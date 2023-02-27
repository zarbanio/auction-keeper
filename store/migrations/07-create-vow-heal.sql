CREATE TABLE if not exists vow_heals
(
    id              SERIAL PRIMARY KEY,
    rad             TEXT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists vow_heal_rad_index on vow_heals(rad);
