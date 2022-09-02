# Currency Allowlist Proposal

Community members, or representatives of projects/protocols are encouraged to submit proposals for allowing currencies to be used in the marketplace.

This document provides instructions for submitting a currency allowlist proposal using the [command-line interface (CLI)](../ledger/infrastructure/interfaces.md#command-line-interface).

## Submitting the proposal via the Command-Line Interface

Before submitting the proposal you will need to create a json file using the template below

***proposal.json:***
```json
{
    "title": "Adding $REGEN as a marketplace currency",
    "description": "This proposal adds $REGEN to the currency allowlist on the Regen Registry marketplace",
    "denom": {
        "bank_denom": "uregen",
        "display_denom": "regen",
        "exponent": 6
    }
}
```
Each field in the json template above are important and should be properly filled out.  Make sure you give the proposal a meaningful title and description.  You will also
be filling out the denom fields described below.

`bank_denom` is the underlying coin denom (i.e. ibc/CDC4587874B85BEA4FCEC3CEA5A1195139799A1FEE711A07D972537E18FD).

`display_denom` is used for display purposes, and serves as the name of the coin denom (i.e. ATOM).

`exponent` is used to relate the `bank_denom` to the `display_denom` and is informational

In our proposal.json example, `bank_denom` is referencing `uregen` (micro-regen) and since regen is the native token of Regen Ledger we do not need to refer to the IBC denom.  If you are proposing to allow any other IBC compatible denom you will need to refer to the IBC denom. The display denom is what users will see when interacting with the currency proposed, which in our example is regen.  

Once the json has been properly edited please use this command to submit the allowlist proposal:
```
regen tx gov submit-proposal allow-denom-proposal proposal.json --deposit=200000000uregen 
```

`proposal.json` refers to the proposal json template above.


## Resources

- [Understanding IBC denoms with Gaia](https://tutorials.cosmos.network/tutorials/understanding-ibc-denoms/)
