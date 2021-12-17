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


## 문자열 조작하기
다음은 strings 패키지에서 제공하는 문자열 조작 함수입니다.

- func Join(a []string, sep string) string: 문자열 슬라이스에 저장된 문자열을 모두 연결
- func Split(s, sep string) []string: 지정된 문자열을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장
- func Fields(s string) []string: 공백을 기존으로 문자열을 쪼개어 문자열 슬라이스로 저장
- func FieldsFunc(s string, f func(rune) bool) []string: 유니코드 검사 함수를 정의한 뒤 특정 유니코드 값을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장
- func Repeat(s string, count int) string: 문자열을 특정 횟수만큼 반복
- func Replace(s, old, new string, n int) string: 문자열에서 특정 문자열을 바꿈
- func Trim(s string, cutset string) string: 문자열 앞 뒤에 있는 특정 문자열 제거
- func TrimLeft(s string, cutset string) string: 문자열 앞에 오는 특정 문자열 제거
- func TrimRight(s string, cutset string) string: 문자열 뒤에 오는 특정 문자열 제거
- func TrimFunc(s string, f func(rune) bool) string: 문자열 정리 함수를 정의하여 문자열 제거
- func TrimLeftFunc(s string, f func(rune) bool) string: 문자열 정리 함수를 정의하여 문자열 앞에 오는 특정 문자열 제거
- func TrimRightFunc(s string, f func(rune) bool) string: 문자열 정리 함수를 정의하여 문자열 뒤에오는 특정 문자열 제거
- func TrimSpace(s string) string: 문자열 앞뒤에오는 공백 문자 제거
- func TrimSuffix(s, suffix string) string: 접미사 제거
- func NewReplacer(oldnew ...string) *Replacer: 문자열 교체 인스턴스를 생성
- func (r *Replacer) Replace(s string) string: 문자열 교체 인스턴스를 사용하여 문자열을 교체
이번에는 문자열을 연결하고, 쪼개고, 반복하고, 바꾸는 방법을 알아보겠습니다.
```
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello,", "world!"}
	fmt.Println(strings.Join(s1, " ")) // Hello, world!

	s2 := strings.Split("Hello, world!", " ")
	fmt.Println(s2[1]) // world!

	s3 := strings.Fields("Hello, world!")
	fmt.Println(s3[1]) // world!

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4) // [Hello Hello]

	fmt.Println(strings.Repeat("Hello", 2)) // HelloHello

	fmt.Println(strings.Replace("Hello, world!", "world", "Go", 1)) // Hello, Go!
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))     // HeGo HeGo
}
```
- strings.Join 함수는 문자열 슬라이스에 저장된 문자열을 모두 연결합니다. 여기서 두 번째 매개변수는 문자열을 연결할 때 사이에 끼워넣을 문자열입니다.
- strings.Split 함수는 지정된 문자열을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장합니다. 예제에서는 “ “ (공백)을 기준으로 문자열을 쪼갭니다.
- strings.Fields 함수는 공백을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장합니다.
- strings.FieldsFunc 함수는 유니코드 검사 함수를 정의한 뒤 특정 유니코드 값을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장합니다. 예제에서는 unicode.Is 함수를 사용하여 한글 유니코드를 기준으로 문자열을 쪼갭니다. 따라서 Hello안녕Hello는 한글 유니코드인 안녕 부분을 기준으로 쪼개져서 [Hello Hello]가 됩니다.
- strings.Repeat 함수는 문자열을 특정 횟수만큼 반복합니다.
- strings.Replace 함수는 문자열에서 특정 문자열을 바꿉니다. 두 번째 매개변수로 바꿀 횟수를 설정할 수 있습니다.
이번에는 문자열을 다듬는(trim) 방법입니다.
```
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Trim("Hello,~~ world!~~", "~~")) // Hello,~~ world!
	fmt.Println(strings.Trim("  Hello, world!  ", " "))  // Hello, world!

	fmt.Println(strings.TrimLeft("~~Hello,~~ world!~~", "~~"))  // Hello,~~ world!~~
	fmt.Println(strings.TrimRight("~~Hello,~~ world!~~", "~~")) // ~~Hello,~~ world!

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	var s = "안녕 Hello 고 Go 언어"
	fmt.Println(strings.TrimFunc(s, f))      //  Hello 고 Go
	fmt.Println(strings.TrimLeftFunc(s, f))  //  Hello 고 Go 언어
	fmt.Println(strings.TrimRightFunc(s, f)) // 안녕 Hello 고 Go

	fmt.Println(strings.TrimSpace("  Hello, world!		")) // Hello, world!

	fmt.Println(strings.TrimSuffix("Hello, world!", "orld!")) // Hello, w
	fmt.Println(strings.TrimSuffix("Hello, world!", "wor"))   // Hello, world!
}
```
- strings.Trim 함수는 문자열의 앞뒤에 있는 특정 문자열을 제거합니다. 단, 문자열 중간에 있는 문자열은 제거하지 않습니다.
- strings.TrimLeft, strings.TrimRight 함수는 문자열 앞이나 뒤에 오는 특정 문자열을 제거합니다. 마찬가지로 문자열 중간에 있는 문자열은 제거하지 않습니다.
- strings.TrimFunc 함수는 문자열 정리 함수를 정의하여 문자열을 정리합니다. 예제에서는 문자열 앞뒤에 오는 한글을 제거합니다.
- strings.TrimLeftFunc, strings.TrimRightFunc 함수는 문자열 정리 함수를 정의하여 문자열 앞이나 뒤에 오는 문자열을 제거합니다.
- strings.TrimSpace 함수는 문자열 앞뒤에 오는 공백 문자를 제거합니다. 공백 문자(스페이스)와 탭 문자를 동시에 제거할 수 있습니다. 따라서 공백 문자와 탭 문자를 제거하기 위해 strings.Trim 함수를 두 번 사용하지 않아도 됩니다. 마찬가지로 문자열 중간에 있는 공백, 탭 문자는 제거하지 않습니다.
- strings.TrimSuffix 함수는 접미사를 제거합니다. 마지막부터 문자열이 일치해야하며 중간만 일치하면 제거하지 않습니다.
다음은 여러 개의 문자열을 동시에 바꾸는 방법입니다.

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")              // 바꿀 문자열을 지정
	fmt.Println(r.Replace("<div><span>Hello, world!</span></div>")) // HTML에서 < >를 &lt; &gt;로 바꿈
}
```
```
&lt;div&gt;&lt;span&gt;Hello, world!&lt;/span&gt;&lt;/div&gt;
```
strings.NewReplacer로 바꿀 문자열을 지정합니다. “기존 문자열”, “새 문자열” 순서로 바꿀 문자열을 여러 개 설정할 수 있습니다. 예제에서는 HTML 문자열에서 태그를 HTML 엔티티 코드(entity code)로 바꿉니다.

## 문자열 변환 함수

다음은 strconv 패키지에서 제공하는 문자열 변환 함수입니다.

- func Atoi(s string) (i int, err error): 숫자로 이루어진 문자열을 숫자로 변환
- func Itoa(i int) string: 숫자를 문자열로 변환
- func FormatBool(b bool) string: 불 값을 문자열로 변환
- func FormatFloat(f float64, fmt byte, prec, bitSize int) string: 실수를 문자열로 변환
- func FormatInt(i int64, base int) string: 부호 있는 정수를 문자열로 변환
- func FormatUint(i uint64, base int) string: 부호 없는 정수를 문자열로 변환
- func AppendBool(dst []byte, b bool) []byte: 불 값을 문자열로 변환하여 슬라이스 뒤에 추가
- func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte: 실수를 문자열로 변환하여 슬라이스 뒤에 추가
- func AppendInt(dst []byte, i int64, base int) []byte: 부호 있는 정수를 문자열로 변환하여 슬라이스 뒤에 추가
- func AppendUint(dst []byte, i uint64, base int) []byte: 부호 없는 정수를 문자열로 변환하여 슬라이스 뒤에 추가
- func ParseBool(str string) (value bool, err error): 불 형태의 문자열을 불로 변환
- func ParseFloat(s string, bitSize int) (f float64, err error): 실수 형태의 문자열을 실수로 변환
- func ParseInt(s string, base int, bitSize int) (i int64, err error): 부호 있는 정수 형태의 문자열을 부호 있는 정수로 변환
- func ParseUint(s string, base int, bitSize int) (n uint64, err error): 부호 없는 정수 형태의 문자열을 부호 없는 정수로 변환
먼저 숫자(정수)를 문자로 문자를 숫자로 변환해보겠습니다.

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var num1 int
	num1, err = strconv.Atoi("100") // 문자열 "100"을 숫자 100으로 변환
	fmt.Println(num1, err)          // 100 <nil>

	var num2 int
	num2, err = strconv.Atoi("10t") // 10t는 숫자가 아니므로 에러가 발생함
	fmt.Println(num2, err)          // 0 strconv.ParseInt: parsing "10t": invalid syntax

	var s string
	s = strconv.Itoa(50) // 숫자 50을 문자열 "50"으로 변환
	fmt.Println(s)       // 50
}
```
```
100 <nil>
0 strconv.ParseInt: parsing "10t": invalid syntax
50
```
- strconv.Atoi 함수는 숫자로 이루어진 문자열을 숫자로 변환합니다. 즉, string에서 int로 변환합니다. 문자열에는 정상적인 숫자가 들어있어야 하며 10t처럼 숫자에 영문자가 섞이거나 아예 숫자가 아니라면 에러가 발생합니다.
- strconv.Itoa 함수는 숫자를 문자열로 변환합니다. 즉, int에서 string으로 변환합니다.
이번에는 실수, 불, 정수를 문자열로 변환해보겠습니다.
```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	s1 = strconv.FormatBool(true)
	fmt.Println(s1) //true: true를 문자열로 변환하여 "true"

	var s2 string
	s2 = strconv.FormatFloat(1.3, 'f', -1, 32) // f, fmt, prec, bitSize
	fmt.Println(s2)                            // 1.3: 1.3을 문자열로 변환하여 "1.3"

	var s3 string
	s3 = strconv.FormatInt(-10, 10)
	fmt.Println(s3) // -10: -10을 10진수로된 문자열로 변환하여 "-10"

	var s4 string
	s4 = strconv.FormatUint(32, 16)
	fmt.Println(s4) // 20: 32를 16진수로된 문자열로 변환하여 "20"
}
```
- strconv.FormatBool 함수는 불 값(true, false)을 문자열로 변환합니다.
- strconv.FormatFloat 함수는 실수를 문자열로 변환합니다. 다음은 함수를 사용할 때 필요한 매개변수입니다.
- - f: 변환할 실수 값입니다.
- - fmt: 실수 형식입니다.
- - - ‘b’: -ddddp±ddd 예) -4503599627370496p-52
- - - ‘e’: -d.dddde±dd 예) 1.23457e+08
- - - ‘E’: -d.ddddE±dd 예) 1.23457E+08
- - - ‘f’: -ddd.dddd 예) 1.352
- - - ‘g’: ‘e’와 같으며 지수 값이 클 때 사용
- - - ‘G’: ‘E’와 같으며 지수 값이 클 때 사용
- - prec: 실수의 정밀도입니다(지수를 제외한 숫자의 자리수). -1을 지정하면 적절한 정밀도로 변환해줍니다.
- - bitSize: 부동소수점 비트 수입니다. 32 또는 64를 지정합니다.
- - strconv.FormatInt, strconv.FormatUint 함수는 부호 있는 정수와 부호 없는 정수를 문자열로 변환합니다. 두 번째 매개변수에 진법을 설정하여 10진수 또는 16진수로된 문자열을 만들 수 있습니다.
다음은 실수, 불, 정수를 문자열로 변환하여 바이트 슬라이스에 추가합니다.

```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s []byte = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(string(s)) // true: true를 문자열로 변환하여 "true"

	s = strconv.AppendFloat(s, 1.3, 'f', -1, 32) // dst, f, fmt, prec, bitSize
	fmt.Println(string(s)) // true1.3: 1.3을 문자열로 변환하여 "1.3", 슬라이스 뒤에 붙여서 true1.3

	s = strconv.AppendInt(s, -10, 10)
	fmt.Println(string(s)) // true1.3-10: -10을 10진수로된 문자열로 변환하여 "-10",
                               // 슬라이스 뒤에 붙여서 true1.3-10

	s = strconv.AppendUint(s, 32, 16)
	fmt.Println(string(s)) // true1.3-1020: 32를 16진수로된 문자열로 변환하여 "20",
                               // 슬라이스 뒤에 붙여서 true1.3-1020
}
```
- strconv.AppendBool 함수는 불 값(true, false)을 문자열로 변환한 뒤 슬라이스 뒤에 추가합니다.
- strconv.FormatFloat 함수는 실수를 문자열로 변환한 뒤 슬라이스 뒤에 추가합니다. 다음은 함수를 사용할 때 필요한 매개변수입니다.
- dst : 변환한 문자열을 추가할 슬라이스입니다.
- - f : 변환할 실수 값입니다.
- - fmt : 실수 형식입니다.
- - - ‘b’ : -ddddp±ddd 예) -4503599627370496p-52
- - - ‘e’ : -d.dddde±dd 예) 1.23457e+08
- - - ‘E’ : -d.ddddE±dd 예) 1.23457E+08
- - - ‘f’ : -ddd.dddd 예) 1.352
- - - ‘g’ : ‘e’와 같으며 지수 값이 클 때 사용
- - - ‘G’ : ‘E’와 같으며 지수 값이 클 때 사용
- - prec : 실수의 정밀도입니다(지수를 제외한 숫자의 자리수). -1을 지정하면 적절한 정밀도로 변환해줍니다.
- - bitSize : 부동소수점 비트 수입니다. 32 또는 64를 지정합니다.
- strconv.FormatInt, strconv.FormatUint 함수는 부호 있는 정수와 부호 없는 정수를 문자열로 변환한 뒤 슬라이스 뒤에 추가합니다. 세 번째 매개변수에 진법을 설정하여 10진수 또는 16진수로된 문자열을 만들 수 있습니다.
반대로 문자열을 실수, 불, 정수로 변환할 수도 있습니다.
```
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var b1 bool
	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err) // true <nil>: 문자열로 "true"를 불로 변환하여 true

	var num1 float64
	num1, err = strconv.ParseFloat("1.3", 64)
	fmt.Println(num1, err) // 1.3 <nil>: 문자열 "1.3"을 실수로 변환하여 1.3

	var num2 int64
	num2, err = strconv.ParseInt("-10", 10, 32)
	fmt.Println(num2, err) // -10 <nil>: 10진수로된 문자열 "-10"을 정수로 변환하여 -10

	var num3 uint64
	num3, err = strconv.ParseUint("20", 16, 32)
	fmt.Println(num3, err) // 32: 16진수로된 문자열 "20"을 정수로 변환하여 32
}
```
- strconv.ParseBool 함수는 불 형태의 문자열을 불로 변환합니다. 즉, string에서 bool로 변환합니다. true, false, TRUE, FALSE, 0, 1, T, F, True, False를 변환할 수 있으며 이외의 문자열은 에러가 발생합니다.
- strconv.ParseFloat 함수는 실수 형태의 문자열을 실수로 변환합니다. 즉, string에서 float64로 변환합니다. 두 번째 매개 변수는 부동소수점 비트 수이며 32 또는 64를 지정할 수 있습니다. 문자열에는 정상적인 형태의 실수가 들어있어야 하며 그렇지 않으면 에러가 발생합니다.
- strconv.ParseInt, strconv.ParseUint 함수는 부호 있는 정수 형태의 문자열을 부호 있는 정수로, 부호 없는 정수 형태의 문자열을 부호 없는 정수로 변환합니다. 즉 string에서 int64로, string에서 uint64로 변환합니다. 문자열에는 정상적인 정수가 들어있어야 하며 영문자가 섞여있거나 아예 숫자가 아니라면 에러가 발생합니다.

