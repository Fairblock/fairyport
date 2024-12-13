#!/bin/zsh

WORKING_DIR=$(pwd)
SCRIPTS_DIR="$WORKING_DIR/integration-test"
TEST_OUT_DIR="$WORKING_DIR/integration-test-out"
FAIRYPORT_CONFIG="$SCRIPTS_DIR/fairyport-test-config.yml"
FAIRYRINGCLIENT_CONFIG="$SCRIPTS_DIR/fairyringclient-test-config.yml"
SHAREGENERATION_CONFIG="$SCRIPTS_DIR/sharegeneration-test-config.yml"

echo "Removing all previous data..."
rm -rf $TEST_OUT_DIR &> /dev/null

. "$SCRIPTS_DIR/killall.sh"

echo "Creating integration test directory..."

if ! mkdir -p $TEST_OUT_DIR 2>/dev/null; then
    echo "Failed to create integration test directory,"
fi

if ! command -v anvil &> /dev/null
then
    echo "Foundry not found, please visit"
    echo "https://github.com/foundry-rs/foundry"
    echo "to install foundry before running intergation test script"
    exit 1
fi

echo "Cloning fairyring & installing it now..."
cd $TEST_OUT_DIR
git clone https://github.com/Fairblock/fairyring.git
cd fairyring
go mod tidy
make install
echo "fairyringd binary installed !"
cd $WORKING_DIR

echo "Installing fairyport..."
cd $WORKING_DIR
go mod tidy
go install
echo "fairyport binary installed !"


echo "Cloning FairyringContract..."
cd $TEST_OUT_DIR
git clone https://github.com/Fairblock/FairyringContract.git
cd FairyringContract
echo "FairyringContract cloned and built !"

# Check if all binaries are installed.
if ! command -v anvil &> /dev/null || ! command -v fairyringd &> /dev/null || ! command -v fairyport &> /dev/null
then
    echo "Failed to find all needed binaries."
    exit 1
fi

echo "Found all needed binaries, starting integration test..."
echo "Starting fairyring devnet..."

cd $WORKING_DIR
. "$SCRIPTS_DIR/start_fairyring.sh"

echo "Fairyring devnet started"

echo "Starting local EVM devnet with anvil..."

anvil --block-time 2 --accounts 3 --balance 100000 --config-out anvil_config.json > $TEST_OUT_DIR/evm_devnet.log 2>&1 &
PID=$!
echo "Starting local node in the background, pid: $PID"
sleep 1
pkey=$(cat anvil_config.json | jq -r '.private_keys[0]')
addr=$(cast wallet address --private-key $pkey)
rm anvil_config.json
echo "Deploying Fairyring Contract on the local EVM devnet..."
cd $TEST_OUT_DIR/FairyringContract

forge install

echo "Done installing Smart Contract dependencies..."
sleep 5

deploy_out=$(forge create src/FairyringContract.sol:FairyringContract --private-key $pkey --json | jq)
# Deploy contract
frc=$(echo $deploy_out | jq -r '.deployedTo')

deploy_out=$(forge create example/NumberGuessing.sol:NumberGuessing --constructor-args 10 "$frc" --private-key $pkey --json | jq)
ng=$(echo $deploy_out | jq -r '.deployedTo')

echo "Fairyring Contract deployed at: $frc"
echo "Number Guessing Contract deployed at: $ng"
echo "Setting up the contract..."

cd $WORKING_DIR

echo "Backupping existing .env to .env.bkup..."
cp .env .env.bkup
rm .env
echo 'COSMOS_MNEMONIC="angry twist harsh drastic left brass behave host shove marriage fall update business leg direct reward object ugly security warm tuna model broccoli choice"' >> .env

FAIRYPORT_HOME=$HOME/.fairyport
if ! mkdir -p $FAIRYPORT_HOME/ 2>/dev/null; then
    echo "Failed to create fairyport folder. Aborting..."
    exit 1
fi
if [ -f $FAIRYPORT_HOME/config.yml ]; then
    echo "Existing fairyport config found!"
    echo "Backing up existing config to config.yml.bkup..."
    cp $FAIRYPORT_HOME/config.yml $FAIRYPORT_HOME/config.yml.bkup
fi
cp $FAIRYPORT_CONFIG "$FAIRYPORT_HOME/config.yml"
sed -i "" "s/0xcA6cC5c1c4Fc025504273FE61fc0E09100B03D98/$frc/g" "${FAIRYPORT_HOME}/config.yml"
fairyport start > "$TEST_OUT_DIR/fairyport.log" 2>&1 &
echo "fairyport started..."


if ! mkdir -p $HOME/.fairyringclient/ 2>/dev/null; then
    echo "Failed to create fairyringclient folder. Aborting..."
    exit 1
fi
if [ -f $HOME/.fairyringclient/config.yml ]; then
    echo "Existing fairyringclient config found!"
    echo "Backing up existing config to config.yml.bkup..."
    cp $HOME/.fairyringclient/config.yml $HOME/.fairyringclient/config.yml.bkup
fi
cp $FAIRYRINGCLIENT_CONFIG $HOME/.fairyringclient/config.yml
fairyringclient start > "$TEST_OUT_DIR/fairyringclient.log" 2>&1 &
echo "fairyringclient started..."
echo "$TEST_OUT_DIR/fairyring/data/fairyring_test_1"

sleep 5

if ! mkdir -p $HOME/.ShareGenerationClient/ 2>/dev/null; then
    echo "Failed to create ShareGenerationClient folder. Aborting..."
    exit 1
fi
if [ -f $HOME/.ShareGenerationClient/config.yml ]; then
    echo "Existing ShareGenerationClient config found!"
    echo "Backing up existing config to config.yml.bkup..."
    cp $HOME/.ShareGenerationClient/config.yml $HOME/.ShareGenerationClient/config.yml.bkup
fi
cp $SHAREGENERATION_CONFIG $HOME/.ShareGenerationClient/config.yml
ShareGenerationClient start > "$TEST_OUT_DIR/ShareGenerationClient.log" 2>&1 &
echo "ShareGenerationClient started..."

cd $WORKING_DIR
./integration-test/start_hermes.sh

echo "hermes relayer started..."

echo "Checking if fairyport relays keys to Cosmos chain correctly..."
sleep 10 # let the chain run
LATEST_HEIGHT=$(fairyringd q pep latest-height)
if [[ "$LATEST_HEIGHT" == *"error"* ]]; then
    echo "fairyport is not working, expected to be able to query latest-height, got '$LATEST_HEIGHT' instead"
    exit 1
fi
echo "Confirmed fairyport works with only Cosmos chain, Got '$LATEST_HEIGHT' latest height from cosmos chain"
echo ""

echo "Start testing fairyport with both Cosmos and EVM chain"
# Test fairyport with EVM:
pkey=${pkey#"0x"}
echo "EVM_PKEY=$pkey" >> .env

echo "Restarting fairyport with EVM key..."
pkill fairyport
fairyport start > "$TEST_OUT_DIR/fairyport.log" 2>&1 &
echo "fairyport restarted."
echo ""

echo "Checking if fairyport relay keys to EVM chain correctly..."
sleep 20
RAND=$(cast call $frc "latestRandomnessHashOnly()(bytes32)")
if [[ "$RAND" == "0x0000000000000000000000000000000000000000000000000000000000000000" ]]; then
    echo "fairyport is not working properly for EVM chain, expected to be able to query non 0 randomness, got '$RAND' instead"
    exit 1
fi
echo "Confirmed fairyport relay keys to EVM FairyringContract, Got $RAND from Contract"
echo ""

echo "Checking if fairyport relays keys to Cosmos chain correctly..."
LATEST_HEIGHT=$(fairyringd q pep latest-height)
if [[ "$LATEST_HEIGHT" == *"error"* ]]; then
    echo "fairyport is not working, expected to be able to query latest-height, got '$LATEST_HEIGHT' instead"
    exit 1
fi
echo "Confirmed fairyport works with only Cosmos chain, Got '$LATEST_HEIGHT' latest height from cosmos chain"
echo ""

echo "fairyport is working properly, cleaning up now..."
./integration-test/killall.sh
./integration-test/cleanup.sh
echo "Done integration test for fairyport."