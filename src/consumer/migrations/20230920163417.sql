-- Create "transaction_dbs" table
CREATE TABLE "public"."transaction_dbs" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "transaction_type" text NULL,
  "location" text NULL,
  "transaction_id" bigint NULL,
  "account_number" bigint NULL,
  "amount" bigint NULL,
  "time" timestamptz NULL,
  "accepted" boolean NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_transaction_dbs_deleted_at" to table: "transaction_dbs"
CREATE INDEX "idx_transaction_dbs_deleted_at" ON "public"."transaction_dbs" ("deleted_at");
