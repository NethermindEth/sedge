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
	"regexp"
	"testing"

	clientsimages "github.com/NethermindEth/sedge/configs/images"
	"github.com/stretchr/testify/assert"
)

var guesser clientsimages.ClientsImagesGuesser

func TestSetImageOrDefault_Execution(t *testing.T) {
	// default clients images
	defaults, err := clientsimages.NewDefaultClientsImages()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		client        Client
		expectedImage *regexp.Regexp
	}{
		{
			client: Client{
				Name: "geth",
				Type: "execution",
			},
			expectedImage: guesser.GethRegexp(),
		},
		{
			client: Client{
				Name: "besu",
				Type: "execution",
			},
			expectedImage: guesser.BesuRegexp(),
		},
		{
			client: Client{
				Name: "nethermind",
				Type: "execution",
			},
			expectedImage: guesser.NethermindRegexp(),
		},
		{
			client: Client{
				Name: "erigon",
				Type: "execution",
			},
			expectedImage: guesser.ErigonRegexp(),
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("", defaults)
			assert.Regexp(t, test.expectedImage, test.client.Image)
		})
	}
}

func TestSetImageOrDefault_Consensus(t *testing.T) {
	// default clients images
	defaults, err := clientsimages.NewDefaultClientsImages()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		client        Client
		expectedImage *regexp.Regexp
	}{
		{
			client: Client{
				Name: "lighthouse",
				Type: "consensus",
			},
			expectedImage: guesser.LighthouseConsensusRegexp(),
		},
		{
			client: Client{
				Name: "prysm",
				Type: "consensus",
			},
			expectedImage: guesser.PrysmConsensusRegexp(),
		},
		{
			client: Client{
				Name: "teku",
				Type: "consensus",
			},
			expectedImage: guesser.TekuConsensusRegexp(),
		},
		{
			client: Client{
				Name: "lodestar",
				Type: "consensus",
			},
			expectedImage: guesser.LodestarConsensusRegexp(),
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("", defaults)
			assert.Regexp(t, test.expectedImage, test.client.Image)
		})
	}
}

func TestSetImageOrDefault_Validator(t *testing.T) {
	// default clients images
	defaults, err := clientsimages.NewDefaultClientsImages()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		client        Client
		expectedImage *regexp.Regexp
	}{
		{
			client: Client{
				Name: "lighthouse",
				Type: "validator",
			},
			expectedImage: guesser.LighthouseValidatorRegexp(),
		},
		{
			client: Client{
				Name: "prysm",
				Type: "validator",
			},
			expectedImage: guesser.PrysmValidatorRegexp(),
		},
		{
			client: Client{
				Name: "teku",
				Type: "validator",
			},
			expectedImage: guesser.TekuValidatorRegexp(),
		},
		{
			client: Client{
				Name: "lodestar",
				Type: "validator",
			},
			expectedImage: guesser.LodestarValidatorRegexp(),
		},
	}
	for _, test := range tests {
		t.Run(test.client.Name, func(t *testing.T) {
			test.client.SetImageOrDefault("", defaults)
			assert.True(t, test.expectedImage.Match([]byte(test.client.Image)))
		})
	}
}

func TestSetImageOrDefault_CustomImage(t *testing.T) {
	// default clients images
	defaults, err := clientsimages.NewDefaultClientsImages()
	if err != nil {
		t.Fatal(err)
	}

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
			test.client.SetImageOrDefault(test.customImage, defaults)
			assert.Equal(t, test.customImage, test.client.Image)
		})
	}
}
