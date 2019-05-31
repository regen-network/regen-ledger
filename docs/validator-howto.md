# Validator How To

Setting up an xrnd validator node requires editing two configuration files in the ~/.xrnd/config directory; config.toml and genesis.json

## config.toml

We will need to know the node-id, the public IP address, and the port.

To determine the node address:
```bash
$ xrnd tendermint show-node-id
0d2e4d4894afd85f1ea5c7256bdfe6386e994af4
```
The output is your node id. You will use this string as the part to the right of the @ in your address.

You will also need to know the public IP address of your host. The non-routable address (eg 172.x.x.x or 192.x.x.x) subnet blocks will _not_ work.

The *nix command ifconfig will not work either.

Hence:
```bash
$ dig @resolver1.opendns.com ANY myip.opendns.com +short
```
This should output your IP address, the one that the public Internet sees.

By Cosmos convention, we will use the port number 26656.

Our address is constructed by combining node-id@IP-address:port

For example: 2e4dds4894afd85f1ea5c7256bdfe6386e994af4@54.83.123.33:26656

Open the file ~/.xrnd/config/config.toml for editing and locate the parameter persistent_peers. It will look something like this:

```bash
persistent_peers = 0d2e4d4894afd85f1ea5c7256bdfe6386e994af4@127.0.0.1:26656,e8799060939c0287c9e4647615b9d5301823aca9@52.90.226.56:26656,6ef5728f20f822fd62599d33dc099c53f32e842e@54.91.81.95:26656
```

Append a comma separator to the end of the line and then enter your address. Copy/Paste is your friend here, like this:
```bash
persistent_peers = "0d2e4d4894afd85f1ea5c7256bdfe6386e994af4@127.0.0.1:26656,e8799060939c0287c9e4647615b9d5301823aca9@52.90.226.56:26656,6ef5728f20f822fd62599d33dc099c53f32e842e@54.91.81.95:26656,2e4dds4894afd85f1ea5c7256bdfe6386e994af4@54.83.123.33:26656"
```
Be careful not to add a newline separator as this will break the parser. Everything on one line.

Add your external (public) IP address here:
```bash
external_address = "tcp://54.83.123.33:26657"
```
Notice that the port number is 26657.

Useful but not strictly necessary, add one or more seeds to help your host find other peers on the network.
```bash
seeds = "7256bdfe6386e992e4dds4894afd854a@33.44.55.66:26656"
```

Save the file, and start the daemon with
```bash
$ xrnd start
```


## Useful commands

If for some reason you trash your database and want to resart in a pristine state:
```bash
$ xrnd unsafe-reset-all
```
This will clean all your data and address book but will not delete your config files.

## AWS notes

By default AWS allows 1024 open files. This needs to be increased to 4096.
```bash
$ ulimit -n 4096
```

Link to AWS-specific notes from Cosmos:
https://forum.cosmos.network/t/using-an-aws-ec2-as-a-cosmos-node/845
Stub for [this story](https://www.pivotaltracker.com/n/projects/2318873/stories/166162869)
