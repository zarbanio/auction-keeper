CREATE TABLE if not exists vow_flogs
(
    id              SERIAL PRIMARY KEY,
    era             BIGINT,
    tx_id           INTEGER UNIQUE REFERENCES transactions (id)
);

CREATE INDEX if not exists vow_flog_era_index on vow_flogs(era);
