CREATE TABLE saldo_history(
    id serial NOT NULL PRIMARY KEY,
    user_id INTEGER,
	change_amount INTEGER,
	description VARCHAR(200),
	source_id INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);