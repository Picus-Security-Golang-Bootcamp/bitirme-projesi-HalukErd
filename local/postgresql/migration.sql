create table products
(
    id         text not null
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text,
    code       text,
    price      numeric
);

create table users
(
    id         text not null
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email      text,
    password   text,
    roles      text
);

alter table users
    owner to "HERDOGAN";

alter table products
    owner to "HERDOGAN";

create unique index idx_users_email
    on users (email);

create index idx_users_deleted_at
    on users (deleted_at);

create index idx_products_deleted_at
    on products (deleted_at);

INSERT INTO public.users (id, created_at, updated_at, deleted_at, email, password, roles) VALUES ('fa2c5c4a-bebc-11ec-8914-ca38388bbb3f', '2022-04-18 02:12:20.198329 +00:00', '2022-04-18 02:12:20.198329 +00:00', null, 'admin', '$2a$14$EQ2czBtRjudsM72Wx2ryL.aWoDNvW159AkmNKloAlJPqUKkVBb/ki', 'admin');

