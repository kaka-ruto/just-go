package main

import "fmt"

const englishPrefixName = "Hello, "

func Hello(name string) string {
	return englishPrefixName + name
}

func main() {
	fmt.Println(Hello("world"))
}
