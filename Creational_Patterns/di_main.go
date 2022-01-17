package main

import (
	"fmt"
	di "my_design_pattern/Creational_Patterns/DI"
)

// A 依赖关系 A -> B -> C
type A struct {
	B   *B
	Num int
}

// NewA NewA
func NewA(b *B) *A {
	return &A{
		B:   b,
		Num: 1,
	}
}

// B B
type B struct {
	C   *C
	Num int
}

// NewB NewB
func NewB(c *C) *B {
	return &B{C: c, Num: 2}
}

// C C
type C struct {
	Num int
}

// NewC NewC
func NewC() *C {
	return &C{
		Num: 3,
	}
}

func main() {
	container := di.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.Num)
		fmt.Printf("%+v: %d", a.B, a.B.Num)
		fmt.Printf("%+v: %d", a.B.C, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
}
