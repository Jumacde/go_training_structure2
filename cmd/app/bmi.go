package main

import "fmt"

type BmiCalculater interface {
	SetBmi(height, weight, bmiResult float32, name string)
	CalcBmi() (string, float32)
}

type Bmi struct {
	height, weight, bmiResult float32
	name                      string
}

func (b *Bmi) SetBmi(height, weight, bmiResult float32, name string) {
	b.height = height
	b.weight = weight
	b.bmiResult = bmiResult
	b.name = name
}

func (b *Bmi) GetBmi() (float32, float32, float32, string) {
	return b.height, b.weight, b.bmiResult, b.name
}

func (b *Bmi) CalcBmi() (string, float32) {
	fmt.Println("BMI calculator")
	fmt.Println("please input your name: ")
	fmt.Scan(&b.name)

	for {
		fmt.Println("please input your height(cm): ")
		if _, err := fmt.Scan(&b.height); err == nil && b.height > 0 {
			break
		}
		fmt.Println("Invalid input. please input your height(cm) again.")
		var dump string
		fmt.Scanln(&dump) // clear buffer
	}

	for {
		fmt.Println("please input your weight(kg): ")
		if _, err := fmt.Scan(&b.weight); err == nil && b.weight > 0 {
			break
		}
		fmt.Println("Invalid input. please input your weight(kg) again.")
		var dump string
		fmt.Scanln(&dump) // clear buffer
	}
	b.bmiResult = b.weight / ((b.height / 100) * (b.height / 100))
	fmt.Printf("%s's bmi is %.2f.", b.name, b.bmiResult)
	return b.name, b.bmiResult

}
