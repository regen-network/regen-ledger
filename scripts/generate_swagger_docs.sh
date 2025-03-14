#!/usr/bin/env bash

set -eo pipefail

SWAGGER_DIR=./app/client/docs
SWAGGER_UI_DIR=${SWAGGER_DIR}/swagger-ui

SDK_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/cosmos-sdk)
IBC_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/ibc-go/v7)

SDK_RAW_URL=https://raw.githubusercontent.com/cosmos/cosmos-sdk/${SDK_VERSION}/client/docs/swagger-ui/swagger.yaml
IBC_RAW_URL=https://raw.githubusercontent.com/cosmos/ibc-go/${IBC_VERSION}/docs/client/swagger-ui/swagger.yaml

SWAGGER_UI_VERSION=4.11.0
SWAGGER_UI_DOWNLOAD_URL=https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_UI_VERSION}.zip
SWAGGER_UI_PACKAGE_NAME=${SWAGGER_DIR}/swagger-ui-${SWAGGER_UI_VERSION}

# install swagger-combine if not already installed
npm list -g | grep swagger-combine > /dev/null || npm install -g swagger-combine --no-shrinkwrap

# install statik if not already installed
go install github.com/rakyll/statik@latest

# download Cosmos SDK swagger yaml file
echo "SDK version ${SDK_VERSION}"
curl -o ${SWAGGER_DIR}/swagger-sdk.yaml -sfL "${SDK_RAW_URL}"

# download IBC swagger yaml file
echo "IBC version ${IBC_VERSION}"
curl -o ${SWAGGER_DIR}/swagger-ibc.yaml -sfL "${IBC_RAW_URL}"

# combine swagger yaml files using nodejs package `swagger-combine`
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ${SWAGGER_DIR}/config.json -f yaml \
  -o ${SWAGGER_DIR}/swagger.yaml \
  --continueOnConflictingPaths true \
  --includeDefinitions true

# if swagger-ui does not exist locally, download swagger-ui and move dist directory to
# swagger-ui directory, then remove zip file and unzipped swagger-ui directory
if [ ! -d ${SWAGGER_UI_DIR} ]; then
  # download swagger-ui
  curl -o ${SWAGGER_UI_PACKAGE_NAME}.zip -sfL ${SWAGGER_UI_DOWNLOAD_URL}
  # unzip swagger-ui package
  unzip ${SWAGGER_UI_PACKAGE_NAME}.zip -d ${SWAGGER_DIR}
  # move swagger-ui dist directory to swagger-ui directory
  mv ${SWAGGER_UI_PACKAGE_NAME}/dist ${SWAGGER_UI_DIR}
  # remove swagger-ui zip file and unzipped swagger-ui directory
  rm -rf ${SWAGGER_UI_PACKAGE_NAME}.zip ${SWAGGER_UI_PACKAGE_NAME}
fi

# move generated swagger yaml file to swagger-ui directory
cp ${SWAGGER_DIR}/swagger.yaml ${SWAGGER_DIR}/swagger-ui/

# update swagger initializer to default to swagger.yaml
# Note: using -i.bak makes this compatible with both GNU and BSD/Mac
sed -i.bak "s|https://petstore.swagger.io/v2/swagger.json|swagger.yaml|" ${SWAGGER_UI_DIR}/swagger-initializer.js

# generate statik golang code using updated swagger-ui directory
statik -src=${SWAGGER_DIR}/swagger-ui -dest=${SWAGGER_DIR} -f -m

# log whether or not the swagger directory was updated
if [ -n "$(git status ${SWAGGER_DIR} --porcelain)" ]; then
  echo "Swagger statik file updated"
else
  echo "Swagger statik file already in sync"
fi
