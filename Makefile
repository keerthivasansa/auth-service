include .env

grpc/generate:
	protoc --go_out=plugins=grpc:. proto/auth.proto

migrate/create:
	goose -dir ${MIGRATION_PATH} ${DB_DRIVER} ${DB_DSN} create ${name} sql

migrate/up:
	goose -dir ${MIGRATION_PATH} ${DB_DRIVER} ${DB_DSN} up

migrate/down:
	goose -dir ${MIGRATION_PATH} ${DB_DRIVER} ${DB_DSN} down

migrate/status:
	goose -dir ${MIGRATION_PATH} ${DB_DRIVER} ${DB_DSN} status