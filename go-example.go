package main

import (
	"fmt"
)

func main() {
	f := func(s string) func() string {
		r := func() string {
			return s
		}
		return r
	}
	hello := f("hello")
	world := f("world")
	fmt.Printf("%v %v\n", hello, world)
	fmt.Printf("%s %s\n", hello(), world())
}
