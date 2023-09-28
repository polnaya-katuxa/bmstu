truncate table games_exact_machines cascade;
truncate table machines_clubs cascade;
truncate table cards cascade;
truncate table attendance cascade;
truncate table staff cascade;
truncate table machines cascade;
truncate table clients cascade;
truncate table loyalty_programs cascade;
truncate table computer_clubs cascade;
truncate table games cascade;

copy games FROM '/labs/lab_01/data/games.csv' DELIMITER ',' CSV HEADER;

copy computer_clubs FROM '/labs/lab_01/data/clubs.csv' DELIMITER ',' CSV HEADER;

copy loyalty_programs FROM '/labs/lab_01/data/loyalty_programs.csv' DELIMITER ',' CSV HEADER;

copy clients FROM '/labs/lab_01/data/clients.csv' DELIMITER ',' CSV HEADER;

copy machines FROM '/labs/lab_01/data/machines.csv' DELIMITER ',' CSV HEADER;

copy staff FROM '/labs/lab_01/data/staff.csv' DELIMITER ',' CSV HEADER;

copy attendance FROM '/labs/lab_01/data/attendances.csv' DELIMITER ',' CSV HEADER;

copy cards FROM '/labs/lab_01/data/cards.csv' DELIMITER ',' CSV HEADER;

copy machines_clubs FROM '/labs/lab_01/data/machines_in_clubs.csv' DELIMITER ',' CSV HEADER;

copy games_exact_machines FROM '/labs/lab_01/data/games_on_machines.csv' DELIMITER ',' CSV HEADER;
