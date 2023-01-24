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

import "github.com/NethermindEth/sedge/configs"

// Client : Struct Represent a client like geth, prysm, etc
type Client struct {
	Name      string
	Type      string
	Image     string
	Endpoint  string
	Supported bool
	Omitted   bool
}

func (c *Client) SetImageOrDefault(image string) {
	switch c.Type {
	case "validator":
		c.setValidatorImage(image)
	}
}

func (c *Client) setValidatorImage(image string) {
	switch c.Name {
	case "lighthouse":
		c.Image = valueOrDefault(image, configs.Lighthouse_ValidatorImage)
	case "prysm":
		c.Image = valueOrDefault(image, configs.Prysm_ValidatorImage)
	case "teku":
		c.Image = valueOrDefault(image, configs.Teku_ValidatorImage)
	case "lodestar":
		c.Image = valueOrDefault(image, configs.Lodestar_ValidatorImage)
	}
}

func valueOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// Clients : Struct Represent a combination of execution, consensus and validator clients
type Clients struct {
	Execution Client
	Consensus Client
	Validator Client
}

type ClientMap map[string]Client

type OrderedClients map[string]ClientMap
