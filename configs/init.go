/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package configs

var networksConfigs map[string]NetworkConfig = map[string]NetworkConfig{
	NetworkMainnet: {
		Name:               NetworkMainnet,
		NetworkService:     "merge",
		GenesisForkVersion: "0x00000000",
		SupportsMEVBoost:   true,
		CheckpointSyncURL:  "https://beaconstate.ethstaker.cc",
		RelayURLs: []string{
			"https://0xa7ab7a996c8584251c8f925da3170bdfd6ebc75d50f5ddc4050a6fdc77f2a3b5fce2cc750d0865e05d7228af97d69561@agnostic-relay.net",
			"https://0x9000009807ed12c1f08bf4e81c6da3ba8e3fc3d953898ce0102433094e5f22f21102ec057841fcb81978ed1ea0fa8246@builder-relay-mainnet.blocknative.com",
			"https://0xad0a8bb54565c2211cee576363f3a347089d2f07cf72679d16911d740262694cadb62d7fd7483f27afd714ca0f1b9118@bloxroute.ethical.blxrbdn.com",
			"https://0x8b5d2e73e2a3a55c6c87b8b6eb92e0149a125c852751db1422fa951e42a09b82c142c3ea98d0d9930b056a3bc9896b8f@bloxroute.max-profit.blxrbdn.com",
			"https://0xb0b07cd0abef743db4260b0ed50619cf6ad4d82064cb4fbec9d3ec530f7c5e6793d9f286c4e082c0244ffb9f2658fe88@bloxroute.regulated.blxrbdn.com",
			"https://0xb3ee7afcf27f1f1259ac1787876318c6584ee353097a50ed84f51a1f21a323b3736f271a895c7ce918c038e4265918be@relay.edennetwork.io",
			"https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net",
			"https://0xa1559ace749633b997cb3fdacffb890aeebdb0f5a3b6aaa7eeeaf1a38af0a8fe88b9e4b1f61f236d2e64d95733327a62@relay.ultrasound.money",
		},
	},
	NetworkGoerli: {
		Name:               NetworkGoerli,
		NetworkService:     "merge",
		GenesisForkVersion: "0x00001020",
		SupportsMEVBoost:   true,
		CheckpointSyncURL:  "https://goerli.checkpoint-sync.ethpandaops.io",
		RelayURLs: []string{
			"https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@builder-relay-goerli.flashbots.net",
			"https://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@bloxroute.max-profit.builder.goerli.blxrbdn.com",
			"https://0x8f7b17a74569b7a57e9bdafd2e159380759f5dc3ccbd4bf600414147e8c4e1dc6ebada83c0139ac15850eb6c975e82d0@builder-relay-goerli.blocknative.com",
			"https://0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12@relay-goerli.edennetwork.io",
			"https://0xb1559beef7b5ba3127485bbbb090362d9f497ba64e177ee2c8e7db74746306efad687f2cf8574e38d70067d40ef136dc@relay-stag.ultrasound.money",
		},
	},
	NetworkSepolia: {
		Name:               NetworkSepolia,
		NetworkService:     "merge",
		GenesisForkVersion: "0x90000069",
		SupportsMEVBoost:   true,
		CheckpointSyncURL:  "https://sepolia.checkpoint-sync.ethpandaops.io",
		RelayURLs: []string{
			"https://0x845bd072b7cd566f02faeb0a4033ce9399e42839ced64e8b2adcfc859ed1e8e1a5a293336a49feac6d9a5edb779be53a@builder-relay-sepolia.flashbots.net",
		},
	},
	NetworkChiado: {
		Name:               NetworkChiado,
		NetworkService:     "merge",
		GenesisForkVersion: "0x0000006f",
		CheckpointSyncURL:  "https://checkpoint.chiadochain.net",
	},
	NetworkGnosis: {
		Name:               NetworkGnosis,
		NetworkService:     "merge",
		GenesisForkVersion: "0x00000064",
		CheckpointSyncURL:  "https://checkpoint.gnosischain.com",
	},
	NetworkHolesky: {
		Name:               NetworkHolesky,
		NetworkService:     "merge",
		GenesisForkVersion: "0x00017000",
		SupportsMEVBoost:   false,
		CheckpointSyncURL:  "https://checkpoint-sync.holesky.ethpandaops.io/",
	},
	NetworkCustom: {
		Name:               NetworkCustom,
		NetworkService:     "merge",
		GenesisForkVersion: "0x00000000", // TODO: only affects keystores generation, ensure the deposit method does not conflict over this.
	},
}

func NetworksConfigs() map[string]NetworkConfig {
	return networksConfigs
}
