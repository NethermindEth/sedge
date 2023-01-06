
#!/bin/bash
echo 'Installing Ubuntu Archive Tools'
git clone https://git.launchpad.net/ubuntu-archive-tools
sudo apt-get install ubuntu-dev-tools -y
cd ubuntu-archive-tools
echo 'Copying Packages'
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=focal sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=kinetic sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=bionic sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=trusty sedge --dry-run
cd ..
echo 'Cleanup'
sudo rm -rf ubuntu-archive-tools
