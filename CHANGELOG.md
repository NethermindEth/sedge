# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Support for Erigon as Execution Client.
- Support for Besu as Execution Client.
- Support for Windows.
- Add search functionality to documentation using Algolia.

### Changed

- Completely replace keystore generation mechanism for Ethereum mainnet.
- Fixed validator restart failure:
  - Validator waits a grace period (2 epochs) before starting.
  - Validator waits for the consensus sync to finish: verifying that the `/eth/v1/node/health` consensus endpoint returns with the HTTP code 200 in its docker-compose `healthcheck` condition so that the validator waits for the consensus be healthy. This replace the older track sync method.
  - Validator service now always restarts unless stopped.

### Fixed

- Fix bad `{{end}}` tag on docker-compose_base template.
- Fix inconsistent behavior of `AssignPorts` function test.

## [v0.5.1] - 2022-12-2

### Fixed

- Fix missing preset flag for Lodestar validator service.
- Support custom configuration for Lighthouse validator-import service.

## [v0.5.0] - 2022-11-29

### Added

- Add checkpoint sync url for Chiado.

### Changed

- Update Gnosis and Chiado networks default clients images to merge ready versions.
- Update client versions.
- Update checkpoint sync url for gnosis.

### Fixed

- Add missing `depends_on` tag to lodestar validator.
- Fix wrong fork version in Gnosis network config.

## [v0.4.0] - 2022-10-25

### Added

- Check for new Sedge releases on GitHub. Sedge will now report if there is a new version released.
- Add `--mev-boost-image` flag to specify which Mev Boost docker image sedge uses.
- Set mev-boost as default on mainnet.
- Documentation for How to Run a Validator Node on Sedge by yourself, once generated.
- Documentation for How to Run a Validator or Full Node on Chiado, Gnosis testnet.
- validator-import service for Lodestar.

### Changed

- Drop support for deprecated networks. (Kiln, Ropsten, Denver).
- Update Chiado testnet configs.
- Sedge uses its own mechanism for keystore generation with `sedge keys` instead of the staking-deposit-cli tool. This
  is experimental and staking-deposit-cli is still being used for mainnet.
- Updated docker image tags for all clients except Teku.

### Fixed

- Print generated files as string instead of []byte.
- keystore_password.txt permissions issue.

## [v0.3.0] - 2022-09-15

### Added

- Add support for gnosis merge testnets named chiado and denver.
- Add support for custom testnets by allowing custom remote config and genesis files.
- Add `--no-validator` flag to exclude the validator node from the full node setup.
- Gnosis Network support.
- Add `--graffiti` to allow overriding the default graffiti used by validators nodes.
- Allow the extra flags to overwrite fixed template images commands flags. Latest flag apparition will be used.
- Add Homebrew installation method.
- Add Checkpoint Sync for Ethereum Mainnet.

### Changed

- Remove Nethermind metrics configuration.
- Prysm and Lodestar can use now an undefined number of Bootnodes.
- Allow providing a number for `--tail` flag of `logs` command, e.g: `sedge logs --tail 50`
- Remove `mainnet` service templates and use `merge` service templates for Mainnet network.
- `checkpoint-sync-url` can now be set separately for every consensus client of a network using `CHECKPOINT_SYNC_URL` as
  env variabe in the consensus client templates.
- Use fixed docker images with latest tested versions.
- Change Prater to Goerli.

### Fixed

- Check for compose plugin installed. The check was running compose with a wrong path flag.
- Errors in validators templates.

## [v0.2.0] - 2022-08-4

### Added

- Support for Goerli/Prater network.
- Docusaurus documentation engine with initial version of docs.
- CI workflows to test and deploy documentation to Github Pages.
- Goerli/Prater support.
- Sepolia support.
- Mev-boost settings for all the currently supported CL clients.
- Relays URLs for all the testnets.
- Mev-boost set on by default in Prater.

### Changed

- Update Readme. Some instructions were moved to the docs. Roadmap updated.
- Updated Teku mev-boost configuration.

## [v0.1.2] - 2022-07-21

### Added

- Table with mev-boost support information on the Readme.
- Table with OS support for dependency installation.
- More instructions and descriptions about the tool on the Readme.

### Changed

- Update Roadmap.

### Fixed

- Fix bad checks for compose availability.
- Stop running Lighthouse with mev-boost settings on Ropsten. Official and stable docker image doesn't support mev-boost
  yet.
- Fix compose installation. Now compose is installed on root user, as sedge runs everything using sudo currently.
- Fix track sync of nodes stoping early.

## [v0.1.1] - 2022-07-20

### Changed

- Update Roadmap.

### Fixed

- Fix error when trying to generate jwtsecret on unexisting folder.

## [v0.1.0] - 2022-07-15

### Added

- Create cli tool able to set up a Ethereum based validator in an on-premise way.
- Generate `docker-compose` scripts and `.env` files for selected clients with the `cli` command.
- Generate keystore folder with the cli using `keys` command. The inner tool used for this is
  the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool.
- Generate `jwtsecret` for post-merge networks.
- Test coverage (unit tests).
- Integrate Kiln network.
- Integrate Ropsten network.
- Integrate MEV-Boost as an option.
