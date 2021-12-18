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
