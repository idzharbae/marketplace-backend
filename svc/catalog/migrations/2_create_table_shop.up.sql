CREATE TABLE shop (
    id serial NOT NULL PRIMARY KEY,
	name VARCHAR(300),
	address VARCHAR(300),
	longitude numeric(28,4),
	latitude numeric(28,4),
	created_at timestamptz,
	updated_at timestamptz
);
ALTER TABLE product
ADD CONSTRAINT shop_fk FOREIGN KEY (shop_id) REFERENCES shop (id);