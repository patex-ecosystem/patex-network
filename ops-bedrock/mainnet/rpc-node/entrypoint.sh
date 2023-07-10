#!/bin/sh
set -exu

VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=/datadir
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GETH_KEYSTORE_DIR="$GETH_DATA_DIR/keystore"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
GETH_SNAPSHOT_FILE_PATH="/mainnet.tar"

CHAIN_ID=789
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"

if [[  -f "$GETH_SNAPSHOT_FILE_PATH" ]] && [[ ! -d "$GETH_KEYSTORE_DIR" ]]; then
echo "$GETH_SNAPSHOT_FILE_PATH snapshot available, processing..."
tar xvf "$GETH_SNAPSHOT_FILE_PATH" -C ./ ;
fi

if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
	echo "$GETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	geth --verbosity="$VERBOSITY" init \
		--datadir="$GETH_DATA_DIR" \
		"$GENESIS_FILE_PATH"
else
	echo "$GETH_CHAINDATA_DIR exists."
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the mainnet.

exec geth \
       --ws \
       --ws.port="$WS_PORT" \
       --ws.addr=0.0.0.0 \
       --ws.origins="*" \
       --http \
       --http.port="$RPC_PORT" \
       --http.addr=0.0.0.0 \
       --http.vhosts="*" \
       --http.corsdomain="*" \
       --authrpc.addr="0.0.0.0" \
       --authrpc.port="8551" \
       --authrpc.vhosts="*" \
       --verbosity=3 \
       --rollup.disabletxpoolgossip=true \
       --nodiscover \
       --syncmode=full \
       --gcmode=archive \
       --maxpeers=0 \
       --datadir="$GETH_DATA_DIR" \
       --txlookuplimit=0 \
       --rollup.sequencerhttp="https://mainnet.patex.io:8545" \
       "$@"
