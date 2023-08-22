-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN provider varchar(25) not null,
                  ADD COLUMN providerId varchar(64) not null,
                  DROP COLUMN email;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN provider,
                  DROP COLUMN providerId,
                  ADD COLUMN email varchar(64) not null;
-- +goose StatementEnd
