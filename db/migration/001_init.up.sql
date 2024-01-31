CREATE TABLE "users" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "balance" bigint NOT NULL
);

CREATE TABLE "transactions" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "from_user_id" UUID NOT NULL,
  "to_user_id" UUID NOT NULL,
  "date" timestamp DEFAULT (current_timestamp),
  "amount" bigint NOT NULL,
  "description" varchar
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");
