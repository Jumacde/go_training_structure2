package main

import "fmt"

type PointChecker interface {
	GetCheckPointer() (int, int, *int)
	exeCheckPointer()
}

type CheckPointer struct {
	//value int
	a, b int
	p    *int
}

func (c *CheckPointer) SetCheckPointer(a, b int, p *int) {
	c.a = a
	c.b = b
	c.p = p
}

func (c *CheckPointer) GetCheckPointer() (int, int, *int) {
	return c.a, c.b, c.p
}

func (c *CheckPointer) exeCheckPointer() {
	var a int = 100
	var b int = 200
	var p *int
	p = &b
	fmt.Printf("p = &b: %v, value: %d\n", p, *p)
	*p = 300
	fmt.Printf("after *p = 300, b: %d\n", b)
	fmt.Printf("local a: %d\n", a)
}
