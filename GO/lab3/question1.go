package lab3

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) readData() {
	fmt.Print("Enter name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter salary: ")
	fmt.Scan(&p.Salary)
}

func (p Person) printData() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Println("Salary:", p.Salary)
}

func main() {
	var firstPerson, secondPerson Person

	firstPerson.readData()
	firstPerson.printData()

	secondPerson.readData()
	secondPerson.printData()
}
