 create table if not exists albums(
    id int auto_increment primary key,
    title varchar(255) not null,
    artist varchar(255) not null,
    year int,
    price real);