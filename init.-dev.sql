create table products (
  id serial primary key,
  name varchar,
  description varchar,
  price decimal,
  quantity integer
);

insert into
  products (name, description, price, quantity)
values
  ("Camiseta", "Preta", 19, 10),
  ("Calça", "Azul", 59, 2);