package main

import "fmt"

func main() {
	words := []string{"go", "is", "fun", "go", "go"}
	m := map[string]int{}
	gCount := map[string]int{}
	for _, w := range words {
		m[w]++
	}

	grades := []string{"A", "B", "A", "F", "C", "B", "A"}
	for _, g := range grades {
		gCount[g]++
	}

	for grade, count := range gCount {
		fmt.Printf("%s : %d\n", grade, count)
	}

	fmt.Println(m)
	fmt.Println(gCount)

}
