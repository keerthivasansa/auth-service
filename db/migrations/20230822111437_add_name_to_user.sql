-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN name varchar(64) not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN name;
-- +goose StatementEnd
