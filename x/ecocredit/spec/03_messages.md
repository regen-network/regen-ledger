# Msg Service

## MsgCreateClass

CreateClass creates a new credit class with an approved list of issuers and optional metadata.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L35-51

## MsgCreateBatch

CreateBatch creates a new batch of credits for an existing credit class. This will create a new batch denom with a fixed supply. Issued credits can be distributed to recipients in either tradable or retired form.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L60-120

## MsgSend

Send sends tradable credits from one account to another account. Sent credits can either be tradable or retired on receipt.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L129-170

## MsgRetire

Retire retires a specified number of credits in the holder's account.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L170-202

## MsgCancel

Cancel removes a number of credits from the holder's account and also deducts them from the tradable supply, effectively cancelling their issuance on Regen Ledger.

+++ https://github.com/regen-network/regen-ledger/blob/master/proto/regen/ecocredit/v1alpha1/tx.proto#L207-227

...

### Create Credit Class

The create credit class operation creates a new credit class.

Arguments:
- the list of issuers of the new credit class
- arbitrary metadata bytes (optional)

The party signing this transaction is the credit designer. This operation will return a new credit class ID.

### Update Credit Class

The update credit class operation will allow for the following to be changed:
- the list of approved issuers 
- the credit designer
- arbitrary metadata bytes attached to the credit

### Issue

The issue operation issues a credit batch of a credit class. It must be signed by an approved issuer of the desired credit class and specify who the receiver of the issued credits will be as well as the number of units to issue in this batch and metadata as described below.

In order to support use cases when credits are to be immediately retired upon issuance, for each account to be issued credits, both an amount of tradeable and retired credit units can be specified.

The arguments for the issue operation are thus:
- Credit Class ID
- Issuer
- Metadata
- List of:
  - Receiving account
  - Tradable units
  - Retired units

This operation will return a new credit batch ID.