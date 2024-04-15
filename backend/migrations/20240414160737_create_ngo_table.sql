-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Ngo (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(128) NOT NULL,
    About VARCHAR(4096) NOT NULL,
    Latitude NUMERIC NOT NULL,
    Longitude NUMERIC NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Ngo;
-- +goose StatementEnd
