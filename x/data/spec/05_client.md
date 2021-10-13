# Client

The `tx` (transaction) and `q` (query) subcommands below give a good illustration
of how this module can be used in practice:

### Data Transactions

```sh 
$ regen tx data --help
# Data transaction subcommands
# 
# Usage:
#   regen tx data [flags]
#   regen tx data [command]
# 
# Available Commands:
#   anchor      Anchors a piece of data to the blockchain based on its secure
#                 hash, effectively providing a tamper resistant timestamp.
#   sign        Sign an arbitrary piece of data on the blockchain.
#   store       Store a piece of data corresponding to a CID on the blockchain.
```

### Data Queries

```sh 
$ regen q data --help
# Querying commands for the data module.
# If a CID is passed as first argument, then this command will query
# timestamp, signers and content (if available) for the given CID. Otherwise,
# this command will run the given subcommand.
# 
# Example (the two following commands are equivalent):
# $ regen query data bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm
# $ regen query data by-cid bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm
# 
# Usage:
#   regen query data [cid] [flags]
#   regen query data [command]
# 
# Available Commands:
#   by-cid      Query for CID timestamp, signers and content (if available)
```
