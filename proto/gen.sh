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

protoc_install_proto_gen_doc() {
  echo "Installing protobuf protoc-gen-doc plugin"
  (go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest 2> /dev/null)
}

echo "Generating go proto code"
cd proto
buf generate --template buf.gen.gogo.yaml

protoc_install_proto_gen_doc

echo "Generating proto docs"
buf generate --template buf.gen.doc.yml

cd ..

go mod tidy

# move proto files to the right places
cp -r github.com/UnUniFi/chain/* ./
rm -rf github.com
rm -rf google.golang.org
