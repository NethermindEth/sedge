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
package prompts

type Prompt interface {
	Passphrase() string
	ExistingVal() int64
	NumberVal() int64
	Eth1Withdrawal() (string, error)
	FeeRecipient() (string, error)

	Select(label string, options ...string) (string, error)
	Confirm(label string) (bool, error)
	Input(label string, required bool) (string, error)
	InputHide(label string) (string, error)
	InputNumber(label string) (int64, error)
	InputFilePath(label string, required bool) (string, error)
}
