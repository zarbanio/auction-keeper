CREATE TABLE if not exists vat_forks
(
    id     		SERIAL PRIMARY KEY,
    ilk  		TEXT,
	src  		TEXT,
	dst  		TEXT,
	dink		NUMERIC(50, 0),
	dart 		NUMERIC(50, 0),
    log_id 		INTEGER UNIQUE REFERENCES evm_logs (id)
);

CREATE INDEX if not exists forks_log_id_index on vat_forks(log_id);
				