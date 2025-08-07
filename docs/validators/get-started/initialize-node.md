# Initialize Node

The following instructions assume that you have already completed the following:

- [Initial Setup](README)
- [Install Regen](install-regen.md)

## Initialize Node

Create the configuration files and data directory by initializing the node. In the following command, replace `[moniker]` with a name of your choice. 

*For Regen Mainnet:*

```bash
regen init [moniker] --chain-id regen-1
```

*For Regen Testnet:*

```bash
regen init [moniker] --chain-id regen-upgrade
```

## Update Genesis

Update the genesis file.

*For Regen Mainnet:*

```bash
curl http://mainnet.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Regen Testnet:*

We should use /genesis_chunk endpoint instead of /genesis because the genesis response is too large.
/genesis_chunked endpoint will return data which is encoded in base64. Combining the outputs of all the chunks and base64 decoding then will produce the genesis file of the network.

```bash
#!/bin/bash

# Output file
output_file="genesis.json"
temp_file="genesis_temp.json"

# Start clean
> "$temp_file"

# Fetch first chunk to get total number of chunks
initial_response=$(curl -s "https://rpc-regen-upgrade.vitwit.com/genesis_chunked?chunk=0")

# Extract total number of chunks
total_chunks=$(echo "$initial_response" | jq -r '.result.total')
if [[ -z "$total_chunks" || "$total_chunks" == "null" ]]; then
  echo "Failed to retrieve total chunk count."
  exit 1
fi

echo "Total chunks to fetch: $total_chunks"

# Loop through all chunks
for ((i=0; i<total_chunks; i++)); do
  echo "Fetching chunk $i..."
  response=$(curl -s "https://rpc-regen-upgrade.vitwit.com/genesis_chunked?chunk=$i")
  chunk_data=$(echo "$response" | jq -r '.result.data')

  if [[ "$chunk_data" == "null" || -z "$chunk_data" ]]; then
    echo "Failed to retrieve chunk $i"
    exit 1
  fi

  # Decode the chunk and append to temp file
  echo "$chunk_data" | base64 --decode >> "$temp_file"
done

# Move the completed file to genesis.json
mv "$temp_file" "$output_file"

echo "Genesis file assembled successfully into $output_file"
```

## Update Peers

Add a seed node for initial peer discovery.

*For Regen Mainnet:*

```bash
PERSISTENT_PEERS="c4460b52c34ad4f12168d05807e998bb8e8b4812@mainnet.regen.network:26656,aebb8431609cb126a977592446f5de252d8b7fa1@regen.rpc.vitwit.com:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Regen Testnet:*

```bash
PERSISTENT_PEERS="fc6cf74d3de04ab0a3836f5adf50968990f2c195@155.138.196.253:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Start Node

The node is now ready to connect to the network:

```bash
regen start
```

## Create a Validator

The next step will be to [create a validator](create-a-validator.md).

## Using State Sync

Also, syncing from genesis will be a slow process. You may want to consider [using state sync](using-state-sync.md).
