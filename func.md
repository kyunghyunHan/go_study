# 👩🏻‍🎓함수 사용

## 💯func 사용
```
package main

import "fmt"

func main() {
	hello() // 컴파일 잘 됨
}

func hello() { // 함수 선언 위치는 중요하지 않음
	fmt.Println("Hello, world!") //Hello, world!
}
```
## 💯매개변수,리턴값

```
func sum(a int, b int) int { // int형 매개변수 a, b 그리고 int 형 리턴값을 가지는 함수 정의
	return a + b
}

func main() {
	r := sum(1, 2)
	fmt.Println(r) // 3
}

```
```
func sum(a int, b int) (r int) {
	r = a + b // 리턴값 변수 r에 값 대입
	return    // 리턴값 변수를 사용할 때는 return 뒤에 변수를 지정하지 않음
}

func main() {
	r := sum(1, 2)
	fmt.Println(r) //3
}
```
- 리턴값의 자료형만 지정할 때와는 달리 괄호 안에 리턴값 변수의 이름과 자료형을 지정! , 이렇게 하면 함수 안에서 리턴값 변수를 사용할 수 있다 그리고 리턴값 변수에 값을 대입한 뒤 함수의 맨 마지막에 return을 사용. - 리턴값 변수를 사용할 때는 return 뒤에 리턴할 변수를 지정하지 않는다

## 💯리턴값 여러개 사용

- Go 언어는 함수에서 값을 여러 개 리턴할 수 있다. 지금까지 다른 언어들은 함수에서 값을 여러 개 받기 위해 매개변수에서 레퍼런스, 포인터 타입을 사용하거나, 클래스 및 구조체의 인스턴스를 활용하는 등 다양한 방법을 사용했다. 하지만 Go 언어는 문법적으로 값을 여러 개 리턴할 수 있으므로 코드가 좀 더 짧고 간단하다

```
]func SumAndDiff(a int, b int) (int, int) { // int형 리턴값이 두 개인 함수 정의
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열. 합과 차를 동시에 리턴
}

func main() {
	sum, diff := SumAndDiff(6, 2) // 변수 두 개에 리턴값 두 개를 대입
	fmt.Println(sum, diff)        // 8 4: 합과 차
}
```

- ( ) (괄호)안에 리턴값 자료형을 , (콤마)로 구분하여 여러 개 지정해주면 된다. 함수 안에서 값을 리턴할 때도 return에 값을 , (콤마)로 구분하여 여러 개 지정한다.

- 리턴값을 받을 때는 sum, diff :=처럼 변수를 콤마로 구분하여 여러 개 지정한다. 여기서 리턴값은 함수에서 지정한 순서대로 리턴된다. 따라서 리턴값을 받을 변수 하나만 지정하면 첫 번째 리턴값만 저장된다

- 첫 번째 리턴값은 생략하고 두 번째 리턴값부터 사용하고 싶다면 _ (밑줄 문자)를 사용한다

```
func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열
}

func main() {
	_, diff := SumAndDiff(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff)           // 4: 차
}
```
다음과 같이 리턴값 여러 개 가운데 특정 값만 생략할 수도 있다 리턴값 다섯 개 중에서 첫 번째, 세 번째, 다섯 번째만 사용하고, 나머지는 밑줄 문자를 사용하여 생략할수있다.

```
func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5 // 리턴할 값을 차례대로 나열
}

func main() {
	a, _, c, _, e := hello() // 2, 4번째 리턴값은 생략
	fmt.Println(a, c, e)     // 1 3 5
}
```
값을 여러 개 리턴할 때도 리턴값에 이름을 정할 수 있다.
```
func SumAndDiff(a int, b int) (sum int, diff int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum = a + b  // 리턴값 변수 sum에 합 대입
	diff = a - b // 리턴값 변수 diff에 차 대입
	return
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff) // 8 4
}
```

## 💯가변인자 사용
- 함수의 매개변수 개수가 정해져있지 않고 유동적으로 변하는 형태를 가변인자라고 한다

```
func sum(n ...int) int { // int형 가변인자를 받는 함수 정의
	total := 0
	for _, value := range n { // range로 가변인자의 모든 값을 꺼냄
		total += value   // 꺼낸 값을 모두 더함
	}

	return total
}

func main() {
	r := sum(1, 2, 3, 4, 5)
	fmt.Println(r) // 15
}
```
- 매개변수의 자료형 앞에 ...을 붙여 가변인자로 지정한다. ...int로 지정했으므로 int 형 값을 여러 개 받을 수 있다. 여기서 가변인자로 받은 변수는 슬라이스 타입이므로 range 키워드를 사용하여 값을 꺼내면 된다.

- 가변인자가 슬라이스 타입이므로 슬라이스를 바로 넘겨줄 수도 있다.

```
func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	r := sum(n...) // ...를 사용하여 가변인자에
                        // 슬라이스를 바로 넘겨줌

	fmt.Println(r) // 15
}
```

## 💯재귀호출
함수에서 자기 자신을 다시 호출하는 함수를 재귀함수(Recursive function)라고 한다
```
package main

import "fmt"

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5)) // 120
}
```

## 💯함수를 변수에 저장
```
func sum(a int, b int) int {
	return a + b
}

func main() {
	var hello func(a int, b int) int = sum // 함수를 저장하는 변수를 선언하고 함수 대입
	world := sum                           // 변수 선언과 동시에 함수를 바로 대입

	fmt.Println(hello(1, 2)) // 3: hello 변수에 저장된 sum 함수 호출
	fmt.Println(world(1, 2)) // 3: world 변수에 저장된 sum 함수 호출
}
```
- 함수를 정의한 뒤 미리 선언한 변수에 대입하기만 하면 함수를 변수에 저장할 수 있다. Go 언어는 변수를 선언하고 대입할 때 자동으로 자료형을 추론할 수 있으므로 :=를 사용하면 간단하게 변수에 함수를 저장할 수 있다.

- 변수뿐만 아니라 슬라이스(배열)에도 함수를 간단하게 저장할 수 있다.

- 슬라이스 = []func(매개변수명 자료형) 리턴값_자료형{함수명1, 함수명2}


```
func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := []func(int, int) int{sum, diff} // 함수를 저장할 수 있는 슬라이스를 생성한 뒤 함수로 초기화

	fmt.Println(f[0](1, 2)) // 3: 배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2)) // -1: 배열의 두 번째 요소로 함수 호출
}
```

- [ ] 뒤에 매개변수의 자료형과 리턴값의 자료형을 지정해야 하며 { } (중괄호) 안에 대입할 함수명을 나열하면 된다.

- 마찬가지로 맵에도 함수를 간단하게 저장할 수 있다.

- 맵 := map[키_자료형]func(매개변수명 자료형) 리턴값_자료형{ “키”: 함수명 }

```
func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := map[string]func(int, int) int{ // 함수를 저장할 수 있는 맵을 생성한 뒤 함수로 초기화
		"sum":  sum,
		"diff": diff,
	}

	fmt.Println(f["sum"](1, 2))  // 3: 맵에 sum 키를 지정하여 함수 호출
	fmt.Println(f["diff"](1, 2)) // -1: 맵에 diff 키를 지정하여 함수 호출
}
```

## 💯익명함수


```
package main

import "fmt"

func main() {
	func() { // 함수에 이름이 없음
		fmt.Println("Hello, world!")
	}()

	func(s string) {   // 익명 함수를 정의한 뒤
		fmt.Println(s)
	}("Hello, world!") // 바로 호출

	r := func(a int, b int) int { // 익명 함수를 정의한 뒤
		return a + b
	}(1, 2)                       // 바로 호출하여 리턴값을 변수 r에 저장

	fmt.Println(r) // 3
}
```
- 익명 함수는 함수를 정의한 뒤 ( ) (괄호)를 사용하여 바로 함수를 호출합니다. 이때도 매개변수 및 리턴값을 사용할 수 있습니다. 이런 익명 함수는 코드양을 줄일 수 있으며 뒤에서 다룰 클로저, 지연 호출(defer), 고루틴(go)에서 주로 사용됩니다.

