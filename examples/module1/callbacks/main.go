package main

func main() {
	DoOperation(1, increase)
	DoOperation(1, decrease)
}

func increase(a, b int) int {
	println("increase result is:", a+b)
	return a + b
}

func DoOperation(y int, f func(int, int) int) {
	f(y, 1)
}

func decrease(a, b int) int {
	println("decrease result is:", a-b)
	return a - b
}
