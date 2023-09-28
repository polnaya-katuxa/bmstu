-- Создание БД

drop database if exists RK2;

create database RK2;

-- Создание таблиц БД

drop table if exists subjects cascade;
drop table if exists departments cascade;
drop table if exists teachers cascade;
drop table if exists ts cascade;

create table subjects (
    id serial primary key,
    name text not null,
    hours int not null check(hours > 0),
    sem int not null check(sem > 0 and sem < 9),
    rate int not null check(rate > 0 and rate < 6)
);

create table departments (
    id serial primary key,
    name text not null,
    info text not null
);

create table teachers (
    id serial primary key,
    fio text not null,
    ph text not null,
    pos text not null,
    dep_id int references departments not null
);

create table ts (
    tid int references teachers not null,
    sid int references subjects not null
);

-- Заполнение таблиц БД

insert into subjects(name, hours, sem, rate) values ('Русский язык', 50, 2, 4);
insert into subjects(name, hours, sem, rate) values ('Математика', 35, 1, 5);
insert into subjects(name, hours, sem, rate) values ('Информатика', 76, 6, 3);
insert into subjects(name, hours, sem, rate) values ('Базы данных', 24, 4, 5);
insert into subjects(name, hours, sem, rate) values ('Сети', 45, 5, 2);
insert into subjects(name, hours, sem, rate) values ('Котики', 58, 4, 1);
insert into subjects(name, hours, sem, rate) values ('Щенки', 33, 3, 4);
insert into subjects(name, hours, sem, rate) values ('Как плавать', 88, 4, 4);
insert into subjects(name, hours, sem, rate) values ('Биология', 90, 2, 5);
insert into subjects(name, hours, sem, rate) values ('Сборка коллайдера', 32, 3, 3);

insert into departments(name, info) values ('Крутая кафедра', 'тут учат как жить красиво и кучеряво');
insert into departments(name, info) values ('Норм кафедра', 'тут учат как жить плохо и кучеряво');
insert into departments(name, info) values ('Очень кафедра', 'тут учат как жить красиво и модно');
insert into departments(name, info) values ('Привет кафедра', 'тут учат как жить стильно и кучеряво');
insert into departments(name, info) values ('Пока кафедра', 'тут учат как жить красиво и здорово');
insert into departments(name, info) values ('Огого кафедра', 'тут учат как жить мощно и кучеряво');
insert into departments(name, info) values ('Моя кафедра', 'тут учат как жить красиво и отлично');
insert into departments(name, info) values ('Наша кафедра', 'тут учат как жить припеваючи и кучеряво');
insert into departments(name, info) values ('Лучшая кафедра', 'тут учат как жить красиво и сильно');
insert into departments(name, info) values ('Плохая кафедра', 'тут учат как жить грациозно и кучеряво');

insert into teachers(fio, ph, pos, dep_id) values ('Семёнов Семён Семонович', 'Доктор наук', 'Профессор', 1);
insert into teachers(fio, ph, pos, dep_id) values ('Голубева Дарина Егоровна', 'Кандидат наук', 'Доцент', 2);
insert into teachers(fio, ph, pos, dep_id) values ('Ефимова Лейла Кирилловна', 'Магистр', 'Ассистент', 3);
insert into teachers(fio, ph, pos, dep_id) values ('Балашова Елизавета Всеволодовна', 'Доктор наук', 'Профессор', 2);
insert into teachers(fio, ph, pos, dep_id) values ('Крюков Тимофей Георгиевич', 'Доктор наук', 'Профессор', 4);
insert into teachers(fio, ph, pos, dep_id) values ('Кочеткова Вера Ярославовна', 'Кандидат наук', 'Доцент', 8);
insert into teachers(fio, ph, pos, dep_id) values ('Назарова Алина Артёмовна', 'Доктор наук', 'Профессор', 9);
insert into teachers(fio, ph, pos, dep_id) values ('Николаев Андрей Александрович', 'Магистр', 'Ассистент', 5);
insert into teachers(fio, ph, pos, dep_id) values ('Абрамова Злата Даниловна', 'Доктор наук', 'Профессор', 5);
insert into teachers(fio, ph, pos, dep_id) values ('Новикова Алёна Арсеновна', 'Магистр', 'Ассистент', 5);

insert into ts(tid, sid) values (1, 2);
insert into ts(tid, sid) values (3, 5);
insert into ts(tid, sid) values (4, 2);
insert into ts(tid, sid) values (1, 7);
insert into ts(tid, sid) values (8, 2);
insert into ts(tid, sid) values (3, 1);
insert into ts(tid, sid) values (10, 9);
insert into ts(tid, sid) values (4, 6);
insert into ts(tid, sid) values (7, 7);
insert into ts(tid, sid) values (7, 6);
insert into ts(tid, sid) values (1, 3);

-- Запросы

-- 2.1) предикат сравнения с квантором
-- Вывести название предмета и его часы, на который выделено часов больше, чем на все предметы,
-- которые ведут доктора наук

select name, hours
from subjects
where hours > all(select hours
                  from ((subjects join ts on subjects.id = ts.sid) as j1 join teachers on j1.tid = teachers.id) as j2
                  where ph = 'Доктор наук');

-- 2.2) агрегатные функции в выражениях столбцов
-- Вывести максимальное и минимальное количество выделенных часов для предметов
-- из 1-3 семестров

select min(hours), max(hours)
from subjects
where sem between 1 and 3;

-- 2.3) создание новой временной локальной таблицы из результата селект
-- Создать временную табличку из преподавателей, работающих на кафедрах,
-- описание которых начинается со слов "тут учат как жить красиво"

drop table if exists teachers_info cascade;

select fio, pos, dep_id
into temporary teachers_info
from teachers
where dep_id in (select id
                 from departments
                 where info like 'тут учат как жить красиво%');

select * from teachers_info;

-- 3) Создать хранимую процедуру, принимающую имя таблицы, которая выводит сведения
-- об индексах этой таблицы в текущей бд, и протестировать.

create or replace procedure get_indexes(table_name text)
as $$
    declare c cursor
        for select *
            from pg_catalog.pg_indexes as pi
            where pi.tablename = table_name;
        r record;
    begin
        open c;
        loop
        fetch c into r;
        exit when not found;
        raise notice 'Found index for "%" in "%": %, %, %, %, %', table_name, current_database(), r.schemaname, r.tablename, r.indexname, r.tablespace, r.indexdef;
        end loop;
        close c;
    end
$$ language plpgsql;

call get_indexes('departments');
call get_indexes('pg_constraint');

