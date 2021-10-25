package main

import "fmt"

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "안녕하세요! %s에 오신 것을 환영합니다!", r.URL.Path[1:])
// }
// func main() {
// 	http.HandleFunc("/", index)
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }
// func ma() {
// 	something.SayHello()
// }
// func min() {
// 	name := "nico"
// 	fmt.Println(name)
// }

// var name bool = false

// //밖에서 는  := 안댐

// func main2() {
// 	name := "hyun" //함수 내부에서만 가능
// 	fmt.Println(name)
// }

// func mutiply(a, b int) int {
// 	return a * b
// }

// func main3() {
// 	fmt.Println(mutiply(2, 2))
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// func main33() {
// 	totalLength, upperNani := lenAndUpper("nico")
// 	fmt.Println(totalLength, upperNani)
// }

// func repeaMe(words string) (length int, upp string) {
// 	defer fmt.Println("im done") //func실행후 실행

// 	return
// }

func superAdd(nunbers ...int) int {
	total := 0
	for index, number := range nunbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}
func main() {
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
