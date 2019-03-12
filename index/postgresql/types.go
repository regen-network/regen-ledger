// Package postgresql implements an Indexer for PostgreSQL that does
// default indexing of blocks and transactions and can be used in
// keepers for custom indexing
package postgresql

import "github.com/regen-network/regen-ledger/index"

// Indexer is a PostgreSQL based Indexer
type Indexer interface {
	index.Indexer

	// AddMigration adds a migration to the list of migrations to be run during
	// InitChain or an upgrade. Migrations will be run in the order
	// they are added
	AddMigration(ddl string)

	// Exec executes a PostgreSQL statement against the database.
	// Can be used in keepers for custom indexing
	Exec(query string, args ...interface{})
}

// InitialSchema is that initial PostgreSQL schema of
const InitialSchema = `
CREATE TABLE block (
  height BIGINT NOT NULL PRIMARY KEY,
  time timestamptz NOT NULL,
  hash bytea
);

CREATE TABLE tx (
  hash bytea NOT NULL PRIMARY KEY,
  block BIGINT NOT NULL REFERENCES block,
  bytes bytea NOT NULL,
  tx_json jsonb,
  code int,
  result bytea
);

CREATE TABLE block_tags (
  block bigint NOT NULL REFERENCES block,
  key text not null,
  value text not null
);

CREATE TABLE tx_tags (
  tx bytea NOT NULL REFERENCES tx,
  key text not null,
  value text not null
);
`
