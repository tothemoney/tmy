## TMY

### Consensus

The TMY network consensus is based on PoA (Proof of Authority) technology. The algorithm for creating a block is quite simple: every 30,000 blocks a list of 21 active validators is created, this list is called an epoch. This list contains validators for the next 30,000 blocks, which will take turns producing blocks in the blockchain and receiving a reward. Validators are added to this list randomly, from another common list of trusted nodes (authorities). Adding or removing members to the list of trusted nodes is done by voting of all existing members. Therefore, the PoA blockchain is protected by trusted nodes that randomly select a trustworthy one and an algorithm that selects only from a list of trusted nodes, and not from a general list of users, as in PoS.

The Proof of Authority model is based on a limited number of block validators and this is what makes it a scalable system. Blocks and transactions are verified by pre-approved participants who act as system moderators.


### Network bandwidth

Thanks to the architecture of the system based on PoA and the white list of trusted nodes that can create blocks, the blockchain has the ability to produce a block every 5-10 seconds, which is a high result compared to existing systems, thereby increasing network throughput. The concept of "block size" is not applicable for this system, instead the concept of "block processing power" is used. Those. the number of transactions (gas) that smart contracts can perform in one block. In the tmy blockchain, this value for a block is 8,000,000 operations. Thus, the throughput of the system is approximately 380 normal transactions (21,000 gas costs a normal transaction) or 150 contract creations within the system (53,000 gas costs a contract creation) per block. Assuming that a block is generated every 5 seconds, the blockchain throughput is approximately 76 transactions per second. It should be clarified that these calculations do not cover the use of smart contracts, including the transfer of tokens and other operations, since it is not realistic to calculate the exact throughput due to the great variety of different options for smart contracts.

### Smart contracts

The tmy blockchain allows you to create smart contracts on its basis, representing a global repository of the state of various programs. It is possible to compare this with many open applications on the computer screen, which are waiting for commands from the user, but distributed in the blockchain. Commands that change the state of a smart contract in the blockchain (variables inside the contract) are called a call and have their cost (gas) based on the number of operations that are carried out within this call, such calls create a transaction that reflects the change in the state of the blockchain. If the command simply reads the existing data from the smart contract, it is free to execute.


### Emission

The tmy blockchain differs from conventional PoA solutions in that it has a block reward. The emission of the system is set as 2 million coins and is calculated for 10 years:
- 1 year 40 000 tmy
- 2 year - 80 000 tmy
- 3 year - 120 000 tmy
- 4 year - 160 000 tmy
- 5 year - 200,000 tmy
- 6 year - 240,000 tmy
- 7 year - 260,000 tmy
- 8 year - 300,000 tmy
- 9 year - 300 000 tmy
- 10 year - 300,000 tmy

### Compatibility

The tmy blockchain is compatible with all ethereum-like systems, web3 solutions, and smart contracts written in tmy are compatible with ethereum, bsc, tron and other cryptocurrencies, whose evm, like tmy, are based on the solidity language.

### Node commands

#### Accounts
First an account must be created
```shell
geth account new 
```
Inter a password on prompt and *remember it(!)*

With the genesis state defined in the above JSON file, you'll need to initialize **every**
`geth` node with it prior to starting it up to ensure all blockchain parameters are correctly
set:

```shell
geth init path/to/genesis.json
```

Therefore, you should download the genesis file and use the command above to init your data directory.  

#### Static nodes
Since the blockchain is closed to arbitrary nodes, the `static-nodes.json` must be set. Currently it looks like below:
```json
[
	"enode://a693fb88c537954fa16011cae2a12f0c9d5d6c99532b9a3ebb816525323c22d9bcf973834fff37f10e3b42e3e63d2b69eecb27c8f141eb4f12190f2ddb626d7a@185.135.80.216:30303?discport=30304"
]
```

The corresponding file must be in the `<data_dir>/geth/` folder before the node is launched.   
*In case of the `static-nodes.json` usage the `--bootnode` is not needed.*

#### Starting up your signer node
If the points above are successfully made, the signer node is ready to be started:
```shell
geth --networkid 8768 \
  --unlock <address-created-above> \
  --mine \
  --miner.threads 1 \
  --miner.gasprice 1 \
  --syncmode full \
  --miner.etherbase <address-created-above> \
  --http.addr "0.0.0.0" \
  --nodiscover
```
`<address-created-above>` is an address of the account which was created of the first step.  
Then account's password must be entered on prompt.  
*Pay attention on to Firewall settings. Port `30303` must be exposed.*

Also, before signing started, the account must be approved by other signers.
