FROM golang:1.23 as builder

ARG BASE_PATH=/root
ARG REGEN_VERSION=v5.1.0  # Default: v5

WORKDIR ${BASE_PATH}

RUN apt-get update && apt-get install -y jq expect

# Install Regen binary
RUN git clone https://github.com/regen-network/regen-ledger.git && \
    cd regen-ledger && \
    git checkout ${REGEN_VERSION} && \
    make build && \
    cp build/regen /usr/local/bin

# Install Cosmovisor
RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0
RUN cp /go/bin/cosmovisor /usr/local/bin/cosmovisor
