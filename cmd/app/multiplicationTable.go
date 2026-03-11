package main

type multiTable interface {
}

type multiplicationTable struct {
	multiplierNum, multiplicandNum int
}

func (m *multiplicationTable) SetMultiplicationTable(multiplierNum, multiplicandNum int) {
	m.multiplierNum = multiplierNum
	m.multiplicandNum = multiplicandNum
}

func (m *multiplicationTable) GetMultiplicationTable() (int, int) {
	return m.multiplierNum, m.multiplicandNum
}

func (m *multiplicationTable) CreateMultiTable() {}
