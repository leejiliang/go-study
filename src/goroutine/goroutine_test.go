package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(j int) { // 开启协程完成任务
			fmt.Printf("result is %d \n", j)
		}(i)
	}
}

func TestGoroutineUnsafe(t *testing.T) {
	var sum = 0
	for i := 0; i < 1000; i++ {
		go func() { // 开启协程完成任务
			sum++
		}()
	}
	time.Sleep(time.Second * 3)
	t.Logf("sum is : %d", sum)
}

func TestGoroutineSafe(t *testing.T) {
	var mux sync.Mutex
	var sum = 0
	for i := 0; i < 1000; i++ {
		go func() { // 开启协程完成任务
			defer func() {
				mux.Unlock() //释放锁
			}()
			mux.Lock() // 加锁
			sum++
		}()
	}
	time.Sleep(time.Second * 3)
	t.Logf("sum is : %d", sum)
}

func TestGoroutineWait(t *testing.T) {
	var mux sync.Mutex
	var wg sync.WaitGroup
	var sum = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { // 开启协程完成任务
			defer func() {
				mux.Unlock() //释放锁
			}()
			mux.Lock() // 加锁
			sum++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("sum is : %d", sum)
}
