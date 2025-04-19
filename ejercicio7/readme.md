1. Crear DB
create database if not exists vi;

2. Crear tabla
create table if not exists albums(id int auto_increment primary key,title varchar(255) not null,artist varchar(255) not null,year int,price real);

3. Ver
show tables

4. insertar data
insert into albums (title, artist, year, price) values 
('The Number of the Beast', 'Iron Maiden', 1982, 25.19),
('Youthanasia', 'Megadeth', 1994, 13.65),
('Master of Puppets', 'Metallica', 1986, 20.97);

5. Ver data 
select * from albums;