# Simple usage with a mounted data directory:
# > docker build -t ununifid .
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid start
FROM golang:1.17-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/UnUniFi/chain

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/ununifid /usr/bin/ununifid

# Run ununifid by default, omit entrypoint to ease using container with ununificli
CMD ["ununifid"]