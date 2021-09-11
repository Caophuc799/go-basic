package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting(n int) string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Holla!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
