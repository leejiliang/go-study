package error_handler

import (
	"errors"
	"fmt"
	"go-study/src/common"
	"os"
	"testing"
)

func TestErrorHandler1(t *testing.T) {
	t.Log("programme start")
	os.Exit(633)
}

func TestErrorHandler2(t *testing.T) {
	t.Log("programme start")
	panic(errors.New("遇到了麻煩。。。"))
}

func TestErrorHandler3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("defer 函數被触发， %s", err)
		}
	}()
	t.Log("programme start")
	panic(errors.New("遇到了麻煩。。。"))
}

func TestCommonUsage(t *testing.T) {
	t.Logf("cal result is : %d", common.Square(8))
}
