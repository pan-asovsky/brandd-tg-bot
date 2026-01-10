package mocks

import (
	"fmt"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type slotLockingMock struct {
	LockOK       bool
	LockErr      error
	IsLockedMap  map[string]bool
	UnlockErr    error
	RefreshErr   error
	AreLockedErr error

	LockCalled    bool
	UnlockCalled  bool
	RefreshCalled bool
}

func NewSlotLockingMock() i.SlotLocking {
	return &slotLockingMock{
		IsLockedMap: make(map[string]bool),
	}
}

func (m *slotLockingMock) Lock(date, time string) (string, bool, error) {
	m.LockCalled = true
	if m.LockErr != nil {
		return "", false, m.LockErr
	}
	return "mock-uid", m.LockOK, nil
}

func (m *slotLockingMock) Unlock(key, u string) error {
	m.UnlockCalled = true
	return m.UnlockErr
}

func (m *slotLockingMock) IsLocked(date, time string) (bool, error) {
	key := m.FormatKey(date, time)
	locked, ok := m.IsLockedMap[key]
	if !ok {
		return false, nil
	}
	return locked, nil
}

func (m *slotLockingMock) RefreshTTL(key string) error {
	m.RefreshCalled = true
	return m.RefreshErr
}

func (m *slotLockingMock) AreLocked(keys ...string) (map[string]bool, error) {
	if m.AreLockedErr != nil {
		return nil, m.AreLockedErr
	}
	res := make(map[string]bool, len(keys))
	for _, k := range keys {
		res[k] = m.IsLockedMap[k]
	}
	return res, nil
}

func (m *slotLockingMock) FormatKey(date, time string) string {
	return fmt.Sprintf("s_lock:%s_%s", date, time)
}
