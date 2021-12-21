# 👩🏻‍🎓ERROR

지금까지 에러가 발생하면 fmt.Println 함수로 출력만 했습니다. 이번에는 fmt.Errorf, log.Fatal 함수를 사용하여 에러를 처리하는 방법에 대해 알아보겠습니다.

다음은 fmt 패키지에서 제공하는 에러 출력 함수입니다.

- func Errorf(format string, a ...interface{}) error: 형식을 지정하여 error 값을 만듦
다음은 log 패키지에서 제공하는 에러 출력 함수입니다.

- func Fatal(v ...interface{}): 에러 로그를 출력하고 프로그램을 완전히 종료
- func Panic(v ...interface{}): 시간과 에러 문자열을 출력한 뒤 패닉을 발생시킴
- func Print(v ...interface{}): 시간과 에러 메시지를 출력하며 프로그램을 종료하지 않음
숫자가 1일때만 Hello 문자열을 리턴하는 helloOne 함수가 있습니다. 여기서 1이 아닐 때는 의도적으로 에러를 발생시켜보겠습니다.

```
package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러를 리턴
}

func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)     // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Fatal(err)   // 에러 문자열이 출력되고 프로그램 종료
	}

	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```
보통 Go 언어의 함수에서는 첫 번째 리턴값으로 결괏값을 사용하고, 두 번째 리턴값으로 에러 값을 사용합니다. 마찬가지로 helloOne 함수도 첫 번째 리턴값은 결과 문자열이고, 두 번째 리턴값은 에러 메시지로 만들었습니다.

```
func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러를 리턴
}
```
error 타입을 리턴할 때는 fmt.Errorf 함수를 사용합니다. 여기에 형식을 지정하여 에러 메시지를 만듭니다. fmt.Errorf 함수의 포맷 사용법은 fmt.Printf 함수와 사용법이 같습니다. 기본 출력 함수와 형식 지정자는 ‘Unit 41 출력 함수 사용하기’를 참조하기 바랍니다.

이제 helloOne 함수를 사용해보겠습니다.
```
func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)     // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Fatal(err)   // 에러 문자열이 출력되고 프로그램 종료
	}

	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```
helloOne 함수에 1을 넣으면 정상적으로 Hello 1 문자열이 리턴되고, error는 nil입니다. 따라서 에러 처리용 if 조건문에 걸리지 않고 넘어갑니다. 하지만 helloOne 함수에 2를 넣으면 에러가 발생합니다. 이때는 log.Fatal 함수로 에러 문자열을 출력하고 프로그램을 완전히 종료합니다(정상 종료이며 실행 파일의 Exit code는 1이 나옵니다).
출력 결과를 보면 다음과 같이 시간과 에러 내용이 출력되지만 프로그램이 완전히 종료되었기 때문에 fmt.Println(s)와 fmt.Println(“Hello, world!”)는 실행되지 않습니다.


```
Hello 1

2015/01/25 09:48:31 2는 1이 아닙니다.
```
프로그램을 정상 종료하지 않고 런타임에 에러를 발생시키려면 다음과 같이 log.Panic 함수를 사용합니다.

```
func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Panic(err) // 에러 문자열이 출력되고 런타임 에러 발생
	}

	// 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```
log.Panic 함수는 시간과 에러 문자열을 출력한 뒤 패닉을 발생시킵니다. **log.Panic 함수 대신 panic 함수를 사용해도 됩니다.

출력 결과를 보면 런타임 패닉이 발생하여 콜스택이 출력됩니다.

```
Hello 1

2015/01/25 10:07:16 2는 1이 아닙니다.
panic: 2는 1이 아닙니다.

goroutine 1 [running]:
runtime.panic(0x483020, 0xc21000a1d0)
        /usr/lib/go/src/pkg/runtime/panic.c:266 +0xb6
log.Panic(0x7f5b80648ef8, 0x1, 0x1)
        /usr/lib/go/src/pkg/log/log.go:307 +0xac
main.main()
        /home/pyrasis/hello_project/src/hello/hello.go:27 +0x263
```
패닉이 발생하더라도 recover 함수로 복구를 사용하면 프로그램을 종료하지 않고 넘어갈 수 있습니다.
```
func main() {
	defer func() {
		// 런타임 에러가 발생하면 recover 함수가 실행됨
		s := recover() // log.Panic 함수에서 출력한 에러 문자열 리턴
		fmt.Println(s) // 에러 문자열 출력
	}()

	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Panic(err) // 에러 문자열이 출력되고 런타임 에러 발생
	}

	// 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}

```
패닉이 발생했을 때 recover 함수는 log.Panic 또는 panic 함수에서 설정한 에러 메시지를 리턴하고, 프로그램을 복구합니다. 자세한 설명은 ‘Unit 27 패닉과 복구 사용하기’를 참조하기 바랍니다.
출력 결과를 보면 프로그램이 종료되지 않고 log.Panic 함수에서 설정한 에러 메시지만 출력됩니다.

```
Hello 1

2는 1이 아닙니다.
2015/01/25 10:11:52 2는 1이 아닙니다.
```
만약 심각한 에러 상황이 아니라면 다음과 같이 log.Print 함수를 사용해도 됩니다.
```
func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Print(err) // 에러가 없으므로 실행되지 않음
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Print(err) // 에러 문자열이 출력됨
	}

	// long.Print 함수는 프로그램을 종료하지 않으므로 이 아래는 계속 실행됨
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```
log.Print 함수는 시간과 에러 메시지를 출력하며 프로그램을 종료하지 않습니다. 단 다음과 같이 프로그램의 실행이 끝난 다음 로그가 출력됩니다.
```
Hello 1


Hello, world!
2015/01/25 10:20:06 2는 1이 아닙니다.
```
## 🔥에러 타입 만들기
에러 타입은 에러를 확장하여 좀 더 자세한 에러 내용을 저장하기 위해 사용합니다. 이번에는 에러 타입을 직접 만들어보겠습니다


```
package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time // 시간
	value int       // 에러를 발생시킨 값
}

func (e HelloOneError) Error() string { // 에러 메시지를 생성하여 리턴하는 Error 함수 구현
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", HelloOneError{time.Now(), n} // 1이 아닐때는 에러 구조체 인스턴스를
                                                // 생성하여 리턴
}

func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Fatal(err) // 에러 문자열이 출력되고 프로그램 종료
	}

	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
```
error 타입이 아니더라도 Error 함수를 구현하면 에러로 사용할 수 있습니다.

```
type HelloOneError struct {
	time  time.Time // 시간
	value int       // 에러를 발생시킨 값
}

func (e HelloOneError) Error() string { // 에러 메시지를 생성하여 리턴하는 Error 함수 구현
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}
```
HelloOneError 구조체에는 시간과 에러를 발생시킨 값을 저장할 수 있습니다. 그리고 Error 함수는 에러 메시지 생성하여 리턴합니다.

helloOne 함수에서는 에러가 발생했을 때 HelloOneError 구조체를 생성하여 리턴합니다. 현재 시간은 time.Now 함수로 구할 수 있습니다.

```
func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", HelloOneError{time.Now(), n} // 1이 아닐때는 에러 구조체 인스턴스를
                                                // 생성하여 리턴
}
```
출력 결과를 보면 Error 함수에서 만든 에러 메시지가 출력됩니다.

```
실행 결과
Hello 1

```
2015/01/25 13:14:59 2015-01-25 13:14:59.3055342 +0900 KST: 2는 1이 아닙니다.
error 타입으로 리턴된 HelloOneError 구조체는 다음과 같이 Type assertion을 해주면 구조체 필드에 접근할 수 있습니다.

```
s, err = helloOne(2)
if err != nil {
	fmt.Println(err.(HelloOneError).value) // 2: Type assertion으로 구조체 필드에 접근
}

```
