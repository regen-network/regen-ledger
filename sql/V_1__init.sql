CREATE TABLE acct (
  id text not null primary key,
);

CREATE TABLE asset (
  id text not null primary key,
  -- issuer text not null references acct(id), ??
  code text,
  supply numeric not null,
);

CREATE TABLE balance (
  acct text not null references acct(id),
  asset text not null references asset(id),
  balance numeric not null default(0)
  primary key(owner_id, asset_id)
);

CREATE FUNCTION transfer(to text, asset text, amount bigint) RETURNS bool AS $$
  $$ language plpgsql volatile;

CREATE TABLE hash_timestamp (
  acct text not null,
  hash text not null,
  block_id bigint not null,
  url text,
  data jsonb -- contains some partial data of the externally referenced data
  primary key (acct, hash)
);

CREATE TABLE esp (
  curator text not null references acct(id),
  name text not null,
  primary key(curator, name)
);

CREATE TYPE esp_schema_type AS ENUM ('json-schema', 'shacl', 'shex');

CREATE TABLE esp_version (
  esp text not null, -- references esp(curator, name)
  version text not null,
  schema jsonb not null,
  schema_type esp_schema_type not null
);

CREATE TABLE esp_report (
  id text not null primary key,
  issuer text not null references acct(id),
  geo geometry('POLYGON') NOT NULL,
  hash text not null,
  url text,
  esp text references esp_version(id),
  data jsonb
);

COMMENT ON COLUMN esp_report(esp) IS 'the name of the ESP used in this report if that can be publicly shared';

COMMENT ON COLUMN esp_report(data) IS 'contains parts of the ESP report which can be publicly shared, and the full report if possible';

CREATE TABLE cls (
  owner text not null,
  name text not null,
  primary key(owner, name)
);

CREATE TYPE cardinality_t AS ENUM ('one', 'many');

CREATE TABLE prop (
  owner text not null,
  name text not null,
  primary key(owner, name),
  type text not null references types(name),
  numeric_min numeric,
  numeric_max numeric,
  cardinality cardinality_t not null
);

CREATE TABLE types (
  name text not null primary key
);

CREATE TABLE offer (
  id uuid not null primary key,
  offerer text not null,
  from_asset text not null,
  to_asset text not null,
  from_amount numeric not null,
  to_amount numeric not null,
  expires timestamptz default null
);

CREATE FUNCTION offer_create(from_asset text, to_asset text, amount_from bigint, amount_to bigint);


