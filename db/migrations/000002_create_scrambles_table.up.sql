CREATE TABLE scrambles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,

    time int NOT NULL,
    scramble text NOT NULL,

    created_on timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_on timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
