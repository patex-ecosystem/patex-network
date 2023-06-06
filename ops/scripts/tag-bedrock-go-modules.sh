#!/usr/bin/env bash

BEDROCK_TAGS_REMOTE="$1"
VERSION="$2"

if [ -z "$VERSION" ]; then
	echo "You must specify a version."
	exit 0
fi

FIRST_CHAR=$(printf '%s' "$VERSION" | cut -c1)
if [ "$FIRST_CHAR" != "v" ]; then
	echo "Tag must start with v."
	exit 0
fi

git tag "pt-bindings/$VERSION"
git tag "pt-service/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-bindings/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-service/$VERSION"

cd pt-chain-ops
go get github.com/patex-ecosystem/patex-network/pt-bindings@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-service@$VERSION
go mod tidy

git add .
git commit -am 'chore: Upgrade pt-chain-ops dependencies'

git tag "pt-chain-ops/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-chain-ops/$VERSION"

cd ../pt-node
go get github.com/patex-ecosystem/patex-network/pt-bindings@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-service@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-chain-ops@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in pt-node/version/version.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade pt-node dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "pt-node/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-node/$VERSION"

cd ../pt-proposer
go get github.com/patex-ecosystem/patex-network/pt-bindings@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-service@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-node@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in pt-proposer/cmd/main.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade pt-proposer dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "pt-proposer/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-proposer/$VERSION"

cd ../pt-batcher
go get github.com/patex-ecosystem/patex-network/pt-bindings@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-service@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-node@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-proposer@$VERSION
go mod tidy

echo Please update the version to ${VERSION} in pt-batcher/cmd/main.go
read -p "Press [Enter] key to continue"

git add .
git commit -am 'chore: Upgrade pt-batcher dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "pt-batcher/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-batcher/$VERSION"

cd ../pt-e2e
go get github.com/patex-ecosystem/patex-network/pt-bindings@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-service@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-node@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-proposer@$VERSION
go get github.com/patex-ecosystem/patex-network/pt-batcher@$VERSION
go mod tidy

git add .
git commit -am 'chore: Upgrade pt-e2e dependencies'
git push $BEDROCK_TAGS_REMOTE
git tag "pt-e2e/$VERSION"
git push $BEDROCK_TAGS_REMOTE "pt-e2e/$VERSION"