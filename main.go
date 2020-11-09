package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type enBot struct{}
type spBot struct{}

func (eb enBot) getGreeting() string {
	return "Hi!"
}

// you can omit the var declaration
// in the receiver (sb spBot) when
// it's not used inside the function
func (spBot) getGreeting() string {
	return "Ola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := enBot{}
	sb := spBot{}
	printGreeting(eb)
	printGreeting(sb)
}
