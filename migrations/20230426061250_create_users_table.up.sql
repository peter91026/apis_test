create table if not exists users(
    user_id serial primary key,
    userneame varchar(50) unique not null,
    password varchar(50) not null,
    email varchar(300)unique not null
);

