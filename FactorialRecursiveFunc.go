package main

import "fmt"

var total = 1

func main() {
	var value int
	fmt.Println("enter a number")
	_, _ = fmt.Scan(&value)
	result(calculate(value))
}

func calculate(number int) int {
	if number == 1 {
		return total
	}
	total *= number
	return calculate(number - 1)
}

func result(t int) {
	fmt.Println(t)
}
