
#!/bin/bash
echo 'Installing Ubuntu Archive Tools'
git clone https://git.launchpad.net/ubuntu-archive-tools
sudo apt-get install ubuntu-dev-tools -y
cd ubuntu-archive-tools
echo 'Copying Packages'
python3 copy-package -y -b -p nethermindeth --ppa-name=sedge -s jammy --to-suite=focal sedge
python3 copy-package -y -b -p nethermindeth --ppa-name=sedge -s jammy --to-suite=kinetic sedge
python3 copy-package -y -b -p nethermindeth --ppa-name=sedge -s jammy --to-suite=bionic sedge
python3 copy-package -y -b -p nethermindeth --ppa-name=sedge -s jammy --to-suite=trusty sedge
cd ..
echo 'Cleanup'
sudo rm -rf ubuntu-archive-tools
