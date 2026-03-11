package main

import "fmt"

type EasyPointers interface {
	GetEasyPointer() (string, uint)
	ShowInfo1()
	ShowInfo2()
	checkP()
}

type EasyPointer struct {
	leaderName string
	age        uint
}

func (ep *EasyPointer) SetEasyPointer(leaderName string, age uint) {
	ep.leaderName = leaderName
	ep.age = age
}

func (ep *EasyPointer) GetEasyPointer() (string, uint) {
	return ep.leaderName, ep.age
}

func (ep *EasyPointer) ShowInfo1() {
	ep.leaderName = "Alice"
	ep.age = 19
	fmt.Printf("%s is leader and she is %d years old.\n", ep.leaderName, ep.age)
	fmt.Printf("After Info1: %s, %d\n", ep.leaderName, ep.age)

}

func (ep *EasyPointer) ShowInfo2() {
	ep.leaderName = "Bob"
	ep.age = 23
	fmt.Printf("%s is leader and he is %d years old.\n", ep.leaderName, ep.age)
	fmt.Printf("After Info1: %s, %d\n", ep.leaderName, ep.age)
}

func (ep *EasyPointer) checkP() {
	if ep.age == 23 {
		fmt.Println("Pointer of Alice-Info is keeping")
	} else {
		fmt.Println("Pointer of Alice-Info isn't saved")
	}
}
