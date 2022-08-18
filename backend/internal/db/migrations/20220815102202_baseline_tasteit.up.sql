CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasteit_user
(
    id   uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    nick TEXT             NOT NULL,
    cid  TEXT             NOT NULL
);

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

CREATE TABLE IF NOT EXISTS recipe
(
    id             uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name           TEXT UNIQUE      NOT NULL,
    unique_name    TEXT UNIQUE      NOT NULL,
    description    TEXT,
    oven_temp      BIGINT,
    estimated_time BIGINT,
    deleted        BOOLEAN          NOT NULL,
    portions       INTEGER          NOT NULL,
    owned_by       uuid REFERENCES owner (id)
);

CREATE TABLE IF NOT EXISTS image
(
    id   uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT             NOT NULL
);


CREATE TABLE IF NOT EXISTS recipe_image
(
    image_id  uuid REFERENCES image (id),
    recipe_id uuid REFERENCES recipe (id),
    PRIMARY KEY (image_id, recipe_id)
);

CREATE TABLE IF NOT EXISTS ingredient
(
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS unit
(
    name TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS recipe_ingredient
(
    id              uuid PRIMARY KEY            NOT NULL DEFAULT uuid_generate_v4(),
    ingredient_name TEXT REFERENCES ingredient (name),
    unit_name       TEXT REFERENCES unit (name) NOT NULL,
    amount          NUMERIC                     NOT NULL,
    number          INTEGER                     NOT NULL,
    is_heading      BOOLEAN                              DEFAULT false NOT NULL,
    recipe_id       uuid REFERENCES recipe (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe_step
(
    recipe_id  uuid REFERENCES recipe (id),
    number     INTEGER               NOT NULL,
    step       TEXT                  NOT NULL,
    is_heading BOOLEAN DEFAULT false NOT NULL,
    PRIMARY KEY (recipe_id, number)
);

CREATE TABLE IF NOT EXISTS recipe_book
(
    id          uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name        TEXT             NOT NULL UNIQUE,
    unique_name TEXT             NOT NULL UNIQUE,
    author      TEXT             NOT NULL,
    deleted     BOOLEAN          NOT NULL,
    owned_by    uuid REFERENCES owner (id)
);

CREATE TABLE IF NOT EXISTS recipe_book_recipe
(
    recipe_book_id uuid REFERENCES recipe_book (id),
    recipe_id      uuid REFERENCES recipe (id),
    PRIMARY KEY (recipe_book_id, recipe_id)
);

CREATE TABLE IF NOT EXISTS recipe_book_image
(
    recipe_book_id uuid REFERENCES recipe_book (id),
    image_id       uuid REFERENCES image (id),
    PRIMARY KEY (recipe_book_id, image_id)
);

CREATE TABLE tag
(
    id          uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name        TEXT             NOT NULL,
    description TEXT             NOT NULL,
    color_red   INTEGER          NOT NULL,
    color_green INTEGER          NOT NULL,
    color_blue  INTEGER          NOT NULL,
    owned_by    uuid REFERENCES owner (id),
    UNIQUE (name)
);

CREATE TABLE recipe_tag
(
    recipe_id uuid REFERENCES recipe (id),
    tag_id    uuid REFERENCES tag (id),
    UNIQUE (recipe_id, tag_id)
);
