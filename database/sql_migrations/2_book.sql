-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE book (
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url VARCHAR(256),
    release_year INTEGER,
    price VARCHAR(256),
    total_page VARCHAR(256),
    thickness VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    category_id INTEGER,
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
            REFERENCES category(id)
)

-- +migrate StatementEnd