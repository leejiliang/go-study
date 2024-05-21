package _interface

import "testing"

type Programme interface {
	sayHello() string
}

type JavaProgrammer struct {
}

func (pg *JavaProgrammer) sayHello() string {
	return "tom"
}

func TestVisitUser(t *testing.T) {
	var programme Programme
	programme = new(JavaProgrammer)
	t.Logf("hello %s", programme.sayHello())
}
