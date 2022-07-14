package main

import (
	"fmt"
	"strconv"
)

func recursion(q int) {
	str := strconv.Itoa(q)
	fmt.Println("123" + str)

	q2, _ := strconv.Atoi(str)
	if q < 2 {
		recursion(q2 - 1) /* 函数调用自身 */
	}
}

func main() {
	recursion(1)
}
