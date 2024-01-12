package main

import "fmt"

func Print() {
	fmt.Println()
	x := X{}
	x.Tes()
}

type S struct {
}

type X struct {
	s S
}

func (x X) Tes() {
	fmt.Println()
	if true {
		x.Test2()
	}
}

func (s S) Test3() {
}

func (x X) Test2() {
	x.s.Test3()
}
