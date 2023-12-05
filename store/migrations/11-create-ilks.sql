CREATE TABLE if not exists ilks (
    "name"                              VARCHAR(255) NOT NULL PRIMARY KEY,
    symbol                              VARCHAR(255) NOT NULL,
    minimum_collateralization_ratio     NUMERIC(50, 18) NOT NULL,
    debt_ceiling                        VARCHAR(255) NOT NULL,
    debt                                VARCHAR(255) NOT NULL,
    annual_stability_fee                NUMERIC(50, 18) NOT NULL,
    dust_limit                          VARCHAR(255) NOT NULL,
    price                               VARCHAR(255) NOT NULL,
    rate                                NUMERIC(78) NOT NULL,
    "join"                              VARCHAR(42) NOT NULL,
    median VARCHAR(42) NOT NULL DEFAULT '0x0000000000000000000000000000000000000000',
    gem VARCHAR(42) NOT NULL DEFAULT '0x0000000000000000000000000000000000000000',
    pip VARCHAR(42) NOT NULL
);
