CREATE TABLE user_account(
    id serial NOT NULL PRIMARY KEY,
    name VARCHAR(200),
	user_name VARCHAR(100) UNIQUE,
	email VARCHAR(100) UNIQUE,
	phone VARCHAR(50),
	password VARCHAR(100),
	type INTEGER,
	created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
	updated_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);