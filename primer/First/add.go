package main

func add(a int, b int) int {
	return a + b
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	a := numbers[1:]
	println("%v", a)
}
