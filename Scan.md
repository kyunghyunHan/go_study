# 출력함수
지금까지 콘솔(터미널)에 문자열을 출력했습니다. 이번에는 콘솔에서 입력을 받아보겠습니다.

다음은 fmt 패키지에서 제공하는 표준 입력 함수입니다.

- func Scan(a ...interface{}) (n int, err error): 콘솔에서 공백, 새 줄로 구분하여 입력을 받음
- func Scanln(a ...interface{}) (n int, err error): 콘솔에서 공백으로 구분하여 입력을 받음
- func Scanf(format string, a ...interface{}) (n int, err error): 콘솔에서 형식을 지정하여 입력을 받음

```
package main

import "fmt"

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2) // fmt.Scan 함수의 두 번째 리턴값은 생략
	fmt.Println("입력 개수:", n)
	fmt.Println(s1, s2)
}
```

fmt.Scan 함수는 표준 입력에서 입력받은 문자, 숫자를 변수에 저장합니다. 입력받을 개수만큼 매개변수에 변수를 넣어주면 되며 반드시 레퍼런스(포인터) 형태로 넣습니다. 그리고 입력이 성공한 개수를 리턴합니다.

소스 파일을 컴파일하여 콘솔(터미널)에서 실행합니다. 그리고 다음과 같이 문자열을 공백으로 구분하여 입력하면 됩니다.

```
$ go build
$ ./scan
Hello, world! (입력)
입력 개수: 2
Hello, world!
```
fmt.Scan 함수는 공백뿐만 아니라 새 줄(Newline)으로도 값을 구분할 수 있으므로 다음과 같이 문자열을 두 줄로 입력해도 됩니다. 만약 fmt.Scan 함수에 매개변수를 4개 넣었다면 4번 입력을 받을 것입니다.
```
$ ./scan
Hello, (입력)
world! (입력)
입력 개수: 2
Hello, world!
```
fmt.Scanln은 한 줄에서 공백으로만 값을 구분하며 엔터키를 입력하여 새 줄로 넘어가면 입력은 종료됩니다
```
var s1, s2 string
n, _ := fmt.Scanln(&s1, &s2)
fmt.Println("입력 개수:", n)
fmt.Println(s1, s2)
```
```
$ go build
$ ./scan
Hello, world! (입력)
입력 개수: 2
Hello, world!
```
```
package main

import "fmt"

func main() {
	var num1, num2 int
	n, _ := fmt.Scanf("%d,%d", &num1, &num2) // 정수형으로 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2)
}
```
```
fmt.Scanf 함수의 첫 번째 매개변수에 형식 지정자를 이용하여 입력받을 형식을 설정합니다. 그리고 두 번째 매개변수부터는 입력받을 변수를 레퍼런스 형식으로 넣습니다. 여기서 % (형식 지정자)개수와 변수의 개수가 같아야 합니다.

소스 파일을 컴파일하여 콘솔(터미널)에서 실행한 뒤 1,2처럼 숫자를 , (콤마)로 구분하여 입력합니다.
```
```
$ go build
$ ./scan
1,2 (입력)
입력 개수: 2
1 2
```
