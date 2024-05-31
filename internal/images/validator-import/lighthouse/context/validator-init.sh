#!/bin/bash
set -e

for key in /keystore/validator_keys/*; do
  if [ -f "$key" ]; then
    echo "Found validator key in $key"
    echo "Importing validator..."
    if [ $VAL_NETWORK = "custom" ]; then
      ./lighthouse account validator import --testnet-dir /network_config --password-file /keystore/keystore_password.txt --keystore $key --reuse-password --datadir /data  
    else
      ./lighthouse account validator import --network $VAL_NETWORK --password-file /keystore/keystore_password.txt --keystore $key --reuse-password --datadir /data
    fi
    echo "Validator imported"
  fi
done
