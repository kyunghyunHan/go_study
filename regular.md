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

## 이메일 주소 검사

다음은 regexp 패키지에서 제공하는 정규표현식 컴파일 함수와 검사 함수입니다.

- func Compile(expr string) (*Regexp, error): 정규표현식을 컴파일하여 Regexp 인스턴스를 생성
- func (re *Regexp) MatchString(s string) bool: 문자열이 Regexp 인스턴스에 정의된 정규표현식에 맞는지 검사
실무에서 많이 사용하는 이메일 주소 정규표현식을 알아보겠습니다.
```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString("hello@mail.example.com"))    // true
	fmt.Println(validEmail.MatchString("hello+kr@example.com"))      // true
	fmt.Println(validEmail.MatchString("hello-world@example.co.kr")) // true
	fmt.Println(validEmail.MatchString("hello_1@example.info"))      // true
	fmt.Println(validEmail.MatchString("gildong.hong@e-xample.com")) // true

	fmt.Println(validEmail.MatchString("@example.com"))            // false
	fmt.Println(validEmail.MatchString("hello@example"))           // fasle
	fmt.Println(validEmail.MatchString("hello@example.cooooooom")) // false
}
```
보통 회원 가입 단계에서 이메일 주소가 정상인지 검사할 때가 많습니다. 이 때는 정규표현식을 사용하면 편리합니다. 먼저 이메일의 규칙은 계정@도메인.최상위도메인 형식이며 계정에 +, -, _등의 기호를 붙이기도 합니다. 그리고 도메인에도 - 기호를 사용할 수 있고, 최상위 도메인이 여러 단계일 수 있습니다. 이런 규칙을 정규표현식으로 표현하면 다음과 같습니다.

```
"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$"
```
부분 부분으로 나누어 생각하면 쉽게 이해할 수 있습니다.

- ^[_a-z0-9+-.]+@는 @를 기준으로 계정을 표현합니다. ^가 붙었으므로 무조건 문자열의 맨 앞에 와야합니다. 그리고 [_a-z0-9+-.]+이므로 영문 소문자, 숫자, _, +, -, .으로 계정이 만들어져 있어야 하고, 한 글자라도 있어야 합니다(+). 즉 예제와 같이 hello, hello+kr, hello-world, hello_1, gildong.hong 등을 표현할 수 있습니다.
- [a-z0-9-]+(\\.[a-z0-9-]+)*는 서브 도메인, 도메인 부분입니다. 이렇게 각 부분이 +, *로 구분되는 부분은 괄호로 묶어줍니다. 영문 소문자와 숫자 -로만 서브 도메인, 도메인이 만들어져 있어야 하고, 한 글자라도 있어야 합니다. 여기서 (\.[a-z0-9-]+)*처럼 괄호로 묶은 뒤 *를 사용하면 이 부분은 있어도 되고 없어도 된다는 뜻입니다. 보통 서브 도메인은 없을 수도 있기 때문입니다. 즉, 예제와 같이 mail.example.com, example.co, e-xample 등을 표현할 수 있습니다.
- (\\.[a-z]{2,4})$는 최상위 도메인(TLD) 부분입니다. 소문자만 사용할 수 있고 문자열의 맨 뒤에 와야 합니다. 여기서 {2,4}처럼 중괄호에 숫자를 지정하여 최소 개수, 최대 개수를 표현할 수 있습니다. 따라서 영문 소문자 2개에서 4개까지 입니다. 즉 예제와 같이 .com, .kr, .info 등을 표현할 수 있습니다(실제로는 4글자 이상인 TLD도 있으므로 상황에 맞게 수정하여 사용합니다).
복잡한 정규표현식은 regexp.Compile 함수로 컴파일 해두었다가 나중에 반복해서 사용할 수 있습니다.

## 문자열 검색
다음은 regexp 패키지에서 제공하는 정규표현식 문자열 검색 함수입니다.

- func (re *Regexp) FindString(s string) string: 정규표현식으로 문자열을 검색한 뒤 찾은 문자열을 리턴
- func (re *Regexp) FindStringSubmatch(s string) []string: 정규표현식으로 문자열을 검색한 뒤 찾은 문자열과 괄호로 구분된 하위 항목을 리턴
- func (re *Regexp) FindStringSubmatchIndex(s string) []int: 정규표현식으로 문자열을 검색한 뒤 찾은 문자열의 위치와 괄호로 구분한 하위 항목의 위치를 리턴
- func (re *Regexp) FindAllString(s string, n int) []string: 정규표현식에 해당하는 모든 문자열을 리턴
- func (re *Regexp) FindAllStringIndex(s string, n int) [][]int: 정규표현식에 해당하는 모든 문자열의 위치를 리턴
이번에는 정규표현식을 사용하여 문자열을 검색해보겠습니다.

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.FindString("hello example.com") // 정규표현식으로 문자열 검색
	fmt.Println(s1)                           // .com: 맨 마지막에 오는 ".영문"이므로 .com

	re2, _ := regexp.Compile("([a-zA-Z]+)(\\.[a-zA-Z]+)$")
	s2 := re2.FindStringSubmatch("hello example.com") // 문자열을 검색한 뒤 찾은 문자열과
                                                          // 괄호로 구분된 하위 항목 리턴
	fmt.Println(s2) // [example.com example .com]: 맨 마지막에 오는 "영문.영문"을 찾고
                        // 하위 항목 example과 .com을 리턴

	n2 := re2.FindStringSubmatchIndex("example.com") // 문자열을 검색한 뒤 찾은 문자열의
                                                         // 위치와 괄호로 구분한 하위 항목의 위치를 리턴
	fmt.Println(n2) // [0 11 0 7 7 11]

	re3, _ := regexp.Compile("\\.[a-zA-Z.]+")
	s3 := re3.FindAllString("hello.co.kr world hello.net example.com", -1)
                        // 정규표현식에 해당하는 모든 문자열 검색
	fmt.Println(s3) // [.co.kr .net .com]: ".영문."에 해당하는 모든 문자열(-1)

	s3 = re3.FindAllString("hello.co.kr world hello.net example.com", 2)
                        // 정규표현식에 해당하는 모든 문자열 검색
	fmt.Println(s3) // [.co.kr .net]:  ".영문."에 해당하는 문자열 2개 리턴

	n3 := re3.FindAllStringIndex("hello.co.kr world hello.net example.com", -1)
                        // 정규표현식에 해당하는 모든 문자열의 위치를 리턴
	fmt.Println(n3) // [[5 11] [23 27] [35 39]]: ".영문."에 해당하는 모든 문자열의 위치
}
```
```
.com
[example.com example .com]
[0 11 0 7 7 11]
[.co.kr .net .com]
[.co.kr .net]
[[5 11] [23 27] [35 39]]
```
다음 함수들은 regexp.Compile 함수에 정규표현식을 설정하여 사용합니다.

- FindString 함수는 정규표현식으로 문자열을 검색한 뒤 찾은 문자열을 리턴합니다.
- FindStringSubmatch 함수는 정규표현식으로 문자열을 검색한 뒤 찾은 문자열과 ( ) (괄호)로 구분한 하위 항목도 함께 리턴합니다. 예제와 같이 hello example.com 문자열에서 ([a-zA-Z]+)(\\.[a-zA-Z]+)$처럼 검색하면 찾은 문자열 example.com과 ([a-zA-Z]+)에 해당하는 example 그리고 (\\.[a-zA-Z]+)$에 해당하는 .com을 함께 리턴합니다.
- FindStringSubmatchIndex 함수는 정규표현식으로 문자열을 검색한 뒤 찾은 문자열의 위치와 괄호로 구분한 하위 항목의 위치를 함께 리턴합니다. 예제의 출력 결과 [0 11 0 7 7 11]과 같이 시작, 끝, 시작, 끝... 형태입니다.
- FindAllString 함수는 문자열에서 정규표현식에 해당하는 모든 문자열을 리턴합니다. 두 번째 매개변수는 찾을 개수이며 앞쪽부터 찾습니다. -1을 설정하면 발견되는 모든 문자열을 리턴하고 2를 설정하면 두 개만 리턴합니다.
- FindAllStringIndex 함수는 문자열에서 정규표현식에 해당하는 모든 문자열의 위치를 리턴합니다. 두 번째 매개변수는 찾을 개수이며 앞쪽부터 찾습니다.

## 문자열 조작하기
다음은 regexp 함수에서 제공하는 정규표현식 문자열 조작 함수입니다.

- func (re *Regexp) ReplaceAllString(src, repl string) string: 정규표현식에 해당하는 문자열을 지정된 문자열과 바꿈
- func (re *Regexp) ReplaceAllLiteralString(src, repl string) string: 문자열을 바꿀 때 정규표현식 기호를 무시
- func (re *Regexp) Split(s string, n int) []string: 정규표현식에 해당하는 문자열을 기준으로 문자열을 쪼갬
정규표현식으로 문자열을 바꿔보겠습니다.
```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("hello example.com", ".net")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s1) // hello example.net: 맨 마지막에 오는 ".영문"을 .net으로 바꿈

	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("Hello, world!", "${2}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s2) // world!: ${2}만 있으므로 두 번째로 찾은 문자열 world!만 남김

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("Hello, world!", "${2} ${1}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s3) // world! Hello,: ${2} ${1}로 설정했으므로 두 문자열의 위치를 바꿈

	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
                        // 찾은 문자열에 ${first}, ${second}로 이름을 정함
	s4 := re4.ReplaceAllString("Hello, world!", "${second} ${first}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s4) // world! Hello,: ${second} ${first}로 설정했으므로
                        // 두 문자열의 위치를 바꿈

	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("Hello, world!", "${second} ${first}")
                        // 문자열을 바꿀 때 정규표현식 기호를 무시함
	fmt.Println(s5) // ${second} ${first}: 정규표현식 기호를 무시했으므로
                        // ${second} ${first}가 그대로 표시됨
}

```
```
hello example.net
world!
world! Hello,
world! Hello,
${second} ${first}
```
다음 함수들은 regexp.Compile 함수에 정규표현식을 설정하여 사용합니다.

- ReplaceAllString 함수는 정규표현식에 해당하는 문자열을 지정된 문자열과 바꿉니다. 첫 번째 매개변수는 검색할 문자열이며 두 번째 매개변수는 바꿀 문자열입니다.
- 정규표현식을 ([a-zA-Z,]+) ([a-zA-Z!]+)처럼 괄호로 구분하여 만든 뒤 ReplaceAllString 함수로 문자열을 검색하면 첫 번째로 찾은 문자열은 ${1}이되고 두 번째로 찾은 문자열은 ${2}가 됩니다. 따라서 ReplaceAllString(“Hello, world!”, “${2}”)와 같이 실행하면 두 번째로 찾은 문자열인 world!만 남깁니다.
- ReplaceAllString(“Hello, world!”, “${2} ${1}”)와 같이 실행하면 첫 번째로 찾은 문자열과 두 번째로 찾은 문자열의 위치를 바꿉니다.
- 정규표현식을 (?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)처럼 만들면 찾은 문자열에 이름을 정할 수 있습니다. ?P<이름> 형식이며 첫 번째로 찾은 문자열은 ${first}가 되고 두 번째로 찾은 문자열은 ${second}가 됩니다. 마찬가지로 ReplaceAllString(“Hello, world!”, “${second} ${first}”)와 같이 실행하면 첫 번째로 찾은 문자열과 두 번째로 찾은 문자열의 위치를 바꿉니다.
- ReplaceAllLiteralString 함수는 문자열을 바꿀 때 정규표현식 기호를 무시합니다. 즉 ${second} ${first}는 동작하지 않고 문자 그대로(Literal) 들어갑니다.
다음은 정규표현식으로 문자열을 분리합니다
	
```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.([a-z]+)\\.")
	s1 := re1.Split("mail.hello.net www.example.com ftp.example.org", -1)
                        // 정규표현식에 해당하는 문자열일 기준으로 모든 문자열을 쪼갬(-1)
	fmt.Println(s1) // [mail net www com ftp org]: ".영문."을 기준으로 쪼갠 결과

	re2, _ := regexp.Compile("\\.([a-z]+)\\.")
	s2 := re2.Split("mail.hello.net www.example.com ftp.example.org", 2)
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s2) // [mail net www.example.com ftp.example.org]: 마지막 문자열을
                        // 제외한 1개 문자열을 쪼갬

	re3, _ := regexp.Compile("\\.([a-z]+)\\.")
	s3 := re3.Split("mail.hello.net www.example.com ftp.example.org", 3)
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s3) // [mail net www com ftp.example.org]: 마지막 문자열을 제외한
                        // 2개 문자열을 쪼갬

	re4, _ := regexp.Compile("\\.([a-z]+)\\.")
	s4 := re4.Split("mail.hello.net www.example.com ftp.example.org", 0)
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s4) // []: 아무 문자열도 쪼개지 않음
}

```
	
Split 함수는 정규표현식에 해당하는 문자열을 기준으로 문자열을 쪼갭니다. 여기서는 정규표현식을 \\.([a-z]+)\\.로 만들었으므로 . (점)으로 시작하고 .으로 끝나는 소문자 1개 이상을 기준으로 쪼갭니다. 따라서 mail.hello.net www.example.com ftp.example.org에서 .hello., .example.은 사라지고 mail net www com ftp org만 남습니다.

두 번째 매개변수는 쪼갤 개수입니다. -1을 지정하면 모든 문자열을 쪼개고, 0을 지정하면 아무것도 쪼개지 않습니다. 1개 이상을 지정하면 문자열 개수 중에서 마지막 문자열은 쪼개지 않습니다. 예제처럼 2를 지정하면 mail.hello.net www.example.com ftp.example.org에서 mail.hello.net만 쪼개고 www.example.com는 쪼개지 않습니다. 따라서 mail net www.example.com ftp.example.org가 됩니다.
	
