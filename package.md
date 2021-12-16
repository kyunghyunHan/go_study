# pakage

Go 언어는 각종 기능과 라이브러리를 패키지로 만들어서 제공합니다.사용하려면 import 사용

```
import "fmt"

```

## 여러패키지 사용
```
import (
	"fmt"
	"runtime"
)
```
## 전역 패키지 사용

import 로 패키지를 사용할떄 이름앞에 .사용하면 전역패키지가 댑니다.
```
import . "fmt" // fmt를 전역 패키지로 가져옴

```

여기서 패키지 여러 개를 전역 패키지로 가져왔을 때 함수, 변수, 상수 이름이 중복될 수 있습니다. 따라서 유닛 테스트 같은 특별한 상황에서만 사용하는 것이 좋습니다.

## 패키지 별칭 사용하기

```
package main

import f "fmt" // fmt를 f로 가져옴

func main() {
	f.Println("Hello, world!")
}
```
import f “fmt” 처럼 패키지 이름 앞에 별칭을 붙여주면 현재 소스 코드 안에서 f.으로 fmt 패키지를 사용할 수 있습니다

```
import _ "fmt" // 패키지를 가져온 뒤 사용하지 않을 때
```
