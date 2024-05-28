package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string `json:"name"`
}

func TestCheckType(t *testing.T) {
	//op := 32
	//checkType(&op)
	//checkTypeAndValue(op)

	reflectObj()
}

func heckType(op interface{}) {
	typeOf := reflect.TypeOf(op)
	switch typeOf.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Print("float")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Print("int")
	default:
		fmt.Print("unknown")
	}
}

func checkTypeAndValue(op interface{}) {
	fmt.Printf("type: %T\n", reflect.TypeOf(op))
	fmt.Printf("type: %T\n", reflect.TypeOf(op).Kind())
	fmt.Printf("value: %T\n", reflect.ValueOf(op))
}

func (t *Student) UpdateName(name string) {
	t.Name = name
}

/*
*
利用反射访问对象的成员
利用反射访问对象的方法
*/
func reflectObj() {
	op := Student{
		Name: "张三",
	}
	name := reflect.ValueOf(op).FieldByName("Name")
	fmt.Printf("name: %v\n", name)

	if byName, b := reflect.TypeOf(op).FieldByName("Name"); b {
		fmt.Printf("json format tag :  %v \n", byName.Tag.Get("json"))
	} else {
		fmt.Printf("reflect no name field\n")
	}

	methodByName := reflect.ValueOf(&op).MethodByName("UpdateName")
	fmt.Printf("methodByName: %v\n", methodByName)
	reflect.ValueOf(&op).MethodByName("UpdateName").Call([]reflect.Value{reflect.ValueOf("赵四")})
	fmt.Printf("new name is %v", reflect.ValueOf(op).FieldByName("Name"))
}
