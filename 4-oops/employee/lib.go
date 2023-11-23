package employee

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) Employee {
	e := Employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}
