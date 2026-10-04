package lab1

import "fmt"

func Calculator() {
	for {
		var choice int
		fmt.Println("1. Integer operations")
		fmt.Println("2. Float operations")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			integerOperations()
		case 2:
			floatOperations()
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func integerOperations() {
	var first, second int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&first)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&second)

	fmt.Printf("Addition: %d\n", first+second)
	fmt.Printf("Subtraction: %d\n", first-second)

	if second == 0 {
		fmt.Println("Division by zero is not valid")
		return
	}
	fmt.Printf("Division: %d\n", first/second)
}

func floatOperations() {
	var first, second float32
	fmt.Print("Enter first float: ")
	fmt.Scan(&first)
	fmt.Print("Enter second float: ")
	fmt.Scan(&second)

	fmt.Printf("Addition: %f\n", first+second)
	fmt.Printf("Subtraction: %f\n", first-second)
	fmt.Printf("Multiplication: %f\n", first*second)

	if second == 0 {
		fmt.Println("Division by zero is not valid")
		return
	}
	fmt.Printf("Division: %f\n", first/second)
}
