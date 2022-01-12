# IBC Transfers

The [Inter-Blockchain Communication (IBC)](https://ibcprotocol.org/) protocol is enabled on Regen Mainnet meaning token holders can transfer tokens between Regen Mainnet and other IBC enabled blockchains.

This document provides instructions for performing an IBC token transfer using either [Keplr Wallet](https://wallet.keplr.app/) or the [command-line interface (CLI)](/regen-ledger/interfaces.html#command-line-interface). 

## Available Channels

IBC transfers are only possible if there is a relayer running that establishes a connection between two blockchains using IBC channels.

The following IBC channels are currently available for transfers to and from Regen Mainnet:

- regen-1 --> osmosis-1 (`channel-1`)
- osmosis-1 --> regen-1 (`channel-8`)
- regen-1 --> cosmoshub-4 (`channel-11`)
- cosmoshub-4 --> regen-1 (`channel-185`)

## Using Keplr Wallet

IBC transfers are possible using [Keplr Wallet](https://wallet.keplr.app/). In order to transfer tokens between two IBC enabled blockchains, follow these steps:

1. Open the Keplr Wallet extension.
2. In the header of the extension, select the blockchain that you want to transfer tokens from (for example, "Osmosis" or "Regen").
3. Click the "Transfer" button under the "IBC Transfer" section. Note that you must have tokens available in order to proceed to the next step.
4. Select the "Destination Chain". If this is your first time using "IBC Transfer" or your first time sending tokens to a specific chain, you'll need to select the "New IBC Transfer Channel" option, then select the "Destination Chain", provide the "Channel ID" (see [available chains](#available-channels) for the "Channel ID"), and click the "Save" button.
5. Finally, add the recipient and a memo (optional) and then click the send button.

## Using The Command-Line Interface

IBC transfers are also possible using the [command-line interface (CLI)](/regen-ledger/interfaces.html#command-line-interface). For instructions on how to install the `regen` binary, check out [Quick Start](/getting-started/).

The following command can be used to perform an IBC transfer:
```
regen tx ibc-transfer transfer <src_port_id> <src_channel_id> <receiver_address> <amount>
```

- `src_port_id` refers to the `port_id` used while establishing the IBC connection between source and destination chains.
- `src_channel_id` refers to the `channel_id` used while establishing the IBC connection between source and destination chains.

The following example performs a transfer from Regen (`regen-1`) to Osmosis (`osmosis-1`):
```
regen tx ibc-transfer transfer transfer channel-1 <receiver_osmosis_address> 1000000uregen --from mykey --chain-id regen-1 --node http://public-rpc.regen.vitwit.com:26657 --fees 5000uregen
```
- In this example, `src_port_id` is `transfer`.
- In this example, `src_channel_id` is `channel-1`.

Transfer back the tokens to source chain (in this case, `regen-1`):

The following command can be used to perform an IBC transfer back to the source chain:
```
osmosisd tx ibc-transfer transfer <dst_port_id> <dst_channel_id> <receiver_address> <amount>
```

First lets query balances on osmosis to check ibc denom for REGEN tokens.
```
osmosisd q bank balances $(osmosisd keys show mykey -a) --node http://143.198.234.89:26657 --chain-id osmosis-1 --node https://rpc-osmosis.keplr.app:443
balances:
- amount: "1000000"
  denom: ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0
- amount: "1000000"
  denom: uion
- amount: "1000000"
  denom: uosmo
pagination:
  next_key: null
  total: "0"
```

```
osmosisd tx ibc-transfer transfer transfer channel-8 <receiver_regen_address> 1000000ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0 --from mykey --chain-id osmosis-1 --node https://osmosis.stakesystems.io:2053
```
- In this example, `dst_port_id` is `transfer`.
- In this example, `dst_channel_id` is `channel-8`.
- In this example, `ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0` is the ibc denom of `uregen` on `osmosis` chain.

## Troubleshooting

IBC is a secure and battle tested protocol.  However, there may be a few cases where your IBC transfers might get stuck and take time to process. If this happens, the IBC protocol ensures that your funds are safe and that all of your pending IBC transfers will eventually be picked up by the relayers.  In the case that they are not picked up, they will be refunded after the packet timeout.
