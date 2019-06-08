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

	if language == spanish {
		return spanishPrefixName + name
	}

	if language == french {
		return frenchPrefixName + name
	}

	return englishPrefixName + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
