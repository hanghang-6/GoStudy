package main

const Prename = "Hello,"

func Hello(name string) string {
	if name == "" {
		return "Hello, World!"
	} else {
		return Prename + " " + name
	}
}
