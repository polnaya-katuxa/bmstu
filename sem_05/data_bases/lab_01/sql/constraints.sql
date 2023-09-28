drop function if exists check_dates;
drop function if exists check_time;
drop function if exists check_registration_date;

create function check_dates(d date, id_club int) returns boolean
as $$
(select establishment_date <= d from computer_clubs where id_club = computer_clubs.id);
$$ language sql;

create function check_time(t1 timestamp without time zone, t2 timestamp without time zone, id_club int) returns boolean
as $$
(select ((is_round_the_clock) or ((t1::time between open_time and close_time) and (t2::time between open_time and close_time) and (t1::date = t2::date))) from computer_clubs where id_club = computer_clubs.id);
$$ language sql;

create function check_registration_date(d date, id_client int) returns boolean
as $$
(select d > birth_date from clients where id_client = clients.id);
$$ language sql;

alter table if exists games
    add constraint valid_year check(release_year > 1960 and release_year <= date_part('year', CURRENT_DATE)),
    add constraint  valid_price check (price > 0);

alter table if exists computer_clubs
    add constraint valid_time check(open_time <= close_time),
    add constraint valid_date check(establishment_date >= '2002-09-23' and establishment_date <= CURRENT_DATE),
    add constraint valid_parking_spot_num check(parking_spot_num >= 0);

alter table if exists loyalty_programs
    add constraint valid_cashback_percent check(cashback_percent >= 0 and cashback_percent < 100),
    add constraint valid_minimum_purchase_sum check(minimum_purchase_sum >= 0);

alter table if exists clients
    add constraint  valid_birth_date check (birth_date < current_date),
    add constraint  valid_phone_number check (phone_number like '+7 ___ ___-__-__');

alter table if exists machines
    add constraint  valid_r check (release_year > 1960 and release_year <= date_part('year', CURRENT_DATE));

alter table if exists staff
    add constraint  valid_birth_date check (birth_date < current_date),
    add constraint  valid_employment_date check (employment_date <= current_date and employment_date > birth_date and check_dates(employment_date, id_club)),
    add constraint  valid_phone_number check (phone_number like '+7 ___ ___-__-__');

alter table if exists attendance
    add constraint valid_rating check (rating between 1 and 5),
    add constraint valid_date check (time_start <= current_timestamp and time_end <= current_timestamp),
    add constraint  valid_time check (time_start < time_end and check_time(time_start, time_end, id_club));

alter table if exists cards
    add constraint  valid_number check (number like '____ ____ ____ ____'),
    add constraint valid_registration_date check (registration_date <= current_date and check_registration_date(registration_date, id_client));
