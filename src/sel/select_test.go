package sel

import (
	"fmt"
	"testing"
	"time"
)

func job1() string {
	time.Sleep(time.Second * 2)
	fmt.Printf("do something in job1")
	return "job1 result"
}

func job2() string {
	time.Sleep(time.Second * 3)
	fmt.Printf("do something in job2")
	return "job2 result"
}

func job2Fast() string {
	time.Sleep(time.Second * 1)
	fmt.Printf("do something in job2")
	return "job2 result"
}

func asyncJob1(job func() string) chan string { // 返回类型是一个异步执行结果获取chan
	//retChan := make(chan string) // 定义chan , 没有buffer
	retChan := make(chan string, 1) // 定义chan , 没有buffer
	go func() {
		ret := job()
		fmt.Println("get job1 result")
		retChan <- ret // 结果放进chan
		fmt.Println("job1 exited")
	}()
	return retChan
}

func TestCsp(t *testing.T) {
	retChan1 := asyncJob1(job1)
	retChan2 := asyncJob1(job2Fast)

	select {
	case ret := <-retChan1:
		t.Logf("get result from chan1: %s", ret)
	case ret := <-retChan2:
		t.Logf("get result from chan2 : %s", ret)
	case <-time.After(time.Millisecond * 500):
		t.Error(" 等待超时")
	}
}
