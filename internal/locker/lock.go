package locker

import (
	"github.com/NethermindEth/sedge/internal/monitoring/locker"
	"github.com/gofrs/flock"
)

type FLock struct {
	locker *flock.Flock
}

func NewFLock() locker.Locker {
	return &FLock{}
}

func (l *FLock) New(path string) locker.Locker {
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
