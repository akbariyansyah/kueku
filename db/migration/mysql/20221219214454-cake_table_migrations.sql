-- +migrate Up
CREATE TABLE cakes (
    id INTEGER PRIMARY KEY,
    title VARCHAR (100) NOT NULL,
    description VARCHAR (100) NOT NULL,
    image VARCHAR (255) NOT NULL,
    rating INTEGER NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

-- +migrate Down
DROP TABLE users;
