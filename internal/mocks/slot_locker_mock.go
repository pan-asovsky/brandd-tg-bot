package mocks

import (
	"fmt"

	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
)

type slotLockerServiceMock struct {
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

func NewSlotLockerMock() isvc.SlotLocker {
	return &slotLockerServiceMock{
		IsLockedMap: make(map[string]bool),
	}
}

func (m *slotLockerServiceMock) Lock(_, _ string) (string, bool, error) {
	m.LockCalled = true
	if m.LockErr != nil {
		return "", false, m.LockErr
	}
	return "mock-uid", m.LockOK, nil
}

func (m *slotLockerServiceMock) Unlock(_, _ string) error {
	m.UnlockCalled = true
	return m.UnlockErr
}

func (m *slotLockerServiceMock) IsLocked(date, time string) (bool, error) {
	key := m.FormatKey(date, time)
	locked, ok := m.IsLockedMap[key]
	if !ok {
		return false, nil
	}
	return locked, nil
}

func (m *slotLockerServiceMock) RefreshTTL(_ string) error {
	m.RefreshCalled = true
	return m.RefreshErr
}

func (m *slotLockerServiceMock) AreLocked(keys ...string) (map[string]bool, error) {
	if m.AreLockedErr != nil {
		return nil, m.AreLockedErr
	}
	res := make(map[string]bool, len(keys))
	for _, k := range keys {
		res[k] = m.IsLockedMap[k]
	}
	return res, nil
}

func (m *slotLockerServiceMock) FormatKey(date, time string) string {
	return fmt.Sprintf("s_lock:%s_%s", date, time)
}
