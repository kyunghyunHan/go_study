# go_study

1. 변수
변수는 Go 키워드 var 를 사용하여 선언한다. var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다. 예를 들어, 아래는 a 라는 정수(int) 변수를 선언한 것이다.

1
var a int  
변수 선언문에서 변수 초기값을 할당할 수도 있다. 즉, float32 타입의 변수 f 에 11.0 이라는 초기값을 할당하기 위해서는 아래와 같이 쓸 수 있다.

1
var f float32 = 11.
일단 선언된 변수는 그 뒤의 문장에서 해당 타입의 값을 할당할 수 있다.

1
2
a = 10
f = 12.0
만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러를 발생시킨다. 따라서 사용되지 않는 변수는 프로그램에서 삭제한다.

동일한 타입의 변수가 복수 개 있을 경우, 변수들을 나열하고 마지막에 타입을 한번만 지정할 수 있다.

1
2
var i, j, k int
복수 변수들이 선언된 상황에서 초기값을 지정할 수 있다. 초기값은 순서대로 변수에 할당된다. 예를 들어, 아래 코드의 경우 i는 1, j는 2, k는 3 이 할당된다.

1
var i, j, k int = 1, 2, 3
변수를 선언하면서 초기값을 지정하지 않으면, Go는 Zero Value를 기본적으로 할당한다. 즉, 숫자형에는 0, bool 타입에는 false, 그리고 string 형에는 "" (빈문자열)을 할당한다.

Go 에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다. 즉, 아래 코드에서 i는 정수형으로 1이 할당되고, s는 문자열로 Hi가 할당된다.

var i = 1 var s = "Hi"
변수를 선언하는 또 다른 방식으로 Short Assignment Statement ( := ) 를 사용할 수 있다. 즉, var i = 1 을 쓰는 대신 i := 1 이라고 var 를 생략하고 사용할 수 있다. 하지만 이러한 표현은 함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다. Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.

2. 상수
상수는 Go 키워드 const 를 사용하여 선언한다. const 키워드 뒤에 상수명을 적고, 그 뒤에 상수타입, 그리고 상수값을 할당한다.

1
2
const c int = 10
const s string = "Hi"
Go 에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다. 즉, 위의 경우 타입 int, string 을 생략하면 Go에서 자동으로 그 타입을 추론하게 된다.

1
2
const c = 10
const s = "Hi"
여러 개의 상수들 묶어서 지정할 수 있는데, 아래와 같이 괄호 안에 상수들을 나열하여 사용할 수 있다.

1
2
3
4
5
const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
)
한가지 유용한 팁으로 상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다. 이 경우 iota가 지정된 Apple에는 0이 할당되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.

1
2
3
4
5
6
const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
)
3. Go 키워드
Go 프로그래밍 언어는 다음과 같은 25개의 예약 키워드들을 갖는다. 이들 Go 키워드들은 변수명, 상수명, 함수명 등의 Identifier로 사용할 수 없다.

1
2
3
4
5
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
