# Simple usage with a mounted data directory:
# > docker build -t ununifid .
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid init
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ununifi:/root/.ununifi ghcr.io/ununifi/ununifid ununifid start
FROM golang:1.17-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 musl-dev

# Set working directory for the build
WORKDIR /go/src/github.com/UnUniFi/chain

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES

# install and setup glibc
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-2.25-r0.apk
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-bin-2.25-r0.apk
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-i18n-2.25-r0.apk
RUN apk add --no-cache glibc-2.25-r0.apk glibc-bin-2.25-r0.apk glibc-i18n-2.25-r0.apk
ENV LD_LIBRARY_PATH /usr/glibc-compat/lib
RUN /usr/glibc-compat/bin/localedef -i en_US -f UTF-8 en_US.UTF-8
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