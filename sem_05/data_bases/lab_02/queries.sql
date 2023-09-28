-- 1) Использовать предикат сравнения
-- Получить список адресов клубов + кол-во парк.мест, где больше 15 парковочных мест
select c.address, c.parking_spot_num
from computer_clubs as c
where c.parking_spot_num > 15;

-- 2) Использовать предикат between
-- Получить список ФИО + дата трудоустройства сотрудников, устроившихся на работу весной 2020 года
select s.surname, s.name, s.patronymic, s.employment_date
from staff as s
where s.employment_date between '2020-03-01' and '2020-05-31';

-- 3) Использовать предикат like
-- Получить список названий игр, имя которых содержит слово "War"
select g.name
from games as g
where g.name like '%War%';

-- 4) Использовать предикат in с вложенным подзапросом
-- Получить список рейтингов посещений с id клиента, совершённых Олегами
select a.rating, a.id_client
from attendance as a
where id_client in (select c.id
                    from clients as c
                    where c.name = 'Олег');

-- 5) Использовать предикат exists с вложенным подзапросом
-- Получить список номеров неактивированных карт клиентов-женщин
select c.number, c.state, c.id_client
from cards as c
where c.state = 'non-activated' and exists(select cl.id
                                           from clients as cl
                                           where c.id_client = cl.id and cl.sex = 'female');

-- 6) Использовать предикат сравнения с квантором
-- Получить список названий игр из Японии + наличие мультиплеерного режима и цена, у которых есть мультиплеерный режим
-- и цена выше, чем у всех Американских мультиплеерных игр с рейтингом по возрасту 3+
-- там самая дорогая стоит 4972.46
select g.name, g.multiplayer_mode, g.price
from games as g
where g.multiplayer_mode = true
    and g.country = 'Japan'
    and g.price > all(select gam.price
                      from games as gam
                      where gam.country = 'United States' and gam.age_rating = 'E' and gam.multiplayer_mode = true);

-- 7) Использовать агрегатные функции в выражениях столбцов
-- Получить максимальную и минимальную цену среди всех клиентов за посещения
select max(total_price) as max_price, min(total_price) as min_price
from (select sum(a.price) as total_price
      from attendance as a
      group by a.id_client) as TotalAttSum;

-- 8) Использовать скалярные подзапросы в выражениях столбцов
-- Получить список id, логинов, средних цен посещения, дат выпуска последней карты клиентов-мужчин
select c.id, c.login,
       (select avg(a.price)
        from attendance as a
        where a.id_client = c.id) as avg_att_price,
       (select max(card.registration_date)
        from cards as card
        where card.id_client = c.id) as latest_card
from clients as c
where c.sex = 'male';

-- 9) Использовать простые выражения case
-- Получить список марок и моделей компьютеров/приставок, где техника марки Apple будет помечена,
-- как хорошая, ПК других марок, как плохие, а приставки как не ПК
select m.brand, m.model,
       case m.brand
            when 'Apple' then 'Nice PC'
            when 'XBox' then 'Not PC'
            when 'PlayStation' then 'Not PC'
            else 'Bad PC'
       end as quality
from machines as m;

-- 10) Использовать поисковые выражения case
-- Получить список названий и компаний-производителей игр, охарактеризовав их цену
select g.name, g.company,
       case
           when g.price < 100 then 'Very cheap'
           when g.price < 1000 then 'Fair'
           when g.price < 2000 then 'Expensive'
           else 'Never buy it'
       end as price
from games as g;

-- 11) Создать новую временную локальную таблицу из результирующего набора данных от select
-- Получить таблицу жанров американских игр с суммарной и средней ценой всех игр этого жанра
drop table if exists price_info;

select g.genre, sum(g.price) as total_price, avg(g.price) as avg_price
into temporary price_info
from games as g
where g.country = 'United States'
group by g.genre;

select * from price_info;

-- 12) Использовать вложенный коррелированный подзапрос в качестве производной таблицы в from
-- Получить список максимальных и минимальных цен за посещения с рейтингом > 3 для всех клубов, где
-- минимальная цена посещения будет > 1000
select id, max_price, min_price
from (select a.id_club as id, max(a.price) as max_price, min(a.price) as min_price
      from attendance as a
      where a.rating > 3
      group by id) as club_prices
where min_price > 1000;

-- Получить цену посещения и id клиента, где цена больше средней цены посещений для данного клиента
select a.price, a.id_client
from attendance as a
where a.price > (select avg(price)
                 from (select a2.price
                       from attendance as a2
                       where a.id_client = a2.id_client) as a3);

-- 13) Использовать вложенные подзапросы с уровнем вложенности 3
-- Получить список названий и цен игр, установленных на компы круглосуточных клубов Ярославля
select g.name, g.price
from games as g
where g.id in (
select gm.id_game
from games_exact_machines as gm
where gm.id_machines_clubs in (
               select mc.id
               from machines_clubs as mc
               where id_club in (select c.id
                                 from computer_clubs as c
                                 where c.is_round_the_clock = true and c.address
                                 like '%Ярославль%')
               )
);

-- 14) Консолидировать данные group by, но без having
-- Получить список дат рождения самых старых уборщиков клубов где количество парк.мест > 10
select min(s.birth_date)
from staff as s
where s.position = 'Уборщик' and s.id_club in (select c.id
                                               from computer_clubs as c
                                               where c.parking_spot_num > 10)
group by s.id_club;

-- 15) Консолидировать данные group by c having
-- Получить список брендов и количества моделей машин, количество моделей у которых больше количества моделей Apple
select m.brand, count(m.model)
from machines as m
where m.brand != 'Apple'
group by m.brand
having count(m.model) > (select count(model)
                         from machines
                         where brand = 'Apple');

-- 16) Сделать insert
insert into loyalty_programs (id, name, design, cashback_percent, minimum_purchase_sum)
values (7, 'Brilliant', 'computers.club/images/brilliant.png', 30, 100000.00)
on conflict do nothing;

-- 17) Сделать insert многострочный, вставляющий в таблицу по вложенному подзапросу
insert into games (id, name, genre, release_year, company, country, age_rating, price, multiplayer_mode)
select (select count(*) + 1
        from games),
       'The Witcher 2077', 'Puzzle',
       (select min(g.release_year)
        from games as g
        where g.genre = 'Racing'),
        'Palmolive', country, age_rating, price, multiplayer_mode
from games
where name like '%Race%'
on conflict do nothing;

-- 18) Сделать update
-- Снизить цены на игры дороже 3000 на 300 единиц
update games
set price = price - 300
where price > 3000.00;

-- 19) Сделать update со скалярным подзапросом в set
-- Установить цену на игры дешевле 50 ед. как среднюю цену на игры жанра спорт
update games
set price = (select avg(price)
             from games
             where genre = 'Sport')
where price < 60.00;

-- 20) Сделать delete
delete from loyalty_programs
where name = 'Brilliant';

-- 21) Сделать delete c вложенным коррелированным подзапросом в where
-- Удалить карты клиента Игоря, который меньше всех в сумме потратил
delete from cards
where id_client in (select id
                    from (clients join (select id_client, sum(price) as sum
                                       from attendance
                                       group by id_client) as sum_price
                    on clients.id = sum_price.id_client) as joined
                    where name = 'Игорь'
                    order by sum
                    limit 1);

-- 22) Использовать простое обобщённое табличное выражение
-- Найти максимальную сумму трат клиента
with client_sums (id, sum) as (
    select id_client, sum(price) as sum
    from attendance
    group by id_client
)
select max(sum)
from client_sums;

-- 23) Использовать рекурсивное обобщённое табличное выражение
drop table if exists managers;
create table managers (
    id int primary key,
    name text not null,
    surname text not null,
    boss_id int references managers(id),
    club_id int references computer_clubs(id)
);

insert into managers (id, name, surname, boss_id, club_id)
values (1, 'Vasya', 'Vasilyev', NULL, NULL), (2, 'Petya', 'Petrov', 1, NULL), (3, 'Misha', 'Michailov', 1, NULL),
       (4, 'Vanya', 'Ivanov', 3, 3), (5, 'Fedya', 'Fedorov', 3, 6);

with recursive managers_tree(id, name, club_id, level) as (
    select id, name, club_id, 0 as level
    from managers
    where boss_id is NULL
    union all
    select m.id, m.name, m.club_id, t.level + 1
    from managers as m join managers_tree as t on m.boss_id = t.id
)

select *
from managers_tree;

-- 24) Оконные функции + min/max/avg over()
-- Вывести клуб, клиента, максимальную/минимальную/среднюю цену за посещение для каждого клиента
select a.id_club,
       a.id_client,
       min(a.price) over(partition by a.id_client),
       max(a.price) over(partition by a.id_client),
       avg(a.price) over(partition by a.id_client)
from attendance as a;

-- Вывести id клуба и id клиента, цену за посещение и пронумеровать записи внутри партиций по клиента
-- в порядке возрастания id клуба и в порядке возрастания цены за посещение
select a.id_club,
       a.id_client,
       a.price,
       row_number() over(partition by a.id_client order by a.id_club),
       row_number() over(partition by a.id_client order by a.price)
from attendance as a;

-- 25) Оконная функция для устранения дублей
-- Получить список уникальных марок компьютеров
select brand
from (select m.brand, row_number() over (partition by m.brand) as rn
from machines as m) as brands
where rn = 1;

--26) игра, установленнная на макс количество машин
with g(name, count) as (
    select t2.name as name, t2.c as count
    from (games as g join (
        select id_game, count(id_machines_clubs) as c
        from games_exact_machines
        group by id_game) as t1 on g.id = t1.id_game) as t2
)

select name, count
from g
where count = (select max(g1.count)
               from g as g1);



