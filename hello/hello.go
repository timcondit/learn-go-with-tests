package main

import "fmt"

// Tiny performance improvement
const englishHelloPrefix = "Hello, "

// Separating "domain" code (sending a string) ..
func Hello(name string) string {
	return englishHelloPrefix + name
}

// .. from side effects (printing to stdout)
func main() {
	fmt.Println(Hello("world"))
}
