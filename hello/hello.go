package main

import "fmt"

const englishPrefixName = "Hello, "
const spanishPrefixName = "Hola, "
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishPrefixName + name
	}

	return englishPrefixName + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
