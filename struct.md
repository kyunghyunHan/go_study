# 구조체 사용 
```
type Rectangle struct {
	width  int
	height int
}
```
{ } (중괄호) 블록 안에는 구조체의 필드(멤버) 목록을 정의합니다. 각 필드를 나타내는 방법은 변수를 선언하는 방법과 같습니다.

구조체 필드에서 자료형이 같은 필드는 한 줄로 나열해도 됩니다.
```
type Rectangle struct {
	width, height int // 자료형이 같은 필드는 한 줄로 나열
}
```
 
- 구조체_포인터 = new(구조체_타입)

```
var rect1 *Rectangle   // 구조체 포인터 선언
rect1 = new(Rectangle) // 구조체 포인터에 메모리 할당

rect2 := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당
```
new 함수로 메모리 공간을 할당할 때는 구조체의 포인터 변수에 할당합니다. var 키워드를 사용하지 않고 :=로 변수를 선언하면서 메모리를 할당할 수 있습니다.

구조체 인스턴스는 생성할 때 값을 초기화할 수 있습니다.

- 구조체_인스턴스 = 구조체_타입{ }
```
var rect1 Rectangle = Rectangle{10, 20} // 구조체 인스턴스를 생성하면서 값 초기화

rect2 := Rectangle{45, 62} // var 키워드 생략 구조체 인스턴스를 생성하면서 값 초기화

rect3 := Rectangle{width: 30, height: 15} // 구조체 필드를 지정하여 값 초기화
```
중괄호 블록 안에 필드 순서대로 값을 나열하면 됩니다. 여기서 필드 명을 생략했을 때는 필드 개수를 모두 채워주어야 합니다. 중괄호 블록 안에 필드명: “값” 형식으로 필드명을 지정해 줄 수도 있습니다. 이때는 필드 개수를 모두 채워주지 않아도 됩니다.
구조체 인스턴스의 필드에 접근할 때는 . (점)을 사용합니다. 마찬가지로 new 함수로 메모리를 할당한 구조체 인스턴스의 필드에 접근할 때도 점을 사용합니다.
```
var rect1 Rectangle // 구조체 인스턴스 생성
var rect2 *Rectangle = new(Rectangle) // 구조체 포인터 선언 후 메모리 할당

rect1.height = 20 // 구조체 인스턴의 필드에 접근할 때 .을 사용
rect2.height = 62 // 구초체 포인터에 접근할 때도 .을 사용

fmt.Println(rect1) // {0 20}: 구조체 인스턴스의 값
fmt.Println(rect2) // &{0 62}: 구조체 포인터이므로 앞에 &가 붙음
```
```
{0 20}
&{0 62}
```
fmt.Println 함수로 구조체 인스턴스나 포인터를 출력해보면 필드의 내용이 그대로 출력됩니다. 이때 구조체 인스턴스는 { } 안에 값이 출력됩니다. 그리고 구조체 포인터는 { } 앞에 메모리 주소를 뜻하는 &가 붙습니다(레퍼런스).

## 구조체 생성자 패턴
new 함수로 구조체의 메모리를 할당하는 동시에 값을 초기화하는 방법은 없습니다. 하지만 다음과 같은 패턴을 사용하여 다른 언어의 생성자(Constructor)를 흉내낼 수 있습니다.
```
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	rect := NewRectangle(20, 10)

	fmt.Println(rect) // &{20 10}
}
```
```
rect := &Rectangle{20, 10} // 구조체를 초기화한 뒤 메모리 주소를 대입

fmt.Println(rect) // &{20 10}
```
```
&{20 10}
```

## 함수에서 구조체 매개변수 사용

```
func rectangleArea(rect *Rectangle) int { // 매개변수로 구조체 포인터를 받음
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{30, 30}

	area := rectangleArea(&rect) // 구조체의 포인터를 넘김

	fmt.Println(area) // 900
}
```
보통 rectangleArea(rect *Rectangle)과 같이 함수의 매개변수에 구조체 포인터를 받습니다. 그리고 값을 넘겨줄 때도 rect 변수에 &를 붙여 주소를 넘겨줍니다.

함수의 매개변수에 구조체 포인터가 아닌 일반적인 형태(구조체 인스턴스)로 넘겨주면 값이 모두 복사되므로 주의합니다

```
func rectangleScaleA(rect *Rectangle, factor int) { // 매개변수로 구조체 포인터를 받음
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

func rectangleScaleB(rect Rectangle, factor int) { // 매개변수로 구조체 인스턴스를 받음
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func main() {
	rect1 := Rectangle{30, 30}
	rectangleScaleA(&rect1, 10) // 구조체의 포인터를 넘김
	fmt.Println(rect1)          // {300 300}: rect1에 바뀐 값이 들어감

	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10) // 구조체 인스턴스를 그대로 넘김
	fmt.Println(rect2)         // {30 30}: rect2는 값이 바뀌지 않음
}
```
```
{300 300}
{30 30}
```
rectangleScaleA는 구조체 포인터를 매개변수로 받았기 때문에 원래 값이 변경되지만, rectangleScaleB는 구조체를 그대로 받으면서 값이 복사되었으므로 원래의 값에는 영향을 미치지 않습니다.

## 구조체 메서드 연결

Go 언어에는 클래스가 없는 대신 구조체에 메서드를 연결할 수 있습니다.

func (리시버명 *구조체_타입) 함수명() 리턴값_자료형 { }
이번에도 앞에서 만든 사각형(Rectangle) 구조체를 사용하겠습니다.
```
type Rectangle struct {
	width  int
	height int
}

//          ↓ 리시버 변수 정의(연결할 구조체 지정)
func (rect *Rectangle) area() int {
	return rect.width * rect.height
	//     ↑  리시버 변수를 사용하여 현재 인스턴스에 접근할 수 있음
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area()) // 200
}
```
함수를 정의할 때 func 키워드와 함수명 사이에 리시버 정의 부분이 추가되었습니다. 이렇게 하면 메서드 안에서는 리시버 변수를 통해 현재 인스턴스의 값에 접근할 수 있습니다. 그리고 구조체 인스턴스에 . (점)을 사용하여 연결된 메서드를 호출합니다.

리시버로 정의한 변수에는 메서드가 포함된 구조체의 인스턴스 포인터가 들어옵니다. 따라서 리시버 변수를 통해서 현재 인스턴스의 필드 값을 가져오거나 변경할 수 있습니다. (즉, 리시버는 C++의 this 포인터 또는 Java의 this 키워드와 비슷합니다).

함수에 구조체 형태의 매개변수를 넘기는 방식과 마찬가지로 리시버 변수를 받는 방법도 포인터와 일반 구조체 방식이 있습니다.

```
//           ↓ 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

//          ↓ 일반 구조체 방식
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func main() {
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1) // {300 300}: rect1에 바뀐 값이 들어감

	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2) // {30 30}: rect2는 값이 바뀌지 않음
}
```
scaleA 메서드는 리시버 변수로 구조체 포인터를 받았기 때문에 원래 값이 변경되지만, scaleB 메서드는 구조체를 그대로 받으면서 값이 복사되었으므로 원래의 값에는 영향을 미치지 않습니다.

메서드를 작성할 때 구조체 인스턴스의 값을 변경한다면 포인터 형태로 받고, 일반적인 상황에서는 리시버 변수를 값 형태로 받는 것이 좀 더 명확합니다.

리시버 변수를 사용하지 않는다면 _ (밑줄 문자)로 변수를 생략할 수 있습니다.
```
func (_ Rectangle) dummy() { // _로 리시버 변수 생략
	fmt.Println("dummy")
}

func main() {
	var r Rectangle
	r.dummy() // dummy
}
```
## 구조체 임베딩 사용
Go 언어는 클래스를 제공하지 않으므로 상속 또한 제공하지 않습니다. 하지만 구조체에서 임베딩(Embedding)을 사용하면 상속과 같은 효과를 낼 수 있습니다.

간단하게 사람(Person)과 학생(Student) 구조체를 정의한 뒤 사람 구조체에는 인사(greeting) 함수를 작성했습니다.

```
type Person struct { // 사람 구조체 정의
	name string
	age  int
}

func (p *Person) greeting() { // 인사(greeting) 함수 작성
	fmt.Println("Hello~")
}

type Student struct {
	p      Person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계
	school string
	grade  int
}

func main() {
	var s Student

	s.p.greeting() // Hello~
}
```
Student 구조체를 보면 p Person 필드가 들어있습니다. 이렇게 되면 구조체가 해당 타입을 가지고 있는(Has-a) 관계가 됩니다. 즉 “학생이 사람을 가지고 있다”가 됩니다. 따라서 greeting 함수를 호출할 때도 s.p.greeting()처럼 p 필드를 통해서 호출합니다.

이번에는 Student 구조체에 Person 구조체를 임베딩합니다.

```
type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	Person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	school string
	grade  int
}

func main() {
	var s Student

	s.Person.greeting() // Hello~
	s.greeting()        // Hello~
}

```
Student 구조체에서 Person 필드를 정의할 때 필드명은 사용하지 않고, 구조체 타입만 지정했습니다. 이렇게 되면 구조체가 해당 타입을 포함하는(Is-a) 관계가 됩니다. 즉 “학생은 사람이다.”가 됩니다. 따라서 greeting 함수를 호출할 때 s.Person.greeting()처럼 Person 구조체 타입을 통해서 호출하거나 s.greeting()처럼 바로 호출할 수 있습니다.

구조체 임베딩을 사용하면 다른 언어의 상속과 동일한 형태가 됩니다. 물론 구조체 안에 여러 개의 구조체를 임베딩하면 다중 상속도 구현할 수 있습니다. 하지만 Go 언어에서는 복잡한 다중 상속 대신 인터페이스를 주로 활용합니다. 인터페이스에 대해서는 뒤에서 자세히 설명하겠습니다.

## 오버라이딩 상황

다음과 같이 Student 구조체도 Person 구조체와 같은 이름의 greeting 메서드를 가지고 있다면 이때는 Student 구조체의 greeting 함수가 오버라이드됩니다.

```
type Student struct {
	Person
	school string
	grade  int
}

func (p *Student) greeting() { // 이름이 같은 메서드를 정의하면 오버라이딩됨
	fmt.Println("Hello Students~")
}

func main() {
	var s Student

	s.Person.greeting() // Hello~
	s.greeting()        // Hello Students~
}
```
부모 구조체의 메서드 이름과 중복된다면 상속 과정의 맨 아래 메서드가 호출됩니다. 즉, 자식 구조체의 메서드가 부모 구조체의 메서드를 오버라이드(Override)하게 됩니다. 부모 구조체의 메서드를 호출하고 싶으면 s.Person.greeting()과 같은 형태로 구조체 타입을 지정하여 호출합니다.
