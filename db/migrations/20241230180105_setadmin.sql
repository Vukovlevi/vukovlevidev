-- +goose Up
-- +goose StatementBegin
UPDATE users SET role = "admin" WHERE username = "vukovlevi";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
