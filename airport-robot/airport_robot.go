package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type GermanGreeter struct{}
type Italian struct{}
type Portuguese struct{}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}

func (GermanGreeter) LanguageName() string {
	return "German"
}
func (GermanGreeter) Greet(name string) string {
	return "Hallo " + name + "!"
}

func (Italian) LanguageName() string {
	return "Italian"
}
func (Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}
func (Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
