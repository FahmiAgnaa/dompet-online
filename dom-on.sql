CREATE TABLE users
(
    id               VARCHAR(100) PRIMARY KEY NOT NULL,
    full_name        VARCHAR(100),
    user_name        VARCHAR(25),
    email            VARCHAR(50) UNIQUE,
    phone_number     VARCHAR(15) UNIQUE,
    password         VARCHAR(100),
    password_confirm VARCHAR(100),
    is_active        BOOLEAN,
    created_at       TIMESTAMP WITH TIME ZONE,
    updated_at       TIMESTAMP,
    deleted_at       TIMESTAMP
);

CREATE TABLE wallets
(
    id            VARCHAR(100) PRIMARY KEY NOT NULL,
    user_id       VARCHAR(100) REFERENCES users (id),
    rekening_user VARCHAR(100),
    balance       BIGINT,
    created_at    TIMESTAMP WITH TIME ZONE,
    updated_at    TIMESTAMP
);

CREATE TYPE payment_method_enum AS ENUM ('Credit Card', 'PayLater', 'Bank Transfer', 'Cash', 'Other');

CREATE TABLE payment_method
(
    id          VARCHAR(100) PRIMARY KEY,
    name        payment_method_enum NOT NULL,
    description VARCHAR(100)
);


CREATE TABLE transactions
(
    id                 VARCHAR(100) PRIMARY KEY NOT NULL,
    user_id            VARCHAR(100) REFERENCES users (id),
    source_wallet_id   VARCHAR(100) REFERENCES wallets (id),
    destination        VARCHAR(100),
    amount             BIGINT,
    description        VARCHAR(100),
    payment_method_id  VARCHAR(100) REFERENCES payment_method (id),
    created_at         TIMESTAMP WITH TIME ZONE
);
