## TMY

Go-Implementation of TMY Network.

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
