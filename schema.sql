
DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL
);

INSERT INTO authors (name) VALUES ('Дмитрий'),('Василий'),('Пётр');
INSERT INTO posts (author_id, title, content, created_at)
VALUES(0, 'Статья', 'Содержание статьи', 0),( 0, 'Статья 2', 'Содержание статьи 2', 11);