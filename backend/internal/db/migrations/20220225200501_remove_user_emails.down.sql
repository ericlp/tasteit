CREATE TABLE IF NOT EXISTS user_email
(
    user_id  uuid REFERENCES tasteit_user (id),
    email    TEXT NOT NULL UNIQUE,
    provider TEXT NOT NULL
);