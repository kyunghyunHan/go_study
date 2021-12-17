# 정규 표현식

정규표현식은 일정한 규칙을 가진 문자열을 표현하는 방법입니다. 많은 문자열 속에서 특정한 규칙을 가진 문자열을 추출하거나, 문자열이 정해진 규칙에 맞는지 판별할 때 사용합니다. Go 언어는 기본 패키지에서 정규표현식을 지원합니다.

다음은 정규표현식의 기본 문법입니다.


|문법	|설명|
|------|---|
|[]	|문자, 숫자 범위를 표현하며 +, -, . 등의 기호를 포함할 수 있습니다.|
|+	|1개 이상의 문자를 표현합니다. 예) a+b는 ab, aab, aaab는 되지만 b는 안됩니다.|
|*	|0개 이상의 문자를 표현합니다. 예) a*b는 b, ab, aab, aaab|
|?	|0개 또는 1개의 문자를 표현합니다. 예) a?b는 b, ab|
|.	|문자 1개만 표현합니다.|
|\\	|정규표현식에서 사용하는 문자를 그대로 표현하려면 앞에 \\를 붙입니다. 예) \\+, \\*|
|^	|해당 문자 범위를 제외합니다.|



|정규표현식	|설명|
|------|---|
|[0-9]|	0부터 9까지 숫자를 표현합니다 |
|[a-z0-9]	|영문 소문자와 숫자를 표현합니다.|
|[a-zA-Z0-9]	|영문 대소문자와 숫자를 표현합니다.|
|[a-zA-Z0-9_]	|영문 대문자, 소문자, 숫자, _를 표현합니다.|
|[^a-z]	|영문 소문자를 제외합니다.|
|[^A-Z]	|영문 대문자를 제외합니다.|
|[^0-9]|숫자를 제외합니다.|
|[A-Fa-f0-9]|	16진수를 표현합니다.|


다음은 regexp 패키지에서 제공하는 정규표현식 검사 함수입니다.

func MatchString(pattern string, s string) (matched bool, err error): 문자열이 정규표현식에 맞는지 검사
간단하게 정규표현식을 사용해보겠습니다.


```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("He", "Hello 100")
	fmt.Println(matched) // true: Hello 100에 He가 포함되므로 true

	matched, _ = regexp.MatchString("H.", "Hi")
	fmt.Println(matched) // true: H 뒤에 글자가 하나가 있으므로 true

	matched, _ = regexp.MatchString("[a-zA-Z0-9]+", "Hello 100")
	fmt.Println(matched) // true: Hello 100은 a부터 z까지, A부터 Z까지 0부터 9까지에 포함되므로 true

	matched, _ = regexp.MatchString("[a-zA-Z0-9]*", "")
	fmt.Println(matched) // true: *는 값이 하나도 안나와도 되므로 true

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "1-1")
	fmt.Println(matched) // true: 1-1는 [0-9]+와 - 그리고 [0-9]+가 합쳐진 것이므로 true

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "1-")
	fmt.Println(matched) // false: 1-는 [0-9]+와 -만 있으므로 false

	matched, _ = regexp.MatchString("[0-9]+/[0-9]*", "1/")
	fmt.Println(matched) // true: 1/는 [0-9]+와 /가 있음. *는 값이 없어도 되므로 true

	matched, _ = regexp.MatchString("[0-9]+\\+[0-9]*", "1+")
	fmt.Println(matched) // true: 1+는 [0-9]+와 +가 있음. *는 값이 없어도 되므로 true

	matched, _ = regexp.MatchString("[^A-Z]+", "hello")
	fmt.Println(matched) // true: hello는 A부터 Z까지에 포함되지 않으므로 true
}
```
regexp.MatchString에 정규표현식과 검사할 문자열을 넣으면 규칙에 맞는지 검사해줍니다.

- He처럼 문자열을 그대로 사용할 수 있습니다.
- . (점)은 글자 하나를 표현합니다. H.이라면 H뒤에 소문자든 대문자든 숫자든 글자 하나만 있으면 됩니다.
- [ ] (대괄호)로 문자, 숫자 범위를 표현합니다. a-z는 모든 소문자, A-Z는 모든 대문자, 0-9는 모든 숫자입니다. 대괄호안에 +, -, . 등의 기호를 포함할 수 있습니다.
- 대괄호 뒤에 +를 붙이면 현재 식이 한 번이라도 만족해야 합니다. *를 붙이면 나와도 되고 안나와도 됩니다.
- 대괄호 바깥에는 -, / 등의 기호를 사용할 수 있고, +, *, . 등 정규표현식에서 사용하는 문자는 앞에 \\를 붙여주어야 합니다.
- [^A-Z]+처럼 대괄호 안에 ^를 넣으면 해당 범위를 제외한다는 뜻입니다. 여기서는 ^A-Z는 모든 대문자를 제외하므로 소문자로만 이루어진 hello는 true가 됩니다.
예제의 정규표현식 규칙과 출력 결과를 비교해보기 바랍니다.

Go 언어는 UTF-8 기반이기 때문에 한글도 정규표현식으로 표현할 수 있습니다.

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("[가-힣]+", "안녕하세요")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("홍[가-힣]+동", "홍길동")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("이재[홍-훈]", "이재홍")
	fmt.Println(matched) // true
}
```
```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var matched, _ = regexp.MatchString("^Hello", "Hello, world!")
	fmt.Println(matched) // true: Hello, world!에서 Hello가 맨 앞에 오므로 true

	matched, _ = regexp.MatchString("^Hello", "Go Hello, world!")
	fmt.Println(matched) // false: Go Hello, world!에서 Hello가 맨 앞에 오지 않으므로 false

	matched, _ = regexp.MatchString("world!$", " Hello, world!")
	fmt.Println(matched) // true: Hello, world!에서 world!가 맨 뒤에 오므로 true

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20)")
	fmt.Println(matched) // true: Hello.SetValue(20)는 .영문(숫자)이므로 true

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20).x")
	fmt.Println(matched) // false: Hello.SetValue(20).x는 .영문(숫자)가 아니므로 false
}
```
- ^는 문자열이 맨 앞에 오는지 검사합니다.
- $는 문자열이 맨 뒤에 오는지 검사합니다.
- \\.[a-zA-Z]+\\([0-9]+\\)$처럼 ( ) (괄호)도 \\를 앞에 붙여서 표현할 수 있습니다. 여기서는 .SetValue(20)처럼 영문으로 된 문자열 뒤에 괄호가 오고, 괄호안에 숫자가 있으면서 맨 뒤에 있는지 검사합니다.
