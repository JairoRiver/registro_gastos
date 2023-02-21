CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "Users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Groups" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_group" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid,
  "group_id" uuid,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Entry" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid,
  "group_id" uuid,
  "type_id" uuid,
  "name" varchar NOT NULL,
  "use_day" date NOT NULL,
  "amount" float NOT NULL,
  "cost" float NOT NULL,
  "cost_indicator" varchar(1),
  "place" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Type" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Users" ("id");

CREATE INDEX ON "Users" ("email");

CREATE INDEX ON "Groups" ("id");

CREATE INDEX ON "user_group" ("id");

CREATE UNIQUE INDEX ON "user_group" ("user_id", "group_id");

CREATE INDEX ON "Entry" ("id");

CREATE INDEX ON "Type" ("id");

ALTER TABLE "Groups" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "user_group" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "user_group" ADD FOREIGN KEY ("group_id") REFERENCES "Groups" ("id");

ALTER TABLE "Entry" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Entry" ADD FOREIGN KEY ("group_id") REFERENCES "Groups" ("id");

ALTER TABLE "Entry" ADD FOREIGN KEY ("type_id") REFERENCES "Type" ("id");
