create table if not exists users (
    id serial,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at timestamp,
    unique (email),
    primary key (id)
);