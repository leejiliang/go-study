package csp

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
	time.Sleep(time.Second * 2)
	fmt.Printf("do something in job2")
	return "job2 result"
}

func asyncJob1() chan string { // 返回类型是一个异步执行结果获取chan
	//retChan := make(chan string) // 定义chan , 没有buffer
	retChan := make(chan string, 1) // 定义chan , 没有buffer
	go func() {
		ret := job1()
		fmt.Println("get job1 result")
		retChan <- ret // 结果放进chan
		fmt.Println("job1 exited")
	}()
	return retChan
}

func TestCsp(t *testing.T) {
	//job1()
	retChan := asyncJob1()
	job2()
	t.Logf("job1 result is : %s", <-retChan) //, 从chan 中获取结果，  如果没有buffer, 在这里才开始执行job1
	t.Logf("all job is done")
}
