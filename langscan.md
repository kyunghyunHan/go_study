# 👩🏻‍🎓문자열 입출력


표준 출력(stdout), 표준 입력(stdin)뿐만 아니라 변수를 문자열로 만들거나 문자열에서 변수로 값을 가져올 수 있다

fmt 패키지에서 제공하는 문자열 입출력 함수

- func Sprint(a ...interface{}) string: 값을 그대로 문자열로 만듦
- func Sprintln(a ...interface{}) string: 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙임
- func Sprintf(format string, a ...interface{}) string: 형식을 지정하여 문자열을 만듦
- func Sscan(str string, a ...interface{}) (n int, err error): 공백, 개행 문자로 구분된 문자열에서 입력을 받음
- func Sscanln(str string, a ...interface{}) (n int, err error): 공백으로 구분된 문자열에서 입력을 받음
- func Sscanf(str string, format string, a ...interface{}) (n int, err error): 문자열에서 형식을 지정하여 입력을 받음
 변수 또는 값을 콘솔(터미널)로 출력하지 않고 문자열을 생성

```
package main

import "fmt"

func main() {
	var s1 string
	s1 = fmt.Sprint(1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만듦
	fmt.Print(s1)

	var s2 string
	s2 = fmt.Sprintln(1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙임
	fmt.Print(s2)

	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 1, 1.1, "Hello, world!") // 형식을 지정하여 문자열로 만듦
	fmt.Print(s3)
}
```
```
1 1.1Hello, world!1 1.1 Hello, world!
1 1.100000 Hello, world!
```
- fmt.Sprint 함수는 주어진 값을 그대로 문자열로 만들어서 리턴한다. 마찬가지로 fmt.Sprintln 함수도 값으로 문자열을 만들지만 문자열 끝에 개행 문자(\n)을 넣어준다

- fmt.Sprintf 함수는 첫 번째 매개변수에 출력할 형식(format)을 지정해준다. 두 번째 매개변수부터는 출력할 값 또는 변수를 나열한다. 여기서 % (형식 지정자)로 값을 출력할 위치와 타입을 정하며 %의 개수와 변수 또는 값의 개수가 같아야 한다. 그리고 출력 포맷에 맞게 문자열을 만들어서 리턴한다.

반대로 문자열에서 입력을 받을 수도 있다.
```
package main

import "fmt"

func main() {
	var num1 int
	var num2 float32
	var s string

	input1 := "1\n1.1\nHello"
	n, _ := fmt.Sscan(input1, &num1, &num2, &s) // 공백, 개행 문자로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n) // 입력 개수: 3
	fmt.Println(num1, num2, s)   // 1 1.1 Hello

	input2 := "1 1.1 Hello"
	n, _ = fmt.Sscanln(input2, &num1, &num2, &s) // 공백으로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)  // 입력 개수: 3
	fmt.Println(num1, num2, s)    // 1 1.1 Hello

	input3 := "1,1.1,Hello"
	n, _ = fmt.Sscanf(input3, "%d,%f,%s", &num1, &num2, &s) // 문자열에서 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n) // 입력 개수: 3
	fmt.Println(num1, num2, s)   // 1 1.1 Hello
}
```
```
입력 개수: 3
1 1.1 Hello
입력 개수: 3
1 1.1 Hello
입력 개수: 3
1 1.1 Hello
```

- fmt.Sscan 함수는 공백이나 개행 문자(\n)로 구분된 문자열에서 입력을 받고, fmt.Sscanln 함수는 공백으로 구분된 문자열에서 입력을 받는다. 첫 번째 매개변수에 문자열을 넣고, 두 번째 매개변수부터는 입력받을 변수를 레퍼런스 형식으로 넣는다.

- fmt.Sscanf 함수는 설정한 형식대로 문자열에서 입력을 받는다. 첫 번째 매개변수에 문자열을 넣고, 두 번째 매개변수에 형식을 설정한다. 그리고 세 번째 매개변수부터는 입력받을 변수를 레퍼런스 형식으로 넣는다. 여기서는 포맷을 %d,%f,%s로 설정했으므로 , (콤마)로 구분된 값을 변수에 저장한다.
