
#!/bin/bash
echo 'Installing Ubuntu Archive Tools'
git clone https://git.launchpad.net/ubuntu-archive-tools
sudo apt-get install ubuntu-dev-tools -y
# echo 'Installing Pip'
# sudo apt install python3-pip  -y
cd ubuntu-archive-tools
echo 'Installing LaunchpadLib'
pip install launchpadlib
echo 'Upgrading LaunchpadLib to the latest version'
pip install launchpadlib --upgrade
echo 'Copying Packages'
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=focal sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=kinetic sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=bionic sedge --dry-run
python3 copy-package -y -b -p nethermind --ppa-name=nethermind -s jammy --to-suite=trusty sedge --dry-run
cd ..
echo 'Cleanup'
sudo rm -rf ubuntu-archive-tools
