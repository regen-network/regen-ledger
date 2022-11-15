FROM golang:1.18-alpine AS build

RUN apk -U upgrade
RUN apk add build-base git linux-headers
WORKDIR /work

COPY go.mod go.sum /work/
COPY api/go.mod api/go.sum /work/api/
COPY errors/go.mod errors/go.sum /work/errors/
COPY types/go.mod types/go.sum /work/types/
COPY x/ecocredit/go.mod x/ecocredit/go.sum /work/x/ecocredit/
COPY x/data/go.mod x/data/go.sum /work/x/data/
COPY x/intertx/go.mod x/intertx/go.sum /work/x/intertx/

RUN go mod download
COPY ./ /work

RUN LEDGER_ENABLED=false make clean build

FROM alpine:3.14 AS run
RUN apk add bash curl jq
COPY ./images/regen-sandbox/bootstrap.sh /usr/bin/bootstrap.sh

VOLUME /regen
COPY --from=build /work/build/regen /usr/bin/
WORKDIR /regen
EXPOSE 26656 26657 9090 1317

ENV BINARY /usr/bin/regen
ENV REGENHOME /regen/.regen

ENTRYPOINT ["/usr/bin/bootstrap.sh"]
STOPSIGNAL SIGTERM
