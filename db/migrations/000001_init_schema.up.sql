CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "date_created" timestamptz NOT NULL DEFAULT (now()),
  "date_modified" timestamptz NOT NULL DEFAULT (now()),
  "account_balance" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "payment_id" bigint,
  "payment_type" varchar,
  "amount" bigint,
  "from_user" bigint NOT NULL,
  "to_user" bigint NOT NULL
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "customer_phone" varchar NOT NULL,
  "amount" bigint,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "mpesa_reference" varchar NOT NULL,
  "failure_reason" varchar,
  "business_id" bigint NOT NULL
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "transactions" ("from_user");

CREATE INDEX ON "transactions" ("to_user");

CREATE INDEX ON "payments" ("business_id");

CREATE INDEX ON "payments" ("status");

CREATE INDEX ON "payments" ("customer_phone");

ALTER TABLE "transactions" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user") REFERENCES "users" ("id");