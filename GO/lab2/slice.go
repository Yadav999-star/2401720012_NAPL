package lab2

import "fmt"

func slice() {
	values := []int{80, 87, 92}
	fmt.Println("Original slice:", values)

	values = append(values, 20)
	fmt.Println("After adding a value:", values)

	values = append(values[:1], values[2:]...)
	fmt.Println("After deleting the second value:", values)

	values[0] = 1
	fmt.Println("After updating the first value:", values)
}
