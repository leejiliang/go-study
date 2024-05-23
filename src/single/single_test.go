package single

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type ObjectMapper struct {
}

var once sync.Once
var objectMapper *ObjectMapper

func createObjectMapper() *ObjectMapper {
	once.Do(func() {
		fmt.Println(" 创建ObjectMapper 对象 ")
		objectMapper = new(ObjectMapper)
	})
	return objectMapper
}

func TestConcurrentCreateObjectMapper(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mapper := createObjectMapper()
			fmt.Printf("对象内存地址： %v \n", unsafe.Pointer(mapper))
			wg.Done()
		}()
	}
	wg.Wait()
}
