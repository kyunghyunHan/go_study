# 문서화

Go 언어는 패키지를 만들면서 문서화도 할 수 있습니다. 다음과 같이 calc 패키지의 Sum 함수에 대한 설명을 작성합니다.
```
// 계산 패키지
package calc

// 두 정수를 더함
func Sum(a int, b int) int {
	return a + b
}
```
package 키워드와 함수 정의 바로 위에 // 주석 형태로 함수 설명을 작성합니다. // 주석을 여러 줄로 작성해도 됩니다.

이제 패키지 및 함수의 정보를 출력해보겠습니다. 다음과 같이 GOPATH 디렉터리에서 godoc 명령을 실행합니다.

```
~$ cd $GOPATH
~/hello_project$ godoc calc
PACKAGE DOCUMENTATION

package calc
    import "calc"

    계산 패키지

FUNCTIONS

func Sum(a int, b int) int
    두 정수를 더함

```
godoc <패키지 이름> 형식입니다. 이렇게 실행하면 소스 코드에 주석 형태로 작성한 설명이 코드와 함께 정리되어 표시됩니다.
 
다음과 같이 godoc <패키지 이름> <함수 이름> 형식으로 특정 함수의 정보를 출력할 수도 있습니다.
```
~$ cd $GOPATH
~/hello_project$ godoc calc Sum
func Sum(a int, b int) int
    두 정수를 더함

```
이번에는 웹 브라우저에서 문서를 볼 수 있도록 웹 서버를 실행해보겠습니다.
```
~$ cd $GOPATH
~/hello_project$ godoc -http=:6060
```
godoc -http=:<포트 번호> 형식입니다. 이렇게 실행하면 설정한 포트로 웹 서버가 실행됩니다.
