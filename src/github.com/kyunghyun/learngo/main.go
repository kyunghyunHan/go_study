package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "안녕하세요! %s에 오신 것을 환영합니다!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
