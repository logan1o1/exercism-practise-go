package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
	name string
}
type Portuguese struct {
	name string
}

func (itln Italian) LanguageName() string {
	return "Italian"
}

func (itl Italian) Greet(name string) string {
	itl.name = name
	return fmt.Sprintf("Ciao %s!", itl.name)
}

func (itln Portuguese) LanguageName() string {
	return "Portuguese"
}

func (itl Portuguese) Greet(name string) string {
	itl.name = name
	return fmt.Sprintf("Olá %s!", itl.name)
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
