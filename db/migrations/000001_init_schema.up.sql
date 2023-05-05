
CREATE TABLE users (
	id BIGSERIAL PRIMARY KEY,
	username text NOT NULL UNIQUE,
	password text NOT NULL	
)
