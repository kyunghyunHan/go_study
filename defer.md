# 👩🏻‍🎓지연호출함수

- defer 함수명()
- defer 함수명(매개변수)

```
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
        //↓ 함수의 호출이 지연됨.
	defer world() // 현재 함수(main())가 끝나기 직전에 호출
	hello()
	hello()
	hello()
}
```
```
Hello
Hello
Hello
world
```
- 여기서 hello 함수보다 world 함수를 먼저 호출했더라도 출력을 해보면 world 함수가 나중에 호출된더. 즉 defer를 사용한 함수는 현재 함수(main())가 끝나기 직전에 호출되므로 다른 함수들 보다 맨 마지막에 호출된다.

```
package main

import (
	"fmt"
	"os"
)

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() // 지연 호출한 file.Close()가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	fmt.Println(string(buf))

	// file.Close() 호출
}

func main() {
	ReadHello()
}
```
os.Open으로 파일을 연 뒤 지연 호출로 file.Close를 호출하면 함수가 끝날 때 무조건 file.Close로 파일을 닫는다. 파일 열기는 성공했지만 중간에 file.Read에서 실패하더라도 지연 호출을 사용했으므로 함수가 종료될 때 file.Close로 파일을 닫는다 특히 이런 방식은 프로그램 흐름상 분기가 많아 에러 처리가 복잡해질때 유용하다
