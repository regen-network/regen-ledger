# Ecocredit Module Overview

::: tip COMING SOON

Expect updates on this page soon...

In the meantime, make sure you're familiar with the basics of the
[Regen Ledger API & CLI](../../api.md), and then let `regen tx ecocredit --help`
and the Ecocredit Module [Protobuf Documentation](./protobuf.md)
be your guide!

:::

### Ecocredit Transactions

```sh
$ regen tx ecocredit --help
# Ecocredit module transactions
#
# Usage:
#   regen tx ecocredit [flags]
#   regen tx ecocredit [command]
#
# Available Commands:
#   cancel        Cancels a specified amount of credits from the account of the transaction author (--from)
#   create_batch  Issues a new credit batch
#   create_class  Creates a new credit class
#   retire        Retires a specified amount of credits from the account of the transaction author (--from)
#   send          Sends credits from the transaction author (--from) to the recipient
#   set_precision Allows an issuer to increase the decimal precision of a credit batch
```
### Ecocredit Queries

```sh
$ regen q ecocredit --help
# Query commands for the ecocredit module
#
# Usage:
#   regen query ecocredit [flags]
#   regen query ecocredit [command]
#
# Available Commands:
#   balance     Retrieve the tradable and retired balances of the credit batch
#   batch_info  Retrieve the credit issuance batch info
#   class_info  Retrieve credit class info
#   precision   Retrieve the maximum length of the fractional part of credits in the given batch
# supply      Retrieve the tradable and retired supply of the credit batch
```
