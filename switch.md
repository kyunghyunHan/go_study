# Switch

다양한 조건을 if, else if 조건문으로 나열하는 것보다 switch 분기문을 사용하면 좀 더 간단하게 조건을 표현할 수 있습니다.

```
switch 변수 {
case 값1:
	// 값1일 때 실행할 코드를 작성합니다.
case 값2:
	// 값2일 때 실행할 코드를 작성합니다.
default:
	// 모든 case에 해당하지 않을 때 실행할 코드를 작성합니다.
}
```

각 case는 처음부터 순서대로 값을 판단하며 값이 일치하면 해당 코드를 실행한 뒤 switch 분기문을 중단합니다. 그리고 모든 case에 해당하지 않을 때는 default:를 실행합니다. 한가지 특이한 점은 Go 언어의 switch 분기문은 C, C++와는 달리 case에서 break 키워드를 생략합니다.

이제 switch 분기문을 사용하여 변수 i가 0, 1, 2, 3, 4일 때 각 숫자를 표시하고 아무것도 해당되지 않을 때는 -1을 표시해보겠습니다.

```
i := 3

switch i {             // 값을 판단할 변수 설정
case 0:                // 각 조건에 일치하는
	fmt.Println(0) // 코드를 실행합니다.
case 1:
	fmt.Println(1)
case 2:
	fmt.Println(2)
case 3:                // 3과 변수의 값이 일치하므로
	fmt.Println(3) // 이 부분을 실행하고 이후 실행을 중단
case 4:
	fmt.Println(4)
default:               // 모든 case에 해당하지 않을 때 실행
	fmt.Println(-1)
}
```

switch에 i를 설정했으므로 i의 값에 따라 각 case가 실행됩니다. 여기서는 i에 3이 들어있으므로 case 3:이 실행되어 3이 출력됩니다.

case는 숫자뿐만 아니라 문자열도 값으로 사용할 수 있습니다.

```
s := "world"

switch s {                   // 값을 판단할 변수 설정
case "hello":                // 각 조건에 일치하는
	fmt.Println("hello") // 코드를 실행합니다.
case "world":                // 문자열 "world"와 변수의 값이 일치하므로
	fmt.Println("world") // 이 부분을 실행하고 이후 실행을 중단
default:
	fmt.Println("일치하는 문자열이 없습니다.")
}
```
결과
```
world
```
## break
```
s := "Hello"
i := 2

switch i {                             // 값을 판단할 변수 설정
case 1:
	fmt.Println(1)
case 2:                                // i가 2이고
	if s == "Hello" {              // s가 "Hello"이면
		fmt.Println("Hello 2") // Hello 2를 출력하고
		break                  // switch 분기문 실행을 중단
	}

	fmt.Println(2)
}
```
결과
```
Hello 2
```



## fallthrough 
특정 case의 문장을 실행한 뒤 다음 case의 문장을 실행하고 싶을 때는 fallthrough 키워드를 사용합니다. 마치 C, C++의 switch 분기문에서 break 키워드를 생략한 것처럼 동작합니다. 단, 맨 마지막 case에는 fallthrough 키워드를 사용할 수 없습니다.

```
i := 3

switch i {                    // 값을 판단할 변수 설정
case 4:                       // 각 조건에 일치하는
	fmt.Println("4 이상") // 코드를 실행합니다.
	fallthrough
case 3:                       // 3과 변수의 값이 일치하므로
	fmt.Println("3 이상") // 이 부분을 실행
	fallthrough           // fallthrough를 사용했으므로 아래 case를 모두 실행
case 2:
	fmt.Println("2 이상") // 실행
	fallthrough
case 1:
	fmt.Println("1 이상") // 실행
	fallthrough
case 0:
	fmt.Println("0 이상") // 실행, 마지막 case에는 fallthrough를 사용할 수 없음
}
```
결과
```
3 이상
2 이상
1 이상
0 이상
```
## 여러 조건을 함께 처리하기

```
i := 3

switch i {
case 2, 4, 6: // i가 2, 4, 6일 때
	fmt.Println("짝수")
case 1, 3, 5: // i가 1, 3, 5일 때
	fmt.Println("홀수")
}
```
## 조건식으로 분기하기
switch 키워드 다음에 판별할 변수를 지정하지 않고 case에서 조건식만으로 문장을 실행할 수도 있습니다. 즉 C, C++에서는 case에서 숫자나 열거형 값만 사용할 수 있었지만 Go 언어에서는 조건식도 사용할 수 있습니다.

```
i := 7

switch {               // case에 조건식을 지정했으므로 판단할 변수는 생략
case i >= 5 && i < 10: // i가 5보다 크거나 같으면서 10보다 작을 때
	fmt.Println("5 이상, 10 미만")
case i >= 0 && i < 5:  // i가 0보다 크거나 같으면서 5보다 작을 때
	fmt.Println("0 이상, 5 미만")
}
```
switch 분기문 안에서 함수를 실행하고 결괏값으로 분기를 할 수 있습니다. 이때는 함수를 호출하고 뒤에 ; (세미콜론)을 붙여줍니다. 또한, case에서는 값으로 분기할 수 없고 조건식만 사용할 수 있습니다.

- math/rand: 무작위(랜덤) 패키지입니다.
- Seed: 시드 값을 설정하는 함수입니다. 여기서는 time 패키지를 사용하여 현재 시간을 설정하였습니다.
- Intn: 랜덤 값을 생성합니다. 랜덤 값의 범위는 0부터 매개변수로 설정한 값까지 입니다.
- time: 시간 패키지입니다.
- Now: 현재 시간을 구하는 함수입니다.
- UnixNano: 유닉스 시간을 나노 초 단위로 리턴합니다.

```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())      // 현재 시간으로 Seed 값 설정
	switch i := rand.Intn(10); {          // rand.Intn 함수를 실행한 뒤 i에 대입
	case i >= 3 && i < 6:                 // i가 3보다 크거나 같으면서 6보다 작을 때
		fmt.Println("3 이상, 6 미만") // 코드 실행
	case i == 9:                          // i가 9일 때
		fmt.Println("9")              // 코드 실행
	default:                              // 모든 case에 해당하지 않을 때
		fmt.Println(i)                // 코드 실행
	}
}
```

