CREATE TABLE file_ownership(
    id serial NOT NULL PRIMARY KEY,
    file_path VARCHAR(200) UNIQUE,
    file_url VARCHAR(200) UNIQUE,
    owner_id INTEGER,
    created_at timestamp,
    updated_at timestamp
);