# Configure a remote xrncli


Before you begin you need to know the IP address:port ond chain-id of the xrnd you are connecting to.

```bash
xrncli config node <host>:<port>

# example: xrncli config node http://172.31.26.37:26657

# NOTE: https is used in the cosomo docs, but I haven't got it working yet. Need to check genesis.json and docs...
```

```bash
xrncli config trust-node true

# Set to true if you trust the full-node you are connecting to, otherwise false
```

```bash
xrncli config chain-id <chain-id-of-xrnd>
````
"/tmp/setup_remote_xrncli.md" 36L, 778C
