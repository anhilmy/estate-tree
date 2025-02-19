-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE estate (
	id serial PRIMARY KEY,
	uuid uuid NOT NULL default uuid_generate_v4(),
	length int NOT NULL,
	width int NOT NULL
);

CREATE TABLE tree (
	id serial PRIMARY KEY,
	uuid uuid NOT NULL default uuid_generate_v4(),
	x_axis  int not null,
	y_axis int not null,
	height int not null
);