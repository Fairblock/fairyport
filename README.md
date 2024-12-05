# **fairyport**

FairyPort is an off-chain service for fetching aggregated keyshares to Destination Chains from FairyRing.

## **Why FairyPort**

To decrypt encrypted transactions on the destination chain, the corresponding aggregated keyshares have to be fetched from the FairyRing Chain to the Destination chain. The straightforward solution of using ICQ (Interchain Query) does not meet the requirements since the verification of blocks on both chains takes up so a lot of blocks, effectively delaying the process by multiple blocks. This opens up avenues for front-running transactions on the Destination Chain since there is a large time gap between Keyshars being published on the FairyRing Chain and the same being registered on the Destination chain.

To solve this issue, FairyPort is introduced

## **Working Principle**

FairyPort is a fairly simple service. It consists of two main components:

1. A websocket client listening for events on the Fairyring Chain
2. A mechanism to make transactions to the Destination Chain

The service constantly listens for Keyshare Aggregation events on the Fairyring Chain. Whenever such a event is recorded in a block, FairyPort makes a transaction to the Destination chain to register the Keyshares. This process completely ignores any kind of block verification. This is because the keyshare being transacted to the Destination Chain can easily be verified on chain and without any further interchain communication. This prevents registration of false or malicious keyshares while also verifying the origin of the KeyShares.

---

## **Installation and Running**

To install FairyPort, simply clone the repo and run `make install` in the root directory. This will install the `fairyport` binary.

To run Fairyport, first make sure you have the correct configuration (discussed in the next section),

You can initialize the default config by `fairyport init` 

Both the FairyRing chain and the Destination Chain should be running for FairyPort to function properly.

To start the service simply run `fairyport start`.

## **Configuration**

The following configuration options are available:

- CosmosRelayConfig
  - Destination Node
  - Metrics Port
  - Derive Path
- EVMRelayTarget
  - Chain RPC
  - Contract Address
- FairyringNodeWS
  - IP
  - Port
  - Protocol

`EVM_PKEY` environment variable is required when relaying to EVM chain.

`COSMOS_MNEMONIC` environment variable is required when relaying to Cosmos chain.

### **Cosmos Relay Config**

| Option           | Description                                                                |
|------------------|----------------------------------------------------------------------------|
| Destination Node | The Destination Cosmos Node to relay to.                                   |
| Metrics Port     | The port that lets prometheus collect metrics                              |
| Derive Path      | The path Fairyport uses to derive the private key from the mnemonic phase. |

### **Destination Node**

| Option         | Description                                                      |
|----------------|------------------------------------------------------------------|
| IP             | The IP address that Destination Cosmos Node.                     |
| Port           | The port that Destination Cosmos Node will use for TendermintRPC |
| Protocol       | The protocol used for communication via the TendermintRPC        |
| gRPC Port      | The port that Destination Node will use for gRPC communication.  |
| Chain ID       | The chain id of the destination cosmos chain.                    |
| Account Prefix | The account prefix of the Destination Cosmos Chain.              |

### **EVM Relay Target**

| Option             | Description                                             |
|--------------------|---------------------------------------------------------|
| Chain RPC          | The WS endpoint to the Destination EVM chain            |
| Contract Address   | The Fairyring Contract address on Destination EVM chain |

### **Fairyring Node WS**

| Option   | Description                                      |
|----------|--------------------------------------------------|
| IP       | The IP address of given Fairyring node endpoint. |
| Port     | The Port of the given Fairyring node endpoint.   |
| Protocol | The Protocol of Fairyring node endpoint.         | 

### **EVM_PKEY** & **COSMOS_MNEMONIC** Environment variable 

1. Pass the environemnt variable to `fairyport` when starting it

```bash
EVM_PKEY=your_hex_private_key fairyport start --config $HOME/.fairyport/config.yml
```

```bash
COSMOS_MNEMONIC="mnemonic phase" fairyport start --config $HOME/.fairyport/config.yml
```

2. Set the environment before running `fairyport`

```bash
export EVM_PKEY=your_hex_private_key
fairyport start --config $HOME/.fairyport/config.yml
```

```bash
export COSMOS_MNEMONIC="mnemonic phase"
fairyport start --config $HOME/.fairyport/config.yml
```

3. Create a `.env` file in the same directory you are running `fairyport`

```bash
EVM_PKEY=your_hex_private_key
COSMOS_MNEMONIC="mnemonic phase"
```

then run `fairyport start --config $HOME/.fairyport/config.yml`