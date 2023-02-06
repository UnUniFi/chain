#!/usr/bin/env bash
# npm i -g swagger-combine

set -eo pipefail

mkdir -p ./tmp-swagger-gen

# generate swagger files (filter query files)
buf generate --template ./proto/buf.gen.swagger.yaml

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
npx swagger-combine ./docs/client/config.json -o ./docs/client/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen
