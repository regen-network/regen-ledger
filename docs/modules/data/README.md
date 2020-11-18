# Data Module Overview

The data module is intended to be a versatile module for anchoring, signing, 
and storing data about ecological health, and ecosystem services on Regen Ledger.

These three functionalities each provide different kinds of information and guarantees
about the data being referenced.

- __Data Anchoring__: Proving that a piece of data was known to exist at a certain point
  in time. This can also be referred to as "secure timestamping".
- __Data Signing__: Asserting to the veracity and validity of a piece of data. Signing
  implies that the contents of the data are generally accepted to be true by the signer.
- __Data Storing__: Storing the raw data itself on the blockchain. This is useful when 
  availability guarantees are necessary. This can also be useful in cases where one
  wants smart contracts to have direct access to the data itself.
  
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

For a guided walk through of how to use some of this module's functionality, check out
the [API and CLI docs](../../api.md), which takes you through the process of setting up
a key pair, getting your node up and running, and anchoring your first CID with the data
module!