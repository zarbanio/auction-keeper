CREATE TABLE if not exists vow_flops
(
    id              SERIAL PRIMARY KEY,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);


