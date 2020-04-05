CREATE TABLE user_account(
    id serial NOT NULL PRIMARY KEY,
    name VARCHAR(200),
	user_name VARCHAR(100),
	email VARCHAR(100),
	phone VARCHAR(50),
	password VARCHAR(100),
	type INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);