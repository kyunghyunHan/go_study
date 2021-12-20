
# 💯goto

- goto 키워드는 정해진 레이블로 곧장 이동!!, 보통 프로그래밍에서 goto 키워드는 권장되지 않지만 일부 상황에서는 코드를 간단하게 만들 수 있으므로 적절히 활용하는 것이 중요하다

- goto 키워드는 goto 레이블 형식으로 사용한다. 여기서 레이블 이름은 변수 이름을 짓는 규칙과 같다

- goto 레이블
레이블:

```
goto LABEL // 이동할 레이블을 지정합니다.

LABEL:
       // 실행할 코드를 작성합니다.
```
- 먼저 다음과 같이 여러 에러 상황이 있을 때 if 조건문으로 에러 처리를 해주면 중복되는 코드가 생긴다
```
package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		fmt.Println("Error 1") // 에러 1 상황
		return
	}

	if a == 2 {
		fmt.Println("Error 2") // 에러 2 상황
		return
	}

	if a == 3 {
		fmt.Println("Error 1") // 에러 1 상황, 중복 코드
		return
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return
}
```
```
package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	if a == 2 {
		goto ERROR2 // 에러 2 상황이면 ERROR2 레이블로 이동
	}

	if a == 3 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return

ERROR1: // 에러 처리 1
	fmt.Println("Error 1")
	return

ERROR2: // 에러 처리 2
	fmt.Println("Error 2")
	return
}
```
먼저 레이블을 사용하여 에러 처리 코드를 작성하였으며, 에러 상황일 때는 매번 에러 처리 코드를 작성하지 않고, goto 키워드를 사용하여 레이블로 이동한 뒤 에러를 처리한다

goto 키워드 바로 뒤에서는 코드가 실행되지 않고 넘어가므로 주의해야한다 



```
package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR // 여기까지만 실행하고 ERROR 레이블로 이동
		b := 1
		fmt.Println(b)
	}

ERROR:
	fmt.Println("Error")
}
```
변수 선언 부분과 fmt.Println(b) 부분은 실행되지 않고 넘어가므로 Error 만 출력!!
