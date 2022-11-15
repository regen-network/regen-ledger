source $(dirname $0)/utils.sh

set -eo pipefail

TX_FLAGS="--from $ADDR1 --yes --fees 5000uregen"

echo "INFO: Creating credit class for bridging..."
regen tx ecocredit create-class $ADDR1 C "Bridging Credit Class" --class-fee 20000000uregen $TX_FLAGS | log_response
CLASS_ID=$(regen q ecocredit classes | jq -r '.classes[-1].id')
echo "INFO:   Credit Class ID: $CLASS_ID"

TEMPDIR=$(mktemp -d)
trap "rm -rf $TEMPDIR" 0 2 3 15

echo "INFO: Bridging credits from polygon and creating new credit batch"
cat > $TEMPDIR/msg_bridge_rcv.json <<EOL
{
  "body": {
    "messages": [
      {
        "@type": "/regen.ecocredit.v1.MsgBridgeReceive",
        "issuer": "$ADDR1",
        "class_id": "$CLASS_ID",
        "project": {
          "reference_id": "VCS-001",
          "jurisdiction": "CA",
          "metadata": "regen:foobar.rdf"
        },
        "batch": {
          "metadata": "regen:batch1.rdf",
          "start_date": "2020-01-01T00:00:00Z",
          "end_date": "2021-01-01T00:00:00Z",
          "recipient": "$ADDR2",
          "amount": "1000"
        },
        "origin_tx": {
          "id": "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc1283434",
          "source": "polygon",
          "contract": "0x0000000000000000000000000000000000000001",
          "note": "first bridge"
        }
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [{
        "denom": "uregen",
        "amount": "5000"
      }],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": []
}
EOL

regen tx sign $TEMPDIR/msg_bridge_rcv.json --from $ADDR1 > $TEMPDIR/msg_bridge_rcv_signed.json
regen tx broadcast $TEMPDIR/msg_bridge_rcv_signed.json | log_response

echo "INFO: Bridging credits back to polygon"
cat > $TEMPDIR/msg_bridge.json <<EOL
{
  "body": {
    "messages": [
      {
        "@type": "/regen.ecocredit.v1.MsgBridge",
        "owner": "$ADDR2",
        "target": "polygon",
        "recipient": "0x0000000000000000000000000000000000000002",
        "credits": [{
          "batch_denom": "$CLASS_ID-001-20200101-20210101-001",
          "amount": "0.000001"
        }]
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [{
        "denom": "uregen",
        "amount": "5000"
      }],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": []
}
EOL

regen tx sign $TEMPDIR/msg_bridge.json --from $ADDR2 > $TEMPDIR/msg_bridge_signed.json
regen tx broadcast $TEMPDIR/msg_bridge_signed.json | log_response
