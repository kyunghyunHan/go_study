# 변수

## 변수선언
```go
var x int
var arr [5]int
var p *int //정수 포인터
```

## 자료형 추론
```go
i:=10 //새로운 정수형
s:="hello" //새로운 문자열
i =20 //i 갑 변경
j=30 //불가능
i="hello" //불가능
i:=30 //불가능 
i:="hi"  //불가능
```
```go
pakage main

impot "fmt"

func main(){
var a= 10
b:="little"
b+="Gophers"
fmt.Println("hello ,playground",a,b)
}
```
