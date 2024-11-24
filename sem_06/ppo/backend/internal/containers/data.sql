insert into users (uuid, login, picture, description, password, mail, balance, enter_time, is_admin) values
('a52b8aea-d751-4933-91bb-691132e3b760', '1', 'picture', 'description', 'password', '1@mail.ru', 100, current_timestamp, false),
('3f010aca-5008-4aa5-a1a3-a061a876783f', '2', 'picture', 'description', 'password', '2@mail.ru', 0, current_timestamp, false),
('77dcd288-79b2-4655-9584-cc9b5329665d', '3', 'picture', 'description', 'password', '3@mail.ru', 0, current_timestamp, false),
('b8b87f28-3cbb-425c-ac4d-52015710d61b', '4', 'picture', 'description', 'password', '4@mail.ru', 0, current_timestamp, false);

insert into subscriptions (reader_id, writer_id) values
('a52b8aea-d751-4933-91bb-691132e3b760', '3f010aca-5008-4aa5-a1a3-a061a876783f'), -- 1 -> 2
('77dcd288-79b2-4655-9584-cc9b5329665d', '3f010aca-5008-4aa5-a1a3-a061a876783f'), -- 3 -> 2
('77dcd288-79b2-4655-9584-cc9b5329665d', 'a52b8aea-d751-4933-91bb-691132e3b760'), -- 3 -> 1
('b8b87f28-3cbb-425c-ac4d-52015710d61b', 'a52b8aea-d751-4933-91bb-691132e3b760'), -- 4 -> 1
('b8b87f28-3cbb-425c-ac4d-52015710d61b', '3f010aca-5008-4aa5-a1a3-a061a876783f'), -- 4 -> 2
('b8b87f28-3cbb-425c-ac4d-52015710d61b', '77dcd288-79b2-4655-9584-cc9b5329665d'); -- 4 -> 3

delete from limits;

insert into limits (uuid, value, bonus) values
('4603afd3-73d2-438a-8e74-7410c18469a8', 1, 100),
('9dc9900a-93e0-4886-a654-09c88cb02f40', 3, 200),
('6cd1dcbc-02db-4980-8fba-eb56bb229fbe', 5, 300);

insert into posts (uuid, content, perms, writer_id, public_time, limit_id) values
('1557a6be-1008-412a-88d9-1f06630d028c', '1', 'paid', '3f010aca-5008-4aa5-a1a3-a061a876783f', current_timestamp, '4603afd3-73d2-438a-8e74-7410c18469a8'),
('21bf7ace-965b-4679-86b8-93a89cba0094', '2', 'free', '77dcd288-79b2-4655-9584-cc9b5329665d', current_timestamp, '4603afd3-73d2-438a-8e74-7410c18469a8'),
('eea07263-3444-40d0-adc4-345ad7728298', '3', 'free', '77dcd288-79b2-4655-9584-cc9b5329665d', current_timestamp, '4603afd3-73d2-438a-8e74-7410c18469a8'),
('c11add9b-207e-4b41-a964-2662ed3cae27', '4', 'paid', '77dcd288-79b2-4655-9584-cc9b5329665d', current_timestamp, '4603afd3-73d2-438a-8e74-7410c18469a8'),
('7e1053d5-f31f-4841-b3f9-d0e47849cfb3', '5', 'free', '77dcd288-79b2-4655-9584-cc9b5329665d', current_timestamp - '3 hours'::interval, '4603afd3-73d2-438a-8e74-7410c18469a8');

insert into reaction_types (uuid, icon) values
    ('bde563b1-66a6-4d00-ac3f-4022be793c81', 'dfkjkejrhf');

insert into reactions (uuid, reaction_type_id, post_id, reactor_id) values
('9a9a818f-f8c8-4e24-8648-60c84c4fdeaa', 'bde563b1-66a6-4d00-ac3f-4022be793c81', 'eea07263-3444-40d0-adc4-345ad7728298', 'b8b87f28-3cbb-425c-ac4d-52015710d61b'), -- 4 -> 3
('6beed1dc-61ac-4439-b927-c7e91c3a3510', 'bde563b1-66a6-4d00-ac3f-4022be793c81', 'eea07263-3444-40d0-adc4-345ad7728298', '3f010aca-5008-4aa5-a1a3-a061a876783f'), -- 2 -> 3
('a15c02f3-0fb8-4380-896e-b835f9542668', 'bde563b1-66a6-4d00-ac3f-4022be793c81', '1557a6be-1008-412a-88d9-1f06630d028c', 'a52b8aea-d751-4933-91bb-691132e3b760'); -- 1 -> 2

