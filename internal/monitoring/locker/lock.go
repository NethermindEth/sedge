package locker

import "github.com/gofrs/flock"

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/Locker.go github.com/NethermindEth/sedge/internal/locker Locker
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
