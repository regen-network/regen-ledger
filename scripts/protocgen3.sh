# this script is for generating protobuf documentation

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do

  # get the version name, e.g. from "./proto/regen/ecocredit/v1alpha1", extract "v1alpha1"
  version=$(basename $dir)

  # get the module name, e.g. from "./proto/regen/ecocredit/v1alpha1", extract "ecocredit"
  module=$(basename $(dirname $dir))

  filename="protobuf_${version}.md"
  destination="./x/${module}/spec"

  # get the parent name, e.g. from "./proto/regen/ecocredit/basket/v1alpha1", extract "ecocredit"
  parent=$(basename $(dirname $(dirname $dir)))

  if [ $parent != "regen" ] ; then
    filename="protobuf_${module}_${version}.md"
    destination="./x/${parent}/spec"
  fi

  # command to generate docs using protoc-gen-doc
  buf protoc \
  -I "proto" \
  -I "third_party/proto" \
  --doc_out=${destination} \
  --doc_opt=docs/markdown.tmpl,${filename} \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done
