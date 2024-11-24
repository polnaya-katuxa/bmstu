-- +goose Up
-- +goose StatementBegin
insert into limits (value, bonus) values
    (5, 10),
    (10, 20),
    (20, 30),
    (50, 40),
    (100, 50);

update posts set limit_id = (
    select l.uuid
    from limits as l
    where l.value = (select min(l2.value)
                     from limits as l2)
    );

insert into reaction_types (icon) values
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/red-heart.png'),
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/fire.png'),
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/disguised-face.png'),
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/broken-heart.png'),
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/poop.png'),
      ('https://2292ce37-f513e8af-f963-4e8e-a185-544861427a71.s3.timeweb.com/postby/sneezing-face.png');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table reaction_types;
delete from limits where value != 2147483647;
-- +goose StatementEnd
