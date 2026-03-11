package main

import "fmt"

type MultiTable interface {
	GetTable() (int, int, [][]int)
	CreateTable() [][]int
}

type MultiplicationTable struct {
	multiplierNum, multiplicandNum int
	mTable                         [][]int
}

func (m *MultiplicationTable) SetTable(multiplierNum, multiplicandNum int, mTable [][]int) {
	m.multiplierNum = multiplierNum
	m.multiplicandNum = multiplicandNum
	m.mTable = make([][]int, 9)
}

func (m *MultiplicationTable) GetTable() (int, int, [][]int) {
	return m.multiplierNum, m.multiplicandNum, m.mTable
}

func (m *MultiplicationTable) CreateTable1() [][]int {
	m.mTable = make([][]int, 10)
	for m.multiplierNum = 0; m.multiplierNum <= 9; m.multiplierNum++ {
		m.mTable[m.multiplierNum] = make([]int, 10)
		for m.multiplicandNum = 0; m.multiplicandNum <= 9; m.multiplicandNum++ {
			m.mTable[m.multiplierNum][m.multiplicandNum] = m.multiplierNum * m.multiplicandNum
		}
	}
	return m.mTable
}

func (m *MultiplicationTable) createTable2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			fmt.Printf("%2d|", i*j)
		}
	}
}
