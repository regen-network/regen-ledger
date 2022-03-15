# Msg Service

## MsgCreateClass

`MsgCreateClass` creates a new credit class with a credit class admin, an approved list of issuers, optional metadata, and a credit type. The party signing this transaction is the credit admin. 

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L92-L105

### Validation:

- `admin` must ba a valid address, and their signature must be present in the transaction
- if the `allowlist_enabled` paramater is set to `true`, `admin` must be on the list of approved credit class creators (the `allowed_class_creators` parameter)
- `issuers` must not be empty and all addresses must be valid addresses 
- `credit_type` (the name of the credit type) must not be empty and on the list of approved credit types (the `credit_types` parameter)
- `metadata` must be less than or equal to 256 bytes

## MsgCreateProject

`MsgCreateProject` creates a new project with a credit class ID, issuer, project location, and optional metadata. The party signing this transaction must be a valid issuer of the project's credit class.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L113-L140

### Validation:

- `issuer` must ba a valid address
- `class_id` must be a valid credit class ID
- `project_location` must be a valid location
- `project_id` if provided, must be a valid project ID
- `metadata` must be less than or equal to 256 bytes

## MsgCreateBatch

`MsgCreateBatch` creates a new batch of credits for an existing credit class. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form. The message must be signed by an approved issuer of the desired credit class.

The message must specify the receiver of the batch of credits as well as the number of units to issue in this batch and metadata.

In order to support use cases when credits are to be immediately retired upon issuance, for each account to be issued credits, both an amount of tradeable and retired credit units can be specified.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L149-L196

### Validation:

- `issuer` must ba a valid address, and their signature must be present in the transaction
- `issuer` must be on the list of approved `issuers` for the given credit class
- `project_id` must be a valid project ID
- `recipient` must ba a valid address
- `tradable_amount` must not be negative
- `retired_amount` must not be negative
- if `retired_amount` is positive, `retirement_location` must be a valid location
- `metadata` must be less than or equal to 256 bytes
- `start_date` must be a valid date
- `end_date` must be a valid date

## MsgSend

`Send` sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L205-L243

### Validation:

- `sender` must ba a valid address, and their signature must be present in the transaction
- `recipient` must ba a valid address
- `credits` must not be empty
- `batch_denom` must be a valid batch denomination
- `tradable_amount` must not be negative
- `retired_amount` must not be negative
- if `retired_amount` is positive, `retirement_location` must be a valid location

## MsgRetire

`Retire` retires a specified number of credits in the holder's account.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L248-L275

### Validation:

- `holder` must ba a valid address, and their signature must be present in the transaction
- `credits` must not be empty
- `batch_denom` must be a valid batch denomination
- `amount` must be positive
- `location` must be a valid location

## MsgCancel

`Cancel` removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L280-L300

### Validation:

- `holder` must ba a valid address, and their signature must be present in the transaction
- `credits` must not be empty
- `batch_denom` must be a valid batch denomination
- `amount` must be positive

## MsgSell

`Sell` creates one or more sell orders (i.e. sell orders are created in batches).

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L353-L385

### Validation:

- `owner` must ba a valid address, and their signature must be present in the transaction
- `batch_denom` must be a valid credit batch denom
- `quantity` must be a positive decimal
- `ask_price` must be a positive integer

## MsgUpdateSellOrders

`UpdateSellOrders` updates one or more sell orders (i.e. sell orders are updated in batches).

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L394-L419

### Validation:

- `owner` must ba a valid address, and their signature must be present in the transaction
- `new_quantity` must be a positive decimal
- `new_ask_price` must be a positive integer

## MsgBuy

`MsgBuy` creates one or more buy orders (i.e. buy orders are created in batches).

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L424-L483

### Validation:

- `buyer` must ba a valid address, and their signature must be present in the transaction
- `quantity` must be a positive decimal
- `bid_price` must be a positive integer

## MsgAllowAskDenom

`AllowAskDenom` is a governance operation which authorizes a new ask denom to be used in sell orders.

+++ https://github.com/regen-network/regen-ledger/tree/f2def5cf4e33c85aa4f336bc3430914d9bed791b/proto/regen/ecocredit/v1alpha2/tx.proto#L494-L508

### Validation:

- `root_address` must be the address of the governance module
- `denom` must be a valid denom
