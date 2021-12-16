# 채널

채널(channel)은 고루틴끼리 데이터를 주고 받고, 실행 흐름을 제어하는 기능입니다. 모든 타입을 채널로 사용할 수 있습니다. 그리고 채널 자체는 값이 아닌 레퍼런스 타입입니다.

먼저 채널을 그림으로 표현하면 다음과 같습니다.
```
고루틴  <-> 채널 <->고루틴
```
- 채널에 값을 보내고 꺼냄

채널은 다음과 같은 형식으로 사용합니다.

make(chan 자료형)
고루틴과 채널을 사용하여 두 정수 값을 더해보겠습니다.
```
package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b          // 채널에 a와 b의 합을 보냄
}

func main() {
	c := make(chan int) // int형 채널 생성

	go sum(1, 2, c)     // sum을 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c            // 채널에서 값을 꺼낸 뒤 n에 대입
	fmt.Println(n)      // 3
}
```
채널을 사용하기 전에는 반드시 make 함수로 공간을 할당해야 합니다. 그리고 이렇게 생성하면 동기 채널(synchronous channel)이 생성됩니다.

다음과 같이 :=를 사용하지 않고, var 키워드로 채널을 선언하고 할당할 수도 있습니다.

```
var c chan int // chan int형 변수 선언
c = make(chan int)
```
sum 함수를 고루틴으로 실행하면서 더할 값과 채널 변수를 넣습니다. 채널을 매개변수로 받는 함수는 반드시 go 키워드를 사용하여 고루틴으로 실행해야 합니다.
```
sum 함수를 고루틴으로 실행하면서 더할 값과 채널 변수를 넣습니다. 채널을 매개변수로 받는 함수는 반드시 go 키워드를 사용하여 고루틴으로 실행해야 합니다.
```
함수에서 채널을 매개변수로 받을 때는 다음과 같은 형식으로 사용합니다.

- 변수명 chan 자료형
```
//                         ↓ int형 채널을 매개변수로 받음
func sum(a int, b int, c chan int) {
	c <- a + b
//        ↑ 채널에 값을 보냄
}
```
채널에 값을 보낼 때는 다음과 같은 형식으로 사용합니다.

채널 <- 값
채널 변수에는 =로 값을 대입하지 않고 <- 연산자를 사용합니다. 여기서는 sum 함수 안에서 a와 b를 더한 값을 채널 c로 보냅니다.

이제 main 함수에서는 채널에서 값을 가져옵니다.

<- 채널
```
n := <-c
```
이번에도 <- 연산자를 사용하지만 순서가 반대로 되어있습니다. 즉 채널 c에서 값을 가져온 뒤 변수 n을 생성하여 대입합니다(fmt.Println(<-c)처럼 변수에 대입하지 않고 바로 사용할 수도 있습니다).

<-c는 채널에서 값이 들어올 때까지 대기합니다. 그리고 채널에 값이 들어오면 대기를 끝내고 다음 코드를 실행합니다. 따라서 채널은 값을 주고 받는 동시에 동기화 역할까지 수행합니다.
요약하자면 다음과 같습니다.

- 채널 <-: 채널에 값을 보냅니다.
- <- 채널: 채널에 값이 들어올 때까지 기다린 뒤 값이 들어오면 값을 가져옵니다.
- 가져온 값은 :=, =를 사용하여 변수에 대입할 수 있습니다.
- 값을 가져오면 다음 코드를 실행합니다.

## 동기채널

이번에는 동기 채널의 특성을 좀 더 자세히 알아보겠습니다. 다음은 고루틴과 메인 함수를 번갈아가면서 실행합니다.

```
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // 동기 채널 생성
	count := 3              // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
			fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
			time.Sleep(1 * time.Second) // 1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done                         // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
		fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
	}
}
```
make(chan bool)처럼 채널을 생성합니다. 여기서는 채널로 값을 주고 받아도 실제로 사용하지 않으므로 자료형은 큰 의미가 없습니다. make 함수에 매개 변수를 하나만 지정했으므로 동기 채널을 생성됩니다. 비동기 채널과 버퍼는 뒤에서 설명하겠습니다.

먼저 고루틴을 생성하고, 반복문을 실행할 때마다 채널 done에 true 값을 보낸 뒤 1초를 기다립니다

```
go func() {
	for i := 0; i < count; i++ {
		done <- true                // 고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
		fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
		time.Sleep(1 * time.Second) // 1초 대기
	}
}()
```
동기 채널이므로 done에 값을 보내면 다른 쪽에서 값을 꺼낼 때까지 대기합니다. 따라서 반복문도 실행되지 않으므로 “고루틴 : 숫자”가 계속 출력되지 않습니다.

이제 메인 함수에서는 반복문을 실행할 때마다 채널 done에서 값을 꺼냅니다.
```
for i := 0; i < count; i++ {
	<-done                         // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
	fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
}
```
<-done 부분에서 채널에 값이 들어올 때까지 대기합니다. 먼저 앞의 고루틴에서 done에 값을 보냈기 때문에 값을 꺼낸 뒤 다음 코드로 진행합니다. 그리고 고루틴 쪽의 대기도 종료되면서 다시 반복문이 실행된 뒤 채널에 값을 보냅니다. 그리고 메인 함수는 채널에서 값을 꺼내고, 다시 고루틴도 채널에 값을 보냅니다. 따라서 다음과 같이 고루틴 → 메인 함수 → 고루틴 → 메인 함수 순서로 실행됩니다.
```
고루틴 :  0
메인 함수 :  0
고루틴 :  1
메인 함수 :  1
고루틴 :  2
메인 함수 :  2
```
동기 채널은 보내는 쪽에서는 값을 받을 때까지 대기하고, 받는 쪽에서는 채널에 값이 들어올 때까지 대기합니다. 따라서, 동기 채널을 활용하면 고루틴의 코드 실행 순서를 제어할 수 있습니다.

## 채널버퍼링
이번에는 채널의 버퍼링에 대해 알아보겠습니다. 다음은 채널의 버퍼가 가득차면 값을 꺼내서 출력합니다.

make(chan 자료형, 버퍼_개수)

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4                 // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 채널에 true를 보냄, 버퍼가 가득차면 대기
			fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
		}
	}()

	for i := 0; i < count; i++ {
		<-done                         // 버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
	}
}

```
채널에 버퍼를 1개 이상 설정하면 비동기 채널(asynchronous channel)이 생성됩니다. 비동기 채널은 보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고 대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기합니다.

고루틴을 생성하고, 반복문을 실행할 때마다 채널 done에 true 값을 보냅니다.
```
go func() {
	for i := 0; i < count; i++ {
		done <- true                // 채널에 true를 보냄, 버퍼가 가득차면 대기
		fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
	}
}()
```
비동기 채널이므로 버퍼가 가득 찰때까지 값을 계속 보냅니다. 여기서는 채널의 버퍼를 2개로 설정했으므로 done에 true를 2번 보낸 뒤 그 다음 루프에서 대기합니다(버퍼가 가득 차면 대기하므로 i가 0, 1일 때 채널에 값을 보낸 뒤 i가 2일때 done <- true에서 대기합니다).

이제 메인 함수에서는 반복문을 실행할 때마다 채널 done에서 값을 꺼냅니다.
```
for i := 0; i < count; i++ {
	<-done                         // 버퍼에 값이 없으면 대기, 값을 꺼냄
	fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
}
```
비동기 채널에 버퍼가 2개이므로 done에는 이미 값이 2개 들어있습니다. 따라서 루프를 두 번 반복하면서 <-done에서 값을 꺼냅니다. 그 다음에는 채널이 비었으므로 실행을 멈추고 대기합니다. 그리고 다시 고루틴 쪽에서 값을 두 번 보내고, 메인 함수에서 두 번 꺼냅니다. 즉 고루틴 → 고루틴 → 메인 함수 → 메인 함수 순서로 실행됩니다.
```
고루틴 :  0
고루틴 :  1
메인 함수 :  0
메인 함수 :  1
고루틴 :  2
고루틴 :  3
메인 함수 :  2
메인 함수 :  3
```

## range,close 사용

이번에는 range 키워드와 close 함수에 대해 알아보겠습니다. 다음은 0부터 4까지 채널에 값을 보내고, 다시 채널에서 값을 꺼내서 출력합니다.
```
package main

import "fmt"

func main() {
	c := make(chan int) // int형 채널을 생성

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄
		}
		close(c)       // 채널을 닫음
	}()

	for i := range c { // range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i) // 꺼낸 값을 출력
	}
}
```
```
0
1
2
3
4
```
for 반복문 안에서 range 키워드를 사용하면 채널이 닫힐 때까지 반복하면서 값을 꺼냅니다. 여기서는 동시에 고루틴 안에서 채널 c에 0부터 4까지 값을 보낸 뒤 close를 사용하여 채널을 닫습니다. 이렇게 하면 range로 0부터 4까지 꺼내고, 값을 출력한 뒤 반복문이 종료됩니다.

다음은 range와 close 함수의 특징입니다.

이미 닫힌 채널에 값을 보내면 패닉이 발생합니다.
채널을 닫으면 range 루프가 종료됩니다.
채널이 열려있고, 값이 들어오지 않는다면 range는 실행되지 않고 계속 대기합니다. 만약 다른 곳에서 채널에 값을 보냈다면(채널에 값이 들어오면) 그때부터 range가 계속 반복됩니다.
정리하자면 range는 채널에 값이 몇 개나 들어올지 모르기 때문에 값이 들어올 때마다 계속 꺼내기 위해 사용합니다.

채널을 가져온 뒤 두 번째 리턴값으로 채널이 닫혔는지 확인할 수 있습니다

```
c := make(chan int, 1)

go func() {
	c <- 1
}()

a, ok := <-c       // 채널이 닫혔는지 확인
fmt.Println(a, ok) // 1 true: 채널은 열려 있고 1을 꺼냄

close(c)           // 채널을 닫음
a, ok = <-c        // 채널이 닫혔는지 확인
fmt.Println(a, ok) // 0 false: 채널이 닫혔음
```
## 보내기 전용 및 받기 전용 채널

보내기 전용 채널과 받기 전용 채널은 값의 흐름이 한 방향으로 고정된 채널입니다.

다음은 0부터 4까지 채널에 값을 보내고, 다시 채널에서 값을 꺼내서 출력합니다. 그리고 반복문이 끝난 뒤 채널에 100을 보낸 뒤 다시 꺼내서 출력합니다.
```
package main

import "fmt"

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100           // 채널에 값을 보냄

	//fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c) // 채널에 값을 꺼냄

	// c <- 1        // 채널에 값을 보내면 컴파일 에러
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
```
```
0
1
2
3
4
100
```
보내기 전용 및 받기 전용 채널은 채널 앞 뒤로 <- 연산자를 붙여서 만듭니다. 보통 함수의 매개변수로 사용하거나, 구조체의 필드로 사용합니다.

- 보내기 전용(send-only): chan<- 자료형 형식입니다. c chan<- int는 int형 보내기 전용 채널 c를 뜻합니다. 보내기 전용 채널은 값을 보낼 수만 있으며 값을 가져오려고 하면 컴파일 에러가 발생합니다.
- 받기 전용(receive-only): <-chan 자료형 형식입니다. c <-chan int는 int형 받기 전용 채널 c를 뜻합니다. 받기 전용 채널은 range 키워드 또는 <- 채널 형식으로 값을 꺼낼 수만 있으며 값을 보내려고 하면 컴파일 에러가 발생합니다.
chan 키워드를 기준으로 <- (화살표)가 붙은 방향을 보면 보내기 전용인지 받기 전용인지 알 수 있습니다. 즉 chan<-은 chan 키워드로 <-가 들어가므로 보내기 전용, <-chan은 chan 키워드에서 <-가 나오고 있으므로 받기 전용 채널입니다.

여기서는 producer 함수는 매개변수로 보내기 전용 채널을 사용하고, consumer 함수는 매개변수로 받기 전용 채널을 사용합니다. 따라서 producer 함수는 값을 보내기만 하고, consumer 함수는 값을 꺼내기만 합니다.

이번에는 채널을 리턴값으로 사용해보겠습니다. 다음은 두 수를 더한 뒤 채널로 리턴합니다.

```
package main

import "fmt"

//                    ↓ 함수의 리턴 값은 int 형 받기 전용 채널
func sum(a, b int) <-chan int {
	out := make(chan int) // 채널 생성
	go func() {
		out <- a + b // 채널에 a와 b의 합을 보냄
	}()
	return out           // 채널 변수 자체를 리턴
}

func main() {
	c := sum(1, 2)   // 채널을 리턴값으로 받아서 c에 대입

	fmt.Println(<-c) // 3: 채널에서 값을 꺼냄
}
```
sum 함수는 받기 전용 채널을 리턴하도록 만들었습니다. 채널을 리턴하려면 먼저 make 함수로 채널을 생성합니다. 그리고 고루틴 안에서 채널에 값을 보낸 뒤 고루틴 바깥에서 채널을 리턴합니다.

sum 함수를 사용하여 채널을 리턴값으로 받았으면 <-c처럼 값을 꺼내면 됩니다.

이번에는 채널만 사용하여 값을 더해보겠습니다.
```
package main

import "fmt"

//                    ↓ 함수의 리턴 값은 int 형 받기 전용 채널
func num(a, b int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		out <- a   // 채널에 a의 값을 보냄
		out <- b   // 채널에 b의 값을 보냄
		close(out) // 채널을 닫음
	}()
	return out // 채널 변수 자체를 리턴
}

//            ↓ 함수의 매개변수는 int형 받기 전용 채널
func sum(c <-chan int) <-chan int {
//                        ↑ 함수의 리턴 값은 int형 받기 전용 채널
	out := make(chan int) // int형 채널 생성
	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼냄
			r = r + i  // 꺼낸 값을 모두 더함
		}
		out <- r           // 더한 결과를 채널에 보냄
	}()
	return out // 채널 변수 자체를 리턴
}

func main() {
	c := num(1, 2) // 1과 2가 들어있는 채널이 리턴됨
	out := sum(c)  // 채널 c를 매개변수에 넘겨서 모두 더함, 더한 값이 들어있는 out 채널을 리턴

	fmt.Println(<-out) // 3: out 채널에서 값을 꺼냄
}
```
```
func num(a, b int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		out <- a   // 채널에 a의 값을 보냄
		out <- b   // 채널에 b의 값을 보냄
		close(out) // 채널을 닫음
	}()
	return out // 채널 변수 자체를 리턴
}
```
이제 채널에는 숫자 두 개가 저장되어 있습니다. 그리고 close 함수로 채널을 닫아서 range 키워드의 반복이 끝나도록 합니다.
다음과 같이 sum 함수에서는 range 키워드로 채널에서 값을 두 개 꺼내서 모두 더합니다. 그리고 더한 값은 리턴용 채널에 보냅니다.
```
func sum(c <-chan int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼냄
			r = r + i  // 꺼낸 값을 모두 더함
		}
		out <- r           // 더한 결과를 채널에 보냄
	}()
	return out // 채널 변수 자체를 리턴
}
```
num 함수가 리턴한 채널에는 1과 2가 들어있습니다. 그리고 이 채널을 sum 함수에 넣어주면 값이 모두 더해집니다. 마지막으로 더한 값은 sum 함수가 리턴한 채널에서 꺼내면 됩니다.
```
c := num(1, 2) // 1과 2가 들어있는 채널이 리턴됨
out := sum(c)  // 채널 c를 매개변수에 넘겨서 모두 더함, 더한 값이 들어있는 out 채널을 리턴

fmt.Println(<-out) // 3: out 채널에서 값을 꺼냄
```
## 셀렉트 사용
Go 언어는 여러 채널을 손쉽게 사용할 수 있도록 select 분기문을 제공합니다.
- select { case <- 채널: 코드 }

```
select {
case <-채널1:
	// 채널1에 값이 들어왔을 때 실행할 코드를 작성합니다.
case <-채널2:
	// 채널2에 값이 들어왔을 때 실행할 코드를 작성합니다.
default:
	// 모든 case의 채널에 값이 들어오지 않았을 때 실행할 코드를 작성합니다.
}
```
select 분기문은 switch 분기문과 비슷하지만 select 키워드 뒤에 검사할 변수를 따로 지정하지 않으며 각 채널에 값이 들어오면 해당 case가 실행됩니다(close 함수로 채널을 닫았을 때도 case가 실행됩니다). 그리고 보통 select를 계속 처리할 수 있도록 for로 반복해줍니다(반복하지 않으면 한 번만 실행하고 끝냅니다).

switch 분기문과 마찬가지로 select 분기문도 default 케이스를 지정할 수 있으며 case에 지정된 채널에 값이 들어오지 않았을 때 즉시 실행됩니다. 단, default에 적절한 처리를 하지 않으면 CPU 코어를 모두 점유하므로 주의합니다.
다음은 채널 2개를 생성하고 100밀리초, 500밀리초 간격으로 숫자와 문자열을 보낸 뒤 꺼내서 출력합니다.
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string 채널 생성

	go func() {
		for {
			c1 <- 10                           // 채널 c1에 10을 보낸 뒤
			time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!를 보낸 뒤
			time.Sleep(500 * time.Millisecond) // 500 밀리초 대기
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:                // 채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
				fmt.Println("c1 :", i) // i 값을 출력
			case s := <-c2:                // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
				fmt.Println("c2 :", s) // s 값을 출력
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
```
```
c2 : Hello, world!
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c2 : Hello, world!
c1 : 10
... (생략)
```

실행을 해보면 select에서 번갈아가면서 10과 Hello, world!가 출력됩니다. 채널 c2에 Hello, world!를 보낸 쪽에서는 500 밀리초 대기하고, 채널 c1에 10을 보낸 쪽에서는 100 밀리초 대기하므로 10이 더 많이 출력됩니다.
case에서는 case i := <-c1:처럼 채널에서 값을 꺼낸 뒤 변수에 바로 저장할 수 있습니다. 만약 꺼낸 값을 사용하지 않는다면 case <-c1:처럼 변수를 생략해도 됩니다.
```
select {
case i := <-c1:                // 채널 c1에 값이 들어왔다면 값을 꺼내서 i에 대입
	fmt.Println("c1 :", i) // i 값을 출력
case s := <-c2:                // 채널 c2에 값이 들어왔다면 값을 꺼내서 s에 대입
	fmt.Println("c2 :", s) // s 값을 출력
}
```
time.After 함수를 사용하면 시간 제한 처리를 할 수 있습니다. time.After는 특정 시간이 지나면 현재 시간을 채널로 보냅니다.
```
select {
case i := <-c1:
	fmt.Println("c1 : ", i)
case s := <-c2:
	fmt.Println("c2 : ", s)
case <-time.After(50 * time.Millisecond): // 50 밀리초 후 현재 시간이 담긴 채널이 리턴됨
	fmt.Println("timeout")
}
```
이처럼 case에서는 time.After와 같이 받기 전용 채널을 리턴하는 함수를 사용할 수 있습니다.

select 분기문은 채널에 값을 보낼 수도 있습니다.

- case 채널 <- 값: 코드
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string 채널 생성

	go func() {
		for {
			i := <-c1                          // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 :", i)             // i 값을 출력
			time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!를 보냄
			time.Sleep(500 * time.Millisecond) // 500 밀리초 대기
		}
	}()

	go func() {
		for { // 무한 루프
			select {
			case c1 <- 10:                 // 매번 채널 c1에 10을 보냄
			case s := <-c2:                // c2에 값이 들어왔을 때는 값을 꺼낸 뒤 s에 대입
				fmt.Println("c2 :", s) // s 값을 출력
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
```
```
c2 : Hello, world!
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c2 : Hello, world!
c1 : 10
... (생략)
```
select 분기문에서 채널에 값을 보내는 case가 있다면 항상 값을 보냅니다. 하지만 채널에 값이 들어왔을 때는 값을 받는 case가 실행됩니다.
여기서는 select에서 매번 채널 c1에 값을 보내지만 채널 c2에 값이 들어오면 c2에서 값을 꺼내서 출력합니다.
```
for { // 무한 루프
	select {
	case c1 <- 10:             // 매번 채널 c1에 10을 보냄
	case s := <-c2:            // c2에 값이 들어왔을 때는 값을 꺼낸 뒤 s에 대입
		fmt.Println("c2 :", s) // s 값을 출력
	}
}
```
예제에서는 보내는 채널과 받는 채널을 두 개 사용했지만 다음과 같이 채널 c1 한 개로 select에서 값을 보내거나 받을 수도 있습니다.
```
있습니다.

c1 := make(chan int) // int형 채널 생성

go func() {
	for {
		i := <-c1                          // 채널 c1에서 값을 꺼낸 뒤 i에 대입
		fmt.Println("c1 :", i)             // i 값을 출력
		time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
	}
}()

go func() {
	for {
		c1 <- 20                           // 채널 c1에 20을 보냄
		time.Sleep(500 * time.Millisecond) // 100 밀리초 대기
	}
}()

go func() {
	for { // 무한 루프
		select {                       // 채널 c1 한 개로 값을 보내거나 받음
		case c1 <- 10:                 // 매번 채널 c1에 10을 보냄
		case i := <-c1:                // c1에 값이 들어왔을 때는 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 :", i) // i 값을 출력
		}
	}
}()

time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
```
```
c1 : 20
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 20
c1 : 10
... (생략)
```
여기서는 매번 채널에 값을 보내지만, select 분기문이 아닌 다른 쪽에서 채널에 값을 보내서 값이 들어왔다면 값을 받는 case가 실행됩니다.

