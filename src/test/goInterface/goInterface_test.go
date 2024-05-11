package goInterface

import (
	"fmt"
	"testing"
)

type Phone interface {
	call()
}

type Nokia struct{}
type Iphone struct{}

func (phone Nokia) call() {
	fmt.Println("this is nokia")
}

func (iphone Iphone) call() {
	fmt.Println("this is iphone")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(Nokia)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
