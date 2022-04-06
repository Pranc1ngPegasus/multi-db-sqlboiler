
-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"(
  "id"         BIGSERIAL,
  "name"       TEXT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "users";
