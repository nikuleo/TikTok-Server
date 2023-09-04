package cache

import (
	"TikTokServer/pkg/tlog"
	"context"
	"fmt"
	"testing"
	"time"
)

func init() {
	tlog.InitLog()
	InitRedis()
}

func TestDistLock(t *testing.T) {
	ctx1, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*20))
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Duration(time.Second*20))
	defer cancel()
	defer cancel2()
	err := Lock(ctx1, "userID:10", "10")
	exited, err := CheckLock("userID:10")
	fmt.Println("第一次加锁：", err)
	fmt.Println("exited: ", exited)
	go func() {
		time.Sleep(time.Second * 4)
		err := UnLock("userID:10", "10")
		fmt.Println("unlock:", err)
	}()
	exited, _ = CheckLock("userID:11")
	fmt.Println("exited:", exited)
	err = Lock(ctx2, "userID:10", "3")
	fmt.Println("第二次加锁：", err)
	// Lock(ctx, "userID:11", "2")
	time.Sleep(time.Second * 10)
	UnLock("userID:10", "3")
	exited, _ = CheckLock("userID:10")
	fmt.Println("exited:", exited)
}
