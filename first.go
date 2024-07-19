package main

import "fmt"

const Prename = "Hello,"

func Hello(name string) string {
	if name == "" {
		return "Hello, World!"
	} else {
		return Prename + " " + name
	}
}

func main() {
	fmt.Println("Hello World")
}
