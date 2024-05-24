package first

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doSomething(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf(" receive result from %d \n", i)
}

func TestGetFirstJobResult(t *testing.T) {
	t.Logf("beautiful runtime goroutine: %d", runtime.NumGoroutine())
	var ch = make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(val int) {
			ret := doSomething(val)
			ch <- ret
		}(i)
	}
	ret := <-ch
	t.Log(ret)
	time.Sleep(time.Second * 1)
	t.Logf("after runtime goroutine: %d", runtime.NumGoroutine())
}
