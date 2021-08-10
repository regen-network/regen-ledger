# Inter-Blockchain Connections (IBC)

## Transferring tokens between the REGEN and other chains

The IBC protocol is active on Regen Ledger and token holders are able to transfer tokens from one IBC enabled chain to another.  This guide provides the configurations for Keplr wallet and includes the command-line interface (CLI) commands that are needed to accomplish and IBC transfer. 

## Keplr Wallet
In order to perform IBC transactions using the Keplr wallet and transfer tokens from the Regen to another chain such as Osmosis, open the keplr extension, click transfer under the IBC Transfer section and simply set the channel-id before clicking send.

## Using the command-line interface
The command for transferring REGEN to OSMOSIS (Regen-1 -> Osmosis-1) via CLI and vice versa is listed including the general commands and examples

**Transferring tokens from the Regen Ledger to Osmosis chain:**

**General command for ibc-transfer**
```sh
$ regen tx ibc-transfer transfer <src_port_id> <src_channel_id> <reciver_address> <amount>
```
Where,
- `src_port_id` refers to the `port_id` used while establishing the ibc connection between source and destination chains. Here for `REGEN<>OSMOSIS`, `port_id` is `transfer`
- `src_channel_id` refers to the `channel_id` used while establishing the ibc connection between source and destination chains. Here for `REGEN<>OSMOSIS`, `src_channel_id` is `channel-1`

**Example:**
```sh
regen tx ibc-transfer transfer transfer channel-1 <osmosis_receiver_address> 1000000uregen --from mykey --chain-id regen-1 --node http://public-rpc.regen.vitwit.com:26657 --fees 5000uregen
```

#### Troubleshooting
IBC is a secure and battle tested protocol on multiple testnets in order to be production ready.  However, there may be a few cases where your IBC transfers might get stuck and take time to process. If this happens, the IBC protocol ensures that your funds are safe and that all of your pending IBC transfers will eventually be picked up by the relayers.  In case they are not picked up they will be refunded after the packet timeout.

