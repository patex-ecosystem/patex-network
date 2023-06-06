FROM --platform=$BUILDPLATFORM golang:1.19.0-alpine3.15 as builder

ARG VERSION=v0.0.0

RUN apk add --no-cache make gcc musl-dev linux-headers git jq bash

# build pt-proposer with the shared go.mod & go.sum files
COPY ./pt-proposer /app/pt-proposer
COPY ./pt-bindings /app/pt-bindings
COPY ./pt-node /app/pt-node
COPY ./pt-service /app/pt-service
COPY ./pt-signer /app/pt-signer
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./.git /app/.git

WORKDIR /app/pt-proposer

RUN go mod download

ARG TARGETOS TARGETARCH

RUN make pt-proposer VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

FROM alpine:3.15

COPY --from=builder /app/pt-proposer/bin/pt-proposer /usr/local/bin

CMD ["pt-proposer"]
