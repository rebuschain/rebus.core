
KEY="mykey"
CHAINID="reb_9999-1"
MONIKER="node-devnet"
KEYRING="test"
# KEYALGO="eth_secp256k1"
KEYALGO="secp256k1"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon
rm -rf ~/.rebusd*

# make install

rebusd config keyring-backend $KEYRING
rebusd config chain-id $CHAINID

# if $KEY exists it should be deleted
rebusd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Rebus (Moniker can be anything, chain-id must be an integer)
rebusd init $MONIKER --chain-id $CHAINID 

# Set gas limit in genesis
cat $HOME/.rebusd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.rebusd/config/tmp_genesis.json && mv $HOME/.rebusd/config/tmp_genesis.json $HOME/.rebusd/config/genesis.json

# Get close date
node_address=$(rebusd keys list | grep  "address: " | cut -c12-)
current_date=$(date -u +"%Y-%m-%dT%TZ")
cat $HOME/.rebusd/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' > $HOME/.rebusd/config/tmp_genesis.json && mv $HOME/.rebusd/config/tmp_genesis.json $HOME/.rebusd/config/genesis.json



if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.rebusd/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.rebusd/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.rebusd/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.rebusd/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
rebusd add-genesis-account $KEY 800000000000000000000000000arebus --keyring-backend $KEYRING

export ACC1=rebus1dl90xa89ljj29mna8uasdxs3nejdw46x8767tc
export ACC2=rebus1scpjcw2ux95e4rf0m42kldkz3z83hu23hsm4wk
export ACC3=rebus1tpr6r224eq2camacpec32nmsuup7yfa6hpy86j
export ACC4=rebus1u05h23z07y4wvhetrz566hjk5d99ep42ammaem
export ACC5=rebus1x9p6dc0pyjvxwz0cptp5aqv5u9mhhl6egh2glh
export ACC6=rebus1t5hfzsn8h6vzcrf3dkwksasgkceg2ejthrqh3r
export ACC7=rebus17eryqu4vsvp6065tz23vfdgpdxztv4ujuyk3jv
export ACC8=rebus1p8sw297thl5jxq9dxetzpsnxeps3daqyh3723j

rebusd add-genesis-account $ACC1 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC2 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC3 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC4 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC5 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC6 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd add-genesis-account $ACC7 25000000000000000000000000arebus --keyring-backend $KEYRING
# rebusd add-genesis-account $ACC8 25000000000000000000000000arebus --keyring-backend $KEYRING
rebusd  add-genesis-account $ACC8  25000000000000000000000000arebus --vesting-amount 25000000000000000000000000arebus --vesting-start-time 1652816751 --vesting-end-time 1666028997

# Update total supply with claim values
validators_supply=$(cat $HOME/.rebusd/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=1000000000000000000000000000
# cat $HOME/.rebusd/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.rebusd/config/tmp_genesis.json && mv $HOME/.rebusd/config/tmp_genesis.json $HOME/.rebusd/config/genesis.json

# Sign genesis transaction
rebusd gentx $KEY 1000000000000000000000arebus --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
rebusd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
rebusd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
# rebusd start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001arebus --json-rpc.api eth,txpool,personal,net,debug,web3
