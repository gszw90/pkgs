package timeLock

import (
	"sync/atomic"
	"time"
)

const (
	_stateUnlock        int32 = 0 // 未上锁状态
	_stateLocked        int32 = 1 // 普通锁状态
	_stateReleaseLocked int32 = 3 // 锁有释放时间的上锁状态
)

type TimeLock struct {
	state int32
}

// TryLockWithTimeout 在一段时间内持续尝试获取锁
func (l *TimeLock) TryLockWithTimeout(t time.Duration) (ok bool) {
	mils := t.Milliseconds()
	for mils > 0 {
		if atomic.CompareAndSwapInt32(&(l.state), _stateUnlock, _stateLocked) {
			ok = true
			break
		}
		time.Sleep(time.Millisecond * 10)
		mils -= 100
	}
	return
}

// LockWithReleaseTime 获取锁,如果锁的持续时间超过t,则自动释放锁
func (l *TimeLock) LockWithReleaseTime(t time.Duration) {
	for {
		if atomic.CompareAndSwapInt32(&(l.state), _stateUnlock, _stateReleaseLocked) {
			go func() {
				<-time.After(t)
				l.Unlock(true)
			}()
			break
		}
	}
}

// Unlock 释放锁
func (l *TimeLock) Unlock(isReleasedLock bool) {
	if isReleasedLock {
		atomic.CompareAndSwapInt32(&(l.state), _stateReleaseLocked, _stateUnlock)
	} else {
		atomic.CompareAndSwapInt32(&(l.state), _stateLocked, _stateUnlock)
	}
}

func NewTimeLock() *TimeLock {
	return &TimeLock{state: _stateUnlock}
}
