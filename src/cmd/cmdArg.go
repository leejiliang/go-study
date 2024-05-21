package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The Greeting object")
}

func main() {
	flag.Parse()
	fmt.Println("Hello ", name)
}
