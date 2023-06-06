GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
VERSION := v0.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

pt-batcher:
	env GO111MODULE=on GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v $(LDFLAGS) -o ./bin/pt-batcher ./cmd

clean:
	rm bin/pt-batcher

test:
	go test -v ./...

lint:
	golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint -e "errors.As" -e "errors.Is"

fuzz:
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelConfig_CheckTimeout ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationZero ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationTimeoutMaxChannelDuration ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationTimeoutZeroMaxChannelDuration ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelCloseTimeout ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelZeroCloseTimeout ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzSeqWindowClose ./batcher
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzSeqWindowZeroTimeoutClose ./batcher

.PHONY: \
	pt-batcher \
	clean \
	test \
	lint \
	fuzz
