# Simple usage with a mounted data directory:
# > docker build -t jpyx .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.jpyx:/jpyx/.jpyx jpyx jpyxd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.jpyx:/jpyx/.jpyx jpyx jpyxd start
FROM golang:1.16-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/lcnem/jpyx

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
  make install

# Final image
FROM alpine:edge

ENV JPYX /jpyx

# Install ca-certificates
RUN apk add --update ca-certificates

RUN addgroup jpyx && \
  adduser -S -G jpyx jpyx -h "$JPYX"

USER jpyx

WORKDIR $JPYX

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/jpyxd /usr/bin/jpyxd

# Run jpyxd by default, omit entrypoint to ease using container with jpyxcli
CMD ["jpyxd"]