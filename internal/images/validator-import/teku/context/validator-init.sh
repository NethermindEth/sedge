#!/bin/bash
set -e

if [[ "$(ls /data/keys)" ]]; then
    echo "Keystore directory is not empty. Skipping." && exit 0
else
    mkdir -p /data/keys /data/passwords
    cd /keystore/validator_keys
    for key in *; do
        FILENAME=`echo ${key} | sed 's/.json//g'`
        cp ${key} "/data/keys/${FILENAME}.json"
        cp ../keystore_password.txt "/data/passwords/${FILENAME}.txt"
        echo "Copying ${key}"
    done
fi

# ensure teku access for new keys
chmod -R 777 /data
