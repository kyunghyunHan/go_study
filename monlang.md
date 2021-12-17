# 문자열

## 문자열 검색하기
문자열을 처리할 때 주로 사용하는 기능은 문자열 검색입니다.

다음은 strings 패키지에서 제공하는 문자열 검색 함수입니다.

func Contains(s, substr string) bool: 문자열이 포함되어 있는지 검색
func ContainsAny(s, chars string) bool: 특정 문자가 하나라도 포함되어 있는지 검색
func ContainsRune(s string, r rune) bool: rune 자료형으로 검색
func Count(s, sep string) int: 문자열이 몇 번 나오는지 구함
func HasPrefix(s, prefix string) bool: 문자열이 접두사인지 판단
func HasSuffix(s, suffix string) bool: 문자열이 접미사인지 판단
func Index(s, sep string) int: 특정 문자열의 위치를 구함
func IndexAny(s, chars string) int: 가장 먼저 나오는 문자의 위치를 구함
func IndexByte(s string, c byte) int: byte 자료형으로 위치를 구함
func IndexRune(s string, r rune) int: rune 자료형으로 위치를 구함
func IndexFunc(s string, f func(rune) bool) int: 검색 함수를 정의하여 위치를 구함
func LastIndex(s, sep string) int: 가장 마지막에 나오는 특정 문자열의 위치를 구함
func LastIndexAny(s, chars string) int: 가장 마지막에 나오는 문자의 위치를 구함
func LastIndexFunc(s string, f func(rune) bool) int: 검색 함수를 정의하여 위치를 구함
먼저 Hello, world! 문자열에서 특정 문자열 및 문자를 찾아보겠습니다.

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))  // true
	fmt.Println(strings.Contains("Hello, world!", "w o")) // false
	fmt.Println(strings.Contains("Hello, world!", "ow"))  // false

	fmt.Println(strings.ContainsAny("Hello, world!", "wo"))  // true
	fmt.Println(strings.ContainsAny("Hello, world!", "w o")) // true
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))  // true

	fmt.Println(strings.Count("Hello Helium", "He")) // 2

	var r rune
	r = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r)) // true

	fmt.Println(strings.HasPrefix("Hello, world!", "He"))   // true
	fmt.Println(strings.HasSuffix("Hello, world!", "rld!")) // true
}
```
- strings.Contains 함수는 문자열에서 특정 문자열이 포함되어 있는지 검색합니다. 즉 검색할 단어 전체가 문자열에 모두 포함되어 있어야하며 정확히 일치해야 합니다. 검색에 성공하면 true를 리턴합니다.
- strings.ContainsAny 함수는 검색할 문자열의 문자가 하나라도 포함되어 있는지 검색합니다. 즉 문자 단위로 검색하며 유니코드 문자코드(Code point) 기준입니다. 검색에 성공하면 true를 리턴합니다.
- strings.Count 함수는 문자열에서 특정 문자열이 몇 번 나오는지 구합니다. 검색할 단어가 정확히 일치해야 합니다.
- strings.ContainsRune 함수는 rune 자료형으로 검색합니다. 특히 알파벳이 아닌 한글, 한자 등을 검색할 때 유용합니다. 검색에 성공하면 true를 리턴합니다.
- strings.HasPrefix, strings.HasSuffix 함수는 문자열에서 특정 문자열이 접미사 또는 접두사인지 판단합니다. 접두사 또는 접미사가 맞다면 true를 리턴합니다.
이번에는 문자열에서 특정 문자열 및 문자의 위치를 알아내보겠습니다.

```
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("Hello, world!", "He"))  // 0: He가 맨 처음에 있으므로 0
	fmt.Println(strings.Index("Hello, world!", "wor")) // 7: wor가 8번째에 있으므로 7
	fmt.Println(strings.Index("Hello, world!", "ow"))  // -1: ow는 없으므로 -1

	fmt.Println(strings.IndexAny("Hello, world!", "eo")) // 1: e가 2번째에 있으므로 1
	fmt.Println(strings.IndexAny("Hello, world!", "f"))  // -1: f는 없으므로 -1

	var c byte
	c = 'd'
	fmt.Println(strings.IndexByte("Hello, world!", c)) // 11: d가 12번째에 있으므로 11
	c = 'f'
	fmt.Println(strings.IndexByte("Hello, world!", c)) // -1: f는 없으므로 -1

	var r rune
	r = '언'
	fmt.Println(strings.IndexRune("고 언어", r)) // 4: "언"이 시작되는 인덱스가 4

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	fmt.Println(strings.IndexFunc("Go 언어", f))  // 3: 한글이 4번째부터 시작하므로 3
	fmt.Println(strings.IndexFunc("Go Language", f)) // -1: 한글이 없으므로 -1

	fmt.Println(strings.LastIndex("Hello Hello Hello, world!", "Hello"))
                                                                // 12: 마지막 Hello가 13번째에 있으므로 12

	fmt.Println(strings.LastIndexAny("Hello, world", "ol")) // 10: 마지막 l이 11번째에 있으므로 10

	fmt.Println(strings.LastIndexFunc("Go 언어 안녕", f)) // 13: 마지막 한글인 '녕'이 시작되는 인덱스가 13
}
```
- strings.Index 함수는 문자열에서 특정 문자열의 위치를 구합니다. 검색할 단어가 정확히 일치해야 합니다. 문자열이 없으면 -1을 리턴합니다.
- strings.IndexAny 함수는 검색할 문자열의 문자 중에서 가장 먼저오는 문자의 위치를 구합니다. 유니코드 단위로 검색합니다. 문자가 없으면 -1을 리턴합니다.
- strings.IndexByte 함수는 byte 자료형으로 위치를 알아냅니다. 문자가 없으면 -1을 리턴합니다.
- strings.IndexRune 함수는 rune 자료형으로 위치를 알아냅니다. 특히 알파벳이 아닌 한글, 한자 등을 검색할 때 유용합니다. 문자가 없으면 -1을 리턴합니다.
- strings.IndexFunc 함수는 검색할 함수를 따로 정의하여 위치를 알아냅니다. 예제에서는 unicode.Is 함수를 사용하여 한글 유니코드로 시작하는 부분의 위치를 알아냅니다. 검색에 실패하면 -1을 리턴합니다.
- strings.LastIndex 함수는 문자열에서 가장 마지막에 나오는 특정 문자열의 위치를 구합니다. 따라서 앞 부분에서 문자열이 나오더라도 뒤에 또 나온다면 가장 마지막에 발견되는 인덱스를 리턴합니다. 검색할 단어가 정확히 일치해야 합니다. 그리고 문자열이 없으면 -1을 리턴합니다.
- strings.LastIndexAny 함수는 검색할 문자열의 문자 중 가장 마지막에 나오는 문자의 위치를 구합니다. 유니코드 단위로 검색합니다. 문자가 없으면 -1을 리턴합니다.
- strings.LastIndexFunc 함수는 검색할 함수를 따로 정의하여 가장 마지막에 나오는 문자의 위치를 알아냅니다. 예제에서는 unicode.Is 함수를 사용하여 가장 마지막에 있는 한글 유니코드의 시작 위치를 알아냅니다.
