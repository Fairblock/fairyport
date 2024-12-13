#!/bin/bash

echo ""
echo "###########################################################"
echo "#           Setting up Relayer between two chain          #"
echo "###########################################################"
echo ""


CHAINID_1=fairyring_test_1
CHAINID_2=fairyring_test_2
CONFIG_FILE=$(pwd)/integration-test/hermes_config.toml
CHAIN_DIR=$(pwd)/integration-test-out

# echo "Creating client on both chain..."
# hermes --config $CONFIG_FILE create client --host-chain $CHAINID_1 --reference-chain $CHAINID_2
# hermes --config $CONFIG_FILE create client --host-chain $CHAINID_2 --reference-chain $CHAINID_1

# echo "Creating connection..."
# hermes --config $CONFIG_FILE create connection --a-chain $CHAINID_2 --a-client 07-tendermint-0 --b-client 07-tendermint-0

echo "Creating channel..."
hermes --config $CONFIG_FILE create channel --new-client-connection --a-chain $CHAINID_2 --b-chain $CHAINID_1 --a-port gov --b-port keyshare --channel-version keyshare-1 --yes
hermes --config $CONFIG_FILE create channel --a-chain $CHAINID_2 --a-connection connection-0 --a-port pep --b-port keyshare --channel-version keyshare-1

