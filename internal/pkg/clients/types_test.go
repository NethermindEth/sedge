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
package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetImageOrDefault_Execution(t *testing.T) {
	tests := []struct {
		client        Client
		expectedImage string
	}{
		{
			client: Client{
				Name: "geth",
				Type: "execution",
			},
			expectedImage: "ethereum/client-go:v1.10.26",
		},
		{
			client: Client{
				Name: "besu",
				Type: "execution",
			},
			expectedImage: "hyperledger/besu:22.10.3",
		},
		{
			client: Client{
				Name: "nethermind",
				Type: "execution",
			},
			expectedImage: "nethermind/nethermind:1.14.7",
		},
		{
			client: Client{
				Name: "erigon",
				Type: "execution",
			},
			expectedImage: "erigon/erigon:v2.29.0",
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("")
			assert.Equal(t, test.expectedImage, test.client.Image)
		})
	}
}

func TestSetImageOrDefault_Consensus(t *testing.T) {
	tests := []struct {
		client        Client
		expectedImage string
	}{
		{
			client: Client{
				Name: "lighthouse",
				Type: "consensus",
			},
			expectedImage: "sigp/lighthouse:v3.3.0",
		},
		{
			client: Client{
				Name: "prysm",
				Type: "consensus",
			},
			expectedImage: "gcr.io/prysmaticlabs/prysm/beacon-chain:v3.2.0",
		},
		{
			client: Client{
				Name: "teku",
				Type: "consensus",
			},
			expectedImage: "consensys/teku:22.11.0",
		},
		{
			client: Client{
				Name: "lodestar",
				Type: "consensus",
			},
			expectedImage: "chainsafe/lodestar:v1.2.2",
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("")
			assert.Equal(t, test.expectedImage, test.client.Image)
		})
	}
}

func TestSetImageOrDefault_Validator(t *testing.T) {
	tests := []struct {
		client        Client
		expectedImage string
	}{
		{
			client: Client{
				Name: "lighthouse",
				Type: "validator",
			},
			expectedImage: "sigp/lighthouse:v3.3.0",
		},
		{
			client: Client{
				Name: "prysm",
				Type: "validator",
			},
			expectedImage: "gcr.io/prysmaticlabs/prysm/validator:v3.2.0",
		},
		{
			client: Client{
				Name: "teku",
				Type: "validator",
			},
			expectedImage: "consensys/teku:22.11.0",
		},
		{
			client: Client{
				Name: "lodestar",
				Type: "validator",
			},
			expectedImage: "chainsafe/lodestar:v1.2.2",
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("")
			assert.Equal(t, test.expectedImage, test.client.Image)
		})
	}
}

func TestSetImageOrDefault_CustomImage(t *testing.T) {
	tests := []struct {
		client      Client
		customImage string
	}{
		{
			client: Client{
				Name: "geth",
				Type: "execution",
			},
			customImage: "my/geth-image:v1.0.0",
		},
		{
			client: Client{
				Name: "lighthouse",
				Type: "consensus",
			},
			customImage: "my/lighthouse-image:v1.0.0",
		},
		{
			client: Client{
				Name: "prysm",
				Type: "validator",
			},
			customImage: "my/prysm-image:v1.0.0",
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault(test.customImage)
			assert.Equal(t, test.customImage, test.client.Image)
		})
	}
}
