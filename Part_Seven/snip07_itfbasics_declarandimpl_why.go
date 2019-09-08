package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int // both `Permanent` & `Contract` now impl this interface (?)
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}
type Contract struct {
	empId    int
	basicpay int
}

func main() {
	pemp1 := Permanent{1, 3000, 20}
	pemp2 := Permanent{2, 4000, 30}
	cemp1 := Contract{3, 3000}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpenses(employees)
}

func (perma Permanent) CalculateSalary() int {
	return perma.basicpay + perma.pf
}

func (contr Contract) CalculateSalary() int {
	return contr.basicpay
}

// CalculateSalary() int
// >> int					Pass a slice which contains both `Permanent` & `Contract` types
// >> CalculateSalary()		Calc the expenses by calling the corresponding methods
func totalExpenses(salCalc []SalaryCalculator) {
	expense := 0
	for _, v := range salCalc {
		expense = expense + v.CalculateSalary()
	}

	fmt.Printf("Total Expense per Month $%d\n", expense)
}
