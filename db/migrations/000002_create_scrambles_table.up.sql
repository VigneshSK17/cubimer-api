CREATE TABLE IF NOT EXISTS scrambles (
    scrambleId BIGSERIAL PRIMARY KEY,
    cube VARCHAR ( 50 ) NOT NULL,
    scrambleStr VARCHAR ( 50 ) NOT NULL,
    time INTEGER NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT now(),

    userId BIGINT NOT NULL REFERENCES users (userId)
);
