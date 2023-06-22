#!/usr/bin/env bash

./scripts/post.sh

# specs directory

cp -r ../specs specs

# modules directory

mkdir -p modules

cp README_modules.md modules/README.md

for D in ../x/*; do
  if [ -d "${D}" ]; then
    rm -rf "modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

(cd .. ; ./scripts/generate_feature_docs.sh)

# TODO: better solution for closing end tag errors
for FILE in modules/**/features/server/*; do
  sed -i 's/<amount>/[amount]/g' $FILE
  sed -i 's/<balance-amount>/[balance-amount]/g' $FILE
  sed -i 's/<balance-before>/[balance-before]/g' $FILE
  sed -i 's/<basket-fee>/[basket-fee]/g' $FILE
  sed -i 's/<batch-start-date>/[batch-start-date]/g' $FILE
  sed -i 's/<bid-amount>/[bid-amount]/g' $FILE
  sed -i 's/<class-fee>/[class-fee]/g' $FILE
  sed -i 's/<credit-amount>/[credit-amount]/g' $FILE
  sed -i 's/<disable-auto-retire>/[disable-auto-retire]/g' $FILE
  sed -i 's/<expiration>/[expiration]/g' $FILE
  sed -i 's/<quantity>/[quantity]/g' $FILE
  sed -i 's/<precision>/[precision]/g' $FILE
  sed -i 's/<retire-on-take>/[retire-on-take]/g' $FILE
  sed -i 's/<retired-credits>/[retired-credits]/g' $FILE
  sed -i 's/<retirement-on-take>/[retirement-on-take]/g' $FILE
  sed -i 's/<token-amount>/[token-amount]/g' $FILE
  sed -i 's/<token-balance>/[token-balance]/g' $FILE
  sed -i 's/<tradable-credits>/[tradable-credits]/g' $FILE
done

# TODO: better solution for closing end tag errors
for FILE in modules/**/features/types/*; do
  sed -i 's/<class-id>/[class-id]/g' $FILE
  sed -i 's/<class-sequence>/[class-sequence]/g' $FILE
  sed -i 's/<country-code>/[country-code]/g' $FILE
  sed -i 's/<credit-type-abbrev>/[credit-type-abbrev]/g' $FILE
  sed -i 's/<exponent-prefix>/[exponent-prefix]/g' $FILE
  sed -i 's/<name>/[name]/g' $FILE
  sed -i 's/<postal-code>/[postal-code]/g' $FILE
  sed -i 's/<project-id>/[project-id]/g' $FILE
  sed -i 's/<project-sequence>/[project-sequence]/g' $FILE
  sed -i 's/<region-code>/[region-code]/g' $FILE
done

# commands directory

mkdir -p commands

cp README_commands.md commands/README.md

go run ../scripts/generate_cli_docs.go

# TODO: better solution for closing end tag errors
for FILE in commands/*; do
  sed -i 's/<appd>/regen/g' $FILE

  sed -i 's/<addr>/[addr]/g' $FILE
  sed -i 's/<chain-id>/[chain-id]/g' $FILE
  sed -i 's/<file>/[file]/g' $FILE
  sed -i 's/<granter>/[granter]/g' $FILE
  sed -i 's/<hash>/[hash]/g' $FILE
  sed -i 's/<name>/[name]/g' $FILE
  sed -i 's/<recipient>/[recipient]/g' $FILE
  sed -i 's/<seq>/[seq]/g' $FILE
  sed -i 's/<sequence>/[sequence]/g' $FILE
done
