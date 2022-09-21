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
-- Inserting data into customers
insert into customers(customer_name) values('Lewis');
insert into customers(customer_name) values('Nikita');
insert into customers(customer_name) values('Johan');
insert into customers(customer_name) values('Verstrappen');
insert into customers(customer_name) values('Sebastian');

-- Inserting data into products
insert into products(name) values('Ikan Tongkol');
insert into products(name) values('Ikan Lele');
insert into products(name) values('Ikan Mas');
insert into products(name) values('Ikan Patin');
insert into products(name) values('Ikan Kakap');

-- Inserting data into orders
insert into orders(customer_id, product_id, order_date, total) values(1, 5, now(), 100);
insert into orders(customer_id, product_id, order_date, total) values(2, 4, now(), 200);
insert into orders(customer_id, product_id, order_date, total) values(3, 3, now(), 300);
insert into orders(customer_id, product_id, order_date, total) values(4, 2, now(), 400);
insert into orders(customer_id, product_id, order_date, total) values(5, 1, now(), 500);
-- or shorthand
insert into orders(customer_id, product_id, order_date, total) 
values(1, 5, now(), 100), (2, 4, now(), 200), (3, 3, now(), 300), (4, 2, now(), 400), (5, 1, now(), 500);