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
package env

import "regexp"

var (
	ReMEV         = regexp.MustCompile(`MEV=(.*)`)
	ReXEEV        = regexp.MustCompile(`XEE_VERSION=(.*)`)
	ReClBOOTNODES = regexp.MustCompile(`CC_BOOTNODES=(.*)`)
	ReElBOOTNODES = regexp.MustCompile(`EC_BOOTNODES=(.*)`)
	ReCHECKPOINT  = regexp.MustCompile(`CHECKPOINT_SYNC_URL=(.*)`)
	ReSPLITTED    = regexp.MustCompile(`EL_NETWORK=(.*)`)
)
