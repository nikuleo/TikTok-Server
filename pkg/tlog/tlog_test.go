package tlog

import (
	"testing"
	"time"
)

func TestTlog(t *testing.T) {

	InitLog()
	defer Sync()
	Info("construction succeeded")

	Infof("Global.Source: '%s'", "xxxx")
	Infof("Global.ChangeMe: '%s', %s", "xxxx", "xxxx")
	Fatalf("Global.Source: '%s'", "xxxx")

	// block for watch test
	time.Sleep(1 * time.Second)
}
