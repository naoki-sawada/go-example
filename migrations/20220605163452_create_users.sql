-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"first_name" character varying (30) NOT NULL,
	"last_name" character varying (30) NOT NULL,
	"email" character varying (256) NOT NULL,
	"birthdate" date NOT NULL,
  PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
