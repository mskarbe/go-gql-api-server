BEGIN;

CREATE DOMAIN CUID AS VARCHAR(25);

-- author
CREATE TABLE IF NOT EXISTS author (
    author_id CUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    description TEXT,
    photo_url TEXT
);

-- book
CREATE TABLE IF NOT EXISTS book (
    book_id CUID PRIMARY KEY,
    year SMALLINT,
    publisher VARCHAR(64),
    description TEXT,
    photo_url TEXT 
);

-- book's author [relationship]
CREATE TABLE IF NOT EXISTS bookAuthor (
    book_id CUID NOT NULL REFERENCES book(book_id),
    author_id CUID NOT NULL REFERENCES author(author_id),
    PRIMARY KEY (book_id, author_id)
);

-- format type
CREATE TABLE IF NOT EXISTS format_type (
    id VARCHAR(10) NOT NULL PRIMARY KEY,
    comment VARCHAR(30)
);

INSERT INTO format_type (id) VALUES ('E-Book'), ('Hardcover'), ('Paperback') ON CONFLICT DO NOTHING;

-- book's format
CREATE TABLE IF NOT EXISTS format (
    format_id CUID PRIMARY KEY,
    book_id CUID NOT NULL REFERENCES book(book_id),
    format_type VARCHAR(10) NOT NULL REFERENCES format_type(id) ON UPDATE CASCADE,
    price FLOAT NOT NULL, 
    supply INT DEFAULT 0 NOT NULL
);

-- category
CREATE TABLE IF NOT EXISTS category (
    category_id CUID PRIMARY KEY,
    comment TEXT
);

-- book category [relationship]
CREATE TABLE IF NOT EXISTS book_category (
    book_id CUID NOT NULL REFERENCES book(book_id),
    category_id CUID NOT NULL REFERENCES category(category_id) ON UPDATE CASCADE,
    PRIMARY KEY (book_id, category_id)
);

COMMIT;
