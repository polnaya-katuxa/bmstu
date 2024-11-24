-- +goose Up
-- +goose StatementBegin

create extension "uuid-ossp";

create type perm as enum (
    'free',
    'paid'
    );

create table users (
                       uuid UUID default uuid_generate_v4() primary key,
                       login text not null,
                       picture text,
                       description text,
                       password text not null,
                       mail text not null check ( mail like '%@%.%' ),
                       balance int default 0 check ( balance >= 0 ),
                       enter_time timestamp not null,
                       is_admin bool default false
);

create table reaction_types (
                                uuid UUID default uuid_generate_v4() primary key,
                                icon text not null
);

create table limits (
                        uuid UUID default uuid_generate_v4() primary key,
                        value int unique not null check ( value >= 0 ),
                        bonus int not null check ( bonus >= 0 )
);

create table posts (
                       uuid UUID default uuid_generate_v4() primary key,
                       content text not null,
                       perms perm default 'free',
                       writer_id UUID references users (uuid) on delete cascade,
                       public_time timestamp not null,
                       limit_id UUID references limits (uuid) on delete cascade
);

create table reactions (
                           uuid UUID default uuid_generate_v4() primary key,
                           reaction_type_id UUID references reaction_types (uuid) on delete cascade,
                           post_id UUID references posts (uuid) on delete cascade,
                           reactor_id UUID references users (uuid) on delete cascade
);

create table subscriptions (
                               uuid UUID default uuid_generate_v4() primary key,
                               writer_id UUID references users (uuid) on delete cascade,
                               reader_id UUID references users (uuid) on delete cascade
);

create table balance_transactions (
                                      uuid UUID default uuid_generate_v4() primary key,
                                      reason text not null,
                                      time timestamp not null,
                                      user_id UUID references users (uuid) on delete cascade,
                                      amount int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists balance_transactions;
drop table if exists subscriptions;
drop table if exists reactions;
drop table if exists posts;
drop table if exists limits;
drop table if exists reaction_types;
drop table if exists users;

drop type if exists perm;
drop extension if exists "uuid-ossp";
-- +goose StatementEnd
