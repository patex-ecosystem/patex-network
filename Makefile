COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
BEDROCK_TAGS_REMOTE?=origin

build: build-go build-ts
.PHONY: build

build-go: submodules pt-node pt-proposer pt-batcher
.PHONY: build-go

build-ts: submodules
	if [ -n "$$NVM_DIR" ]; then \
		. $$NVM_DIR/nvm.sh && nvm use; \
	fi
	yarn install
	yarn build
.PHONY: build-ts

submodules:
	# CI will checkout submodules on its own (and fails on these commands)
	if [ -z "$$GITHUB_ENV" ]; then \
		git submodule init; \
		git submodule update; \
	fi
.PHONY: submodules

pt-bindings:
	make -C ./pt-bindings
.PHONY: pt-bindings

pt-node:
	make -C ./pt-node pt-node
.PHONY: pt-node

pt-batcher:
	make -C ./pt-batcher pt-batcher
.PHONY: pt-batcher

pt-proposer:
	make -C ./pt-proposer pt-proposer
.PHONY: pt-proposer

mod-tidy:
	# Below GOPRIVATE line allows mod-tidy to be run immediately after
	# releasing new versions. This bypasses the Go modules proxy, which
	# can take a while to index new versions.
	#
	# See https://proxy.golang.org/ for more info.
	export GOPRIVATE="github.com/patex-ecosystem/patex-network" && go mod tidy
.PHONY: mod-tidy

clean:
	rm -rf ./bin
.PHONY: clean

sepolia-up:
	@bash ./ops-bedrock/patex-sepolia/node-up.sh
.PHONY: sepolia-up

sepolia-down:
	@(cd ./ops-bedrock/patex-sepolia && GENESIS_TIMESTAMP=$(shell date +%s) docker-compose stop)
.PHONY: sepolia-down

sepolia-clean:
	rm -rf ./.patex-sepolia
	cd ./ops-bedrock/patex-sepolia && docker-compose down
	docker image ls 'patex-sepoli*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-bedrock --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: sepolia-clean

sepolia-logs:
	@(cd ./ops-bedrock/patex-sepolia && docker-compose logs -f)
	.PHONY: sepolia-logs

test-unit:
	make -C ./pt-node test
	make -C ./pt-proposer test
	make -C ./pt-batcher test
	yarn test
.PHONY: test-unit

# Remove the baseline-commit to generate a base reading & show all issues
semgrep:
	$(eval DEV_REF := $(shell git rev-parse develop))
	SEMGREP_REPO_NAME=patex-ecosystem/patex-network semgrep ci --baseline-commit=$(DEV_REF)
.PHONY: semgrep

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules


tag-bedrock-go-modules:
	./ops/scripts/tag-bedrock-go-modules.sh $(BEDROCK_TAGS_REMOTE) $(VERSION)
.PHONY: tag-bedrock-go-modules

update-pt-geth:
	./ops/scripts/update-pt-geth.py
.PHONY: update-pt-geth
