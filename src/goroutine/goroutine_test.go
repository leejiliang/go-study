package goroutine

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(j int) { // 开启协程完成任务
			fmt.Printf("result is %d \n", j)
		}(i)
	}
}
