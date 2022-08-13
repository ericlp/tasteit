ALTER TABLE tag
    ADD COLUMN created_by uuid REFERENCES tasteit_user (id),
    DROP COLUMN owned_by;

ALTER TABLE recipe_book
    ADD COLUMN created_by uuid REFERENCES tasteit_user (id),
    DROP COLUMN owned_by;

ALTER TABLE recipe
    ADD COLUMN created_by uuid REFERENCES tasteit_user (id),
    DROP COLUMN owned_by;

ALTER TABLE tasteit_user
    ADD COLUMN name TEXT NOT NULL default 'tasteit_user',
    DROP COLUMN cid,
    DROP COLUMN nick;

DROP TABLE IF EXISTS user_owner;

DROP TABLE IF EXISTS owner;