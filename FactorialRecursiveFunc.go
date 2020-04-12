package main

import "fmt"

func main() {
	var value int
	fmt.Println("enter a number")
	_, _ = fmt.Scan(&value)
	result(calculate(value))
}

func calculate(number int) int {
	if number == 1 {
		return 1
	}
	return number * calculate(number-1)
}

func result(t int) {
	fmt.Println(t)
}
