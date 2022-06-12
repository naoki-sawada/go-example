-- +goose Up
-- +goose StatementBegin
CREATE TABLE "organization_user" (
	"organization_id" uuid NOT NULL REFERENCES "organizations" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
	"user_id" uuid NOT NULL REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
	PRIMARY KEY ("organization_id", "user_id"),
	UNIQUE ("user_id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "organization_user";
-- +goose StatementEnd
