CREATE TABLE IF NOT EXISTS user_email
(
    user_id  BIGINT REFERENCES tasteit_user (id),
    email    TEXT NOT NULL UNIQUE,
    provider TEXT NOT NULL
);

INSERT INTO user_email(user_id, email, provider)
SELECT id, email, provider
FROM tasteit_user;

ALTER TABLE tasteit_user
    DROP COLUMN email,
    DROP COLUMN provider;