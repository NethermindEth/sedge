#!/bin/bash
#exit when any command fails
set -e
cd $GITHUB_WORKSPACE
mkdir -p build/package/debian/src/github.com/NethermindEth/sedge/
rsync -aq . build/package/debian/src/github.com/NethermindEth/sedge/ --exclude build/ --exclude .git/ --exclude docs/ --exclude scripts/
cd build/package/debian/src/github.com/NethermindEth/sedge/ && go mod vendor
cd $GITHUB_WORKSPACE

echo "sedge ($VERSION) jammy; urgency=medium

  * Sedge ($VERSION release)

 -- Nethermind <devops@nethermind.io>  $( date -R )" > $GITHUB_WORKSPACE/build/package/debian/debian/changelog

cd build/package/debian
debuild -S -uc -us
cd ..
echo 'Signing package'
debsign -p 'gpg --batch --yes --no-tty --pinentry-mode loopback --passphrase-file /tmp/PASSPHRASE' -S -k$PPA_GPG_KEYID sedge_${VERSION}_source.changes
echo 'Uploading'
dput -s ppa:nethermindeth/sedge sedge_${VERSION}_source.changes
echo "Publishing Sedge to PPA complete"
echo 'Cleanup'
rm -r sedge_$VERSION*
