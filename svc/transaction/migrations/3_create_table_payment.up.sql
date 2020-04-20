CREATE TABLE payment(
    id serial NOT NULL PRIMARY KEY,
    order_id INTEGER,
    amount INTEGER,
    method INTEGER,
    status INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
)