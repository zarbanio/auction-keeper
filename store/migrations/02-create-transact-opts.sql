CREATE TABLE if not exists transact_opts
(
    id                  BIGINT PRIMARY KEY,
    from                TEXT,
	nonce               BIGINT,
	value               BIGINT,
	gasPrice            BIGINT,
	           BIGINT,
	gasTipCap           BIGINT,
	gasLimit            BIGINT,
	noSend              BOOLEAN
);

CREATE INDEX if not exists transact_opts_id_index ON transact_opts (id);
CREATE INDEX if not exists transact_opts_from_index ON transact_opts (from);
