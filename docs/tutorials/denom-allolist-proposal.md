# Currency Allow list Proposals

Community members, or representatives of projects/protocols are encouraged to submit proposals for allowing currencies to be used in the marketplace.

This document provides instructions for submitting a currency allowlist proposal using the [command-line interface (CLI)](../ledger/infrastructure/interfaces.html#command-line-interface).

## Submitting the proposal via the Command-Line Interface

Before submitting the proposal you will need to format a json using the template below

***proposal.json:***
```
{
    "title": "some title",
    "description": "some description",
    "denom": {
        "bank_denom": "uregen",
        "display_denom": "regen",
        "exponent": 6
    }
}
```
`bank_denom` is the underlying coin denom (i.e. ibc/CDC4587874B85BEA4FCEC3CEA5A1195139799A1FEE711A07D972537E18FD).

`display_denom` is used for display purposes, and serves as the name of the coin denom (i.e. ATOM).

`exponent` is used to relate the `bank_denom` to the `display_denom` and is informational

The following command can be used to submit the allowlist proposal:
```
regen tx gov submit-proposal allow-denom-proposal <proposal.json> --deposit=200000000uregen 
```

`proposal.json` refers to the proposal json template above.



## Resources

- [Understanding IBC denoms with Gaia](https://tutorials.cosmos.network/tutorials/understanding-ibc-denoms/)
