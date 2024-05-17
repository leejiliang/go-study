package _interface

import (
	"fmt"
	"testing"
)

type Phone interface {
	call(num string) string
}

type Nokia struct {
	brand string
}

type Iphone struct {
	color string
}

func (ph Nokia) call(num string) string {
	fmt.Printf("this is a %s , num from %s\n", ph.brand, num)
	return "call end"
}

func (p Iphone) call(num string) string {
	fmt.Printf("this is a %s , num from %s\n", p.color, num)
	return "call end"
}

func TestInterfaceMethods(t *testing.T) {
	var phone Phone
	phone = new(Nokia)
	fmt.Println(phone.call("123456"))

	phone = Nokia{brand: "n93"}
	fmt.Println(phone.call("123456"))

	phone = new(Iphone)
	fmt.Println(phone.call("78910j"))
	phone = Iphone{color: "red"}
	fmt.Println(phone.call("123456"))
}
