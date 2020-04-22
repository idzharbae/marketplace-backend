CREATE TABLE payment(
    id serial NOT NULL PRIMARY KEY,
    order_id INTEGER,
    amount INTEGER,
    payment_method INTEGER,
    payment_status INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
)