# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.1] - 2022-07-15
### Added
- Create cli tool able to set up a Ethereum based validator in an on-premise way.
- Generate `docker-compose` scripts and `.env` files for selected clients with the `cli` command.
- Generate keystore folder with the cli using `keys` command. The inner tool used for this is the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool.
- Generate `jwtsecret` for post-merge networks.
- Test coverage (unit tests)
- Integrate Kiln network
- Integrate Ropsten network
- Integrate MEV-Boost as an option