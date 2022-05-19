
# 💯Go 공부 및 정리 <img width="100" alt="스크린샷 2021-11-24 오후 12 28 54" src="https://user-images.githubusercontent.com/88940298/143169831-99d59153-1735-49a6-8d7c-d3768dfc9400.png">
## 🌱변수 
## 🌱상수 const

- go에서는 타입을 정하지 않으면 go에서 알아서 타입을 지정해준다  
- go가 다른언어랑 다른 점은 클래스가 없다는 것이다.메서드라고 표현하는 것도 표현 방식이 다른 뿐 일반 함수와 유사하게 표시된다  
- go는 리턴 값을 여러개의 값으로 돌릴 수 있다.  
- go&go-routine go 키워드는 멀티스레드를 사용할때 유용하며 동시성을 가진다.  
- go get go get을 통하여 다른 패키지를 다운받을 수 있다.얻어온 패키지는 GOPATH로 다운로드 된다.GOPATH는 작업공간이라고도 하며 사용자마다 별도로 관리된다  
- go mod 사용자 정의 모듈을 만들떄 사용한다. example.com폴더로 진입하여 go mod init example.com명령어를 통해 모듈을 정의한다  
<img width="355" alt="스크린샷 2021-10-27 오후 2 11 44" src="https://user-images.githubusercontent.com/88940298/139003610-95bf570c-fa06-4959-815c-5e2bea14f2c7.png">

## 🌱바이트
byte에는 보통 16진수, 문자 값으로 저장합니다. 실무에서는 바이너리 파일에서 데이터를 읽고 쓰거나, 데이터를 암호화, 복호화할 때 주로 사용!
```
var num1 byte = 10   // 10진수 저장
var num2 byte = 0x32 // 16진수 저장
var num3 byte = 'a'  // 문자 저장
```


## 🌱룬 
rune은 유니코드(UTF-8) 문자 코드를 저장할 때 사용합니다. ' ' (작은 따옴표)로 묶어주어야 하며 문자 그대로 저장하거나 \u 또는 \U를 사용하여 유니코드 문자 코드로 저장할 수도 있다. 여기서 \U는 값을 16진수 8자리로 맞춰주어야 한다
```
var r1 rune = '한'
var r2 rune = '\ud55c'     // 한
var r3 rune = '\U0000d55c' // 한
```


## 🌱숫자 연산. 
숫자 연산에는 덧셈(+), 뺄셈(-), 곱셈(*), 나눗셈(/), 나머지(%), 시프트(<<, >>), 비트 반전(^) 연산자를 사용할 수 있다
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

## 🌱오버플로우와 언더플로우. 
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
## 🌱문자열 길이. 
```
var s1 string = "한글"
var s2 string = "Hello"

fmt.Println(len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
fmt.Println(len(s2)) // 5: 알파벳 5글자이므로 5
```

## 🌱열거형. 
## 🌱[패키지](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/package.md)
## 🌱[fmt](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/fmt.md)
## 🌱[if문](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/if.md)
## 🌱[for문](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/for.md) 
## 🌱[goto](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/goto.md) 
## 🌱[Switch](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/switch.md) 
## 🌱[Array](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/arr.md)
## 🌱[Slice](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/slice.md)
## 👩🏻‍🎓[Map](https://github.com/kyunghyunHan/go_study/blob/8b7e2994f1f700cd0a1ef59e9e980166b545c5c0/map.md)
## 👩🏻‍🎓[func](https://github.com/kyunghyunHan/go_study/blob/2ec57d0e5802105739988cedd75e60055ba6600e/func.md)
## 🌱[closure](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/closure.md)
## 👩🏻‍🎓[defer](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/defer.md)
## 👩🏻‍🎓[panic](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/panic.md)
## 🌱[Pointers](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/pointer.md)
## 👩🏻‍🎓[struct](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/struct.md)
## 👩🏻‍🎓[interface](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/interface.md)
## 👩🏻‍🎓[gorutine](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/goroutine.md)
## 🎓🔥[chanl](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/channel.md)
## 👩🏻‍🎓[Mutex](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/mutex.md)
## 🎓[reflection](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/reflection.md)
## 🔥[인터넷 소스사용](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/iunternetsource.md)
## ❓[패키지 생성](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/packageadd.md)
## 🧢[문서화](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/forder.md)
## ✨[입력함수 사용](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Scan.md)
## 👬[문자열 입출력](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/langscan.md)
## 🌱[파일 입출력](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/filescan.md)
## 👨‍❤️‍💋‍👨[유니코드](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/UTF-8.md)
## ✨[문자열](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/monlang.md)
## 📖[정규표햔](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/regular.md)
## 🖼[파일 처리](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/file.md)
## 🍎[입출력 인터페이스](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/ioReader.md)
## 🪞[JSON](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/json.md)
## 😀[압축](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/gzip.md)
## ☕️[암호화](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/hash.md)
## 👮🏼‍♀️[정렬](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Sort.md)
## ♥️[컨테이너](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/container.md)
## ♥️[TCP프로토콜](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/TCP.md)
## 🫔[PRC프로토콜](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/RPC.md)
## 📆[명령줄](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/ARG.md)
## 📖[에러](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Error.md)
## 💯[단위테스트](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/TEST.md)
## 😀캡슐화 Getter/setter. 
## 🌱[서버 연결](https://github.com/kyunghyunHan/go_study/blob/f29db2308b47119deb3f708d129a7fa0df3fbeea/golangMysql.md)
## 🌱[Mysql 연결]()

