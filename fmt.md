# 💯fmt 출력함수

종류	설명
- %v	= 모든 값을 사용할 수 있는 기본 지정자이며 슬라이스, 맵, 포인터, 구조체, 인터페이스 등에 사용된다 %+v처럼 +를 붙이면 구조체의 필드도 함께 표시된다
```
fmt.Printf("%+v\n", data) // {a:1 b:2}
```
- %#v = 타입과 값을 함께 표현한다
```
fmt.Printf("%#v\n", m)    // map[string]int{"Hello":1}
fmt.Printf("%#v\n", a)    // []int{1, 2, 3}
fmt.Printf("%#v\n", data) // main.Data{a:1, b:2}
```
- %T	 = 타입을 표현한다
```
fmt.Printf("%T\n", num1) // int
fmt.Printf("%T\n", m)    // map[string]int
fmt.Printf("%T\n", a)    // []int
fmt.Printf("%T\n", data) // main.Data
```
- %%	 = 아무 값도 표현하지 않고, % 문자를 그대로 출력
- %t	 = 불을 표현
- %b	 = 2진수를 표현
- %c	 = 유니코드로 된 문자 하나를 표현
```
fmt.Printf("%c %c\n", 'a', '한') // a 한
```
- %d	 = 10진수를 표현
- %o	 = 8진수를 표현
- %q	 = 문자열을 escape하여 표현(특수문자를 문자 그대로 표현하기 위해 문자 앞에 역슬래시를 붙이는 작업을 escape라고 한다
```
fmt.Printf("%q\n", `"Hello, world!" "안녕"`) // "\"Hello, world!\" \"안녕\""
```
- %x	 = 숫자와 영문 소문자로 16진수를 표현
- %X	 = 숫자와 영문 대문자로 16진수를 표현
- %U	 = 유니코드 형식으로 표현
```
fmt.Printf("%U\n", '\ud55c') // U+D55C
```
- %b	 = 실수를 2의 제곱으로 표현
```
var num2 float32 = 3.2
fmt.Printf("%b\n", num2) // 13421773p-22
```
- %e	 = 실수를 지수 표기법으로 표현하고 e는 소문자이다
```
var num2 float32 = 3.2
fmt.Printf("%e\n", num2) // 3.200000e+00
```
- %E	 = 실수를 지수 표기법으로 표현하고 E는 대문자이다
```
var num2 float32 = 3.2
fmt.Printf("%E\n", num2) // 3.200000E+00
```
- %f
- %F	 = 실수를 소수점 형식으로 표현
- %g
- %g	 = 실수에서 지수가 크면 %e, %E로 표현하고, 지수가 작으면 %f, %F로 표현
```
fmt.Printf("%g %g\n", 1234567.1234567, 1.2) // 1.2345671234567e+06 1.2
```
- %s	 = 문자열을 표현
- %p	 = 포인터, 채널을 표현하며 앞에 0x가 붙는다

