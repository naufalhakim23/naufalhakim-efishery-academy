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
-- Creating table orders
create table if not exists orders (
	id serial not null primary key,
	customer_id int not null references customers(id),
	product_id int not null references products(id),
	order_date timestamp not null,
	total float not null
);
