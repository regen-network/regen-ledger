## regen tx ecocredit create-basket

Creates a bank denom that wraps credits

### Synopsis

Creates a bank denom that wraps credits

Parameters:
		name: the name used to create a bank denom for this basket token.

Flags:
		exponent: the exponent used for converting credits to basket tokens and for bank
			denom metadata. The exponent also limits the precision of credit amounts
			when putting credits into a basket. An exponent of 6 will mean that 10^6 units
			of a basket token will be issued for 1.0 credits and that this should be
			displayed as one unit in user interfaces. It also means that the maximum
			precision of credit amounts is 6 decimal places so that the need to round is
			eliminated. The exponent must be >= the precision of the credit type at the
			time the basket is created.
		disable-auto-retire: disables the auto-retirement of credits upon taking credits
			from the basket. The credits will be auto-retired if disable_auto_retire is
			false unless the credits were previously put into the basket by the address
			picking them from the basket, in which case they will remain tradable.
		credit-type-abbreviation: filters against credits from this credit type abbreviation (e.g. "BIO").
		allowed_classes: comma separated (no spaces) list of credit classes allowed to be put in
			the basket (e.g. "C01,C02").
		min-start-date: the earliest start date for batches of credits allowed into the basket.
		start-date-window: the duration of time (in seconds) measured into the past which sets a
			cutoff for batch start dates when adding new credits to the basket.
		basket-fee: the fee that the curator will pay to create the basket. It must be >= the
			required Params.basket_creation_fee. We include the fee explicitly here so that the
			curator explicitly acknowledges paying this fee and is not surprised to learn that the
			paid a big fee and didn't know beforehand.
		description: the description to be used in the basket coin's bank denom metadata.

```
regen tx ecocredit create-basket [name] [flags]
```

### Examples

```

		$regen tx ecocredit create-basket HEAED
			--from regen...
			--exponent=3
			--credit-type-abbreviation=FOO
			--allowed_classes="class1,class2"
			--basket-fee=100regen
			--description="any description"
		
```

### Options

```
  -a, --account-number uint               The account number of the signing account (offline mode only)
      --allowed-classes strings           comma separated (no spaces) list of credit classes allowed to be put in the basket (e.g. "C01,C02")
      --basket-fee string                 the fee that the curator will pay to create the basket (e.g. "20regen")
  -b, --broadcast-mode string             Transaction broadcasting mode (sync|async|block) (default "sync")
      --credit-type-abbreviation string   filters against credits from this credit type abbreviation (e.g. "C")
      --description string                the description to be used in the bank denom metadata.
      --disable-auto-retire               dictates whether credits will be auto-retired upon taking
      --dry-run                           ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
      --exponent string                   the exponent used for converting credits to basket tokens
      --fee-account string                Fee account pays fees for the transaction instead of deducting from the signer
      --fees string                       Fees to pay along with transaction; eg: 10uatom
      --from string                       Name or address of private key with which to sign
      --gas string                        gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float              adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string                 Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only                     Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                              help for create-basket
      --keyring-backend string            Select keyring's backend (os|file|kwallet|pass|test|memory) (default "os")
      --keyring-dir string                The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                            Use a connected Ledger device
      --minimum-start-date string         the earliest start date for batches of credits allowed into the basket (e.g. "2012-01-01")
      --node string                       <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string                       Note to add a description to the transaction (previously --memo)
      --offline                           Offline mode (does not allow any online functionality
  -o, --output string                     Output format (text|json) (default "json")
  -s, --sequence uint                     The sequence number of the signing account (offline mode only)
      --sign-mode string                  Choose sign mode (direct|amino-json), this is an advanced feature
      --start-date-window uint            sets a cutoff for batch start dates when adding new credits to the basket (e.g. 1325404800)
      --timeout-height uint               Set a block timeout height to prevent the tx from being committed past a certain height
  -y, --yes                               Skip tx broadcasting prompt confirmation
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [regen tx ecocredit](regen_tx_ecocredit.md)	 - Ecocredit module transactions

###### Auto generated by spf13/cobra on 17-Mar-2022
