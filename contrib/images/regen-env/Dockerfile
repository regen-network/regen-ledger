FROM golang:1.17-alpine AS build
RUN apk add build-base git linux-headers
WORKDIR /work

COPY go.mod go.sum /work/
COPY types/go.mod types/go.sum /work/types/
COPY x/ecocredit/go.mod x/ecocredit/go.sum /work/x/ecocredit/
COPY x/group/go.mod x/group/go.sum /work/x/group/
COPY x/data/go.mod x/data/go.sum /work/x/data/
COPY orm/go.mod orm/go.sum /work/orm/

RUN go mod download
COPY ./ /work

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta5/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep d16a2cab22c75dbe8af32265b9346c6266070bdcf9ed5aa9b7b39a7e32e25fe0

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make clean build


FROM alpine:3.14 AS run
RUN apk add bash curl jq
COPY contrib/images/regen-env/wrapper.sh /usr/bin/wrapper.sh

VOLUME /regen
COPY --from=build /work/build/regen /regen/
WORKDIR /regen
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start", "--log_format", "plain"]
STOPSIGNAL SIGTERM

