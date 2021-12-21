 # 👩🏻‍🎓고루틴
 
 - 고루틴(Goroutine)은 함수를 동시에 실행시키는 기능!, 다른 언어의 스레드 생성 방법보다 문법이 간단하고, 스레드보다 운영체제의 리소스를 적게 사용하므로 많은 수의 고루틴을 쉽게 생성할 수 있다
```
package main

import "fmt"

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	go hello() // 함수를 고루틴으로 실행

	fmt.Scanln() // main 함수가 종료되지 않도록 대기
}
```
- 함수를 호출할 때 앞에 go 키워드를 붙이면 해당 함수는 고루틴으로 실행된다. 여기서는 hello 함수를 고루틴으로 실행한다
- 고루틴으로 hello 함수를 실행하면 main 함수와 동시에 실행되기 때문에 hello 함수 안의 fmt.Println이 호출되기 전에 main 함수가 종료되어버린다. 따라서 fmt.Scanln을 사용하여 main 함수가 종료되지 않도록 대기합니다(엔터를 입력하면 대기를 끝냅니다).

- 반복문을 사용하여 고루틴을 100개 생성

```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)          // 랜덤한 숫자 생성
	time.Sleep(time.Duration(r)) // 랜덤한 시간 동안 대기
	fmt.Println(n)               // n 출력
}

func main() {
	for i := 0; i < 100; i++ { // 100번 반복하여
		go hello(i)        // 고루틴 100개 생성
	}

	fmt.Scanln()
}
```
- 고루틴 100개를 동시에 실행하면서 각 고루틴은 랜덤한 시간동안 대기한다
- rand.Intn 함수에 100을 넣으면 0부터 100까지 랜덤한 숫자를 리턴한다
- time.Sleep 함수를 사용하여 설정한 시간동안 대기한다
- 여기서 time.Sleep 함수에 값을 넣을 때는 time.Duration 타입으로 변환해준다

```
 time.Duration
1초를 각 단위로 표현하면 다음과 같습니다.

나노초: 1000000000 (1/1,000,000,000초)
마이크로초: 1000000 (1/1,000,000초)
밀리초: 1000 (1/1,000초)
time.Duration은 int64 타입이며 나노초 단위입니다. Go 언어에서는 각 초단위로 상수를 제공하므로 곱셈을 통해 시간을 표현할 수 있습니다.

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
10초: time.Second * 10
35마이크로초: time.Microsecond * 35
120밀리초: time.Millisecond * 120
고루틴을 종료하려면 함수안에서 return으로 빠져나오거나 runtime.Goexit 함수를 사용하면 됩니다. 단, 리턴값은 사용할 수 없습니다(return에 값이나 변수를 지정하더라도 무시됩니다).
```

## 💯멀티코어 
- Go 언어는 CPU 코어를 한 개만 사용하도록 설정되어 있다 다음은 시스템의 모든 CPU 코어를 사용하는 방법이다

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) { // 익명 함수를 고루틴으로 실행
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
```
- runtime.NumCPU 함수로 현재 시스템의 CPU 코어 개수를 구한 뒤 runtime.GOMAXPROCS 함수에 설정해준다. 이렇게 하면 모든 CPU 코어에서 고루틴을 실행할 수 있다

- runtime.GOMAXPROCS 함수는 CPU 코어 개수를 구하지 않고, 특정 값을 설정해도 된다. 그리고 runtime.GOMAXPROCS 함수에 0을 넣으면 설정 값은 바꾸지 않으며 현재 설정 값만 리턴한다

## 💯클로저 고루틴실행
	
  
  - 함수 안에서 클로저를 정의한 뒤 고루틴으로 실행할 수 있다. 예제의 출력 결과를 좀 더 확실하게 표현하기 위해 CPU 코어는 하나만 사용하도록 설정한다
  
  ```
  package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // CPU를 하나만 사용

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) {          // 익명 함수를 고루틴으로 실행(클로저)
			fmt.Println(s, n) // s와 매개변수로 받은 n 값 출력
		}(i)                      // 반복문의 변수는 매개변수로 넘겨줌
	}

	fmt.Scanln()
}
  ```
  ```
  Hello, world! 0
Hello, world! 1
Hello, world! 2
.. 생략 ..
Hello, world! 97
Hello, world! 98
Hello, world! 99
  ```
  - 클로저를 고루틴으로 실행할 때 반복문 안에서 변수 사용에 주의해야 한다. 예제에서는 반복문으로 증가하는 i를 클로저에서 그대로 사용하지 않고, 매개변수로 넘겨주었다.
  -  일반 클로저는 반복문 안에서 순서대로 실행되지만 고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행된다.
  
  ```
  func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문의 변수를
                                          // 클로저에서 바로 출력
		}()
	}

	fmt.Scanln()
}

  ```
  ```
  Hello, world! 100
Hello, world! 100
Hello, world! 100
... 생략 ...
Hello, world! 100
Hello, world! 100
Hello, world! 100
  ```
  - 이렇게 되면 변수 i는 0부터 99까지 증가하고, 다시 100이 되면서 반복문이 끝난다. 고루틴은 반복문이 완전히 끝난 다음에 생성되므로 고루틴이 생성된 시점의 변수 i의 값은 100이다 따라서 모두 Hello, world! 100이 출력된다(CPU 코어를 하나만 사용했을 때의 상황이며, CPU 코어를 여러 개 사용하면 결과는 조금 달라지지만 의도했던대로 0부터 99까지 빠짐없이 출력되지는 않는다).

- 클로저를 고루틴으로 실행할 때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨준다. 즉 매개변수로 넘겨주는 시점에 해당 변수의 값이 복사되므로 고루틴이 생성될 때 그대로 사용할 수 있다. 또한, CPU 코어를 하나만 사용하든 여러 개 사용하든 상관없이 반복문에 의해 바뀌는 변수는 매개변수로 넘겨주어야 한다.
