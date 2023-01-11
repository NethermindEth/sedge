
#!/bin/bash
echo 'Installing Ubuntu Archive Tools'
git clone https://git.launchpad.net/ubuntu-archive-tools
sudo apt-get install ubuntu-dev-tools -y
cd ubuntu-archive-tools
echo 'Copying Packages'
python3 copy-package -y -b -p nethermindeth --ppa-name=nethermind -s jammy --to-suite=focal nethermind --dry-run
python3 copy-package -y -b -p nethermindeth --ppa-name=nethermind -s jammy --to-suite=kinetic nethermind --dry-run
python3 copy-package -y -b -p nethermindeth --ppa-name=nethermind -s jammy --to-suite=bionic nethermind --dry-run
python3 copy-package -y -b -p nethermindeth --ppa-name=nethermind -s jammy --to-suite=trusty nethermind --dry-run
cd ..
echo 'Cleanup'
sudo rm -rf ubuntu-archive-tools
