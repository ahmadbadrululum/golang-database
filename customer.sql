alter table customer
add column email varchar(100),
add column balance int default 0,
add column status double default 0.0,
add column created_at timestamp default current_timestamp,
add column birth_date date,
add column married boolean default false;

create table user (
    username varchar(100) not null,
    password varchar(100) not null,
    primary key (username)
);