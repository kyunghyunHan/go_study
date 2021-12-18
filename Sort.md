# 정렬

데이터를 처리하다보면 정렬을 자주 하게 됩니다. 이번에는 Go 언어에서는 다양한 데이터를 정렬 할 수 있도록 패키지를 제공해줍니다.

다음은 sort 패키지에서 제공하는 정렬 함수와 타입입니다.

- func Sort(data Interface): 데이터를 오름차순으로 정렬
- func Reverse(data Interface) Interface: 데이터를 내림차순으로 정렬
- type IntSlice []int: int 정렬 인터페이스
- type Float64Slice []float64: float64 정렬 인터페이스
- type StringSlice []string: string 정렬 인터페이스
먼저 정수, 실수, 문자열 슬라이스의 요소를 정렬해보겠습니다.

```
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}

	sort.Sort(sort.IntSlice(a))               // int 슬라이스를 오름차순으로 정렬
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a))) // int 슬라이스를 내림차순으로 정렬
	fmt.Println(a)

	sort.Sort(sort.Float64Slice(b))               // float64 슬라이스를 오름차순으로 정렬
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.Float64Slice(b))) // float64 슬라이스를 내림차순으로 정렬
	fmt.Println(b)

	sort.Sort(sort.StringSlice(c))               // string 슬라이스를 오름차순으로 정렬
	fmt.Println(c)
	sort.Sort(sort.Reverse(sort.StringSlice(c))) // string 슬라이스를 내림차순으로 정렬
	fmt.Println(c)
}
```
```
[3 5 6 7 10]
[10 7 6 5 3]
[1.3 4.2 5.5 7.6 9.9]
[9.9 7.6 5.5 4.2 1.3]
[Andrew John Maria]
[Maria John Andrew]
```
데이터가 담긴 슬라이스를 sort.Sort 함수에 넣으면 정렬이 됩니다. 기본적으로 오름차순 정렬이며 내림차순 정렬을 하려면 sort.Reverse 함수를 사용하면 됩니다.

|자료형	|정렬 인터페이스|
|---|------------|
|int|	sort.IntSlice|
|float64	|sort.Float64Slice|
|string	|sort.StringSlice|


정수, 실수, 문자열과 같은 기본 자료형은 다음과 같이 슬라이스를 바로 넣을 수 있는 함수를 제공합니다. 이 함수들은 기본적으로 오름차순 정렬을 합니다.

```
sort.Ints(a)
fmt.Println(a) // [3 5 6 7 10]

sort.Float64s(b)
fmt.Println(b) // [1.3 4.2 5.5 7.6 9.9]

sort.Strings(c)
fmt.Println(c) // [Andrew John Maria]
```
## 정렬 인터페이스를 이용하여 구조체 정렬하기
다음은 정렬 인터페이스의 정의입니다.
```
type Interface interface {
	Len() int           // 데이터의 길이를 구함
	Less(i, j int) bool // 대소관계를 판단. i가 작으면 true를 리턴하도록 구현
	Swap(i, j int)      // Less 메서드에서 true가 나오면 두 데이터의 위치를 바꿈
}
```

|메서드	|설명|
|---|----|
|Len	|데이터의 개수(길이)를 구합니다.|
|Less	|대소관계를 판단합니다. 두 데이터 i, j를 받은 뒤 i가 작으면 true를 리턴하도록 구현합니다.|
|Swap	|Less 메서드에서 true가 나오면 두 데이터의 위치를 바꿉니다.|

이제 정렬 인터페이스를 구현하여 구조체가 담긴 슬라이스를 정렬해보겠습니다. 다음은 학생을 이름순, 점수순으로 정렬합니다.

```
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type Students []Student

func (s Students) Len() int {
	return len(s) // 데이터의 길이를 구함. 슬라이스이므로 len 함수를 사용
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name // 학생 이름순으로 정렬
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // 두 데이터의 위치를 바꿈
}

type ByScore struct { // 점수순 정렬을 위해 구조체 정의
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score // 학생 점수순으로 정렬
}

func main() {
	s := Students{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	sort.Sort(s) // 이름을 기준으로 오름차순 정렬
	fmt.Println(s)

	sort.Sort(sort.Reverse(ByScore{s})) // 점수를 기존으로 내림차순 정렬
	fmt.Println(s)
}
```
```
[{Andrew 72.6} {John 93.1} {Maria 89.3}]
[{John 93.1} {Maria 89.3} {Andrew 72.6}]
```

학생의 이름과 점수를 저장하는 구조체를 정의하고, 학생을 저장하는 슬라이스도 정의합니다.

```
type Student struct {
	name  string
	score float32
}

type Students []Student
```
이제 Students 슬라이스를 정렬하려면 Students에 sort.Interface를 구현해야 합니다. 따라서 sort.Interface는 다음과 같이 Len, Less, Swap 메서드를 만족해야합니다.

```
func (s Students) Len() int {
	return len(s) // 데이터의 길이를 구함. 슬라이스이므로 len 함수를 사용
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name // 학생 이름순으로 정렬
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // 두 데이터의 위치를 바꿈
}
```
Students 타입의 Less 메서드는 학생 이름을 기준으로 정렬하도록 구현했습니다. 점수를 기준으로 정렬하려면 다음과 같이 구조체와 Less 메서드를 정의합니다.

```
type ByScore struct { // 점수순 정렬을 위해 구조체 정의
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score // 학생 점수순으로 정렬
}
```
ByScore 구조체에는 Students 타입을 포함시켰습니다. 따라서 Students 타입에서 구현했던 Len, Less, Swap 메서드를 그대로 따릅니다. 하지만 점수를 기준으로 정렬해야 하므로 ByScore 구조체에서 Less 메서드를 다시 구현합니다. 즉, Go 언어에서는 이런 방식으로 상속(확장)을 구현합니다.

이제 학생 데이터를 준비한 뒤 이름순, 점수순(내림차순)으로 정렬합니다.

```
s := Students{
	{"Maria", 89.3},
	{"Andrew", 72.6},
	{"John", 93.1},
}

sort.Sort(s) // 이름을 기준으로 오름차순 정렬
fmt.Println(s)

sort.Sort(sort.Reverse(ByScore{s})) // 점수를 기존으로 내림차순 정렬
fmt.Println(s)
```
Students 타입은 sort.Interface를 구현했으므로 sort.Sort 함수에 그대로 넣을 수 있습니다. 그리고 점수순으로 정렬하려면 슬라이스를 ByScore 타입으로 다시 초기화해줍니다. 이렇게 하면 이름순으로 정렬하는 Students 타입의 Less 메서드를 사용하지 않고, 점수순으로 정렬하는 ByScore 타입의 Less 메서드를 사용하게 됩니다. 즉, 데이터를 둘러싸는 타입을 바꿔주면서 기능(메서드)을 바꾸는 방식입니다.

## 정렬 키로 정렬하기
이번에는 인터페이스를 교체하는 방법 대신 정렬 키 함수를 사용하는 방법에 대해 알아보겠습니다. 다음은 정렬 키의 기본 형태입니다.

```
type Data struct { // 정렬할 데이터
}

type By func(s1, s2 *Data) bool // 각 상황별 정렬 함수를 저장할 타입

func (by By) Sort(data []Data) { // By 함수 타입에 메서드를 붙임
	sorter := &dataSorter{   // 데이터와 정렬 함수로 sorter 초기화
		data: data,
		by:   by,
	}
	sort.Sort(sorter) // 정렬
}

type dataSorter struct {
	data []Data                  // 데이터
	by   func(s1, s2 *Data) bool // 각 상황별 정렬 함수
}

func (s *dataSorter) Len() int { // 데이터 길이를 구하는 메서드
}

func (s *studentSorter) Less(i, j int) bool { // 대소관계를 판단하는 메서드
}

func (s *studentSorter) Swap(i, j int) { // 데이터 위치를 바꾸는 메서드
}
```
이제 정렬 키를 사용하여 학생의 이름순, 점수순으로 정렬을 해보겠습니다.
```
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type By func(s1, s2 *Student) bool // 각 상황별 정렬 함수를 저장할 타입

func (by By) Sort(students []Student) { // By 함수 타입에 메서드를 붙임
	sorter := &studentSorter{       // 데이터와 정렬 함수로 sorter 초기화
		students: students,
		by:       by,
	}
	sort.Sort(sorter) // 정렬
}

type studentSorter struct {
	students []Student                  // 데이터
	by       func(s1, s2 *Student) bool // 각 상황별 정렬 함수
}

func (s *studentSorter) Len() int {
	return len(s.students) // 슬라이스의 길이를 구함
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j]) // by 함수로 대소관계 판단
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i] // 데이터 위치를 바꿈
}

func main() {
	s := []Student{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	name := func(p1, p2 *Student) bool { // 이름 오름차순 정렬 함수
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool { // 점수 오름차순 정렬 함수
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool { // 점수 내림차순 정렬 함수
		return !score(p1, p2)
	}

	By(name).Sort(s) // name 함수를 사용하여 이름 오름차순 정렬
	fmt.Println(s)

	By(reverseScore).Sort(s) // reverseScore 함수를 사용하여 점수 내림차순 정렬
	fmt.Println(s)
}
```

```
[{Andrew 72.6} {John 93.1} {Maria 89.3}]
[{John 93.1} {Maria 89.3} {Andrew 72.6}]
```
문법이 다소 생소해보이지만 간단한 내용입니다. 먼저 By라는 함수 타입을 만듭니다. By는 *Student 타입의 매개변수 s1, s2를 받고 bool을 리턴하는 함수 타입일 뿐 실제 함수가 아닙니다. 나중에 이름, 점수 순으로 정렬하는 실제 함수를 By 타입으로 변환하게 됩니다.
```
type By func(s1, s2 *Student) bool // 각 상황별 정렬 함수를 저장할 타입
```
이제 By 함수 타입에 메서드를 붙입니다. 즉, 함수 타입도 구조체, 슬라이스와 마찬가지로 메서드를 붙일 수 있습니다. 그리고 Sorter 구조체도 정의합니다.
```
func (by By) Sort(students []Student) { // By 함수 타입에 메서드를 붙임
	sorter := &studentSorter{       // 데이터와 정렬 함수로 sorter 초기화
		students: students,
		by:       by,
	}
	sort.Sort(sorter) // 정렬
}

type studentSorter struct {
	students []Student                  // 데이터
	by       func(s1, s2 *Student) bool // 각 상황별 정렬 함수
}
```
- By 타입에 Sort 메서드를 구현합니다. 정렬할 슬라이스와 By 타입의 함수(이름, 점수순 정렬 키 함수)를 studentSorter 구조체에 넣은 뒤 sort.Sort 함수로 정렬합니다.
- studentSorter 구조체는 정렬할 데이터인 Student 슬라이스와 정렬 키 함수를 필드로 가지고 있습니다. 이 슬라이스와 키 함수는 Len, Less, Swap 메서드에서 사용됩니다.
실제 정렬을 수행하는 studentSorter 구조체에 sort.Interface를 구현합니다.
```
func (s *studentSorter) Len() int {
	return len(s.students) // 슬라이스의 길이를 구함
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j]) // by 함수로 대소관계 판단
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i] // 데이터 위치를 바꿈
}
```
- Len: 데이터의 개수(길이)를 구합니다. len(s.students)처럼 studentSorter에 저장된 슬라이스의 길이를 구합니다.
- Less: 대소관계를 판단합니다. 두 데이터 i, j를 받은 뒤 i가 작으면 true를 리턴하도록 구현합니다. 여기서는 부등호로 데이터를 비교하지 않고, studentSorter에 저장된 by(정렬 키) 함수를 사용합니다.
- Swap: Less 메서드에서 true가 나오면 두 데이터의 위치를 바꿉니다.
이제 학생 데이터를 준비하고, 정렬 키 함수를 정의한 뒤 이름 오름차순, 점수 내림차순으로 정렬합니다.

```
s := []Student{
	{"Maria", 89.3},
	{"Andrew", 72.6},
	{"John", 93.1},
}

name := func(p1, p2 *Student) bool { // 이름 오름차순 정렬 함수
	return p1.name < p2.name
}
score := func(p1, p2 *Student) bool { // 점수 오름차순 정렬 함수
	return p1.score < p2.score
}
reverseScore := func(p1, p2 *Student) bool { // 점수 내림차순 정렬 함수
	return !score(p1, p2)
}

By(name).Sort(s) // name 함수를 사용하여 이름 오름차순 정렬
fmt.Println(s)

By(reverseScore).Sort(s) // reverseScore 함수를 사용하여 점수 내림차순 정렬
fmt.Println(s)
```
name, score, reverseScore 정렬 키 함수를 정의합니다. name 함수는 이름을 오름차순으로 정렬하고, score 함수는 점수를 오름차순으로 정렬합니다. 이때 reverseScore 함수는 내림차순 정렬을 할 것이므로 score 함수를 호출하여 ! 연산자로 결과를 뒤집어주면 됩니다.

Sort 메서드를 사용할 수 있도록 name, reverseScore 함수를 By 타입으로 변환합니다. 그리고 Sort 함수에 학생 데이터를 넣어서 정렬합니다. 즉 함수의 타입을 변환하여 기능(Sort 메서드)을 사용하고, 내부적으로 함수 자기 자신을 호출하여 정렬을 수행하는 방식입니다.

이처럼 Go 언어는 하나의 데이터 타입을 여러 인터페이스로 바꿔가면서 OOP를 구현합니다.

함수의 타입을 바꾸는 방식을 좀 더 기본적으로 표현하면 다음과 같습니다.
```
var b By = name
b.Sort(s)
fmt.Println(s)
```
By 타입으로 변수를 선언한 뒤 name 함수를 대입해도 똑같이 동작합니다. By 타입은 Sort 메서드만 붙어있을 뿐 name 함수와 매개변수, 리턴값 자료형이 같으므로 당연히 대입할 수 있습니다. 또한, By(name)은 int(10.1), float32(10), []byte(“Hello”)과 같은 문법입니다.
