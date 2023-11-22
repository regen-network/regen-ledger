source $(dirname $0)/utils.sh

TX_FLAGS="--from $ADDR1 --yes --fees 5000uregen"

echo "INFO: Anchoring dataset: regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"
regen tx data anchor regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf $TX_FLAGS | log_response

echo "INFO: Attesting dataset: regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf"
regen tx data attest regen:13toVgf5UjYBz6J29x28pLQyjKz5FpcW3f4bT5uRKGxGREWGKjEdXYG.rdf $TX_FLAGS | log_response

echo "INFO: Defining resolver http://resolver.mydataservice.com"
regen tx data define-resolver "http://resolver.mydataservice.com" $TX_FLAGS | log_response

echo "INFO: Registering dataset to resolver http://resolver.mydataservice.com"
TEMPDIR=$(mktemp -d)
trap "rm -rf $TEMPDIR" 0 2 3 15

cat > $TEMPDIR/content.json <<EOL
{
  "content_hashes": [
    {
      "graph": {
        "hash": "YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE=",
        "digest_algorithm": "DIGEST_ALGORITHM_BLAKE2B_256",
        "canonicalization_algorithm": "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
        "merkle_tree": "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED"
      }
    }
  ]
}
EOL

regen tx data register-resolver 1 $TEMPDIR/content.json $TX_FLAGS | log_response
