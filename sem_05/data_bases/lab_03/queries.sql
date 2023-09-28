-- Скалярная функция
-- Найти среднюю цену на игры старше заданного года
drop function if exists avg_price_arter_year;
create or replace function avg_price_arter_year(year int) returns numeric
as $$
    begin
        return (select avg(price) as a
                from games
                where release_year >= year);
    end;
$$
language plpgsql;

select avg_price_arter_year(2002);

-- Подставляемая табличная функция
-- Вывести уникальные имена клиентов, начинающиеся на заданную подстроку
drop function if exists names_by_letter;
create or replace function names_by_letter(letter text) returns table(n text)
as $$
    begin
        return query (select distinct c.name
                from clients as c
                where c.name like (letter || '%'));
    end;
$$
language plpgsql;

select names_by_letter('К');

-- Многооператорная табличная функция
-- Повысить рейтинги всех посещений на 1 и деактивировать все карты заданного клиента по id,
-- вывести его login и телефон
drop function if exists hate_person;
create or replace function hate_person(id_bad_person int) returns table(log text, phone text)
as $$
    begin
        update attendance set rating = rating + 1 where id_client = id_bad_person and rating < 5;
        update cards set state = 'non-activated' where id_client = id_bad_person;
        return query (select login, phone_number
                      from clients
                      where id = id_bad_person);
    end;
$$
language plpgsql;

select hate_person(202);

-- Рекурсивная функция
-- Вычислить факториал заданного числа
drop function if exists f_func;
create or replace function f_func(n int) returns int
as $$
    begin
        if (n = 0)
            then return 1;
        else return n * f_func(n-1);
        end if;
    end;
$$
language plpgsql;

select f_func(6);

-- Хранимая процедура без параметров или с параметрами
-- Повысить цену на указанное число единиц на все паззлы
drop procedure if exists puzzle_up;
create or replace procedure puzzle_up(up numeric)
as $$
    begin
        update games set price = price + up where genre = 'Puzzle';
    end;
$$ language plpgsql;

call puzzle_up(20);

-- Рекурсивная хранимая процедура или хранимая процедура с
-- рекурсивным ОТВ
-- Вывести ммена n первых игр
drop procedure if exists f_proc;
create or replace procedure f_proc(n int, cur int)
as $$
    begin
        if (cur <=  n)
        then raise notice 'Game %: %', cur, (select g.name
                                             from games as g
                                             where g.id = cur);
        call f_proc(n, cur + 1);
        end if;
    end;
$$ language plpgsql;

call f_proc(10, 1);

-- Хранимая процедура с курсором
-- Активировать карты клиентов, посетивших клубы сети в указанный период
drop procedure if exists h;
create or replace procedure h(start_d date, end_d date)
as $$
declare c cursor
        for select * from attendance as a
        where date(a.time_start) between start_d and end_d;
        r record;
begin
    open c;
    loop
        fetch c into r;
        exit when not found;
        update cards
        set state = 'activated' where id_client = r.id_client;
    end loop;
    close c;
end
$$ language plpgsql;

call h('1990-01-01', '1990-01-31');

-- Хранимая процедура доступа к метаданным
-- Вывод названий текстовых аттрибутов таблицы machines
drop procedure if exists get_attr;
create or replace procedure get_attr()
as $$
declare c cursor
        for select * from information_schema.columns
        where table_name = 'machines' and data_type = 'text';
        r record;
begin
    open c;
    loop
        fetch c into r;
        exit when not found;
        raise notice 'machines column: %', r.column_name;
    end loop;
    close c;
end
$$ language plpgsql;

call get_attr();

-- Триггер AFTER
-- При изменении цены на аркады так, что они стоят дороже 2000 единиц,
-- отключать мультиплеерный режим
create or replace function price_check() returns trigger
as $$
begin
    raise notice 'old = %', old;
    raise notice 'new = %', new;

    update games g set multiplayer_mode = false where old.id = g.id;

    raise notice 'no multi for % with id %, old price: %, new price: %', old.name, old.id, old.price, new.price;

    return new;
end
$$ language plpgsql;

drop trigger if exists update_games on games;

create trigger update_games
    after update of price on games
    for each row
    when(new.price > 2000 and new.genre = 'Arcade')
    execute function price_check();

update games g
set price = price + 10
where g.id = 3;

-- Триггер INSTEAD OF
-- При попытке добавить сотрудника бармена с именем Алексей выдавать ошибку
drop view if exists staff_names;
create view staff_names as
select *
from staff;

create or replace function staff_check() returns trigger
as $$
begin
    if (new.position = 'Бармен' and new.name = 'Алексей')
    then raise exception 'no barmens named Alexey!!!!1!!1';
    else
        raise notice 'this Alexey is ok bc he is %', new.position;
        insert into staff values (new.*);
    end if;

    return new;
end
$$ language plpgsql;

drop trigger if exists insert_staff on staff;

create trigger insert_staff
    instead of insert on staff_names
    for each row
    execute function staff_check();

insert into staff_names(id, id_club, name, surname, patronymic, birth_date, sex, phone_number, employment_date, position)
values ((select count(*) + 1 from staff), 1, 'Алексей', 'Алёшин', 'Алексеевич', '1995-02-03', 'male', '+7 964 666-66-16', '2018-01-01', 'Уборщик')
on conflict do nothing;
insert into staff_names(id, id_club, name, surname, patronymic, birth_date, sex, phone_number, employment_date, position)
values ((select count(*) + 1 from staff), 1, 'Алексей', 'Алёшин', 'Алексеевич', '1995-02-03', 'male', '+7 964 666-66-16', '2018-01-01', 'Бармен');

-- Хранимая процедура без курсора
-- Активировать карты клиентов, посетивших клубы сети в указанный период
drop procedure if exists my_proc;
create or replace procedure my_proc(start_d date, end_d date)
as $$
    begin
        update cards set state = 'activated' where id_client in (
            select a.id_client
            from attendance as a
            where date(a.time_start) between start_d and end_d
            );
    end;
$$ language plpgsql;

call h('1990-01-01', '1990-01-31');

call my_proc('1990-01-01', '1990-01-31');