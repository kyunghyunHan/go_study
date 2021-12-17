# 유니코드 함수 사용하기



다음은 Unicode 패키지에서 제공하는 함수입니다.

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
유니코드는 전 세계 모든 문자를 표현할 수 있습니다. 따라서 문자를 종류별로 처리할 수 있도록 범위가 지정되어 있습니다. 다음은 대표적인 범위 테이블(unicode.RangeTable)입니다.

- unicode.Latin: 라틴 문자, 로마자, 영문자
- unicode.Hangul: 한글
- unicode.Han: 한자
- unicode.Hiragana, unicode.Katakana: 일본어 히라가나, 카타카나
- unicode.Is 함수는 범위 테이블을 지정하여 문자가 해당 범위 테이블에 포함되는지 확인할 수 있습니다. 예제에서는 문자가 한글, 한자, 영문자인지 확인해보았습니다. 그리고 unicode.In 함수는 문자가 여러 범위 테이블 중에 포함되는지 확인할 수 있습니다.

이번에는 유니코드 문자의 특성을 확인해보겠습니다.

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
- unicode.IsGraphic 함수는 값이 화면에 표시될 수 있는지 확인합니다. 숫자, 영문자, 한글 등은 화면에 글꼴로 표시되므로 true를 리턴하지만 \n같은 제어 문자는 글꼴로 표시되지 않으므로 false를 리턴합니다.
- unicode.IsLetter 함수는 값이 문자인지 확인합니다. 문자라면 true, 문자가 아니라면 false를 리턴합니다.
- unicode.IsDigit 함수는 값이 숫자인지 확인합니다. 숫자라면 true, 숫자가 아니라면 false를 리턴합니다.
- unicode.IsControl 함수는 값이 제어 문자인지 확인합니다. \n같은 제어 문자는 true를 리턴하며 제어 문자가 아니라면 false를 리턴합니다.
- unicode.IsMark 함수는 값이 마크인지 확인합니다. 마크라면 true, 마크가 아니라면 false를 리턴합니다. \u17c9의 모양은 다음 주소에서 확인할 수 있습니다.
- http://www.charbase.com/17c9-unicode-khmer-sign-muusikatoan
- unicode.IsPrint 함수는 값이 Go 언어에서 출력할 수 있는지 확인합니다. 출력할 수 있으면 true, 출력할 수 없으면 false입니다.
- unicode.IsPunct 함수는 값이 문장 부호인지 확인합니다. 문장 부호라면 true, 문장 부호가 아니라면 false를 리턴합니다.
- unicode.IsSpace 함수는 값이 공백(스페이스, 탭, 개행)인지 확인합니다. 공백이라면 true, 공백이 아니라면 false를 리턴합니다.
- unicode.IsSymbol 함수는 값이 심볼인지 확인합니다. ♥같은 심볼은 true를 리턴하며 심볼이 아니라면 false를 리턴합니다.
- unicode.IsUpper, unicode.IsLower 함수는 값이 대문자 또는 소문자인지 확인합니다. 대문자 또는 소문자가 맞다면 true, 아니면 false를 리턴합니다.
