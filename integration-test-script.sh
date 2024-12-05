#!/bin/zsh

WORKING_DIR=$(pwd)
TEST_OUT_DIR="$WORKING_DIR/integration-test-dir"
FAIRYPORT_CONFIG=integration-test-config.yml

echo "Removing all previous data..."
rm -rf $TEST_OUT_DIR &> /dev/null

echo "Terminating any running anvil instances..."
killall anvil

echo "Terminating any running fairyringd instances..."
killall fairyringd

echo "Terminating any running fairyport instances..."
killall fairyport

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

if ! command -v fairyport &> /dev/null
then
    echo "The fairyport binary is not installed."
    echo "Installing it now..."
    cd $WORKING_DIR
    go mod tidy
    go install
    echo "fairyport binary installed !"
fi


echo "Cloning FairyringContract..."
cd $TEST_OUT_DIR
git clone https://github.com/Fairblock/FairyringContract.git
cd FairyringContract
echo "FairyringContract cloned and built !"

# Check if all binaries are installed.
if ! command -v anvil &> /dev/null || ! command -v fairyringd &> /dev/null || ! command -v fairyport &> /dev/null &> /dev/null
then
    echo "Failed to find all needed binaries."
    exit 1
fi

echo "Found all needed binaries, starting integration test..."
echo "Starting fairyring devnet..."

cd $TEST_OUT_DIR
CHAIN_DIR=$TEST_OUT_DIR
CHAINID=fairyring_devnet
BINARY=fairyringd

VAL_MNEMONIC_1="clock post desk civil pottery foster expand merit dash seminar song memory figure uniform spice circle try happy obvious trash crime hybrid hood cushion"

WALLET_MNEMONIC_1="banner spread envelope side kite person disagree path silver will brother under couch edit food venture squirrel civil budget number acquire point work mass"
WALLET_MNEMONIC_2="veteran try aware erosion drink dance decade comic dawn museum release episode original list ability owner size tuition surface ceiling depth seminar capable only"
RLY_MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAINID ..."
$BINARY init devnet --home $CHAIN_DIR/$CHAINID --default-denom ufairy --chain-id=$CHAINID &> /dev/null

echo "Adding genesis accounts..."
echo $VAL_MNEMONIC_1 | $BINARY keys add val1 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend test
echo $WALLET_MNEMONIC_1 | $BINARY keys add wallet1 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend test
echo $WALLET_MNEMONIC_2 | $BINARY keys add wallet2 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend test
echo $RLY_MNEMONIC_1 | $BINARY keys add rly1 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend test

VAL1_ADDR=$($BINARY keys show val1 --home $CHAIN_DIR/$CHAINID -a --keyring-backend test)
WALLET1_ADDR=$($BINARY keys show wallet1 --home $CHAIN_DIR/$CHAINID -a --keyring-backend test)
WALLET2_ADDR=$($BINARY keys show wallet2 --home $CHAIN_DIR/$CHAINID -a --keyring-backend test)
RLY1_ADDR=$($BINARY keys show rly1 --home $CHAIN_DIR/$CHAINID -a --keyring-backend test)

$BINARY genesis add-genesis-account $VAL1_ADDR 1000000000000ufairy --home $CHAIN_DIR/$CHAINID --keyring-backend test
$BINARY genesis add-genesis-account $WALLET1_ADDR 1000000000000ufairy --home $CHAIN_DIR/$CHAINID --keyring-backend test
$BINARY genesis add-genesis-account $WALLET2_ADDR 1000000000000ufairy --home $CHAIN_DIR/$CHAINID --keyring-backend test
$BINARY genesis add-genesis-account $RLY1_ADDR 1000000000000ufairy --home $CHAIN_DIR/$CHAINID --keyring-backend test

echo "Creating and collecting gentx..."
$BINARY genesis gentx val1 100000000000ufairy --home $CHAIN_DIR/$CHAINID --chain-id $CHAINID --keyring-backend test
$BINARY genesis collect-gentxs --home $CHAIN_DIR/$CHAINID &> /dev/null
echo "Changing defaults and ports in app.toml and config.toml files..."

sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/cors = false/cors = true/g' $CHAIN_DIR/$CHAINID/config/app.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID/config/app.toml
sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = "0ufairy"/g' $CHAIN_DIR/$CHAINID/config/app.toml

sed -i -e 's/"max_deposit_period": "172800s"/"max_deposit_period": "10s"/g' $CHAIN_DIR/$CHAINID/config/genesis.json
sed -i -e 's/"voting_period": "172800s"/"voting_period": "10s"/g' $CHAIN_DIR/$CHAINID/config/genesis.json
sed -i -e 's/"reward_delay_time": "604800s"/"reward_delay_time": "0s"/g' $CHAIN_DIR/$CHAINID/config/genesis.json

sed -i -e 's/"trusted_addresses": \[\]/"trusted_addresses": \["'"$VAL1_ADDR"'","'"$RLY1_ADDR"'"\]/g' $CHAIN_DIR/$CHAINID/config/genesis.json
TRUSTED_PARTIES='{"client_id": "07-tendermint-0", "connection_id": "connection-0", "channel_id": "channel-0"}'

sed -i -e 's/"trusted_counter_parties": \[\]/"trusted_counter_parties": \['"$TRUSTED_PARTIES"'\]/g' $CHAIN_DIR/$CHAINID/config/genesis.json
sed -i -e 's/"key_expiry": "100"/"key_expiry": "50"/g' $CHAIN_DIR/$CHAINID/config/genesis.json

echo "Starting $CHAINID in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/$CHAINID.log"
$BINARY start --log_level info --log_format json --home $CHAIN_DIR/$CHAINID --pruning=nothing > $CHAIN_DIR/$CHAINID.log 2>&1 &

echo "Fairyring devnet started"

echo "Starting local EVM devnet with anvil..."

anvil --accounts 3 --balance 100000 --config-out anvil_config.json > evm_devnet.log 2>&1 &
PID=$!
echo "Starting local node in the background, pid: $PID"
sleep 1
pkey=$(cat anvil_config.json | jq -r '.private_keys[0]')
addr=$(cast wallet address --private-key $pkey)

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
echo 'COSMOS_MNEMONIC="banner spread envelope side kite person disagree path silver will brother under couch edit food venture squirrel civil budget number acquire point work mass"' >> .env
echo "EVM_PKEY=$pkey" >> .env

sed -i -e 's#"0xcA6cC5c1c4Fc025504273FE61fc0E09100B03D98"#"'"$frc"'"#g' $FAIRYPORT_CONFIG

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $HOME/.fairyport/ 2>/dev/null; then
    echo "Failed to create fairyport folder. Aborting..."
    exit 1
fi

# check if config exists, backup and copy new config if it does

if [ -f $HOME/.fairyport/config.yml ]; then
    echo "Existing fairyport config found!"
    echo "Backing up existing config to config.yml.bkup..."
    cp $HOME/.fairyport/config.yml $HOME/.fairyport/config.yml.bkup
fi

cp $FAIRYPORT_CONFIG $HOME/.fairyport/config.yml
fairyport start > "$TEST_OUT_DIR/fairyport.log" 2>&1 &
echo "Fairyport started..."