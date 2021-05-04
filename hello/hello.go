package main

import "fmt"

// Separating "domain" code (sending a string) ..
func Hello(name string) string {
	return "Hello, " + name
}

// .. from side effects (printing to stdout)
func main() {
	fmt.Println(Hello("world"))
}
