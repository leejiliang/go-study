package point

import "testing"

func TestGetVarPoint(t *testing.T) {
	str1 := "Hello World"
	t.Log(&str1) // & 取内存地址符号
}

func TestDeclarePoint(t *testing.T) {
	// * 号用于指定变量是作为一个指针
	var ip *int     //指向整形
	var fp *float32 // 指向浮点型
	t.Log(ip, fp)
}

func TestUsePoint(t *testing.T) {
	// 定义指针变量
	var strp *string
	var helStr = "Hello World"
	// 为指针变量赋值
	strp = &helStr
	t.Log("strp 内存地址:", strp)
	t.Log("helStr:", &helStr)

	// 访问指针地址指向的地址的值
	t.Log("strp 指针值: ", *strp)
}

func TestNilPtr(t *testing.T) {
	var ptr *int
	t.Log("ptr value", ptr)

	// 空指针判断
	if ptr == nil {
		t.Log("ptr is nil")
	}
}
