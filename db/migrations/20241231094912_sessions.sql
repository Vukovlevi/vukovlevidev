-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions (
    session_token VARCHAR(32),
    userId INTEGER,
    expires DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM sessions;
DROP TABLE sessions;
-- +goose StatementEnd
