package main

import "fmt"

const englishPrefixName = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return englishPrefixName + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
