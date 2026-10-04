package lab2

func Vowels(str string) int {
	countstr := 0
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			countstr++
		}
	}
	return countstr
}
