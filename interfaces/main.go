package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot()
	sb := spanishBot()
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eg englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very custom logic
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// Very custom logic
	return "Hola"
}
