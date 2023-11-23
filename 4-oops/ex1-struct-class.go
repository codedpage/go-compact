package main

import (
	"compact/4-oops/employee"
	"fmt"
)

func main() {
	ex1()
	ex2()
	ex3()
}

type Employee1 struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func ex1() {
	e := employee.Employee{
		FirstName:   "Rahul",
		LastName:    "kumar",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}

func ex2() {

	e := employee.New("Manoj", "kumar", 30, 10)
	e.LeavesRemaining()
}

func ex3() {
	e := Employee1{"aa", "aaa", 30, 20}          //local struct : no need key
	e1 := employee.Employee{"bb", "bbb", 30, 20} //package struct : literal uses unkeyed fields - vscode warning
	e2 := employee.Employee{FirstName: "cc"}     //package struct : it is okay

	fmt.Println(e)
	fmt.Println(e1, e2)

}
