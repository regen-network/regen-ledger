source $(dirname $0)/utils.sh

set -e

TX_FLAGS="--from $ADDR1 --yes --fees 5000uregen"
echo "INFO: Creating Credit Class - C01"
regen tx ecocredit create-class $ADDR1 C "Test Credit Class" --class-fee 20000000uregen $TX_FLAGS | log_response

echo "INFO: Creating project C01-001"
regen tx ecocredit create-project C01 US "Horsetail Ranch" $TX_FLAGS | log_response

echo "INFO: Creating credit batch C01-001-20200101-20210101-001"
TEMPDIR=$(mktemp -d)
trap "rm -rf $TEMPDIR" 0 2 3 15

cat > $TEMPDIR/batch.json <<EOL
{
  "project_id": "C01-001",
  "issuer": "$ADDR1",
  "issuance": [
    {
      "recipient": "$ADDR1",
      "tradable_amount": "1000",
      "retired_amount": "500",
      "retirement_jurisdiction": "US-WA"
    },
    {
      "recipient": "$ADDR2",
      "tradable_amount": "1000",
      "retired_amount": "500",
      "retirement_jurisdiction": "US-OR"
    }
  ],
  "metadata": "regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf",
  "start_date": "2020-01-01T00:00:00Z",
  "end_date": "2021-01-01T00:00:00Z",
  "open": false
}
EOL

regen tx ecocredit create-batch $TEMPDIR/batch.json $TX_FLAGS | log_response


echo "INFO: Creating NCT basket (with C01 as allowed credit class)"
regen tx ecocredit create-basket NCT --credit-type-abbrev C --allowed-classes C01 --basket-fee 20000000uregen --description "Testing NCT Basket" $TX_FLAGS | log_response
