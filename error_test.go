package log_test

import (
	"fmt"
	log "github.com/subiz/log"
	"testing"
	"time"
)

func A() error {
	err := B()
	return err
}

func B() error {
	err := CCCCCC()
	return err
}

func CCCCCC() error {
	err := DDDDDD()
	return err
}

func DDDDDD() error {
	err := E()
	return err
}

func E() error {
	err := log.EServer(nil, log.M{"number3": "thanh"})
	return err
}

func TestError(t *testing.T) {
	err := log.EAccountLocked("thanh") // A()
	fmt.Println("EEEEEE", err.Error())
	time.Sleep(20 * time.Second)
}

func TestWrap(t *testing.T) {
	var err error = log.EAccountLocked("thanh") // A()
	log.WrapStack(err, 0)
	time.Sleep(20 * time.Second)
}
