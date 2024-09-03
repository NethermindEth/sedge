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
package locker

import "github.com/gofrs/flock"

//go:generate mockgen -package=sedge_mocks -destination=../../../mocks/locker.go github.com/NethermindEth/sedge/internal/monitoring/locker Locker
type Locker interface {
	New(path string) Locker
	Lock() error
	Unlock() error
	Locked() bool
}

type FLock struct {
	locker *flock.Flock
}

func NewFLock() Locker {
	return &FLock{}
}

func (l *FLock) New(path string) Locker {
	l.locker = flock.New(path)
	return l
}

func (l *FLock) Lock() error {
	return l.locker.Lock()
}

func (l *FLock) Unlock() error {
	return l.locker.Unlock()
}

func (l *FLock) Locked() bool {
	return l.locker.Locked()
}
