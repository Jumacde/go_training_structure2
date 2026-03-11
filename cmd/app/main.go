package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	choice := 0

	fmt.Println("please choice to execute programm.")
	for {
		fmt.Println("\n1: MultiplicationTable, 2: CheckPointer, 3:BMI-calculator, 4: basic-calculate, 0: end. ")
		if _, err := fmt.Scan(&choice); err == nil {
			if choice == 0 {
				println("bye bye.")
				break
			}

			switch choice {
			case 1:
				fmt.Println("execute MultiplicationTable")
				mul := &MultiplicationTable{}
				mul.SetTable(0, 0, nil)
				//mulTable := mul.CreateTable1()
				//fmt.Println(mulTable)
				mul.createTable2()
				continue
			case 2:
				fmt.Println("execute CheckPointer")
				cp := &CheckPointer{}
				cp.SetCheckPointer(0, 0, nil)
				cp.exeCheckPointer()
				continue
			case 3:
				fmt.Println("execute BMI-calculator")
				bmi := &Bmi{}
				bmi.SetBmi(0.0, 0.0, 0.0, "")
				bmi.CalcBmi()
				continue
			case 4:
				fmt.Println("execute basic-calculate")
				calc := &Calc{}
				calc.SetCalc(0.0, 0.0, 0.0)
				calc.ChoiceOperation()
			}

		}
		fmt.Println("Invalid input. please input a number again.")
		var dump string
		fmt.Scanln(&dump) // clear buffer
	}

}
