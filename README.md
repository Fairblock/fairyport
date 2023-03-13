# **fairyport**

FairyPort is an off-chain service for fetching aggregated keyshares to Destination Chains from FairyRing.

## **Why FairyPort**

To decrypt encrypted transactions on the destination chain, the corresponding aggregated keyshares have to be fetched from the FairyRing Chain to the Destination chain. The straightforward solution of using ICQ (Interchain Query) does not meet the requirements since the verification of blocks on both chains takes up so a lot of blocks, effectively delaying the process by multiple blocks. This opens up avenues for frontruning transactions on the Destination Chain since there is a large time gap between Keyshars being published on the FairyRing Chain and the same being registered on the Destination chain.

To solve this issue, FairyPort is introduced

## **Working Principle**

FairyPort is a fairly simple service. It consists of two main components:

1. A websocket client listening for events on the Fairyring Chain
2. A mechanism to make transactions to the Destination Chain

The service constantly listens for Keyshare Aggregation events on the Fairyring Chain. Whenever such a event is recorded in a block, FairyPort makes a transaction to the Destination chain to register the Keyshares. This process completely ignores anykind of block verfication. This is because the keyshare being transacted to the Destination Chain can easily be verified on chain and without any further interchain communication. This prevents registration of false or malicious keyshares while also verifying the origin of the KeyShares.

---

## **Installation and Running**

To install FairyPort, simply clone the repo and run `go install` in the root directory. This will install the `fairyport` binary.

To run Fairyport, first make sure you have the correct configuration (discussed in the next section). Both the FairyRing chain and the Destination Chain should be running for FairyPort to function.

To start the service simply run `fairyport start`.

## **Configuration**

The following configuration options are available:

- FairyRingNode
- DestinationNode
- Mnemonic

### **FairyRing Node**

| Option    | Description                                                                      |
|-----------|----------------------------------------------------------------------------------|
| ip        | The IP address that Fairy Ring Node will use.                                    |
| port      | The port that Fairy Ring Node will use for TendermintRPC                         |
| protocol  | The protocol used for communication via the TendermintRPC                        |
| grpcport  | The port that Fairy Ring Node will use for gRPC communication with other nodes.  |

### **Destination Node**

| Option    | Description                                                                      |
|-----------|----------------------------------------------------------------------------------|
| ip        | The IP address that Destination Node will use.                                   |
| port      | The port that Destination Node will use for TendermintRPC                        |
| protocol  | The protocol used for communication via the TendermintRPC                        |
| grpcport  | The port that Destination Node will use for gRPC communication with other nodes. |

### **Mnemonic**

|  Option   | Description                                                                      |
|-----------|----------------------------------------------------------------------------------|
| Mnemonic  | The seed phrase used to generate the private key for the account responsible for making transactions to the Destination Chain|
