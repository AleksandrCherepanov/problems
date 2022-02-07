package main

import (
	"fmt"

	"github.com/AleksandrCherepanov/problems/two_sum"
)

func main() {
	numbers := []int{2, 5, 7, 8, 3, 6, 9}
	target := 11
	fmt.Println(two_sum.Execute(numbers, target))
}
