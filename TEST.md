# 👩🏻‍🎓단위테스트

날이 갈 수록 소프트웨어는 덩치가 커지고 점점 복잡해져갑니다. 특히 서비스 환경이나 전체적인 실행 과정 중에서 버그를 잡아내기는 정말 어려운 일입니다. 이러한 상황을 개선하기 위해 소프트웨어를 잘게 쪼개어 테스트하는 유닛 테스트 기법이 등장했습니다. 큰 구조 속에서 버그를 찾는 일보다 작은 단위에서 버그를 찾고 고치는 일이 훨씬 쉽기 때문입니다.

Go 언어는 자체적으로 단위 테스트를 제공합니다. 따라서 개발 과정 중에 자연스럽게 단위 테스트를 구성할 수 있습니다. 또한, 컴파일러에 명령이 내장되어 있으므로 따로 설치나 설정을 하지 않고도 편리하게 사용할 수 있습니다.

이번 예제는 ‘Unit 39 패키지 만들기’에서 만들었던 calc 패키지와 sum 함수를 사용하겠습니다. 다음 내용을 GOPATH/src/calc/sum.go 파일로 저장합니다. GOPATH를 설정하는 방법은 ‘Unit 3 기본 디렉터리 설정하기’를 참조하기 바랍니다.

```
package calc

func Sum(a int, b int) int {
	return a + b
}
```
이제 테스트 코드를 작성해보겠습니다. 다음 내용을 GOPATH/src/calc/sum_test.go 파일로 저장합니다. 파일 이름은 테스트할 소스 파일 이름에 _test.go를 붙이므로 다른 테스트 코드를 작성할 때도 이 규칙을 따르도록 합니다.
```
package calc

import "testing" // 테스트 코드는 항상 testing 패키지를 가져옴

func TestSum(t *testing.T) { // 함수 이름은 Test로 시작,
                             // Test 다음의 함수 첫 글자는 대문자, 매개변수는 *testing.T
	r := Sum(1, 2) // 덧셈 함수로 1과 2를 더하면 결과는 3이 나와야 함
	if r != 3 {    // 3이 아닐 때는 테스트 실패
		t.Error("결괏값이 3이 아닙니다. r=", r) // 에러 발생
	}
}
```

단위 테스트를 수행하는 함수는 다음 규칙을 지켜야 합니다. 그렇지 않으면 컴파일러에서 테스트 코드를 인식하지 못합니다.

- 테스트 함수는 TestSum처럼 항상 Test로 시작합니다.
- Test 다음에 함수 이름이 오며 함수 이름의 첫 글자는 항상 대문자로 만들어야 합니다. 예) TestSum, TestReadData, TestConvert
- 항상 *testing.T 타입의 매개변수를 받습니다.
테스트 코드를 작성하는 방법은 간단합니다. 데이터를 함수에 넣은 뒤 나와야 하는 결괏값과 맞는지만 판별하면 됩니다. 여기서는 Sum 함수가 두 숫자를 더하므로 1과 2를 더한 값인 3이 나와야 정상입니다. 만약 3이 나오지 않는다면 함수가 잘못 구현되었으므로 t.Error 함수를 사용하여 에러를 출력합니다.

콘솔(터미널)에서 GOPATH/src/calc 디렉터리로 이동한 뒤 다음 명령을 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였습니다.

```
~$ cd $GOPATH/src/calc
~/hello_project/src/calc$ go test
PASS
ok      calc    0.088s
```

go test 형식입니다. 명령 실행이 끝난 뒤 결과가 PASS로 나오면 테스트에 통과한 것입니다. 만약 테스트에 통과하지 못하면 FAIL과 함께 앞에서 설정한 에러 메시지가 출력됩니다.

이번에는 테스트 케이스를 좀 더 많이 준비해보겠습니다.
```
package calc

import "testing"

type TestData struct { // 테스트 데이터 구조체
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{ // 테스트 데이터 슬라이스
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) {
	for _, d := range testdata {               // 반복문으로 데이터를 꺼낸 뒤
		r := Sum(d.argument1, d.argument2) // Sum 함수를 테스트
		if r != d.result {
			t.Errorf(
				"%d + %d의 결과값이 %d(이)가 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}
```
테스트 데이터 구조체를 작성합니다. 이 구조체의 형태는 테스트할 함수에 따라 달라집니다. 여기서는 덧셈 함수를 위한 테스트 데이터이므로 int 형으로 된 argument1, argument2, result를 필드로 가지고 있습니다.

슬라이스를 생성하여 테스트 데이터를 채워넣습니다. 테스트 데이터는 함수에서 발생할 수 있는 다양한 예외 상황을 고려하여 준비합니다. 여기서는 덧셈 함수가 매우 간단하여 예외 상황이 별로 없지만 실무에서는 다양한 예외 상황들이 나올 것입니다.

테스트 함수에서는 for 반복문을 사용하여 슬라이스에서 데이터를 꺼낸 뒤 Sum 함수에 넣습니다. 그리고 결괏값을 테스트 데이터와 비교하여 잘못된 결과이면 에러를 출력합니다.

Go 언어의 기본 패키지를 보면 테스트 코드를 작성하는 방법이 아주 잘 나와있으므로 참고하기 바랍니다.

## ✏️벤치마크 테스트 사용하기
테스트 코드는 함수에 매개변수를 넣고 결괏값이 정상적으로 동작하는지만 검사합니다. 벤치마크 테스트는 성능을 측정하는 기능입니다.

덧셈 함수의 성능을 측정해보겠습니다. 다음 내용을 GOPATH/src/calc/sum_test.go 파일로 저장합니다.
```
package calc

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

```
벤치마크를 수행하는 함수는 다음 규칙을 지켜야 합니다. 그렇지 않으면 컴파일러에서 벤치마크 테스트 코드를 인식하지 못합니다.

- 테스트 함수는 BenchmarkSum처럼 항상 Benchmark로 시작합니다.
- Benchmark 다음에 함수 이름이 오며 함수 이름의 첫 글자는 항상 대문자로 만들어야 합니다. 예) BenchmarkSum, BenchmarkData, BenchmarkConvert
- 항상 *testing.B 타입의 매개변수를 받습니다.
벤치마크 테스트는 b.N에 정의된 값 만큼 for 반복문을 실행합니다. 그리고 반복문 안에서 성능을 측정할 함수를 호출합니다.

콘솔(터미널)에서 GOPATH/src/calc 디렉터리로 이동한 뒤 다음 명령을 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였습니다.
```
~$ cd $GOPATH/src/calc
~/hello_project/src/calc$ go test -bench .
testing: warning: no tests to run
PASS
BenchmarkSum    2000000000               0.81 ns/op
ok      calc    1.805s
```
go test -bench . 형식입니다. 즉, -bench . 옵션을 붙여주면 Benchmark로 시작하는 함수를 자동으로 실행해줍니다.

Sum 함수를 20억 회 실행하였고 for 반복문 한 번에 0.81 나노초(ns)가 걸렸다는 뜻입니다. 여기서는 덧셈 함수가 매우 간단하여 벤치마크 결과가 다소 의미가 없지만 실무에서는 유용하게 활용될 것입니다.


