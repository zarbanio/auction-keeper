CREATE TABLE if not exists transactions
(
    id      				BIGINT PRIMARY KEY,
	hash    				TEXT,
	from    				TEXT,
	cost 					BIGINT,
	data 					TEXT,
	value 					BIGINT,
	gas_price				BIGINT,
	gas 					INT,
	nonce 					INT,
	to 						TEXT,
	v 						TEXT,
	r 						TEXT,
	s 						TEXT,
	block_timestamp 		BIGINT,
	block_number	 		BIGINT,
	cumulative_gas_used 	INT,
	block_hash				TEXT,
	status  				INT,
);

CREATE INDEX if not exists transactions_id_index ON transactions (id);
CREATE INDEX if not exists transactions_timestamp_index ON transactions (time);
CREATE INDEX if not exists transactions_hash_index ON transactions (hash);
