package cache

import (
	"context"
	"errors"
	"time"
)

const (
	// 过期时间
	ttl = 30 * time.Second
	// 重新获取锁间隔
	tryLockInterval = time.Second
	// 解锁脚本
	unlockScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end`
)

var (
	// ErrLockFailed 加锁失败
	ErrLockFailed = errors.New("lock failed")
	// ErrTimeout 加锁超时
	ErrTimeout = errors.New("timeout")
)

func Lock(ctx context.Context, key, value string) error {
	err := TryLock(key, value)
	if err == nil {
		return nil
	}
	if !errors.Is(err, ErrLockFailed) {
		return err
	}
	// 加锁失败，不断尝试
	ticker := time.NewTicker(time.Duration(tryLockInterval))
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return ErrTimeout
		case <-ticker.C:
			// 重新尝试加锁
			err := TryLock(key, value)
			if err == nil {
				return nil
			}
			if !errors.Is(err, ErrLockFailed) {
				return err
			}
		}
	}
}

func TryLock(key, value string) error {
	success, err := RdbDistLock.SetNX(Ctx, key, value, time.Duration(ttl)).Result()
	if err != nil {
		return err
	}
	// 加锁失败
	if !success {
		return ErrLockFailed
	}
	return nil
}

func UnLock(key, value string) error {
	err := RdbUnlockScript.Run(Ctx, RdbDistLock, []string{key}, value).Err()
	return err
}

func CheckLock(key string) (bool, error) {
	exists, err := RdbDistLock.Exists(Ctx, key).Result()
	if err != nil {
		return false, err
	}
	if exists > 0 {
		return true, nil
	}
	return false, nil
}
