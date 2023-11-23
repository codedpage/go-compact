package main

import "fmt"

type permanentEmp struct {
	empid    int
	basicpay int
	pf       int
}

type contractEmp struct {
	empid    int
	basicpay int
}

func (p permanentEmp) calSalary() int {
	return p.basicpay + p.pf
}

func (p contractEmp) calSalary() int {
	return p.basicpay
}

type salaryCalculate interface {
	calSalary() int
}

func totalExp(s []salaryCalculate) {
	exp := 0
	for _, v := range s {
		exp = exp + v.calSalary()
	}

	fmt.Println(exp)
}

func main() {
	p1 := permanentEmp{1, 5000, 50}
	p2 := permanentEmp{2, 6000, 60}
	p3 := contractEmp{3, 2000}

	fmt.Println(p1.calSalary(), p2.calSalary(), p3.calSalary())
	emp := []salaryCalculate{p1, p2, p3}
	totalExp(emp)
}

//https://go.dev/play/p/r_DpKDmEHv-