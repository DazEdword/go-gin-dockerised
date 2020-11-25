CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);

INSERT INTO users VALUES (1, 'alice', 'password1', 'alice@localhost.com');
INSERT INTO users VALUES (2, 'bob', 'password2', 'bob@localhost.com');
INSERT INTO users VALUES (3, 'rich', 'password3', 'rich@localhost.com');
