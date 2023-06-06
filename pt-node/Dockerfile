FROM --platform=$BUILDPLATFORM golang:1.19.9-alpine3.16 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build pt-node with the shared go.mod & go.sum files
COPY ./pt-node /app/pt-node
COPY ./pt-chain-ops /app/pt-chain-ops
COPY ./pt-service /app/pt-service
COPY ./pt-bindings /app/pt-bindings
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/pt-node

RUN go mod download

ARG TARGETOS TARGETARCH

RUN make pt-node VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

FROM alpine:3.16

COPY --from=builder /app/pt-node/bin/pt-node /usr/local/bin

CMD ["pt-node"]
