-- migrate:up

CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY,
	username text NOT NULL UNIQUE,
	password text NOT NULL	
)

-- migrate:down

DROP TABLE IF EXISTS users;
