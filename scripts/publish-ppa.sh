#!/bin/bash
#exit when any command fails
set -e
cd /home/runner/work/sedge/sedge/sedge
go install github.com/golang/mock/mockgen@v1.6.0
go generate ./...
mkdir -p build/package/debian/src/github.com/NethermindEth/sedge/
rsync -aq . build/package/debian/src/github.com/NethermindEth/sedge/ --exclude build/ --exclude .git/ --exclude docs/
cd build/package/debian/src/github.com/NethermindEth/sedge/ && go mod vendor
cd /home/runner/work/sedge/sedge/sedge

export SVERSION=${VERSION#v}
echo "sedge ($SVERSION) jammy; urgency=medium

  * Sedge ($SVERSION release)

 -- Nethermind <devops@nethermind.io>  $( date -R )" > /home/runner/work/sedge/sedge/sedge/build/package/debian/debian/changelog

cd build/package/debian
debuild -S -uc -us
cd ..
echo 'Signing package'
debsign -p 'gpg --batch --yes --no-tty --pinentry-mode loopback --passphrase-file /tmp/PASSPHRASE' -S -k$PPA_GPG_KEYID sedge_${SVERSION}_source.changes
echo 'Uploading'
dput -f ppa:nethermindeth/sedge sedge_${SVERSION}_source.changes
echo "Publishing Sedge to PPA complete"
echo 'Cleanup'
rm -r sedge_$SVERSION*
