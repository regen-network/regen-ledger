# Tendermint Postgres Indexer

Tendermint provides an option for node operators to [index transactions and events](https://docs.tendermint.com/v0.34/app-dev/indexing-transactions.html). Transactions and events can either be indexed in a `kv` indexer or `psql` indexer.

This document provides instructions for setting up a `psql` indexer.

## Dependencies

- Docker `>=20.10`
- Postgres `>=v14.5`
- Tendermint `>=v0.34.15`

### Repository Setup

The first step is to create a new repository that will be the home of the indexer.

```sh
mkdir indexer
```

```sh
cd indexer
```

```sh
git init
```

### Schema Setup

The next step will be creating the schema for the database. We will be using the tendermint schema available [here](https://github.com/tendermint/tendermint/blob/v0.34.15/state/indexer/sink/psql/schema.sql).

```sh
curl https://raw.githubusercontent.com/tendermint/tendermint/v0.34.15/state/indexer/sink/psql/schema.sql > schema.sql
```

The above schema creates three `VIEWS`:

- `event_attributes`
- `block_events`
- `tx_events`

### Docker Setup

The next step will be creating a `Dockerfile` in the root of the repository with the following content:

```dockerfile
FROM postgres:14.5-alpine

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB regen_events
COPY schema.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

CMD ["postgres"]
```

### Postgres Setup

Now that we have the schema and dockerfile setup, the next step is building the docker image and starting the postgres server.

Build the docker image:

```sh
docker build -t "regen-events-indexer" .
```

Start the postgres server:

```sh
docker run -it -d -p 5432:5432 -e -v ./pg-volume:/var/lib/postgresql/data regen-events-indexer:latest
```

### Regen Setup

The next step is to test the indexer using a local test network.

The following documentation will help you get started:

- [Install Regen](https://docs.regen.network/ledger/get-started/#building-from-source)
- [Local Testnet](https://docs.regen.network/ledger/get-started/local-testnet.html#create-accounts)

## Postgres indexer configuration

To enable the indexer, go to `[tx_index]` section in `config.toml` file. You can find the `config.toml` file in `~/.regen/config`.

Set `indexer` to `psql`:

```toml
indexer = "psql"
```

Set `psql-conn` to the location of the postgres database:

```toml
psql-conn = "postgresql://postgres:postgres@127.0.0.1:5432/regen_events?sslmode=disable"
```

Start (or restart) the local testnet:

```sh
regen start
```

Perform some transactions on the local testnet.

### Performing Queries

:::warning
Searching is not enabled for the `psql` indexer type via Tendermint's RPC and therefore the following queries will fail:

- `TxByEvents` RPC call
- `TxByHash` RPC call
- `tx-by-events` REST endpoint
- `tx-by-hash` REST endpoint
:::

Log in to postgres container:

```sh
docker exec -tiu postgres <container-id>  psql
```

Change database to `regen_events`:

```
\c regen_events
```

List all database tables:

```
\dt
```

The above command should return the following output:

```
           List of relations
 Schema |    Name    | Type  |  Owner   
--------+------------+-------+----------
 public | attributes | table | postgres
 public | blocks     | table | postgres
 public | events     | table | postgres
 public | tx_results | table | postgres
(4 rows)
```

The following query will get all events from the `ecocredit` module:

```sql
SELECT *
FROM event_attributes
WHERE type LIKE 'regen.ecocredit.%';
```
    
The following query will get all events from the `data` module:

```sql
SELECT *
FROM event_attributes
WHERE type LIKE 'regen.data.%';
```

The following query will get all block events:

```sql
SELECT blocks.rowid as block_id, height, chain_id, type, key, composite_key, value
FROM blocks JOIN event_attributes ON (blocks.rowid = event_attributes.block_id)
WHERE event_attributes.tx_id IS NULL; 
```

The following query will get all events from modules using legacy events (e.g. the `bank` module):
    
```sql
SELECT DISTINCT tx_id
FROM attributes JOIN events ON events.rowid=attributes.event_id
WHERE attributes.value LIKE '%bank%';
```
