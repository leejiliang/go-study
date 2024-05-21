package _interface

import "testing"

type Programer interface {
	sayHello() string
}

type JavaProgramer struct {
}

func (pg *JavaProgramer) sayHello() string {
	return "tom"
}

func TestVisitUser(t *testing.T) {
	var programe Programer
	programe = new(JavaProgramer)
	t.Logf("hello %s", programe.sayHello())
}
