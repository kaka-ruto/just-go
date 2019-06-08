package main

import "fmt"

const englishPrefixName = "Hello, "
const spanishPrefixName = "Hola, "
const frenchPrefixName = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefixName
	case french:
		prefix = frenchPrefixName
	default:
		prefix = englishPrefixName
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
