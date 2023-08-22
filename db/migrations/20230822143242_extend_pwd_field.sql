-- +goose Up
-- +goose StatementBegin
ALTER TABLE users MODIFY COLUMN password varchar(100) not null; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users MODIFY COLUMN password varchar(64) not null; 
-- +goose StatementEnd
