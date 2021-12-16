## panic

프로그램이 잘못되어 에러가 발생한 뒤 종료되는 상황을 패닉이라고 합니다. 다음과 같이 배열의 크기보다 큰 인덱스에 접근했을 때 발생하는 에러가 대표적입니다.

```
a := [...]int{1, 2}

for i := 0; i < 3; i++ {  // i는 0, 1, 2까지 증가
	fmt.Println(a[i]) // 배열의 인덱스는 0, 1까지만 있으므로 배열을 벗어난 접근을 하게 됨
}
```
```
1
2
panic: runtime error: index out of range
```
잘못된 부분으로 인해 발생하는 Go 언어 런타임 에러뿐만 아니라 panic 함수를 사용하면 사용자가 직접 에러를 발생시킬 수도 있습니다.
```
func main() {
	panic("Error !!")
	fmt.Println("Hello, world!") // 실행되지 않음
}
```
```
panic: Error !!
```
panic 함수는 문법적인 에러는 아니지만 로직에 따라 에러로 처리하고 싶을 때 사용합니다.

recover 함수를 사용하면 패닉이 발생했을 때 프로그램이 바로 종료되지 않고 예외 처리를 할 수 있으며 다른 언어의 try catch 구문과 비슷하게 동작합니다.

- 변수 := recover()

```
package main

import "fmt"

func f() {
	defer func() {         // recover 함수는 지연 호출로 사용해야 함
		s := recover() // 패닉이 발생해도 프로그램을 종료하지 않음, panic 함수에서 설정한 에러 메시지를 받아옴
		fmt.Println(s)
	}()

	panic("Error !!!") // panic 함수로 에러 메시지 설정, 패닉 발생
}

func main() {
	f()

	fmt.Println("Hello, world!") // 패닉이 발생했지만 계속 실행됨
}
```
```
Error !!!
Hello, world!
```

recover 함수는 panic 함수에서 설정한 에러 메시지를 받아올 수 있습니다. 그리고 recover 함수는 반드시 지연 호출 함수로 사용해야 합니다. 그렇지 않으면 프로그램이 바로 종료 되어버리기 때문에 recover 함수가 실행되지 않습니다. 여기서는 panic 함수에서 설정한 에러 메시지 “Error !!!”를 recover 함수로 받아온 뒤 출력합니다. 그리고 recover 함수로 복구를 했으므로 “Hello, world!”도 정상적으로 출력됩니다.

실제 런타임 에러 상황에서도 recover 함수를 사용하면 프로그램이 종료되지 않고 계속 실행합니다.

먼저 recover 함수를 사용하지 않아서 프로그램이 종료되는 상황입니다.

```
package main

import "fmt"

func f() {
	a := [...]int{1, 2}      // 정수가 2개 저장된 배열

	for i := 0; i < 5; i++ { // 배열 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, world") // 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 부분은 실행되지 않음
}
```
```
1
2
panic: runtime error: index out of range
```
배열 크기를 벗어난 접근을 했기 때문에 런타임 에러가 발생하여 프로그램이 종료됩니다. 하지만 다음과 같이 recover 함수를 사용하면 런타임 에러(패닉)가 발생하더라도 프로그램을 종료하지 않고 계속 실행합니다.
```
package main

import "fmt"

func f() {
	defer func() {
		s := recover()   // recover 함수로 런타임 에러(패닉) 상황을 복구
		fmt.Println(s)
	}()

	a := [...]int{1, 2}      // 정수가 2개 저장된 배열

	for i := 0; i < 5; i++ { // 배열 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, world!") // 런타임 에러가 발생했지만 recover 함수로 
	                             // 복구되었기 때문에 이 부분은 정상적으로 실행됨
}
실행
```
```
1
2
runtime error: index out of range
Hello, world!
```
배열을 벗어난 접근을 하여 런타임 에러가 발생했지만 recover 함수가 패닉 상황을 복구했습니다. 따라서 프로그램이 종료되지 않고 계속 실행되므로 뒤에 오는 fmt.Println도 정상적으로 실행됩니다.
