-- +goose Up
-- +goose StatementBegin
CREATE TABLE "organizations" (
	"id" uuid NOT NULL DEFAULT gen_random_uuid(),
	"name" character varying (100) NOT NULL,
	PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "organizations";
-- +goose StatementEnd
