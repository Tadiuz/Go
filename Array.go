package main

import (
	"fmt"
)

func main() {

	//grades := [3]int{97, 85, 93}
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	var students [3]string
	fmt.Printf("students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("students: %v\n", students)
	fmt.Printf("students: %v\n", len(students))

}
