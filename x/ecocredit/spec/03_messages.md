# Msg Service

## MsgCreateClass

`MsgCreateClass` creates a new credit class with a credit class designer, an approved list of issuers, optional metadata, and a credit type. The party signing this transaction is the credit designer. 

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L35-L51

### Arguments:

- `designer` - The address of the account that created the credit class. The `designer` has special permissions to change the list of issuers and to perform other administrative operations.
- `issuers` - The addresses of the accounts that are approved as credit issuers.
- `metadata` - Any arbitrary metadata bytes (optional).
- `credit_type` - The type of credit (e.g. "carbon", "biodiversity").

### Response:

This operation will return a new credit class ID.

`MsgCreateClassResponse` is the response type.

- `class_id` - The unique ID of the newly created credit class.

## MsgCreateBatch

`MsgCreateBatch` creates a new batch of credits for an existing credit class. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form. The message must be signed by an approved issuer of the desired credit class.

The message must specify the receiver of the batch of credits as well as the number of units to issue in this batch and metadata.

In order to support use cases when credits are to be immediately retired upon issuance, for each account to be issued credits, both an amount of tradeable and retired credit units can be specified.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L60-L120

### Arguments:

- `issuer` - The address of the credit batch issuer.
- `class_id` - The unique ID of the credit class.
- `issuance` - The credits issued in the batch.
  - `recipient` - The address of the account that received the credit issuance.
  - `tradable_amount` - The number of credits in this issuance that can be traded by this recipient. Decimal values are acceptable.
  - `retired_amount` - The number of credits in this issuance that are effectively retired by the issuer on receipt. Decimal values are acceptable.
  - `retirement_location` - the location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country_code>[-<sub_national_code>[ <postal_code>]], with the first two fields conforming to ISO 3166-2, and postal_code being up to 64 alphanumeric characters.
- `metadata` - Any arbitrary metadata attached to the credit batch.
- `start_date` - The beginning of the period during which this credit batch was quantified and verified.
- `end_date` - The end of the period during which this credit batch was quantified and verified.
- `project_location` - The location of the project backing the credits in this batch. It is a string of the form <country_code>[-<sub_national_code>[ <postal_code>]], with the first two fields conforming to ISO 3166-2, and postal_code being up to 64 alphanumeric characters. country_code is required, while sub_national_code and postal_code can be added for increasing precision.

### Response:

This operation will return a new credit batch ID.

`MsgCreateBatchResponse` is the response type.

- `batch_denom` - The unique denomination ID of the newly created batch.

## MsgSend

Send sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L129-L170

### Arguments:

- `sender` - The address of the account sending credits.
- `recipient` - The address of the account receiving credits.
- `credits` - The credits being sent.
  - `batch_denom` - The unique ID of the credit batch.
  - `tradable_amount` - The number of credits in this transfer that can be traded by the recipient. Decimal values are acceptable within the precision returned by Query/Precision.
  - `retired_amount` - The number of credits in this transfer that are effectively retired by the issuer on receipt. Decimal values are acceptable within the precision returned by Query/Precision.
  - `retirement_location` - The location of the beneficiary or buyer of the retired credits. This must be provided if retired_amount is positive. It is a string of the form <country_code>[-<sub_national_code>[ <postal_code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters.
  
### Response:

`MsgSendResponse` is the response type.

- empty

## MsgRetire

Retire retires a specified number of credits in the holder's account.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L170-L202

### Arguments:

- `holder` - The credit holder address.
- `credits` - The credits being retired.
  - `batch_denom` - The unique ID of the credit batch.
  - `amount` - The number of credits being retired. Decimal values are acceptable within the precision returned by Query/Precision.
- `location` - The location of the beneficiary or buyer of the retired credits. It is a string of the form <country_code>[-<sub_national_code>[ <postal_code>]], with the first two fields conforming to ISO 3166-2, and postal-code being up to 64 alphanumeric characters.

### Response:

`MsgRestireResponse` is the response type.

- empty

## MsgCancel

Cancel removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L207-L227

### Arguments:

- `holder` - The credit holder address.
- `credits` - The credits being cancelled.
  - `batch_denom` - The unique ID of the credit batch.
  - `batch_denom` - The number of credits being cancelled. Decimal values are acceptable within the precision returned by Query/Precision.
  
### Response:

`MsgRestireResponse` is the response type.

- empty
