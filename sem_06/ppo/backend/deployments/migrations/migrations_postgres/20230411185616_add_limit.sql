-- +goose Up
-- +goose StatementBegin
insert into limits (value, bonus) values
(2147483647, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from limits where value = 2147483647;
-- +goose StatementEnd
