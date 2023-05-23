package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func PrintAthfan() string {
	return "Athfan"
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// In our function signature we have made a named return value (prefix string)
	// This will create a variable called prefix in your function
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return // Naked return
}

func main() {
	fmt.Println(PrintAthfan())
	fmt.Println(Hello("Athfan", ""))
}