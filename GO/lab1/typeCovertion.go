package lab1

import "fmt"

func TypeConvertion() {
	fmt.Println("Hello, Ankit Yadav")

	var number int
	fmt.Println("Default integer value:", number)

	wholeNumber := 32
	decimalNumber := float32(wholeNumber)
	fmt.Println("Integer converted to float:", decimalNumber)

	letterCode := 65
	letter := string(rune(letterCode))
	fmt.Println("Integer converted to character:", letter)
}
