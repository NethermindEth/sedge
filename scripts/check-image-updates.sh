cd templates/envs/mainnet/consensus/
LIGHTHOUSE_CURR=$(cat lighthouse.tmpl | grep -o -P '(?<=:).*?(?={)')
LODESTAR_CURR=$(cat lodestar.tmpl | grep -o -P '(?<=:).*?(?={)')
PRYSM_CURR=$(cat prysm.tmpl | grep -o -P '(?<=:).*?(?={)')
TEKU_CURR=$(cat teku.tmpl | grep -o -P '(?<=:).*?(?={)')
cd ../execution
NETH_CURR=$(cat nethermind.tmpl | grep -o -P '(?<=:).*?(?={)')
GETH_CURR=$(cat geth.tmpl | grep -o -P '(?<=:).*?(?={)')
cd ../../gnosis/
PRYSM_VAL_CURR=$(cat validator/prysm.tmpl | grep -o -P '(?<=:).*?(?={)')
PRYSM_BCN_CURR=$(cat consensus/prysm.tmpl | grep -o -P '(?<=:).*?(?={)')
LIGHTHOUSE_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/sigp/lighthouse/releases/latest | jq -r ".tag_name")
LODESTAR_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ChainSafe/lodestar/releases/latest | jq -r ".tag_name")
PRYSM_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/prysmaticlabs/prysm/releases/latest | jq -r ".tag_name")
PRYSM_BEACON_LATEST=$(curl -H "Authorization: Bearer $PAT" https://api.github.com/orgs/gnosischain/packages/container/gbc-prysm-beacon-chain/versions | jq -r '.[0].metadata.container.tags[0]')
PRYSM_VALIDATOR_LATEST=$(curl -H "Authorization: Bearer $PAT" https://api.github.com/orgs/gnosischain/packages/container/gbc-prysm-validator/versions | jq -r '.[0].metadata.container.tags[0]')
TEKU_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ConsenSys/teku/releases/latest | jq -r ".tag_name")
NETH_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/NethermindEth/nethermind/releases/latest | jq -r ".tag_name")
GETH_LATEST=$(curl -H "Authorization: Bearer $PAT" -sL https://api.github.com/repos/ethereum/go-ethereum/releases/latest | jq -r ".tag_name")
cd ..

if [[ $LIGHTHOUSE_CURR < $LIGHTHOUSE_LATEST ]]; then
    echo "New version of Lighthouse is available. Current version: $LIGHTHOUSE_CURR, new version: $LIGHTHOUSE_LATEST"
    for i in '**/**/lighthouse.tmpl'; do  
    sed -i "s/$LIGHTHOUSE_CURR/$LIGHTHOUSE_LATEST/g" $i; 
    done
fi
if [[ $LODESTAR_CURR < $LODESTAR_LATEST ]]; then
    echo "New version of Lodestar is available. Current version: $LODESTAR_CURR, new version: $LODESTAR_LATEST"
    for i in '**/**/lodestar.tmpl'; do  
    sed -i "s/$LODESTAR_CURR/$LODESTAR_LATEST/g" $i; 
    done
fi
if [[ $PRYSM_CURR < $PRYSM_LATEST ]]; then
    echo "New version of Prysm is available. Current version: $PRYSM_CURR, new version: $PRYSM_LATEST"
    for i in '**/**/prysm.tmpl'; do  
    sed -i "s/$PRYSM_CURR/$PRYSM_LATEST/g" $i;
    done
fi
if [[ $PRYSM_BCN_CURR < $PRYSM_BEACON_LATEST ]]; then
    echo "New version of Prysm beacon is available. Current version: $PRYSM_BCN_CURR, new version: $PRYSM_BEACON_LATEST"
    for i in '**/**/prysm.tmpl'; do  
    sed -i "s/$PRYSM_BCN_CURR/$PRYSM_BEACON_LATEST/g" $i;
    done
fi
if [[ $PRYSM_VAL_CURR < $PRYSM_VALIDATOR_LATEST ]]; then
    echo "New version of Prysm validator is available. Current version: $PRYSM_VAL_CURR, new version: $PRYSM_VALIDATOR_LATEST"
    for i in '**/**/prysm.tmpl'; do  
    sed -i "s/$PRYSM_VAL_CURR/$PRYSM_VALIDATOR_LATEST/g" $i;
    done
fi
if [[ $TEKU_CURR < $TEKU_LATEST ]]; then
    echo "New version of Teku is available. Current version: $TEKU_CURR, new version: $TEKU_LATEST"
    for i in '**/**/teku.tmpl'; do  
    sed -i "s/$TEKU_CURR/$TEKU_LATEST/g" $i; 
    done
fi
if [[ $NETH_CURR < $NETH_LATEST ]]; then
    echo "New version of Nethermind is available. Current version: $NETH_CURR, new version: $NETH_LATEST"
    for i in '**/**/nethermind.tmpl'; do  
    sed -i "s/$NETH_CURR/$NETH_LATEST/g" $i; 
    done
fi
if [[ $GETH_CURR < $GETH_LATEST ]]; then
    echo "New version of Geth is available. Current version: $GETH_CURR, new version: $GETH_LATEST"
    for i in '**/**/geth.tmpl'; do  
    sed -i "s/$GETH_CURR/$GETH_LATEST/g" $i; 
    done
fi