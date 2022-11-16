source $(dirname $0)/utils.sh

set -e

TX_FLAGS="--from $ADDR1 --yes --fees 5000uregen"

echo "INFO: Creating credit class..."
regen tx ecocredit create-class $ADDR1 C "Test Credit Class" --class-fee 20000000uregen $TX_FLAGS | log_response

CLASS_ID=$(regen q ecocredit classes | jq -r '.classes[-1].id')
echo "INFO:   Credit Class ID: $CLASS_ID"

echo "INFO: Creating project $CLASS_ID-001"
regen tx ecocredit create-project $CLASS_ID US "Horsetail Ranch" $TX_FLAGS | log_response

echo "INFO: Creating credit batch $CLASS_ID-001-20200101-20210101-001"
TEMPDIR=$(mktemp -d)
trap "rm -rf $TEMPDIR" 0 2 3 15

cat > $TEMPDIR/batch.json <<EOL
{
  "project_id": "$CLASS_ID-001",
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


echo "INFO: Creating NCT basket (with $CLASS_ID as allowed credit class)"
regen tx ecocredit create-basket NCT --credit-type-abbrev C --allowed-classes $CLASS_ID --basket-fee 20000000uregen --description "Testing NCT Basket" $TX_FLAGS | log_response
