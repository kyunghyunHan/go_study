# 👩🏻‍🎓리플렉션

- 리플렉션은 실행 시점(Runtime, 런타임)에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능이다. 리플렉션은 Java, C#처럼 가상 머신 위에서 실행되는 언어나 Python, Ruby 등의 스크립트 언어에서 주로 사용하였다. 마찬가지로 Go 언어도 기본 패키지에서 리플렉션을 제공한다

간단하게 변수와 구조체의 타입을 표시해보겠습니다.

```
package main

import (
	"fmt"
	"reflect"
)

type Data struct { // 구조체 정의
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // int: reflect.TypeOf로 자료형 이름 출력

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s)) // string: reflect.TypeOf로 자료형 이름 출력

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f)) // float32: reflect.TypeOf로 자료형 이름 출력

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // main.Data: reflect.TypeOf로 구조체 이름 출력
}
```
```
int
string
float32
main.Data
```
- reflect.TypeOf 함수를 사용하면 일반 자료형이나 구조체의 타입을 알 수 있다. 여기서는 int, string, float32 형 변수의 자료형이 출력된다. 마찬가지로 구조체도 타입을 알아낼 수 있는데 Data 구조체는 main 패키지 안에 속해있기 때문에 main.Data로 나온다.

- 리플렉션으로 변수의 타입뿐만 아니라 값에 대한 상세한 정보도 얻어올 수 있다.
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)  // f의 타입 정보를 t에 저장
	v := reflect.ValueOf(f) // f의 값 정보를 v에 저장

	fmt.Println(t.Name())                    // float64: 자료형 이름 출력
	fmt.Println(t.Size())                    // 8: 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // true: 자료형 종류를 알아내어 
                                                 // reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64)   // false: 자료형 종류를 알아내어 reflect.Int64와 비교

	fmt.Println(v.Type())                    // float64: 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // true: 값이 담긴 변수의 자료형 종류를 
                                                 // 알아내어 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64)   // false: 값이 담긴 변수의 자료형 종류를 
                                                 // 알아내어 reflect.Int64와 비교
	fmt.Println(v.Float())                   // 1.3: 값을 실수형으로 출력
}
```
```
float64
8
true
false
float64
true
false
1.3
```
- reflect.TypeOf 함수로 float64 변수의 타입 정보 reflect.Type를 얻어왔다. 타입 정보로는 타입의 이름, 변수(타입)의 크기 등을 알 수 있으며 Kind 함수를 사용하면 상수 형식으로된 타입 종류도 알 수 있다.

- reflect.ValueOf 함수로 float64 변수의 값 정보 reflect.Value를 얻어오면 타입 정보, 타입 종류, 변수에 저장된 값을 알 수 있다. 여기서 변수가 float64라면 v.Float(), int라면 v.Int(), string이라면 v.String()과 같이 각 타입에 맞는 함수를 사용하면 변수에 저장된 값을 가져올 수 있다.

## 💯구조체 태그 가져오기 
```
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"` // 구조체에 태그 설정
	age  int    `tag1:"나이" tag2:"Age"`  // 구조체에 태그 설정
}

func main() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2")) // true 이름 Name

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2")) // true 나이 Age
}
```
```
true 이름 Name
true 나이 Age
```
- 구조체 필드의 태그는 `태그명:"내용"` 형식으로 지정한다. 태그는 문자열 형태이며 문자열 안에 " " (따옴표)가 포함되므로 ` ` (백쿼트)로 감싸준다. 여러 개를 지정할 때는 공백으로 구분해준다.

- reflect.TypeOf 함수에 구조체 인스턴스를 넣으면 reflect.Type이 리턴된다. 여기서 FieldByName 함수에 구조체의 필드 이름을 지정하면 필드 정보를 얻을 수 있다. 두 번째 리턴값은 해당 이름으로 필드가 존재하는지 여부이다. 필드 정보를 얻은 뒤에는 name.Tag.Get(“tag1”)과 같이 태그를 가져올 수 있다.

## 💯포인터와 인터페이스의 값 가져오기 
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))               // *int
	fmt.Println(reflect.ValueOf(a))              // <*int Value>
	fmt.Println(reflect.ValueOf(a).Int())        // 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())       // <int Value>
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 1

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))         // int
	fmt.Println(reflect.ValueOf(b))        // <int Value>
	fmt.Println(reflect.ValueOf(b).Int())  // 1
	fmt.Println(reflect.ValueOf(b).Elem()) // 런타임 에러
}
```
- 포인터는 일반 변수와는 다르게 값을 가져오려면 reflect.ValueOf 함수로 값 정보 reflect.Value를 얻어온 뒤 다시 Elem 함수로 값 정보를 가져와야 한다. 그리고 변수의 타입에 맞는 Int, Float, String 등의 함수로 값을 가져온다.

- 여기서는 int 포인터 a의 값 정보에서 바로 Int 함수로 값을 가져오려면 런타임 에러가 발생한다. 따라서 Elem 함수로 포인터의 메모리에 저장된 실제 값 정보를 가져와야 한다.

- 빈 인터페이스 b에 1을 대입하면 타입 정보는 int이고 값 정보는 int이다. 따라서 인터페이스의 값을 가져오려면 변수 타입에 맞는 Int, Float, String 등의 함수를 사용하면 된다.

## 💯동적으로 함수 생성

- 리플렉션을 사용하여 동적으로 함수를 만들어내는 방법

- 먼저 reflect.MakeFunc 함수를 사용하는 방법이다.
```
package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value { // 매개변수와 리턴값은 반드시 
                                               // []reflect.Value를 사용
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	var hello func() // 함수를 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem() // hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), h) // h의 함수 정보를 생성

	fn.Set(v) // hello의 값 정보인 fn에 h의 함수 정보 v를 설정하여 함수를 연결

	hello()
}
```
```
Hello, world!
```
- 먼저 매개변수와 리턴값으로 []reflect.Value를 사용하는 함수 h를 정의한다. 이 함수에서 실제 기능을 구현하며 여기서는 Hello, world!를 출력했다. 리플렉션으로 생성하는 함수는 반드시 []reflect.Value를 매개변수로 받고 리턴값으로 사용해야 한다.

- 이제 함수를 담을 수 있는 변수를 hello를 선언한다. 그리고 reflect.ValueOf 함수에 hello의 주소(레퍼런스)를 넘긴 뒤 Elem 함수로 실제 값 정보를 가져온다. 그 다음에는 reflect.MakeFunc(fn.Type(), h)과 같이 함수 h를 넣어서 함수 정보를 생성한다. 마지막으로 fn.Set(v)와 같이 hello의 값 정보인 fn에 함수 정보 v를 설정하여 함수를 연결해준다.
모든 과정을 마친 뒤 hello 함수를 호출하면 Hello, world!가 출력된다.. 즉 hello 함수는 h 함수를 이용하여 동적으로 생성된 함수이다. 하지만 Hello, world!를 출력하기에는 복잡하기만 하다,

- 동적 함수 생성을 좀더 제대로 활용하기 위해 함수 하나로 정수, 실수, 문자열 더하기를 모두 처리하는 함수를 생성

- 먼저 두 값을 더하는 sum 함수를 구현
```
...

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

...
```

- 리플렉션으로 생성될 함수이므로 매개변수와 리턴값의 타입은 []reflect.Value이다. 그러므로 두 값을 더할 때 첫 번째 매개변수는 args[0], 두 번째 매개변수는 args[1]이 된다.여기서는 두 값의 타입 종류를 비교해서 다르면 nil을 리턴을 하도록 구현햇다. 하지만 구현하기에 따라서 float32와 int32를 더하여 float64를 만들 수도 있고, string과 int32를 더하여 string을 만들 수도 있다.

- switch 분기문으로 타입 종류를 구분하고, 같은 종류의 타입은 묶어서 처리한다.. return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}와 같이 타입 종류에 맞게 변수에서 값을 꺼내서 더하고, []reflect.Value 슬라이스로 만들어서 리턴한다.

- 이제 매개변수로 받은 변수에 함수를 연결해주는 함수를 구현
```
...

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

...
```
- 빈 인터페이스로 매개변수를 받은 뒤 reflect.ValueOf 함수와 Elem 함수를 사용하여 실제 값 정보를 가져온다. 이때 매개변수로 들어올 값은 함수를 저장할 수 있는 변수이다. 그 다음에는 reflect.MakeFunc(fn.Type(), sum)과 같이 함수 sum을 넣어서 함수 정보를 생성, 마지막으로 fn.Set(v)와 같이 fptr의 값 정보인 fn에 함수 정보 v를 설정하여 함수를 연결해준다

이렇게 하면 매개변수, 리턴값 자료형의 형태가 다양한 함수를 동적으로 연결시킬 수 있다.

-  main 함수에서 함수를 저장할 수 있는 변수를 선언, 함수 모양을 하고 있지만 실제 함수가 정의되어 있지 않기 때문에 호출은 할 수 없는 상태이다.
```
...

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

...
}

```
- 두 값을 더하므로 매개변수는 두 개이며 각 타입별로 선언한다. 여기서 리턴값은 크기가 가장 큰 타입을 사용한다. 즉 정수는 int64, 실수는 float64를 사용해야 한다.

- 앞에서 선언한 변수의 레퍼런스(포인터)를 makeSum 함수에 넣어주면 완전한 형태의 함수가 생성된다.
```
func main() {
...

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
```
```
3
5.599999904632568
Hello, world!
```
- intSum 함수는 정수를 더하고, floatSum 함수는 실수를 더하고, stringSum 함수는 문자열을 붙인다. 함수는 세 종류지만 실제로는 앞에서 구현한 sum 함수 하나로 처리된다. 이처럼 리플렉션의 동적 함수 생성 기능을 활용하면 타입별로 함수를 여러 벌 구현하지 않아도 된다.

-  전체 소스 코드
```
package main

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
```
