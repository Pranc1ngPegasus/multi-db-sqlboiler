
-- +migrate Up
CREATE TABLE IF NOT EXISTS "accounts"(
  "id" BIGSERIAL,
  "user_id" BIGINT NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "deleted_at" TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS idx_user_id_on_accounts ON "accounts" (user_id);

-- +migrate Down
DROP TABLE IF EXISTS "accounts";
