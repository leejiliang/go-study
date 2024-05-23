package cancel_context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func produceData(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func isChanCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():

		return true
	default:
		return false
	}
}

func receiveData(ch chan int, ctx context.Context, group *sync.WaitGroup) {
	for tt := 0; tt < 5; tt++ {
		group.Add(1)
		go func() {
			for {
				if !isChanCancel(ctx) {
					if i, open2 := <-ch; open2 {
						fmt.Printf("receive data : %d \n", i)
						time.Sleep(time.Millisecond * 1000)
					} else {
						fmt.Println("producer closed channel")
						break
					}
				} else {
					fmt.Println("任务取消")
					break
				}
			}
			group.Done()
		}()
	}

}

func TestCancelJob(t *testing.T) {
	var dataCh = make(chan int, 10)
	var ctx, cancel = context.WithCancel(context.Background())
	var wg sync.WaitGroup
	produceData(dataCh)
	receiveData(dataCh, ctx, &wg)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}
