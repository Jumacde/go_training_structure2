package main

import (
	"fmt"
	"math"
)

type Calculator interface {
	GetCalc() (float32, float32)
	ChoiceOperation()
}

type Calc struct {
	x, y, result float32
}

func (cal *Calc) SetCalc(x, y, result float32) {
	cal.x = x
	cal.y = y
	cal.result = result
}
func (cal *Calc) GetCalc() (float32, float32, float32) {
	return cal.x, cal.y, cal.result
}

func (cal *Calc) Calc_add() float32 {
	cal.result = cal.x + cal.y
	fmt.Printf("%.2f + %.2f = %.2f", cal.x, cal.y, cal.result)
	return cal.result
}

func (cal *Calc) Calc_sub() float32 {
	cal.result = cal.x - cal.y
	fmt.Printf("%.2f - %.2f = %.2f", cal.x, cal.y, cal.result)
	return cal.result
}

func (cal *Calc) Calc_mul() float32 {
	cal.result = cal.x * cal.y
	fmt.Printf("%.2f * %.2f = %.2f", cal.x, cal.y, cal.result)
	return cal.result
}

func (cal *Calc) Calc_div() float32 {
	cal.result = cal.x / cal.y
	fmt.Printf("%.2f / %.2f = %.2f", cal.x, cal.y, cal.result)
	return cal.result
}

func (cal *Calc) Calc_mod() float32 {
	floMod := math.Mod(float64(cal.x), float64(cal.y))
	cal.result = float32(floMod)
	fmt.Printf("%.2f mod %.2f = %.2f", cal.x, cal.y, cal.result)
	return cal.result
}

func (cal *Calc) ChoiceOperation() {
	operations := 0

	for {
		fmt.Println("\nplease choice a arithmetic operations:")
		fmt.Println("1: x + y, 2: x - y, 3: x * y, 4: x / y, 5: x mod y, 0: end")
		if _, err := fmt.Scan(&operations); err == nil {
			// input 0 == end program.
			if operations == 0 {
				println("bye bye.")
				break
			}
			// input between 1 and 4. you can input every digit for x and y.
			if operations >= 1 && operations <= 5 {
				fmt.Print("x = ")
				fmt.Scan(&cal.x)
				fmt.Print("y = ")
				fmt.Scan(&cal.y)

			}

			switch operations {
			case 1:
				cal.Calc_add()
				continue
			case 2:
				cal.Calc_sub()
				continue
			case 3:
				cal.Calc_mul()
				continue
			case 4:
				cal.Calc_div()
				continue
			case 5:
				cal.Calc_mod()
				continue
			}
			fmt.Println("Invalid input. please input your height(cm) again.")
			var dump string
			fmt.Scanln(&dump) // clear buffer
		}

	}
}
