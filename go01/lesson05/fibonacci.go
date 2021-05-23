package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var in string

	fmt.Println("Введите целое число:")
	fmt.Scanln(&in)
	num, _ := strconv.Atoi(in)
	printFib(fibr, 0, num)
}

func fibr(n int) int {
	if n < 2 { return 1 }
	return fibr(n-2) + fibr(n-1)
}

type fibfunc func(int) int

func printFib(fib fibfunc, a, b int) {
	for i := a; i < b; i++ {
		fmt.Printf("%d, ", fib(i))
	}
	fmt.Println("...")
}