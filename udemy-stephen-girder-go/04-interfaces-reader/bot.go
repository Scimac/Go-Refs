package main

type englishBot struct {
	greeting string
}

type spanishBot struct {
	greeting string
}

type bot interface {
	// receiver function used by multiple structs
	// this is how we define interfaces in GO
	getGreeting() string
}

func BotFun() {
	eb := englishBot{greeting: "Hello!"}
	sb := spanishBot{greeting: "Hola!"}

	// unnecessary duplication of logic and code
	printEnglishGreeting(eb)
	printSpanishGreeting(sb)

	// using interfaces to remove duplication of code and logic
	// since all functions in GO need to have a defined type
	// we can use interfaces to define a common type
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return eb.greeting
}

func (sb spanishBot) getGreeting() string {
	return sb.greeting
}

func printEnglishGreeting(eb englishBot) {
	println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	println(sb.getGreeting())
}

func printGreeting(b bot) {
	println("This is a interfaced function")
	println(b.getGreeting())
}
