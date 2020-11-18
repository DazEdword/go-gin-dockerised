BEGIN;

CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS contacts(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   surname VARCHAR (50),
   address VARCHAR (300),
   email VARCHAR (300)
);

COMMIT;
