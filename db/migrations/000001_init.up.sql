BEGIN;

CREATE DOMAIN CUID AS VARCHAR(25);

-- author
CREATE TABLE IF NOT EXISTS author (
    author_id CUID PRIMARY KEY,
    full_name VARCHAR(64) NOT NULL,
    description TEXT,
    photo_url TEXT
);

-- book
CREATE TABLE IF NOT EXISTS book (
    book_id CUID PRIMARY KEY,
    title TEXT,
    year SMALLINT,
    publisher VARCHAR(64),
    description TEXT,
    photo_url TEXT 
);

-- book's author [relationship]
CREATE TABLE IF NOT EXISTS book_author (
    book_id CUID NOT NULL REFERENCES book(book_id),
    author_id CUID NOT NULL REFERENCES author(author_id),
    PRIMARY KEY (book_id, author_id)
);

-- format type
CREATE TABLE IF NOT EXISTS format_type (
    id VARCHAR(20) NOT NULL PRIMARY KEY,
    comment VARCHAR(30)
);

INSERT INTO format_type (id) VALUES ('E-Book'), ('Hardcover'), ('Paperback') ON CONFLICT DO NOTHING;

-- book's format
CREATE TABLE IF NOT EXISTS format (
    format_id VARCHAR(20) PRIMARY KEY,
    book_id CUID NOT NULL REFERENCES book(book_id),
    format_type VARCHAR(10) NOT NULL REFERENCES format_type(id) ON UPDATE CASCADE,
    price FLOAT NOT NULL, 
    supply INT DEFAULT 0 NOT NULL
);

-- category
CREATE TABLE IF NOT EXISTS category (
    category_id VARCHAR(20) PRIMARY KEY,
    comment TEXT
);

-- book category [relationship]
CREATE TABLE IF NOT EXISTS book_category (
    book_id CUID NOT NULL REFERENCES book(book_id),
    category_id VARCHAR(20) NOT NULL REFERENCES category(category_id) ON UPDATE CASCADE,
    PRIMARY KEY (book_id, category_id)
);

COMMIT;
