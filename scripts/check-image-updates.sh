function get-field() {
    # $1 - field name
    cat configs/client_images.yaml | yq "$1"
}

function update-field() {
    # $1 - field name
    # $2 - new value
    echo "Updating $1 to $2 in configs/client_images.yaml"
    yq eval -i "$1 = $2" configs/client_images.yaml
}

function update-client() {
    # $1 client name
    # $2 client type
    # $3 client path
    # $4 latest version

    CURRENT_VERSION=$(get-field "$3.version")
    if [[ $CURRENT_VERSION < $4 ]] ; then
        echo "New version of $1 $2 client is available. Current version: $CURRENT_VERSION, new version: $4"
        update-field "$3.version" "\"$4\""
        echo ""
    fi
}

# Geth
GETH_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ethereum/go-ethereum/releases/latest | jq -r ".tag_name")
update-client "Geth" "execution" ".execution.geth" "$GETH_LATEST_VERSION"

# Besu
BESU_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/hyperledger/besu/releases/latest | jq -r ".tag_name")
update-client "Besu" "execution" ".execution.besu" "$BESU_LATEST_VERSION"

# Netehrmind
NETHERMIND_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/NethermindEth/nethermind/releases/latest | jq -r ".tag_name")
update-client "Nethermind" "execution" ".execution.nethermind" "$NETHERMIND_LATEST_VERSION"

# Erigon
ERIGON_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ledgerwatch/erigon/releases/latest | jq -r ".tag_name")
update-client "Erigon" "execution" ".execution.erigon" "$ERIGON_LATEST"

# Lighthouse
LIGHTHOUSE_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/sigp/lighthouse/releases/latest | jq -r ".tag_name")
update-client "Lighthouse" "beacon-chain" ".beacon-chain.lighthouse" "$LIGHTHOUSE_LATEST_VERSION"
update-client "Lighthouse" "validator" ".validator.lighthouse" "$LIGHTHOUSE_LATEST_VERSION"

# Lodestar
LODESTAR_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ChainSafe/lodestar/releases/latest | jq -r ".tag_name")
update-client "Lodestar" "beacon-chain" ".beacon-chain.lodestar" "$LODESTAR_LATEST_VERSION"
update-client "Lodestar" "validator" ".validator.lodestar" "$LODESTAR_LATEST_VERSION"

# Teku
TEKU_LATEST_VERSION=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ConsenSys/teku/releases/latest | jq -r ".tag_name")
update-client "Teku" "beacon-chain" ".beacon-chain.teku" "$TEKU_LATEST_VERSION"
update-client "Teku" "validator" ".validator.teku" "$TEKU_LATEST_VERSION"

## Prysm
PRYSM_LATEST_BEACON_VERSION=$(curl -H "Authorization: Bearer $PAT" https://api.github.com/orgs/gnosischain/packages/container/gbc-prysm-beacon-chain/versions | jq -r '.[0].metadata.container.tags[0]')
update-client "Prysm" "beacon-chain" ".beacon-chain.prysm" "$PRYSM_LATEST_VERSION"
PRYSM_LATEST_VALIDATOR_VERSION=$(curl -H "Authorization: Bearer $PAT" https://api.github.com/orgs/gnosischain/packages/container/gbc-prysm-validator/versions | jq -r '.[0].metadata.container.tags[0]')
update-client "Prysm" "validator" ".validator.prysm" "$PRYSM_LATEST_VALIDATOR_VERSION"
