CREATE TABLE product(
    id serial NOT NULL PRIMARY KEY,
	shop_id INTEGER,
	name VARCHAR(400),
	slug VARCHAR(400),
	quantity INTEGER,
	price_per_kg INTEGER,
	stock_kg numeric(10,2),
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);