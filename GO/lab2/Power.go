package lab2


func Power(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result = result * a
	}
	return result
}
