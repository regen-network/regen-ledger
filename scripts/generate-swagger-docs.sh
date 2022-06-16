#!/usr/bin/env bash

set -eo pipefail

SWAGGER_DIR=./client/docs
SWAGGER_UI_DIR=${SWAGGER_DIR}/swagger-ui

SDK_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/cosmos-sdk)
IBC_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/ibc-go/v2)

SDK_RAW_URL=https://raw.githubusercontent.com/cosmos/cosmos-sdk/${SDK_VERSION}/client/docs/swagger-ui/swagger.yaml
IBC_RAW_URL=https://raw.githubusercontent.com/cosmos/ibc-go/${IBC_VERSION}/docs/client/swagger-ui/swagger.yaml

SWAGGER_UI_VERSION=4.11.0
SWAGGER_UI_DOWNLOAD_URL=https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_UI_VERSION}.zip
SWAGGER_UI_PACKAGE_NAME=${SWAGGER_DIR}/swagger-ui-${SWAGGER_UI_VERSION}

# download Cosmos SDK swagger yaml file
echo "SDK version ${SDK_VERSION}"
wget "${SDK_RAW_URL}" -O ${SWAGGER_DIR}/swagger-sdk.yaml

# download IBC swagger yaml file
echo "IBC version ${IBC_VERSION}"
wget "${IBC_RAW_URL}" -O ${SWAGGER_DIR}/swagger-ibc.yaml

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
  wget ${SWAGGER_UI_DOWNLOAD_URL} -O ${SWAGGER_UI_PACKAGE_NAME}.zip
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
  echo "\033[91mSwagger updated\033[0m"
else
  echo "\033[92mSwagger in sync\033[0m"
fi
