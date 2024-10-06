package materi3

import (
	"fmt"
)

func ForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}
}

func ForLoopWithBreak() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}
}

func ForLoopWithContinue() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Perulangan ke-", i)
	}
}

func ForLoopWithLabel() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Println("Perulangan ke-", i)
		}
	}
}

func ForLoopWithRange() {
	var names = [4]string{"Muhammad", "Fauzan", "Rizki", "Nur"}

	for i, name := range names {
		fmt.Println("Index ke-", i, "=", name)
	}
}

func ForLoopWithMap() {
	var person = map[string]string{
		"name":  "Muhammad Fauzan Rizki Nur",
		"grade": "A",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

func ForLoopWithSlice() {
	var fruits = []string{"apple", "grape", "banana"}

	for i, fruit := range fruits {
		fmt.Println("Index ke-", i, "=", fruit)
	}
}
