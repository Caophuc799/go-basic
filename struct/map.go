package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "red",
	}
	colors["green"] = "green"

	fmt.Println(colors)

	animal := make(map[int]string)

	animal[10] = "Hello"
	animal[11] = "Hello"
	delete(animal, 10)
	fmt.Println(animal[10], animal)
	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("key ", key, "value", value)
	}
}
