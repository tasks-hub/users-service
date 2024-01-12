-- +goose Up
-- +goose StatementBegin
CREATE DATABASE users_db;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE users_db;
-- +goose StatementEnd
