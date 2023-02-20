-- migrate:up

CREATE TABLE IF NOT EXISTS users (
    userId BIGSERIAL PRIMARY KEY,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    password VARCHAR ( 50 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS scrambles (
    scrambleId BIGSERIAL PRIMARY KEY,
    cube VARCHAR ( 50 ) NOT NULL,
    scrambleStr VARCHAR ( 50 ) NOT NULL,
    time INTEGER NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),

    userId BIGINT NOT NULL REFERENCES users (userId)
);


-- migrate:down
DROP TABLE users;
