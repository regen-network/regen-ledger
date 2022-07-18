FROM golang:1.18-alpine AS build
RUN apk add build-base git linux-headers
WORKDIR /work

COPY go.mod go.sum /work/
COPY api/go.mod api/go.sum /work/api/
COPY mocks/wasmd/go.mod mocks/wasmd/go.sum /work/mocks/wasmd/
COPY types/go.mod types/go.sum /work/types/
COPY x/ecocredit/go.mod x/ecocredit/go.sum /work/x/ecocredit/
COPY x/data/go.mod x/data/go.sum /work/x/data/

RUN go mod download
COPY ./ /work

RUN LEDGER_ENABLED=false make clean build


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

