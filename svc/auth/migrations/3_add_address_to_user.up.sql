ALTER TABLE user_account
    ADD COLUMN province VARCHAR(200),
    ADD COLUMN city VARCHAR(200),
    ADD COLUMN detail_address VARCHAR(400),
    ADD COLUMN zip_code INTEGER,
    ADD COLUMN description VARCHAR(400);