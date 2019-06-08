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

	prefix := englishPrefixName

	switch language {
	case spanish:
		prefix = spanishPrefixName
	case french:
		prefix = frenchPrefixName
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
