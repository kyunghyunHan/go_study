# 👩🏻‍🎓유니코드 함수 사용하기



Unicode 패키지에서 제공하는 함수

- func Is(rangeTab *RangeTable, r rune) bool: 문자가 지정한 범위 테이블에 포함되는지 확인
- func In(r rune, ranges ...*RangeTable) bool: 문자가 여러 범위 테이블 중에 포함되는지 확인
- func IsGraphic(r rune) bool: 값이 화면에 표시될 수 있는지 확인
- func IsLetter(r rune) bool: 값이 문자인지 확인
- func IsDigit(r rune) bool: 값이 숫자인지 확인
- func IsControl(r rune) bool: 값이 제어 문자인지 확인
- func IsMark(r rune) bool: 값이 마크인지 확인
- func IsPrint(r rune) bool: 값이 Go 언어에서 출력할 수 있는지 확인
- func IsPunct(r rune) bool: 값이 문장 부호인지 확인
- func IsSpace(r rune) bool: 값이 공백인지 확인
- func IsSymbol(r rune) bool: 값이 심볼인지 확인
- func IsUpper(r rune) bool: 값이 대문자인지 확인
- func IsLower(r rune) bool: 값이 소문자인지 확인
먼저 문자가 유니코드인지 확인하는 방법입니다.

```
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1)) // true: r1은 한글이므로 true
	fmt.Println(unicode.Is(unicode.Latin, r1))  // false: r1은 라틴 문자가
	                                            // 아니므로 false

	var r2 rune = '漢'
	fmt.Println(unicode.Is(unicode.Han, r2))    // true: r2는 한자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r2)) // false: r2는 한글이 아니므로 false

	var r3 rune = 'a'
	fmt.Println(unicode.Is(unicode.Latin, r3))  // true: r3은 라틴 문자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r3)) // false: r3은 한글이 아니므로 false

	fmt.Println(unicode.In(r1, unicode.Latin, unicode.Han, unicode.Hangul)) // true: r1은 한글이므로 true
}
```
유니코드는 전 세계 모든 문자를 표현할 수 있다. 따라서 문자를 종류별로 처리할 수 있도록 범위가 지정되어 있다. 다음은 대표적인 범위 테이블(unicode.RangeTable)이다.

- unicode.Latin: 라틴 문자, 로마자, 영문자
- unicode.Hangul: 한글
- unicode.Han: 한자
- unicode.Hiragana, unicode.Katakana: 일본어 히라가나, 카타카나
unicode.Is 함수는 범위 테이블을 지정하여 문자가 해당 범위 테이블에 포함되는지 확인할 수 있다. 예제에서는 문자가 한글, 한자, 영문자인지 확인하였다. 그리고 unicode.In 함수는 문자가 여러 범위 테이블 중에 포함되는지 확인할 수 있다

 유니코드 문자의 특성

```
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsGraphic('1'))  // true: 1은 화면에 표시되는 숫자이므로 true
	fmt.Println(unicode.IsGraphic('a'))  // true: a는 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('한')) // true: '한'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('漢')) // true: '漢'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('\n')) // false: \n 화면에 표시되는 문자가 아니므로 false

	fmt.Println(unicode.IsLetter('a')) // true: a는 문자이므로 true
	fmt.Println(unicode.IsLetter('1')) // false: 1은 문자가 아니므로 false

	fmt.Println(unicode.IsDigit('1'))     // true: 1은 숫자이므로 true
	fmt.Println(unicode.IsControl('\n'))  // true: \n은 제어 문자이므로 true
	fmt.Println(unicode.IsMark('\u17c9')) // true: \u17c9는 마크이므로 true

	fmt.Println(unicode.IsPrint('1')) // true: 1은 Go 언어에서 표시할 수 있으므로 true
	fmt.Println(unicode.IsPunct('.')) // true: .은 문장 부호이므로 true

	fmt.Println(unicode.IsSpace(' '))   // true: ' '는 공백이므로 true
	fmt.Println(unicode.IsSymbol('♥')) // true: ♥는 심볼이므로 true

	fmt.Println(unicode.IsUpper('A')) // true: A는 대문자이므로 true
	fmt.Println(unicode.IsLower('a')) // true: a는 소문자이므로 true
}
```
- unicode.IsGraphic 함수는 값이 화면에 표시될 수 있는지 확인. 숫자, 영문자, 한글 등은 화면에 글꼴로 표시되므로 true를 리턴하지만 \n같은 제어 문자는 글꼴로 표시되지 않으므로 false를 리턴.
- unicode.IsLetter 함수는 값이 문자인지 확인. 문자라면 true, 문자가 아니라면 false를 리턴
- unicode.IsDigit 함수는 값이 숫자인지 확인. 숫자라면 true, 숫자가 아니라면 false를 리턴.
- unicode.IsControl 함수는 값이 제어 문자인지 확인. \n같은 제어 문자는 true를 리턴하며 제어 문자가 아니라면 false를 리턴.
- unicode.IsMark 함수는 값이 마크인지 확인. 마크라면 true, 마크가 아니라면 false를 리턴. \u17c9의 모양은 다음 주소에서 확인
- http://www.charbase.com/17c9-unicode-khmer-sign-muusikatoan
- unicode.IsPrint 함수는 값이 Go 언어에서 출력할 수 있는지 확인. 출력할 수 있으면 true, 출력할 수 없으면 false
- unicode.IsPunct 함수는 값이 문장 부호인지 확인 문장 부호라면 true, 문장 부호가 아니라면 false를 리턴.
- unicode.IsSpace 함수는 값이 공백(스페이스, 탭, 개행)인지 확인. 공백이라면 true, 공백이 아니라면 false를 리턴
- unicode.IsSymbol 함수는 값이 심볼인지 확인. ♥같은 심볼은 true를 리턴하며 심볼이 아니라면 false를 리턴
- unicode.IsUpper, unicode.IsLower 함수는 값이 대문자 또는 소문자인지 확인. 대문자 또는 소문자가 맞다면 true, 아니면 false를 리턴


## 👩🏻‍🎓UTF-8 함수 사용하기

- UTF-8은 유니코드를 저장하거나 전송할 때 사용하는 인코딩 방식 중의 하나이다(UTF-8 이외에도 UTF-7, UTF-16, UTF-32 등 다양한 방식이 있다). Go 언어에서는 UTF-8을 주로 사용!!

다음은 unicode/utf8 패키지에서 제공하는 함수

- func RuneLen(r rune) int: 문자의 바이트 수를 구함
- func RuneCountInString(s string) (n int): 문자열의 실제 길이를 구함
- func DecodeRune(p []byte) (r rune, size int): byte 슬라이스에서 첫 글자를 디코딩함
- func DecodeLastRune(p []byte) (r rune, size int): byte 슬라이스에서 마지막 글자를 디코딩함
- func DecodeRuneInString(s string) (r rune, size int): 문자열에서 첫 글자를 디코딩함
- func DecodeLastRuneInString(s string) (r rune, size int): 문자열에서 마지막 글자를 디코딩함
- func Valid(p []byte) bool: byte 슬라이스가 UTF-8이 맞는지 확인
- func ValidRune(r rune) bool: rune 변수에 저장된 값이 UTF-8이 맞는지 확인
- func ValidString(s string) bool: 문자열이 UTF-8이 맞는지 확인
영문자를 저장할 때는 문자 하나가 1바이트를 차지, 하지만 한글, 한자, 일본어 같은 문자는 1바이트로 저장할 수가 없기 때문에 보통 2바이트로 저장(Multi Byte Character Set, MBCS). UTF-8은 가변 길이 문자 인코딩 방식이라 문자를 저장할 때 1바이트에서 4바이트까지 사용하며 한글은 3바이트로 저장

먼저 한글 글자 하나의 길이(바이트 수)를 구현

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "한"
	fmt.Println(len(s)) // 3: 한글은 3바이트로 저장하므로 3

	var r rune = '한'
	fmt.Println(utf8.RuneLen(r)) // 3: 한글은 3바이트로 저장하므로 3
}
```

- len 함수를 사용하면 string 변수에 저장된 한글의 바이트 수를 구할 수 있다. 그리고 utf8.RuneLen 함수를 사용하면 rune 변수에 저장된 한글 글자 하나의 바이트 수를 구할 수 있다. 여기서 한글 글자 하나는 3바이트를 차지하므로 3이 출력.

이제 한글 문자열에서 바이트 수가 아닌 실제 길이(글자 개수)를 구현

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(s1)) // 5: "안녕하세요"의 실제 길이는 5
}
```
- utf8.RuneCountInString 함수를 사용하면 문자열의 실제 글자 개수를 구할 수 있다.

 한글 문자열에서 글자를 디코딩

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", r, size) // 안 3: "안녕하세요"의 첫 글자를 디코딩하여 '안', 바이트 수 3

	r, size = utf8.DecodeRune(b[3:]) // '안'의 길이가 3이므로 인덱스 3부터 부분 슬라이스를 만들면 "녕하세요"가 됨
	fmt.Printf("%c %d\n", r, size)   // 녕 3: "녕하세요"를 첫 글자를 디코딩하여 '녕', 바이트 수 3

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r, size) // 요 3: "안녕하세요"의 마지막 글자를 디코딩하여 '요', 바이트 수 3

	// '요'의 길이가 3이므로 // 문자열 길이-3을 하여 부분 슬라이스를 만들면
	// "안녕하세"가 됨
	r, size = utf8.DecodeLastRune(b[:len(b)-3])

	fmt.Printf("%c %d\n", r, size) // 세 3: "안녕하세"의 마지막 글자를 디코딩하여 '세', 바이트 수 3
}
```
```
안 3
녕 3
요 3
세 3
```

- utf8.DecodeRune 함수는 byte 슬라이스에서 첫 글자를 디코딩하여 리턴하고, 디코딩된 바이트 수도 리턴, 글자 하나만 디코딩하며 다음 글자를 디코딩하려면 b[3:]과 같이 부분 슬라이스로 만들어주면 된다.
- utf8.DecodeLastRune 함수는 byte 슬라이스에서 마지막 글자를 디코딩하여 리턴하고, 디코딩된 바이트 수도 리턴,. 글자 하나만 디코딩하며 다른 글자를 디코딩하려면 b[:len(b)-3]과 같이 부분 슬라이스로 만들어주면된다.
문자열의 첫 글자와 마지막 글자를 구하려면 어떻게 해야 할까? 다음은 영문 문자열의 첫 글자와 마지막 글자를 구한다

```
package main

import "fmt"

func main() {
	s := "Hello, world!"

	fmt.Printf("%c\n", s[0])        // H: 인덱스 0이 첫 번째 글자
	fmt.Printf("%c\n", s[len(s)-1]) // !: 문자열 길이에서 1을 뺀 인덱스가 마지막 글자
}
```
- 영문 문자열은 인덱스 0이 첫 번째 글자이고 문자열 길이에서 1을 뺀 인덱스(인덱스는 0부터 시작하므로)가 마지막 글자이다.

- 한글 문자열은 UTF-8에서 3바이트로 저장되므로 인덱스로 접근하면 한글이 정상적으로 출력되지 않는다.

```
package main

import "fmt"

func main() {
	s := "안녕하세요"

	fmt.Printf("%c\n", s[0])        // 3바이트 중 1바이트만 출력하므로 한글이 정상적으로 출력되지 않음
	fmt.Printf("%c\n", s[len(s)-1]) // 3바이트 중 1바이트만 출력하므로 한글이 정상적으로 출력되지 않음
}
```
따라서 한글 문자열의 첫 글자와 마지막 글자를 구하려면 다음과 같이 utf8.DecodeRuneInString, utf8.DecodeLastRuneInString 함수를 사용하면 된다.

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "안녕하세요"

	r, _ := utf8.DecodeRuneInString(s) // UTF-8 문자열의 첫 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)              // 안: 문자열의 첫 글자

	r, _ = utf8.DecodeLastRuneInString(s) // UTF-8 문자열의 첫 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)                 // 요: 문자열의 마지막 글자
}
```
- utf8.DecodeRuneInString 함수는 UTF-8 문자열에서 첫 글자를 디코딩하여 리턴하고, 디코딩된 바이트 수도 리턴
- utf8.DecodeLastRuneInString 함수는 UTF-8 문자열에서 마지막 글자를 디코딩하여 리턴하고, 디코딩된 바이트 수도 리턴
값이나 문자열이 UTF-8이 맞는지 확인하는 방법은 다음과 같다

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var b1 []byte = []byte("안녕하세요")
	fmt.Println(utf8.Valid(b1)) // true: "안녕하세요"는 UTF-8이 맞으므로 true
	var b2 []byte = []byte{0xff, 0xf1, 0xc1}
	fmt.Println(utf8.Valid(b2)) // false: 0xff 0xf1 0xc1은 UTF-8이 아니므로 false

	var r1 rune = '한'
	fmt.Println(utf8.ValidRune(r1)) // true: '한'은 UTF-8이 맞으므로 true
	var r2 rune = 0x11111111
	fmt.Println(utf8.ValidRune(r2)) // false: 0x11111111은 UTF-8이 아니므로 false

	var s1 string = "한글"
	fmt.Println(utf8.ValidString(s1)) // true: "한글"은 UTF-8이 맞으므로 true
	var s2 string = string([]byte{0xff, 0xf1, 0xc1})
	fmt.Println(utf8.ValidString(s2)) // false: 0xff 0xf1 0xc1은 UTF-8이 아니므로 false
}
```
- utf8.Valid 함수는 byte 슬라이스가 UTF-8이 맞는지 확인. UTF-8이 맞으면 true, 아니면 false를 리턴
- utf8.ValidRune 함수는 rune 변수에 저장된 값이 UTF-8이 맞는지 확인. UTF-8이 맞으면 true, 아니면 false를 리턴
- utf8.ValidString 함수는 문자열이 UTF-8이 맞는지 확인. UTF-8이 맞으면 true, 아니면 false를 리턴
