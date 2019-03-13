package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//very custom logic for generatin an english greeting
	return "Hi there!"
}

func (pb portugueseBot) getGreeting() string {
	return "Ola!"
}
