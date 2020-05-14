CREATE TABLE review(
    id serial NOT NULL PRIMARY KEY,
	user_id INTEGER,
	product_id INTEGER,
	shop_id INTEGER,
	title VARCHAR(200),
	content VARCHAR(600),
	photo_url VARCHAR(200),
	rating NUMERIC(10, 2),
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);