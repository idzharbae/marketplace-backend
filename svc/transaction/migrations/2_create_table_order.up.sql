CREATE TABLE orders(
    id serial NOT NULL PRIMARY KEY,
    product_id INTEGER ARRAY,
    shop_id INTEGER,
    user_id INTEGER,
    total_price INTEGER,
    status INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);