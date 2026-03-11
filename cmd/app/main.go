package main

import "fmt"

func main() {
	fmt.Println("")
	/*
		mul := &MultiplicationTable{}
		mul.SetTable(0, 0, nil)
		//mulTable := mul.CreateTable1()
		//fmt.Println(mulTable)
		mul.createTable2()
	*/
	cp := &CheckPointer{}
	cp.SetCheckPointer(0, 0, nil)
	cp.exeCheckPointer()

}
