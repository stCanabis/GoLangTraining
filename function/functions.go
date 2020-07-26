package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("sum", sum(slice1...))
	fmt.Println(errorTest())
}

func sum(data ...int) int {
	var result int
	for _, val := range data {
		result += val
	}
	return result
}

func errorTest() (int, error) {
	return 32, fmt.Errorf("some error need be fixed")
}
