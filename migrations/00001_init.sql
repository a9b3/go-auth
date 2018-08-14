-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "user" (
  id int NOT NULL,
  email text,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "user";
