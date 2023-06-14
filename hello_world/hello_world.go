package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello"
	spanishPrefix = "Hola"
	frenchPrefix  = "Bonjur"
)

func Hello(name string, lang string) (greeting string) {
	prefix := greetingPrefix(lang)
	if name == "" {
		name = "world"
	}
	return prefix + ", " + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
