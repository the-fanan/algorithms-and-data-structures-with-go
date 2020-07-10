package main

import (
	"fmt"
)

func main(){
	fmt.Println(getImportance([]*Employee{
		&Employee{
			Id: 1,
			Importance: 5,
			Subordinates: []int{2,3},
		},
		&Employee{
			Id: 2,
			Importance: 3,
			Subordinates: []int{},
		},
		&Employee{
			Id: 3,
			Importance: 3,
			Subordinates: []int{},
		},
	},1))
}

/**
	You are given a data structure of employee information, which includes the employee's unique id, their importance value and their direct subordinates' id.

	For example, employee 1 is the leader of employee 2, and employee 2 is the leader of employee 3. They have importance value 15, 10 and 5, respectively. Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3 is also a subordinate of employee 1, the relationship is not direct.

	Now given the employee information of a company, and an employee id, you need to return the total importance value of this employee and all their subordinates.
*/
type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mapper := make(map[int]int)
	totalImportance := 0
	for i,e := range employees {
		mapper[e.Id] = i
	}
	totalImportance += getImportanceHelper(&employees, id, &mapper)
	return totalImportance
}

func getImportanceHelper(employees *[]*Employee, id int, mapper *map[int]int) int {
  employee := (*employees)[(*mapper)[id]]
	totalImportance := employee.Importance
	if employee.Subordinates != nil && len(employee.Subordinates) > 0 {
		for _,v := range employee.Subordinates {
			totalImportance += getImportanceHelper(employees, v, mapper)
		}
	}
	return totalImportance
}