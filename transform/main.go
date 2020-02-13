package main

import (
	"fmt"
	"strings"
)

func main() {
	animals := []string{"DoG", "cat", "she3P", "SNAKE"}
	fmt.Println("Animals: ", animals)

	transformedAnimals := applyTransform(animals, strings.ToUpper)
	fmt.Println("Transformed Animals: ", transformedAnimals)
}

type converter func(string) string

func applyTransform(list []string, action converter) []string {
	for i, item := range list {
		list[i] = action(item)
	}
	return list
}
