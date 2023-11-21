CREATE TABLE if not exists median_prices (
    id              SERIAL PRIMARY KEY,
    val             NUMERIC(50, 0),
    age             INTEGER,
    log_id          INTEGER UNIQUE REFERENCES evm_logs (id)
);
