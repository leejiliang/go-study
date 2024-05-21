package extend

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) sayHi(v string) {
	fmt.Printf(".....%s", v)
}

func (p *Pet) runFast() {
	fmt.Printf("pet run ...")
}

type Cat struct {
	Pet
}

func (c *Cat) sayHi(v string) {
	fmt.Printf("this is a cat: %s", v)
}

func TestExtend(t *testing.T) {
	var cat = new(Cat)
	cat.runFast()
	cat.sayHi("miao")
	t.Log("done")
}
