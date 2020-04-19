CREATE TABLE cart(
    id serial NOT NULL PRIMARY KEY,
    product_id INTEGER,
    user_id INTEGER,
    quantity_kg NUMERIC(10, 2),
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);