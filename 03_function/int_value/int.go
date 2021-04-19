package main

import "fmt"

func getSum(int1, int2 int) int {
	return int1 - int2
}

func main() {
	fmt.Println(getSum(100, 50))
}
