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
proto_dirs=$(find ./ -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    buf generate --template buf.gen.yaml $file
  done
done

# cd ..

# command to generate docs using protoc-gen-doc
# buf protoc \
#   -I "proto" \
#   -I "proto-thirdparty" \
#   --doc_out=./docs/core \
#   --doc_opt=/docs/protodoc-markdown.tmpl,proto-docs.md. \
#   $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

cd ..

go mod tidy

# move proto files to the right places
cp -r github.com/UnUniFi/chain/* ./
rm -rf github.com
rm -rf google.golang.org
