-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (id varchar(64) not null primary key, email varchar(120) not null, password varchar(64) not null);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
