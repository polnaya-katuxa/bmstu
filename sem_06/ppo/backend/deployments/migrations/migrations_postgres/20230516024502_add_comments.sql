-- +goose Up
-- +goose StatementBegin
create table comments (
                           uuid UUID default uuid_generate_v4() primary key,
                           public_time timestamp not null,
                           content text not null,
                           post_id UUID references posts (uuid) on delete cascade,
                           commentator_id UUID references users (uuid) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists comments;
-- +goose StatementEnd
