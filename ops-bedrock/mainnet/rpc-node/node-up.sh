#!/usr/bin/env bash

# This script starts a local mainnet using Docker Compose. We have to use
# this more complicated Bash script rather than Compose's native orchestration
# tooling because we need to start each service in a specific order, and specify
# their configuration along the way. The order is:
#
# 4. Start L2
# 5. Get the genesis
# 6. Generate the rollup driver's config using the genesis hashes and the
#    timestamps recovered in step 4 as well as the address of the PatexPortal
#    contract deployed in step 3.
# 7. Start the rollup driver.
# 8. Start the L2 output submitter.
#
# The timestamps are critically important here, since the rollup driver will fill in
# empty blocks if the tip of L1 lags behind the current timestamp. This can lead to
# a perceived infinite loop. To get around this, we set the timestamp to the current
# time in this script.
#
# This script is safe to run multiple times. It stores state in `.mainnet`, and
# contracts-bedrock/deployments/mainnetL1.
#
# Don't run this script directly. Run it using the makefile, e.g. `make mainnet-up`.
# To clean up your mainnet, run `make mainnet-clean`.

set -eu

NETWORK=mainnet
MAINNET="$PWD/.mainnet"
L2_URL="http://localhost:9545"

PT_GETH_GENESIS_URL="https://mainnet.patex.io/genesis.json"
PT_GETH_SNAPSHOT_URL="https://mainnet.patex.io/snapshots/mainnet.tar"

# Helper method that waits for a given URL to be up. Can't use
# cURL's built-in retry logic because connection reset errors
# are ignored unless you're using a very recent version of cURL
function wait_up {
  echo -n "Waiting for $1 to come up..."
  i=0
  until curl -s -f -o /dev/null "$1"
  do
    echo -n .
    sleep 0.25

    ((i=i+1))
    if [ "$i" -eq 300 ]; then
      echo " Timeout!" >&2
      exit 1
    fi
  done
  echo "Done!"
}

mkdir -p ./.mainnet

# Download genesis file if not exists
if [ ! -f "$MAINNET/genesis.json" ]; then
  wget -O "$MAINNET"/genesis.json "$PT_GETH_GENESIS_URL"
fi
# Download snapshot file if not exists
if [ ! -f "$MAINNET/mainnet.tar" ]; then
  wget -O "$MAINNET"/mainnet.tar "$PT_GETH_SNAPSHOT_URL"
fi

# Generate jwt if not exists
if [ ! -f "$MAINNET/jwt.txt" ]; then
  openssl rand -hex 32 > "$MAINNET"/jwt.txt
fi

# Generate p2p key if not exists
if [ ! -f "$MAINNET/ptnode_p2p_priv.txt" ]; then
  openssl rand -hex 32 > "$MAINNET"/ptnode_p2p_priv.txt
fi

# Bring up L2.
(
  cd ops-bedrock/mainnet/rpc-node
  echo "Bringing up L2..."
  DOCKER_BUILDKIT=1 docker-compose build --progress plain
  docker-compose up -d m_l2
  wait_up $L2_URL
)



# Bring up pt-node
(
  cd ops-bedrock/mainnet/rpc-node
  echo "Bringing up pt-node..."
  docker-compose up -d m_pt-node
)

echo "Patex Mainnet node is ready."
