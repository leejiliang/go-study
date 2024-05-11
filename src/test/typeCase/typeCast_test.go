package typeCase

import (
	"strconv"
	"testing"
)

func TestTypeCase(t *testing.T) {
	var a int = 10
	var b float32 = float32(a)
	t.Logf("b value is %f", b)

	var str1 string = "100v"
	var num1 int
	var err error
	num1, err = strconv.Atoi(str1)
	t.Logf("num1 is %d, error info %v", num1, err)

	var num2 int = 200
	str2 := strconv.Itoa(num2)
	t.Logf("str2 is %v", str2)

	var i interface{} = "hello world"
	str3, ok3 := i.(string)
	t.Logf("str is %v, error info %v", str3, ok3)
}
