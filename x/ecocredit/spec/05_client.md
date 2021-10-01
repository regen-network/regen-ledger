# Client

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
#   cancel         Cancels a specified amount of credits from the account of the transaction author (--from)
#   create-batch   Issues a new credit batch
#   create-class   Creates a new credit class
#   gen-batch-json Generates JSON to represent a new credit batch for use with create-batch command
#   retire         Retires a specified amount of credits from the account of the transaction author (--from)
#   send           Sends credits from the transaction author (--from) to the recipient
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
#   batch-info  Retrieve the credit issuance batch info
#   batches     List all credit batches in the given class with pagination flags
#   class-info  Retrieve credit class info
#   classes     List all credit classes with pagination flags
#   supply      Retrieve the tradable and retired supply of the credit batch
#   types       Retrieve the list of credit types
```
