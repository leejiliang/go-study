package cancel

import (
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

func isChanCancel(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}

func receiveData(ch chan int, cancelCh chan struct{}, group *sync.WaitGroup) {
	for tt := 0; tt < 5; tt++ {
		group.Add(1)
		go func() {
			for {
				if !isChanCancel(cancelCh) {
					if i, open2 := <-ch; open2 {
						fmt.Printf("receive data : %d \n", i)
						time.Sleep(time.Millisecond * 2000)
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
	var cancelCh = make(chan struct{}, 1)
	var wg sync.WaitGroup
	produceData(dataCh)
	receiveData(dataCh, cancelCh, &wg)
	time.Sleep(time.Second * 1)
	close(cancelCh)
	wg.Wait()
}
