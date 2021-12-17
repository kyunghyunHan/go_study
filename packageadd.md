# 패키지 생성

이번에는 패키지를 직접 만들어서 사용해보겠습니다. GOPATH를 설정하지 않았다면 GOPATH부터 설정합니다. GOPATH를 설정하는 방법은 ‘Unit 3 기본 디렉터리 설정하기’를 참조하기 바랍니다.

먼저 새 패키지를 만든다면 디렉터리 구조와 소스 파일은 다음과 같은 모양이 됩니다.
```
package calc

func Sum(a int, b int) int { // 외부에서 사용할 수 있도록 함수의 첫 글자는 영문 대문자로
	return a + b
}
```
소스 파일의 첫 줄에서 package calc로 설정하여 현재 파일이 calc 패키지에 포함된다는 것을 알려줍니다.

패키지 안에서 함수, 변수, 상수의 이름을 정하는 방법은 두 가지가 있습니다.

첫 글자를 영문 소문자로 지정하면 패키지 안에서만 사용할 수 있습니다. 즉 외부에서 사용할 수 없습니다. 예) sum, max, hello
첫 글자를 영문 대문자로 지정하면 외부에서 사용할 수 있습니다. 예) Sum, Max, Hello
여기서 지정한 함수 이름은 Sum, 즉 대문자로 시작했으므로 외부에서 사용할 수 있습니다.

이제 calc 패키지를 사용해보겠습니다. main 패키지의 main 함수가 있는 소스 파일(GOPATH/src/hello/hello.go)에서 다음과 같이 작성합니다.

```
package main

import (
	"calc" // calc 패키지 가져오기
	"fmt"
)

func main() {
	fmt.Println(calc.Sum(1, 2)) // calc 패키지의 Sum 함수  사용
}
```
import로 calc 패키지를 가져온 뒤 calc.Sum과 같이 패키지의 함수를 사용하면 됩니다.

만약 패키지 디렉터리가 GOPATH/src/hello/calc라면 다음과 같이 사용합니다

```
import "hello/calc"
```
즉 기준이 되는 디렉터리는 GOPATH/src입니다.

패키지를 컴파일하여 라이브러리로 만들려면 go install 명령을 사용합니다. 다음과 같이 패키지가 들어있는 GOPATH/src/calc 디렉터리에서 go install 명령을 실행합니다.
```
~$ cd $GOPATH/src/calc
~/hello_project/src/calc$ go install
```
다음과 같이 go install <패키지 이름> 형식으로 명령을 실행하면 디렉터리의 위치와는 상관 없이 패키지를 컴파일할 수 있습니다.
```
~$ go install calc
```
GOPATH/pkg/운영체제_아키텍처 디렉터리에 calc.a 라이브러리 파일이 생성되어 있습니다.
```
리눅스: GOPATH/pkg/linux_amd64
Mac OS X: GOPATH/pkg/darwin_amd64
Windows: GOPATH/pkg/windows_amd64/calc.a
```
