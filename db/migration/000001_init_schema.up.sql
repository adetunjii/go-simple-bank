
CREATE TABLE "accounts"
(
    "id"         bigserial PRIMARY KEY,
    "owner"      varchar     NOT NULL,
    "balance"    bigint      NOT NULL,
    "currency"   varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions"
(
    "id"         bigserial PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "amount"     bigint NOT NULL,
    "currency"          varchar NOT NULL,
    "transaction_type"  varchar NOT NULL,
    "status"            varchar NOT NULL,
    "reference"         varchar NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transfers"
(
    "id"              bigserial PRIMARY KEY,
    "source_account_id" bigint,
    "destination_account_id"   bigint,
    "amount"          bigint NOT NULL,
    "currency"        varchar NOT NULL,
    "reference"       varchar NOT NULL,
    "created_at"      timestamptz DEFAULT (now())
);


ALTER TABLE "transactions"
    ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("source_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
    ADD FOREIGN KEY ("destination_account_id") REFERENCES "accounts" ("id");


CREATE
INDEX ON "accounts" ("owner");

CREATE
INDEX ON "transactions" ("account_id");

CREATE
INDEX ON "transfers" ("source_account_id");

CREATE
INDEX ON "transfers" ("destination_account_id");

CREATE
INDEX ON "transfers" ("source_account_id", "destination_account_id");
