# interface

- 메서드의 집합체 
```
type hello interface { // 인터페이스 정의
}

func main() {
	var h hello    // 인터페이스 선언
	fmt.Println(h) // <nil>: 빈 인터페이스이므로 nil이 출력됨
}
```
인터페이스를 선언하는 방법은 변수를 선언하는 방법과 같습니다. 즉, var 변수명 인터페이스 형식으로 선언합니다. 그리고 인터페이스는 다른 자료형과 동일하게 함수의 매개변수, 리턴값, 구조체의 필드로 사용할 수 있습니다.

여기서는 빈 인터페이스를 정의하고 선언해보았습니다. 빈 인터페이스기 때문에 할 수 있는 것이 없으며 fmt.Println으로 출력해봐도 nil이 나옵니다.

이제 메서드를 가지는 인터페이스를 정의해보겠습니다.

- type 인터페이스명 interface { 메서드 }

```
type 인터페이스명 interface {
	메서드1() 리턴값_자료형
	메서드2()                  // 리턴값이 없을 때
}
```

{ } (중괄호) 블록안에 메서드 이름, 매개변수 자료형, 리턴값 자료형을 지정하여 한 줄 씩 나열합니다. 여기서는 , (콤마)로 구분하지 않습니다.

다음은 int 자료형에 메서드를 연결하고, 인터페이스로 해당 메서드를 호출합니다.
```
package main

import "fmt"

type MyInt int // int형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스를 통하여 MyInt의 Print 메서드 호출
}
```
먼저 type 새_자료형 자료형 형식으로 기존 자료형을 새 자료형으로 정의할 수 있습니다. 여기서는 기본 자료형인 int에 메서드를 연결하기 위해 MyInt를 새로 정의했습니다. 그 다음 MyInt에 Print 메서드를 연결하고, 현재 변수의 값을 출력하도록 구현하였습니다.
이제 인터페이스를 정의합니다. 보통 인터페이스의 이름은 ~er 형태로 짓습니다(예: Reader, Writer). 우리는 Print 함수를 가지는 Printer 인터페이스를 정의했습니다.

```
type Printer interface {
	Print()
}
```
main 함수 안에서 인터페이스를 선언한 뒤 변수를 대입합니다. 그리고 인터페이스에 . (점)을 사용하여 메서드를 호출합니다. 여기서는 p.Print()처럼 인터페이스의 Print 메서드를 호출했지만 실제로는 MyInt의 Print 메서드가 호출됩니다. 즉 인터페이스에 담긴 실제 타입(자료형, 구조체)의 메서드가 호출됩니다.

이제 인터페이스를 제대로 활용하기 위해 각기 다른 자료형 두 개를 인터페이스 한 개에 담아보겠습니다. 다음은 int 자료형과 사각형 구조체의 내용을 출력하고, int 자료형과 사각형 구조체의 인스턴스를 담을 수 있는 인터페이스를 정의한 예제입니다.
```
package main

import "fmt"

type MyInt int // int 형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(r.width, r.height)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = r     // r을 인터페이스 p에 대입
	p.Print() // 10 20: 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출
}
```
MyInt를 정의하고, 너비와 높이를 가지는 사각형 구조체 Rectangle을 정의했습니다. 그리고 MyInt, Rectangle 모두 자신의 내용을 출력하는 Print 함수를 구현했습니다. 이제 두 타입 모두 똑같은 Print 함수를 가지고 있습니다(여기서 똑같은 함수라는 것은 함수 이름, 매개변수 자료형, 리턴값 자료형이 모두 같은 상태를 뜻합니다).

MyInt 자료형, Ractangle 구조체, Printer 인터페이스를 그림으로 표현하면 다음과 같습니다.

이제 인터페이스를 사용해보겠습니다.

```
func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = r     // r을 인터페이스 p에 대입
	p.Print() // 10 20: 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출
}
```
Printer 인터페이스 p에 MyInt 형 변수 i를 대입했습니다. 마찬가지로 다시 p에 Rectangle 인스턴스를 대입했습니다. 이렇게 전혀 다른 타입을 인터페이스에 대입할 수 있습니다. 즉 인터페이스는 자료형이든 구조체든 타입에 상관없이 메서드 집합만 같으면 동일한 타입으로 봅니다.

인터페이스의 메서드를 호출하면 인터페이스에 담긴 실제 타입의 메서드가 호출되므로 MyInt 형 변수를 대입한 뒤 Print 함수를 호출하면 5가 출력되고, Rectangle 인스턴스를 대입하면 10과 20이 출력됩니다.

인터페이스를 선언하면서 초기화하려면 다음과 같이 :=를 사용하면 됩니다. 인터페이스에는 ( ) (괄호)를 사용하여 변수나 인스턴스를 넣어줍니다.

```
var i MyInt = 5
r := Rectangle{10, 20}

p1 := Printer(i) // 인터페이스를 선언하면서 i로 초기화
p1.Print()       // 5

p2 := Printer(r) // 인터페이스를 선언하면서 r로 초기화
p2.Print()       // 10 20
```
다음과 같이 배열(슬라이스) 형태로도 인터페이스를 초기화 할 수 있습니다.
```
var i MyInt = 5
r := Rectangle{10, 20}

pArr := []Printer{i, r} // 슬라이스 형태로 인터페이스 초기화
for index, _ := range pArr {
	pArr[index].Print() // 슬라이스를 순회하면서 Print 메서드 호출
}

for _, value := range pArr {
	value.Print() // 슬라이스를 순회하면서 Print 메서드 호출
}
```
## 덕타이핑
이렇게 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판단하는 방식을 덕 타이핑(Duck typing)이라 합니다. 이 용어는 다음과 같은 덕 테스트(오리 테스트)에서 유래되었습니다.

“만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라 부르겠다.”

Go 언어로는 다음과 같이 덕 타이핑을 구현할 수 있습니다

```
package main

import "fmt"

type Duck struct { // 오리(Duck) 구조체 정의
}

func (d Duck) quack() {     // 오리의 quack 메서드 정의
	fmt.Println("꽥~!") // 오리 울음 소리
}

func (d Duck) feathers() { // 오리의 feathers 메서드 정의
	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

type Person struct { // 사람(Person) 구조체 정의
}

func (p Person) quack() {                           // 사람의 quack 메서드 정의
	fmt.Println("사람은 오리를 흉내냅니다. 꽥~!") // 사람이 오리 소리를 흉내냄
}

func (p Person) feathers() { // 사람의 feathers 메서드 정의
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface { // quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feather 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feather 메서드 호출
}
```
Quacker 인터페이스는 quack, feathers 함수를 정의하고 있습니다. 그리고 오리(Duck)와 사람(Person) 모두 quack, feathers 함수를 구현했습니다.

실제로 사용할 때는 다음과 같이 inTheForest 함수에서 Quacker 인터페이스를 매개변수로 받습니다. 여기에 오리든 사람이든 inTheForest 함수에 넣을 수 있으며 quack, feathers 메서드를 호출합니다. 오리는 진짜 “꽥~”하고 울것이고, 사람이라면 오리를 흉내내어 소리를 낼 것입니다.

```
func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feather 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feather 메서드 호출
}
```
```
꽥~!
오리는 흰색과 회색 털을 가지고 있습니다.
사람은 오리를 흉내냅니다. 꽥~!
사람은 땅에서 깃털을 주워서 보여줍니다.
```
오리든 사람이든 상관없이 꽥(quack)과 깃털(feathers) 메서드만 가졌다면 모두 같은 인터페이스로 처리할 수 있습니다.

타입이 특정 인터페이스를 구현하는지 검사하려면 다음과 같이 사용합니다.

- interface{}(인스턴스).(인터페이스)

```
var donald Duck

if v, ok := interface{}(donald).(Quacker); ok {
	fmt.Println(v, ok)
}
```
```
{} true
```
Duck 타입의 인스턴스 donald를 빈 인터페이스에 넣은 뒤 Quacker 인터페이스와 같은지 확인합니다. 첫 번째 리턴값은 검사했던 인스턴스이며 두 번째 리턴값은 인스턴스가 해당 인터페이스를 구현하고 있는지 여부입니다. 구현하고 있다면 true 그렇지 않으면 false입니다.
## 빈 인터페이스 사용

인터페이스에 아무 메서드도 정의되어 있지 않으면 모든 타입을 저장할 수 있습니다.
```
func f1(arg interface{}) { // 모든 타입을 저장할 수 있음
}
```
```
type Any interface{} // 인터페이스에 메서드를 지정하지 않음

func f2(arg Any) {   // 모든 타입을 저장할 수 있음
}
```
빈 인터페이스 타입은 함수의 매개변수, 리턴값, 구조체의 필드로 사용할 수 있습니다.

이제 모든 타입을 받아서 내용을 출력하는 함수를 만들어보겠습니다.

```
v
package main

import (
	"fmt"
	"strconv"
)

//                      ↓ 빈 인터페이스를 사용하여 모든 타입을 받음
func formatString(arg interface{}) string {
	//       ↓ 인터페이스에 저장된 타입에 따라 case 실행
	switch arg.(type) {
	case int:                      // arg가 int형이라면
		i := arg.(int)         // int형으로 값을 가져옴
		return strconv.Itoa(i) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32:                  // arg가 float32형이라면
		f := arg.(float32)     // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
                                       // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64:                  // arg가 float64형이라면
		f := arg.(float64)     // float64형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64)
                                       // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string:                   // arg가 string이라면
		s := arg.(string)      // string으로 값을 가져옴
		return s               // string이므로 그대로 리턴
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))
}
```
```
1
2.5
Hello, world!
```
formatString 함수의 매개변수를 보면 타입을 interface{}로 지정하였습니다. 이렇게 하면 모든 타입을 처리할 수 있습니다. * 인터페이스에 저장된 타입을 알아내려면 switch 분기문 안에서 arg.(type)처럼 인터페이스 변수에 .(type)을 사용합니다. 단 이 방법은 switch 분기문 안에서만 사용할 수 있고, 일반적인 방법으로는 사용할 수 없습니다. * 타입에 따라 case로 나눕니다. * 빈 인터페이스 변수는 그대로 사용할 수 없으므로 arg.(int), arg.(float32), arg.(float64), arg.(string)처럼 타입을 지정하여 값을 가져옵니다. 이렇게 타입을 원하는 형태로 바꾸는 작업을 Type assertion이라고 합니다. * 각 타입에 맞게 strconv.Itoa, strconv.FormatFloat 함수를 사용하여 값을 문자열로 만듭니다. string의 값은 문자열이므로 변환하지 않고 그대로 리턴합니다.

일반 자료형뿐만 아니라 구조체 인스턴스 및 포인터도 빈 인터페이스로 넘길 수 있습니다.
```
package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Person 구조체 정의
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:               // arg의 타입이 Person이라면
		p := arg.(Person)  // Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	case *Person:              // arg의 타입이 Person 포인터라면
		p := arg.(*Person) // *Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(&Person{"John", 12}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString(andrew))
}
```
```
Maria 20
John 12
Andrew 35
```
switch arg.(type) { }으로 인터페이스에 저장된 타입을 알아낸 뒤 각 구조체 타입별로 case를 만들어 처리합니다. 여기서 구조체를 그대로 넘겨줬다면 case Person:으로 처리하고, 구조체의 포인터를 넘겨줬다면 case *Person:으로 처리합니다. 마찬가지로 구조체일 때는 arg.(Person), 포인터일 때는 arg.(*Person)으로 값을 가져옵니다.

인터페이스에 저장된 타입이 특정 타입인지 검사하려면 다음과 같이 사용합니다.
```
var t interface{}
t = Person{"Maria", 20}

if v, ok := t.(Person); ok {
	fmt.Println(v, ok)
}
```
```
{Maria 20} true
```
인터페이스.(타입) 형식입니다. 첫 번째 리턴값은 해당 타입으로 된 값이며 두 번째 리턴값은 타입이 맞는지 여부입니다. 타입이 일치하면 true 그렇지 않으면 false입니다.
