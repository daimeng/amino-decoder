FROM golang:1.15

# Injest build args from Makefile
ARG BINARY
ARG GITHUB_USERNAME
ARG GOARCH
ENV BINARY=${BINARY}

# Install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates \
  && rm -rf /var/lib/apt/lists/*

# Set working directory for the build
WORKDIR /go/src/github.com/${GITHUB_USERNAME}/${BINARY}

# Add source files
COPY . .

# Make the binary
RUN make install

WORKDIR /root

# Run ${BINARY} by default
CMD ${BINARY}
