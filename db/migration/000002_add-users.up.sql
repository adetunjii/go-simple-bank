CREATE TABLE "users"
(
    "username"        varchar PRIMARY KEY,
    "password"        varchar NOT NULL,
    "full_name"       varchar NOT NULL,
    "email"           varchar UNIQUE NOT NULL,
    "created_at"      timestamptz DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");


-- Creates a unique composite index on the owner and currency. This means that a user cannot have an account with the same currency twice --
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");