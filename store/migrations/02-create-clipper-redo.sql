
CREATE TABLE if not exists clipper_redoes
(
    id              SERIAL PRIMARY KEY,
    sale_id         INT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists clipper_redoes_sale_id_index on clipper_redoes(sale_id);
