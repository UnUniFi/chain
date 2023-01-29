#!/usr/bin/env bash
# https://docs.buf.build/installation/
# https://github.com/grpc-ecosystem/grpc-gateway#installation
# Note: go version 1.18, buf version 1.13.1
# go mod tidy
# go install \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
#     google.golang.org/protobuf/cmd/protoc-gen-go \
#     google.golang.org/grpc/cmd/protoc-gen-go-grpc

set -eo pipefail

protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the cosmos-sdk folder."
    return 1
  fi

  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
}

protoc_gen_gocosmos

echo "buf generate start"
cd proto
buf generate --template buf.gen.yaml

cd ..

echo "doc generate"
# command to generate docs using protoc-gen-doc
protoc \
  -I "proto" \
  -I "proto-thirdparty" \
  --doc_out=./docs/core \
  --doc_opt=./docs/protodoc-markdown.tmpl,proto-docs.md:./ \
  ./proto/*/*.proto

go mod tidy

# move proto files to the right places
cp -r github.com/UnUniFi/chain/* ./
rm -rf github.com
rm -rf google.golang.org
