package main

import "fmt"

func main() {
	fmt.Printf("")
	mul := &MultiplicationTable{}
	mul.SetTable(0, 0, nil)
	//mulTable := mul.CreateTable1()
	//fmt.Println(mulTable)
	mul.createTable2()

}
