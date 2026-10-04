package lab2

import "fmt"

func MapQuestion() {
	scores := map[string]int{
		"Java":   90,
		"Python": 85,
		"SQL":    70,
	}

	fmt.Println("Starting scores:", scores)

	scores["English"] = 95
	fmt.Println("After adding English:", scores)

	delete(scores, "Python")
	fmt.Println("After removing Python:", scores)

	fmt.Println("Current subjects and scores:")
	for subject, score := range scores {
		fmt.Printf("%s: %d\n", subject, score)
	}

	score, found := scores["Physics"]
	fmt.Printf("Physics found: %t, score: %d\n", found, score)
}
