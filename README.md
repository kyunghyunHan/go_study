 ![다운로드](https://user-images.githubusercontent.com/88940298/146376240-aef62ec1-d091-4a99-9bc1-90852b2db0f6.png)

# Go  -공부 및 정리-

<img width="200" alt="스크린샷 2021-11-24 오후 12 28 54" src="https://user-images.githubusercontent.com/88940298/143169831-99d59153-1735-49a6-8d7c-d3768dfc9400.png">


## 가비지 컬렉션
- Go 언어는 가비지 컬렉션(Garbage Collection, GC)을 제공합니다.
- C, C++는 메모리를 할당하면 반드시 해제를 해주어야 합니다.  로직 작성보다 메모리 관리에 더 많은 노력과 시간을 소모하고 있어서 생산성이 많이 떨어집니다. 이후 메모리를 알아서 관리해주는 가비지 컬렉션 기술이 나왔고, 이 기술을 사용한 Java와 C#이 등장했습니다. 마찬가지로 Python, Ruby, JavaScript 등의 스크립트 언어들도 각 언어의 가상 머신에서 가비지 컬렉션 기술을 사용하고 있습니다.

- Go 언어는 메모리를 관리해주는 가비지 컬렉터(Garbage Collector)가 실행 파일안에 내장되어 있습니다. 가상 머신 위에서 실행되는 언어들처럼 가상 머신이 메모리 관리를 해주는 것과 차이가 있습니다. 
- Go 언어는 C, C++ 실행 파일 방식의 간결함과 가상 머신의 가비지 컬렉션 기능을 함께 가지고 있습니다.

## 병행성 
Go 언어는 go 키워드를 통해 함수 여러 개를 동시에 실행할 수 있습니다. 이렇게 실행된 함수를 고루틴(Goroutine)이라고 하는데 스레드와는 차이점이 있습니다. 스레드는 운영체제의 커널에서 제공하는 리소스이기 때문에 많이 생성할수록 부담이 커집니다. 그림 1-5와 같이 Go 언어는 적정량의 스레드를 생성해서 고루틴을 처리합니다. 또한, 최대 프로세서(코어) 개수 설정에 따라 멀티코어도 지원됩니다(현재 프로세서의 개수를 구할 수 있으며, 일부 프로세서만 사용할 수도 있습니다)
<img width="400" alt="스크린샷 2021-10-26 오전 12 39 35" src="https://user-images.githubusercontent.com/88940298/138727383-95f932d6-ffde-467d-bd05-5e3806c4906f.png">


go 에서는 ;(세미클론이 필요없다) 선택적으로 사용할 수는 있음  
go에서는 패키지 개념 패키지 내부에서는 변수,상수,함수,구조체,메서드가 선언되어 있다.  
모든 go파일은 패키지 내부에속해야한다 main패키지가 아닌 다른코드를 실행시킬수 없음  

import"fmt" 다른 패키지를 포함하기 위해 사용 go 의패키지는 src에 모여있으며 pkg폴더에는 그러한 패키지가 컴파일된 소스코드가 모여있다  
변수 선언 var 사용 함수안에서는 단축변수선언 사용가능 :=  

변수를 여러개를 선언하고 초기화할떄에는 반드시 선언하는 변수와 초기화하는 값의 개수가 같아야함 타입은 달라도 됨  

go에서는 패키지,타입과 똑같은 이름의 변수를 사용하여도 에러가 나지 않는다  



## 변수 

var 변수1, 변수2 자료형 = 초깃값1, 초깃값2 : 여러개의 변수를 

```
var x, y int = 30, 50       // x = 30, y = 50
var age, name = 10, "Maria" // age = 10, name = "Maria"
```
```
a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false
```
## 상수 const

go에서는 타입을 정하지 않으면 go에서 알아서 타입을 지정해준다  
 
go가 다른언어랑 다른 점은 클래스가 없다는 것이다.메서드라고 표현하는 것도 표현 방식이 다른 뿐 일반 함수와 유사하게 표시된다  
go는 리턴 값을 여러개의 값으로 돌릴 수 있다.  
defer키워드는 go에서 지원하는 특이한 키워드중 하나 이다. defer 을 통해  먼저 예약이 가능 하며 , defer은 나중에 설정된 defer부터 실행된다.  

go&go-routine go 키워드는 멀티스레드를 사용할때 유용하며 동시성을 가진다.  
 
go get go get을 통하여 다른 패키지를 다운받을 수 있다.얻어온 패키지는 GOPATH로 다운로드 된다.GOPATH는 작업공간이라고도 하며 사용자마다 별도로 관리된다  

go mod 사용자 정의 모듈을 만들떄 사용한다. example.com폴더로 진입하여 go mod init example.com명령어를 통해 모듈을 정의한다  


<img width="355" alt="스크린샷 2021-10-27 오후 2 11 44" src="https://user-images.githubusercontent.com/88940298/139003610-95bf570c-fa06-4959-815c-5e2bea14f2c7.png">

## 바이트
byte에는 보통 16진수, 문자 값으로 저장합니다. 실무에서는 바이너리 파일에서 데이터를 읽고 쓰거나, 데이터를 암호화, 복호화할 때 주로 사용합니다.
```
var num1 byte = 10   // 10진수 저장
var num2 byte = 0x32 // 16진수 저장
var num3 byte = 'a'  // 문자 저장
```


## 룬 
rune은 유니코드(UTF-8) 문자 코드를 저장할 때 사용합니다. ' ' (작은 따옴표)로 묶어주어야 하며 문자 그대로 저장하거나 \u 또는 \U를 사용하여 유니코드 문자 코드로 저장할 수도 있습니다. 여기서 \U는 값을 16진수 8자리로 맞춰주어야 합니다
```
var r1 rune = '한'
var r2 rune = '\ud55c'     // 한
var r3 rune = '\U0000d55c' // 한
```


## 숫자 연산. 
숫자 연산에는 덧셈(+), 뺄셈(-), 곱셈(*), 나눗셈(/), 나머지(%), 시프트(<<, >>), 비트 반전(^) 연산자를 사용할 수 있습니다.  
```
var num1 uint8 = 3
var num2 uint8 = 2

fmt.Println(num1 + num2)  // 5
fmt.Println(num1 - num2)  // 1
fmt.Println(num1 * num2)  // 6
fmt.Println(num1 / num2)  // 1
fmt.Println(num1 % num2)  // 1
fmt.Println(num1 << num2) // 12
fmt.Println(num1 >> num2) // 0
fmt.Println(^num1)        // 252: 비트 반전 연산자
```

## 오버플로우와 언더플로우. 
각 자료형에서 저장할 수 있는 최대 크기를 넘어서면 오버플로우(Overflow), 최소 크기보다 작아지면 언더플로우(Underflow)  
```
package main

import "fmt"
import "math"

func main() {
	var num1 uint8 = math.MaxUint8 + 1   // 오버플로우 컴파일 에러 발생
	var num2 uint16 = math.MaxUint16 + 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = math.MaxUint32 + 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = math.MaxUint64 + 1 // 오버플로우 컴파일 에러 발생
}
```
## 문자열 길이. 
```
var s1 string = "한글"
var s2 string = "Hello"

fmt.Println(len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
fmt.Println(len(s2)) // 5: 알파벳 5글자이므로 5
```

## 열거형. 

```
const (
	Sunday       = 0
	Monday       = 1
	Tuesday      = 2
	Wednesday    = 3
	Thursday     = 4
	Friday       = 5
	Saturday     = 6
	numberOfDays = 7
)
```

## [패키지](https://github.com/kyunghyunHan/go_study/blob/0c7e77e48504492e73e7c61e257d35ef441d6890/package.md)


## [fmt](https://github.com/kyunghyunHan/go_study/blob/b51d98381a325e779dbc4ed371d2f258cda180b6/fmt.md)
출력함수. 

## [if문](https://github.com/kyunghyunHan/go_study/blob/969ebb1c93c8cb200082b8e62919c6ebc8a7e02d/if.md)
조건문을 사용할떄 변수를 함께 선언할 수있다.  

swtich 조건문과 마찬가지로 사용법은 switch,case와 비슷하지만 break문이 필요없다 .새로운 키워드인 fallthrough를 선택했다.fallthrough를사용해야 다음으로 넘어 갈 수있다.또한 다른 언어와 다르게 조건을 생략하고 case를 사용할 수 잇다.  

## [for문](https://github.com/kyunghyunHan/go_study/blob/667aa33064651847371e1e867db3ee4581c2c736/for.md) 

반복문을 사용할 떄는 for키워드를 이용한다. go 언어에는 while문이 없기 떄문에 for문에서 모든 반복을 해결한다.  
무한루프의 경우 아무것도 안주면 되어 타언어에 비해 심플하다 배열이나 슬라이드,문을 순회할떄 range키워드를 사용하여 순회 할 수 있다. 키 인덱스가 들어가며 두번쨰에는 값이 들어간다.  
반복문을 제어하는 방법으로는 break,continue가 존재한다.break를 사용하면 반복문에서 즉시 탈출하며,continue를사용하면 그 이후 표현식은 무시하고 다시 조건문으로 돌아간다  

## [goto](https://github.com/kyunghyunHan/go_study/blob/c3ff62212208d1d65ac838d443c7c22fc92ebcd2/goto.md) 


## [Switch](https://github.com/kyunghyunHan/go_study/blob/da340d061c02304fafdbe324439eb4a1f91b7ace/switch.md) 

## [Array](https://github.com/kyunghyunHan/go_study/blob/e7e3267cf53cf89426232efa34635a7d47fd2727/arr.md)
## [Slice](https://github.com/kyunghyunHan/go_study/blob/20b7a158827c881e3c9055357917f4c969c1e736/slice.md)

제로갑:배열만 선언해놓고, 값을 할당하지 않은 경우 각 타입에 맞는 제로값이 설정된다  
인덱스 제로베이스 : 0 부터 시작 length-1까지 해당  
슬라이스:크기가 정해지지 않은 배열이며 길이가 동적으로 변할수 있다.  
go의 슬라이스는 배열의 뷰(view)일뿐 다른 의미는 가지지 않는다 .  
make()빌트인 함수 사용하여 생성  
또한 슬라이스를 변경하면 원본 배열도 바뀌며 원본 배열을 바꾸면 슬라이스도 변경된다.  
append()슬라이스에 값을 추가할떄 사용   

## [Map](https://github.com/kyunghyunHan/go_study/blob/main/map.md)
<img width="461" alt="스크린샷 2021-10-27 오후 2 35 41" src="https://user-images.githubusercontent.com/88940298/139005924-44e7384d-59ac-48a1-a11d-bdac7f2309c1.png">

맵은 자바스크립트의 객체 리터럴과 비슷하게 키와 값을 가질 수 있는 자료구도이다  
맵을 사용하면 배열과 슬라이스를 사용했을떄 거처야할 일부 번거로운 연산과정을 거칠 필요가 없이 키와 값이 매핑되어 있기 떄문에 접근이나 값을 수정하는 것이 수월하다.  
데이터가 순서대로 정렬되지 않는 특징이 있어 정려로딘 데이터를 원한다면 별도의 정렬 함수를 사용해야한다.  
delete()함수를 사용하면 맵에서 특정 키에 해당하는 값을 제거 할수 있다  
for~range를 사용하여 순회할수 있으며, 순서가 지정되있지 않기 떄문에 무엇이 먼저 나올수 있을지 모른다는것점이 있다.    
<img width="430" alt="스크린샷 2021-10-27 오후 2 53 48" src="https://user-images.githubusercontent.com/88940298/139007738-0376e5af-079c-44c8-ae20-228d1847a9f3.png">
<img width="467" alt="스크린샷 2021-10-27 오후 3 09 01" src="https://user-images.githubusercontent.com/88940298/139009338-e0c7544e-8e2a-4d2b-831b-f5cbd68e4c7d.png">  


## [func]()
## [closure]()
## [defer]()
## [panic]()


## [Pointers](https://github.com/kyunghyunHan/go_study/blob/6e6df9cdca4d7e9a5015e4da88af43d8daee8f0a/pointer.md)
<img width="171" alt="스크린샷 2021-10-27 오후 2 20 51" src="https://user-images.githubusercontent.com/88940298/139004540-59a4b021-cb73-4120-aa5c-1f8cf0c42b13.png">
& :메모리의 주소값
* : 주소에 담긴 값을 살펴볼수 잇음

<img width="340" alt="스크린샷 2021-10-27 오후 2 26 28" src="https://user-images.githubusercontent.com/88940298/139005018-256a8d69-4597-40b3-aa49-0cd0cdc570c3.png">



## [struct](https://github.com/kyunghyunHan/go_study/blob/20b7a158827c881e3c9055357917f4c969c1e736/struct.md)

go 에는 클래스가 없다.  객체라는 단어는 거의 사용하지 않으며 c언어처럼 구조체라는 존재가 있다.  
구조체는 클래스처럼 객체를 찍어내기위한 판 이라 접근하면 좋다.  
struct라는 키워드를 사용하며 person이라는구조체를 선언하고 main에서
구조체를 타입으로 갖는 변수를 선언하고 값을 할당하는 모습을 보여주고 있다.  
메서드는 일반적은 메서드파라메터외에 리시버파라메터라는것을 사용한다.이는 go의 특징이다.  
func다음에 파라메타 같은 형태로 나타낸 것을 볼수 있는데 리시버파라메타이다.this키워드를 대신하며 this에비해 타입을 명시해줄수 있다.  
또한 리시버파라메타에 대해 포인터 자료형으로 주었는데 포인터로 주지 않으면 값복사가 되어 실제로 원본이 변화하지 않게 된다.따라서 포인터 리시버파라메타를 주어 원본도 바뀔수 있게 한것이다  
## [interface]()
## [gorutine]()
## [chanl]()
## [Mutex]()
## [reflection]()
## [인터넷 소스사용]()
## [패키지 생성]()
## [문서화]()
## [입력함수 사용]()
## [문자열 입출력]()
## [파일 입출력]()
## [유니코드]()
## [문자열]()
## [정규표햔]()
## [파일 처리]()
## [입출력 인터페이스]()
## [JSON]()
## [압축]()
## [암호화]()
## [정렬]()
## [컨테이너]()
## [TCP프로토콜]()
## [PRC프로토콜]()
## [명령줄]()
## [에러]()
## [단위테스트]()
## 캡슐화 Getter/setter. 
Main패키지에서 접근 가능한것은 대문자로 되어있는 메서드이며 setter를 만들떄 항상 Set*의 형태로한다
또한 New로 시작하는 것은 생성자라 생각하면 좋다


## [서버 연결](https://github.com/kyunghyunHan/go_study/blob/f29db2308b47119deb3f708d129a7fa0df3fbeea/golangMysql.md)
## [Mysql 연결]()

