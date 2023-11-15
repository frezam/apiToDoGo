-- Table: users
CREATE TABLE IF NOT EXISTS users
(
    id       UUID PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    email    VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Table: tasks
CREATE TABLE IF NOT EXISTS tasks
(
    id          UUID PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    user_id     UUID REFERENCES users (id)
);