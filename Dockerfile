# Simple usage with a mounted data directory:
# > docker build -t ununifid .
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid start
FROM golang:1.19-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/UnUniFi/chain

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add $PACKAGES

# install and setup glibc
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.4/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.4/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 682a54082e131eaff9beec80ba3e5908113916fcb8ddf7c668cb2d97cb94c13c
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep ce3d892377d2523cf563e01120cb1436f9343f80be952c93f66aa94f5737b661
ARG arch=x86_64
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a

# for cosmwasm build option
RUN BUILD_TAGS=muslc LINK_STATICALLY=true make install

RUN apk add --update util-linux
RUN whereis ununifid

# Final image
FROM alpine:3.15

# Install ca-certificates
RUN apk add --update ca-certificates

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/ununifid /usr/bin/ununifid

# Run ununifid by default, omit entrypoint to ease using container with ununificli
CMD ["ununifid"]
