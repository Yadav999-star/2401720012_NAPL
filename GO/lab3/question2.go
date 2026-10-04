package lab3

import "fmt"
import "time"

type Student struct {
	Name  string
	Age   int
	Marks float32
}

func PointersModify(number *int) {
	*number = 100
}

func QuestionPointers() {
	number := 10
	pointer := &number

	fmt.Println("Value:", number)
	fmt.Println("Address:", pointer)
	fmt.Println("Value through pointer:", *pointer)

	PointersModify(&number)
	fmt.Println("Value after changing through pointer:", number)

	student := Student{}
	student.Name = "Sidharth"
	student.Age = 20
	student.Marks = 90
	fmt.Println("Student:", student)
}
