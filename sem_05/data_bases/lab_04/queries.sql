create extension if not exists plpython3u;

-- Защита: попытка добавить клиента -> проверять, существуют ли такие ФИО и дата, если он есть,
-- обновить только остальные поля
drop function if exists check_one;
create or replace function check_one(n text, s text, p text, bd date) returns bool
as $$
    begin
        return (select count(c.id) > 0 as a
                from clients as c
                where c.name = n and c.surname = s and c.birth_date = bd and c.patronymic = p);
    end;
$$
language plpgsql;

create or replace function do_clients_check()
returns trigger
language plpython3u
as $$
    name = TD["new"]["name"]
    surname = TD["new"]["surname"]
    patr = TD["new"]["patronymic"]
    bird = TD["new"]["birth_date"]
    id = TD["new"]["id"]
    sex = TD["new"]["sex"]
    phone_number = TD["new"]["phone_number"]
    login = TD["new"]["login"]

    query1 = f"""
    insert into clients values ('{TD['new']['id']}', '{TD['new']['name']}', '{TD['new']['surname']}', '{TD['new']['patronymic']}'
    , '{TD['new']['birth_date']}', '{TD['new']['sex']}', '{TD['new']['phone_number']}', '{TD['new']['login']}')
    """
    query2 = f"""
    update clients
    set sex = '{sex}', phone_number = '{phone_number}', login = '{login}'
    where name = '{name}' and surname = '{surname}' and birth_date = '{bird}' and patronymic = '{patr}'
    """
    query3 = f"""
    select check_one('{name}', '{surname}', '{patr}', '{bird}')
    """

    res1 = plpy.execute(query3)
    plpy.notice(res1[0]['check_one'])
    if (res1[0]['check_one']):
        plpy.notice("update")
        res2 = plpy.execute(query2)
    else:
        plpy.notice("insert")
        res = plpy.execute(query1)

    return None
$$;

drop view if exists clients2;
create view clients2 as
select *
from clients;

drop trigger if exists insert_cli on clients;

create trigger insert_cli
    instead of insert on clients2
    for each row
    execute function do_clients_check();

insert into clients2(id, name, surname, patronymic, birth_date, sex, phone_number, login)
values ((select count(*) + 1 from clients), 'Евгений', 'Щербаков', 'Артемович', '2002-06-02', 'male', '+7 964 668-66-16', 'kfnaedjfhjfhe');

insert into clients2(id, name, surname, patronymic, birth_date, sex, phone_number, login)
values ((select count(*) + 1 from clients), 'Илья', 'Щербаков', 'Артёмович', '2003-06-02', 'male', '+7 964 669-66-16', 'kfnaergkljekrwgjhkafhe');


-- 1. Скалярная функция CLR.
-- Средняя цена на игры младше определённого года
drop function if exists avg_price_arter_year;
create or replace function avg_price_arter_year(year int)
returns numeric
language plpython3u
as $$
    query = f"""
    select avg(price) as a
    from games
    where release_year >= {year}
    """
    res = plpy.execute(query)
    if res:
        return res[0]['a']
    return 0
$$
;

select avg_price_arter_year(2002);

-- 2. Пользовательская агрегатная функция CLR.
-- Средняя цена на игры младше определённого года c помощью агрегации
drop function if exists get_avg_price_arter_year;
create or replace function get_avg_price_arter_year(year int)
returns numeric
language plpython3u
as $$
    query = f"""
    select price as a
    from games
    where release_year >= {year}
    """
    res = plpy.execute(query)
    qlen = len(res)
    s = 0
    for x in res:
        s += x['a']
    s = s / qlen
    return s
$$;

select get_avg_price_arter_year(2002);

-- 3. Табличная функция CLR.
-- Получить табличку из имени, фамилии и возраста для клиентов с именами на заданную букву
drop function if exists get_names_by_letter;
create or replace function get_names_by_letter(letter text) returns table(n text, s text, a int)
language plpython3u
as $$
    import datetime
    def get_age(date):
        bday = datetime.datetime.strptime(date, "%Y-%m-%d")
        today = datetime.date.today()
        age = today.year - bday.year - ((today.month, today.day) < (bday.month, bday.day))
        return age
    query = f"""
    select distinct c.name, c.surname, c.birth_date
    from clients as c
    where c.name like ('{letter}%')
    """
    res = plpy.execute(query)
    if res:
        for client in res:
            yield (client['name'], client['surname'], get_age(client['birth_date']))
    return "","",0
$$;

select * from get_names_by_letter('К');

-- 4. Хранимая процедура CLR.
-- Поднять цену на заданную величину на все игры-паззлы
drop procedure if exists do_puzzle_up;
create or replace procedure do_puzzle_up(up numeric)
language plpython3u
as $$
    query = f"""
    update games set price = price + {up} where genre = 'Puzzle'
    """
    res = plpy.execute(query)
$$;

call do_puzzle_up(20);

-- 5. Триггер CLR.
-- Реакция на вставку Алексея
create or replace function do_staff_check()
returns trigger
language plpython3u
as $$
    query = f"""
    insert into staff values ('{TD['new']['id']}', '{TD['new']['id_club']}', '{TD['new']['name']}', '{TD['new']['surname']}'
    , '{TD['new']['patronymic']}', '{TD['new']['birth_date']}', '{TD['new']['sex']}', '{TD['new']['phone_number']}'
    , '{TD['new']['employment_date']}', '{TD['new']['position']}')
    """

    position = TD["new"]["position"]
    name = TD["new"]["name"]

    if (position == 'Бармен' and name == 'Алексей'):
        plpy.notice("no barmens named Alexey!!!!1!!1")
    else:
        plpy.notice("this Alexey is ok")
        res = plpy.execute(query)

    return None
$$;

drop view if exists staff_names;
create view staff_names as
select *
from staff;

drop trigger if exists insert_staff on staff;

create trigger insert_staff
    instead of insert on staff_names
    for each row
    execute function do_staff_check();

insert into staff_names(id, id_club, name, surname, patronymic, birth_date, sex, phone_number, employment_date, position)
values ((select count(*) + 1 from staff), 1, 'Алексей', 'Алёшин', 'Алексеевич', '1995-02-03', 'male', '+7 964 666-66-16', '2018-01-01', 'Уборщик')
on conflict do nothing;
insert into staff_names(id, id_club, name, surname, patronymic, birth_date, sex, phone_number, employment_date, position)
values ((select count(*) + 1 from staff), 1, 'Алексей', 'Алёшин', 'Алексеевич', '1995-02-03', 'male', '+7 964 666-66-16', '2018-01-01', 'Бармен');

-- 6. Тип данных CLR.
-- Тип данных под соответствие логина клиента проценту его скидки
drop type if exists clients_loyalty cascade;
create type clients_loyalty as
(
    login text,
    cashback_percent int
);

create or replace function get_clients_loyalty()
returns setof clients_loyalty
language plpython3u
as $$
	query = """
	select t2.login, t2.cashback_percent
	from (((clients join cards as c on clients.id = c.id_client) as t1
	    join loyalty_programs as l on l.id = t1.loyalty_program_id)) as t2
	"""
	res = plpy.execute(query)
	return ([row for row in res])
$$;

select * from get_clients_loyalty();