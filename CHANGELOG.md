# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Support for Goerli/Prater network
- Docusaurus documentation engine with initial version of docs
- CI workflows to test and deploy documentation to Github Pages
- Goerli/Prater support
- Sepolia support
- Mev-boost settings for all the currently supported CL clients
- Relays URLs for all the testnets
- Mev-boost set on by default in Prater

### Changed
- Update Readme. Some instructions were moved to the docs. Roadmap updated.
- Updated Teku mev-boost configuration

## [0.1.2] - 2022-07-21

### Added
- Table with mev-boost support information on the Readme
- Table with OS support for dependency installation
- More instructions and descriptions about the tool on the Readme

### Changed 
- Update Roadmap.

### Fixed
- Fix bad checks for compose availability.
- Stop running Lighthouse with mev-boost settings on Ropsten. Official and stable docker image doesn't support mev-boost yet.
- Fix compose installation. Now compose is installed on root user, as sedge runs everything using sudo currently.
- Fix track sync of nodes stoping early.

## [0.1.1] - 2022-07-20
### Changed
- Update Roadmap.

### Fixed
- Fix error when trying to generate jwtsecret on unexisting folder.

## [0.1.0] - 2022-07-15
### Added
- Create cli tool able to set up a Ethereum based validator in an on-premise way.
- Generate `docker-compose` scripts and `.env` files for selected clients with the `cli` command.
- Generate keystore folder with the cli using `keys` command. The inner tool used for this is the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool.
- Generate `jwtsecret` for post-merge networks.
- Test coverage (unit tests).
- Integrate Kiln network.
- Integrate Ropsten network.
- Integrate MEV-Boost as an option.
