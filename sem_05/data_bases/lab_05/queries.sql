-- 1. Из таблиц базы данных, созданной в первой лабораторной работе,
-- извлечь данные в XML (MSSQL) или JSON(Oracle, Postgres).
copy (select row_to_json(c)
from clients c)
to '/labs/lab_05/data/clients.json';

-- 2. Выполнить загрузку и сохранение XML или JSON файла в таблицу.
-- Созданная таблица после всех манипуляций должна соответствовать
-- таблице базы данных, созданной в первой лабораторной работе.
drop table if exists clients_json;
create table clients_json (
    id           serial primary key,
    name         text     not null,
    surname      text     not null,
    patronymic   text     not null,
    birth_date   date     not null,
    sex          sex_type not null,
    phone_number text     not null unique,
    login        text     not null unique
);

drop table if exists json_clients_info;
create table json_clients_info (
    info json
);

copy json_clients_info from '/labs/lab_05/data/clients.json';
select * from json_clients_info;

select * from json_clients_info, json_populate_record(null::clients_json, info);

truncate table clients_json;

insert into clients_json
select id, name, surname, patronymic, birth_date, sex, phone_number, login
from json_clients_info as j, json_populate_record(null::clients_json, j.info);

select * from clients_json;

-- 3. Создать таблицу, в которой будет атрибут(-ы) с типом XML или JSON,
-- или добавить атрибут с типом XML или JSON к уже существующей таблице.
-- Заполнить атрибут правдоподобными данными с помощью команд INSERT или UPDATE.
alter table attendance add column feedback jsonb;

update attendance
set feedback = case
    when rating = 1 then '{"staff": "they are rude and dirty", "machines": "so old that even tetris does not work on them"}'::jsonb
    when rating = 2 then '{"staff": "they are dirty", "atmosphere": "staying home is better", "machines": "only tetris works on them"}'::jsonb
    when rating = 3 then '{"atmosphere": "i fell asleep so idk", "machines": "somebody was playing tetris, so i suppose they are ok"}'::jsonb
    when rating = 4 then '{"staff": "they do not know funny jokes", "atmosphere": "better then staying at home", "machines": "vs code works well"}'::jsonb
    when rating = 5 then '{"staff": "my brother works here", "atmosphere": "i want to stay here forever", "machines": "i played cyberpunk 2077 and it did not even warm up"}'::jsonb
    else '{"staff": "-", "atmosphere": "-"}'::jsonb
    end;

-- 4. Выполнить следующие действия:
-- 4.1. Извлечь XML/JSON фрагмент из XML/JSON документа
select id_client, id_club, rating, price, feedback->'machines' as m
from attendance
where id_client = 200 and rating = 5 and id_club = 511;

-- 4.2. Извлечь значения конкретных узлов или атрибутов XML/JSON
-- документа
select id_client, id_club, rating, price, feedback->>'machines' as m
from attendance
where id_client = 200 and rating = 5 and id_club = 511;

-- 4.3. Выполнить проверку существования узла или атрибута
select id_client, id_club, rating, price
from attendance
where feedback ? 'staff';

-- 4.4. Изменить XML/JSON документ
update attendance
set feedback = feedback || '{"games":"no normal games only zmeyka"}'::jsonb
where rating = 1 and id_client = 304;

select id_client, rating, price, feedback
from attendance
where rating = 1 and id_client = 304;

-- 4.5. Разделить XML/JSON документ на несколько строк по узлам
select * from  json_each((select distinct feedback from attendance where id_client = 200 and rating = 5 and id_club = 511)::json);

-- защита
drop function if exists hate_person;
create or replace function hate_person(id_bad_person int, id_end int) returns table(info json)
as $$
    begin
        update attendance set rating = rating + 1 where id_client between id_bad_person and id_end and rating < 5;
        update cards set state = 'non-activated' where id_client between id_bad_person and id_end ;
        return query (select row_to_json(row)
                      from (
                          select login, phone_number as info
                          from clients
                          where id between id_bad_person and id_end
                           ) as row);
    end;
$$
language plpgsql;

copy (
    select hate_person(202, 206)
    )
to '/labs/lab_05/data/haters.json';

--select array_to_json(array((select distinct feedback from attendance where id_client = 200)));

-- поля: логин, фамилия, список адресов посещённых данным клиентом клубов
-- copy (select row_to_json(clients_att_info) from (select
--     c.login, c.surname,
--     array_to_json(
--         array(
--             (select ca.address
--              from (attendance as a join computer_clubs cc on a.id_club = cc.id) as ca
--              where ca.id_client = c.id)
--         )
--     ) as ca_exact
-- from clients as c) clients_att_info) to '/labs/lab_05/data/clients_att_info.json';