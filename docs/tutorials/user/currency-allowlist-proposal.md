# Currency Allowlist Proposal

Allowed currencies are currencies (or "denoms") that can be used for the asking price by owners of ecosystem service credits when listing credits for sale in the marketplace.

Community members or representatives of projects/protocols are encouraged to submit proposals to add allowed currencies. See [Governance Guidelines](https://github.com/regen-network/governance#guidelines) for additional information about submitting proposals and to ensure the recommended steps are taken for a successful proposal.

This document provides instructions for submitting a currency allowlist proposal using the [command-line interface (CLI)](../../ledger/infrastructure/interfaces.md#command-line-interface). For instructions on how to install the `regen` binary, see [Install Regen](../../ledger/get-started/README.md).

## Create Proposal

The first step is to create a json file that includes information about the proposal and the currency that will be added to the list of allowed currencies if the proposal passes.

Create a `proposal.json` file using the following example (note that the name and location of the file is not significant as long as you use the same name and location when submitting the proposal):

```json
{
  "title": "Add $REGEN to the currency allowlist",
  "description": "This proposal adds $REGEN to the currency allowlist",
  "messages": [
    {
      "@type": "/regen.ecocredit.marketplace.v1.MsgAddAllowedDenom",
      "authority": "regen10d07y265gmmuvt4z0w9aw880jnsr700j9qceqh",
      "bank_denom": "uregen",
      "display_denom": "regen",
      "exponent": 6
    }
  ]
}
```

Each field in the json file is required and should be properly filled out.

Make sure you give the proposal a meaningful title and description. The description should provide a rationale as to why this currency should be added to the list and can be written in either plain text or markdown (see [Proposal #15](https://wallet.keplr.app/chains/regen/proposals/15) for an example of a detailed description).

- `authority` is the address of the gov module on Regen Mainnet (you can verify this is the correct address with `regen q auth module-account gov`).

- `bank_denom` is the denom that will be added to the list. In the example above, `uregen` (i.e. "micro regen") is being added. For any denom that is not native to the network, the IBC denom is required (e.g. `ibc/CDC4587874B85BEA4FCEC3CEA5A1195139799A1FEE711A07D972537E18FD`).

- `display_denom` is the display name of the denom (e.g. `regen` or `atom`). The display denom is what users will see when interacting with the proposed currency in the marketplace.

- `exponent` is used to relate the `bank_denom` to the `display_denom` and is informational. For example, `1000000uregen` is equal to `1regen` and therefore the exponent is `6`.

Keep in mind that non-native tokens represented by an IBC denom only enables a currency from a specific source. For example, `atom` transferred from one chain is different from `atom` transferred from another chain, i.e. each of these represents `atom` but they have different IBC denoms.

## Submit Proposal

Once the json file has been created, you can use the following command to submit the proposal:

```bash
regen tx gov submit-proposal proposal.json --deposit=200000000uregen --from <key-name> --fees <fee-amount>
```

- `proposal.json` refers to the json file from the previous step, which can be deleted once the proposal has been submitted.

Additional flags may be required depending on your configuration. For example, you may need to add a `--node`, `--chain-id`, and `--keyring-backend` if these options are not preconfigured.

You can also submit the proposal without a deposit (i.e. without `--deposit`) or with an amount that is less than the full deposit and that can then be filled by another account or with another transaction. The voting period will only start once the full deposit has been received.

## Additional Resources

- [Understanding IBC Denoms](https://tutorials.cosmos.network/tutorials/understanding-ibc-denoms/)
