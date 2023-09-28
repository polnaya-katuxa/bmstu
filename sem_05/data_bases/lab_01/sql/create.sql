drop table if exists games_exact_machines;
drop table if exists machines_clubs;
drop table if exists cards;
drop table if exists attendance;
drop table if exists staff;
drop table if exists machines;
drop table if exists clients;
drop table if exists loyalty_programs;
drop table if exists computer_clubs;
drop table if exists games;

drop type if exists age_category;
drop type if exists sex_type;
drop type if exists machine_type;
drop type if exists state_type;

create type age_category as enum (
    'EC', -- 3+
    'E', -- 6+
    'E10+', -- 10+
    'T', -- 13+
    'M', -- 17+
    'AO', -- 18+
    'RP', -- PENDING
    'RP 17+' -- PENDING 17+
    );

create type sex_type as enum (
    'male',
    'female'
    );

create type machine_type as enum (
    'computer',
    'game console'
    );

create type state_type as enum (
    'activated',
    'non-activated'
    );

create table games
(
    id               serial primary key,
    name             text           not null,
    genre            text           not null,
    release_year     int            not null,
    company          text,
    country          text,
    age_rating       age_category   not null,
    price            numeric(10, 2) not null,
    multiplayer_mode boolean        not null default false
);

create table computer_clubs
(
    id                 serial primary key,
    address            text    not null unique,
    open_time          time    not null default '09:00',
    close_time         time    not null default '22:00',
    establishment_date date    not null,
    parking_spot_num   int     not null default 0,
    is_round_the_clock boolean not null default false
);

create table loyalty_programs
(
    id                   serial primary key,
    name                 text           not null unique,
    design               text           not null default 'computer.club/images/default_design.png',
    cashback_percent     int            not null unique,
    minimum_purchase_sum numeric(10, 2) not null unique
);

create table clients
(
    id           serial primary key,
    name         text     not null,
    surname      text     not null,
    patronymic   text     not null,
    birth_date   date     not null,
    sex          sex_type not null,
    phone_number text     not null unique,
    login        text     not null unique
);

create table machines
(
    id           serial primary key,
    brand        text         not null,
    model        text         not null,
    country      text         not null,
    release_year int          not null,
    type         machine_type not null
);

create table staff
(
    id              serial primary key,
    id_club         int references computer_clubs (id) on delete cascade,
    name            text     not null,
    surname         text     not null,
    patronymic      text     not null,
    birth_date      date     not null,
    sex             sex_type not null,
    phone_number    text     not null,
    employment_date date     not null,
    position        text     not null
);

create table attendance
(
    time_start timestamp without time zone,
    id_club    int references computer_clubs (id) on delete cascade,
    id_client  int references clients (id) on delete cascade,
    time_end   timestamp without time zone not null,
    rating     int,
    price      numeric(10, 2)              not null default 400.00,
    primary key (time_start, id_club, id_client)
);

create table cards
(
    number             text primary key,
    loyalty_program_id int references loyalty_programs (id) on delete cascade,
    id_client          int references clients (id) on delete cascade,
    registration_date  date       not null,
    state              state_type not null default 'non-activated'
);

create table machines_clubs
(
    id         serial primary key,
    id_club    int references computer_clubs (id) on delete cascade,
    id_machine int references machines (id) on delete cascade
);

create table games_exact_machines
(
    id_game           int references games (id) on delete cascade,
    id_machines_clubs int references machines_clubs (id) on delete cascade,
    primary key (id_game, id_machines_clubs)
);