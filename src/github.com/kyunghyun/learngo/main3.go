package main

import (
	"fmt"
	"strings"
)

func mutiply(a, b int) int {
	return a * b
}

func main3() {
	fmt.Println(mutiply(2, 2))
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main33() {
	totalLength, upperNani := lenAndUpper("nico")
	fmt.Println(totalLength, upperNani)
}

func repeaMe(words string) (length int, upp string) {
	defer fmt.Println("im done") //func실행후 실행

	return
}

func superAdd(nunbers ...int) int {
	total := 0
	for index, number := range nunbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}
func manw() {
	superAdd(1, 2, 3, 4, 5)
}
