CREATE TABLE IF NOT EXISTS owner
(
    id      uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name    TEXT NOT NULL,
    is_user BOOLEAN,
    UNIQUE (name, is_user)
);

CREATE TABLE IF NOT EXISTS user_owner
(
    owner_id        uuid REFERENCES owner (id),
    tasteit_user_id uuid REFERENCES tasteit_user (id),
    PRIMARY KEY (owner_id, tasteit_user_id)
);