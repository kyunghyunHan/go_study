package main

import "fmt"

var name bool = false

//밖에서 는  := 안댐

func main2() {
	name := "hyun" //함수 내부에서만 가능
	fmt.Println(name)
}
