package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func addData(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
}

func receiveData(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("reveive data: %d \n", i)
				time.Sleep(time.Millisecond * 100)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	var cChan = make(chan int, 10)
	addData(cChan)
	receiveData(cChan, &wg)
	t.Log(" executed over")
	wg.Wait()
}
