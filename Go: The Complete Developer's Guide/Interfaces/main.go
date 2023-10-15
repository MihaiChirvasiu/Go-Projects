package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func ( /*eb*/ englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// we do not use eb in this function so it can be ommitted
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
