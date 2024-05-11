package goMap

import "testing"

func TestInitMap(t *testing.T) {
	// 使用make函数定义map
	//var idNameMap = make(map[int]string, 10) // 10表示初始容量， 会自动扩容
	// 字面量创建Map
	map2 := map[int]string{1: "a", 2: "b", 3: "c"}
	// 获取元素
	t.Logf("map2 key 2 value is %v", map2[2])
	// 修改元素
	map2[3] = "hello"
	t.Logf("map2 key 3 value is %v", map2[3])
	// 获取长度
	t.Logf("map2 len %d", len(map2))
	// 遍历map
	for k, v := range map2 {
		t.Logf("map2 key %d value is %v", k, v)

	}
	// 删除元素
	delete(map2, 1)
	t.Logf("map2 key 1 value is  %v", map2[1])
	// 掺入元素
	map2[4] = "four"
	t.Logf("map2 key 4 valye is  %v", map2[4])
	// 判断map中是否存在某个key
	val, exits := map2[5]
	t.Logf("key 5 value is %v, 是否存在%v", val, exits)
	val, exits = map2[4]
	t.Logf("key 4 value is %v, 是否存在%v", val, exits)
}
