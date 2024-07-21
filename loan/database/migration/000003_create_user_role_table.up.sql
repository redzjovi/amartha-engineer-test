CREATE TABLE user_role (
    user_id INTEGER NOT NULL,
    role_code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, role_code)
);
