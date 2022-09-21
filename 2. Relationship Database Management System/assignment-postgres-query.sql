-- Creating table customers
create table if not exists customers (
	id serial not null primary key,
	customer_name char(50) not null
);
-- Creating table products
create table if not exists products (
	id serial not null primary key,
	name char(50) not null
);