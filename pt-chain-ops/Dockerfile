FROM golang:1.19.9-alpine3.15 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

COPY ./pt-chain-ops /app/pt-chain-ops
COPY ./pt-bindings /app/pt-bindings
COPY ./pt-node /app/pt-node
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/pt-chain-ops

RUN make pt-migrate

FROM alpine:3.15

COPY --from=builder /app/pt-chain-ops/bin/pt-migrate /usr/local/bin

ENTRYPOINT ["pt-migrate"]
