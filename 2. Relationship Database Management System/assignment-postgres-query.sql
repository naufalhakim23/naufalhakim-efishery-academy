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

-- Update customers table
update customers set customer_name = 'Lewis Hamilton' where id = 1;
update customers set customer_name = 'Nikita Mazepin' where id = 2;
update customers set customer_name = 'Johan Sebastian' where id = 3;
update customers set customer_name = 'Max Verstrappen' where id = 4;
update customers set customer_name = 'Sebastian Vettel' where id = 5;

-- Update products table
update products set name = 'Ikan Bandeng' where id = 1;
update products set name = 'Ikan Tuna' where id = 2;
update products set name = 'Ikan Gurame' where id = 3;
update products set name = 'Ikan Kembung' where id = 4;
update products set name = 'Ikan Nila' where id = 5;

-- Update orders table
update orders set total = 1000 where id = 1;
update orders set product_id = 1 where id = 2; -- execute this to clear out foreign key constraint in orders table (product_id)
update orders set customer_id = 1 where id = 2; -- execute this to clear out foreign key in orders table to customers table
update orders set product_id = 1 where id = 3; -- execute this to clear out foreign key constraint in orders table (product_id)
update orders set customer_id = 1 where id = 3; -- execute this to clear out foreign key in orders table to customers table
update orders set order_date = now() where id = 4;
update orders set total = 5000 where id = 5;

-- Delete data from orders table
delete from orders where id = 2;
delete from orders where total = 5000;

-- Delete data from customers table
delete from customers where id = 2;
delete from customers where customer_name = 'Johan Sebastian';

-- Delete data from products table
delete from products where name = 'Ikan Kembung';
delete from products where id = 3;