# Use Go 1.23 as the base image for building and setting up the environment
FROM golang:1.23 as builder

# Define a build argument for the base path
ARG BASE_PATH=/mnt/nvme

# Set the working directory using the base path
WORKDIR ${BASE_PATH}

# Update package lists and install dependencies (jq)
RUN apt-get update && apt-get install -y jq

# Clone the regen repository and checkout the specific version (this can be parameterized)
ARG REGEN_VERSION=v6.0.0-rc5
RUN git clone https://github.com/regen-network/regen-ledger.git && \
    cd regen-ledger && \
    git checkout ${REGEN_VERSION}

# Build the regen binary using the Makefile
RUN cd ${BASE_PATH}/regen-ledger && make build && cp build/regen /usr/local/bin