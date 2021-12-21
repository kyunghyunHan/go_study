# 👩🏻‍🎓명령줄 

프로그램을 만들다보면 실행할 때 옵션을 설정하고 싶을 때가 있습니다. Go 언어에서는 명령줄에서 설정한 옵션을 간단하게 사용할 수 있습니다.

다음은 os 패키지에서 제공하는 변수입니다.

- var Args []string: 명령줄에서 설정한 옵션
다음 내용을 GOPATH/src/cmdargs/cmdargs.go 파일로 저장합니다. GOPATH를 설정하는 방법은 ‘Unit 3 기본 디렉터리 설정하기’를 참조하기 바랍니다.

```
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}
```
Go 언어의 main 함수에는 매개변수가 없습니다. 그래서 명령줄에서 설정한 옵션을 가져오려면 os.Args 슬라이스를 사용합니다.

이제 소스 파일을 컴파일한 뒤 실행해봅니다. 콘솔(터미널)에서 GOPATH/src/cmdargs 디렉터리로 이동한 뒤 cmdargs 파일을 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였습니다.

```
~$ cd $GOPATH/src/cmdargs
~/hello_project/src/cmdargs$ ./cmdargs -hello=1 -world=2
[./cmdargs -hello=1 -world=2]
```
실행 파일을 실행하면 현재 실행 파일과 설정한 옵션이 출력됩니다. os.Args의 첫 번째 요소는 항상 현재 실행 파일의 경로이며 두 번째 요소부터 옵션입니다(Windows에서는 현재 실행 파일의 경로는 절대 경로입니다).

os.Args의 문자열을 일일이 분석하여 사용하기에는 매우 불편합니다. 그래서 Go 언어에서는 명령줄 옵션을 편리하게 사용할 수 있도록 flag 패키지를 제공합니다.

다음은 flag 패키지에서 제공하는 함수입니다.

- func String(name string, value string, usage string) *string: 명령줄 옵션을 받은 뒤 문자열로 저장
- func Int(name string, value int, usage string) *int: 명령줄 옵션을 받은 뒤 정수로 저장
- func Float64(name string, value float64, usage string) *float64: 명령줄 옵션을 받은 뒤 실수로 저장
- func Bool(name string, value bool, usage string) *bool: 명령줄 옵션을 받은 뒤 불로 저장
- func Parse(): 명령줄 옵션(os.Args)의 내용을 각 자료형별로 분석
- func NFlag() int: 명령줄 옵션의 개수를 리턴
- var Usage = func() {}: 명령줄 옵션의 기본 사용법 출력
다음 내용을 GOPATH/src/cmdflag/cmdflag.go 파일로 저장합니다.

```
package main

import (
	"flag"
	"fmt"
)

func main() {
	title := flag.String("title", "", "영화 이름")      // 명령줄 옵션을 받은 뒤 문자열로 저장
	runtime := flag.Int("runtime", 0, "상영 시간")      // 명령줄 옵션을 받은 뒤 정수로 저장
	rating := flag.Float64("rating", 0.0, "평점")       // 명령줄 옵션을 받은 뒤 실수로 저장
	release := flag.Bool("release", false, "개봉 여부") // 명령줄 옵션을 받은 뒤 불로 저장

	flag.Parse() // 명령줄 옵션의 내용을 각 자료형별로 분석

	if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
		flag.Usage()   // 명령줄 옵션 기본 사용법 출력
		return
	}

	fmt.Printf(
		"영화 이름: %s\n상영 시간: %d분\n평점: %f\n",
		*title, // 포인터이므로 값을 꺼낼 때는 역참조 사용
		*runtime,
		*rating,
	) // 명령줄 옵션으로 받은 값을 출력

	if *release == true {
		fmt.Println("개봉 여부: 개봉")
	} else {
		fmt.Println("개봉 여부: 미개봉")
	}
}
```
- 간단하게 영화 정보를 명령줄 옵션으로 넘겨주고 화면에 출력하는 예제입니다. 먼저 명령줄 옵션으로 받을 항목들을 설정합니다. flag.String, flag.Int, flag.Float64, flag.Bool과 같이 타입별로 사용할 수 있는 함수가 준비되어 있습니다. 매개변수는 옵션명, 기본값, 설명 순이며 옵션 값이 저장될 포인터가 리턴됩니다(flag.StringVar처럼 뒤에 Var가 붙은 함수는 첫 번째 매개변수로 옵션 값을 저장할 포인터를 받을 뿐 사용 방법은 같습니다).

- flag.Parse() 함수를 실행하면 각 옵션 값 포인터에 명령줄에서 설정한 옵션 값이 저장됩니다. 그리고 flag.NFlag() 함수로 설정된 옵션의 개수를 알 수 있습니다. 만약 설정된 옵션이 없다면 flag.Usage() 함수로 옵션 사용법을 출력합니다.

- 이제 소스 파일을 컴파일한 뒤 실행해봅니다. 콘솔(터미널)에서 GOPATH/src/cmdflag 디렉터리로 이동한 뒤 cmdflags 파일을 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다).
```
~$ cd $GOPATH/src/cmdflag
~/hello_project/src/cmdflag$ ./cmdflag -title=인셉션 -rating=8.8 -runtime=148 -release=true
영화 이름: 인셉션
상영 시간: 148분
평점: 8.800000
개봉 여부: 개봉
```
cmdflag 파일을 실행할 때 아무 옵션을 설정하지 않으면 다음과 같이 옵션 사용법이 출력됩니다.
```
~$ cd $GOPATH/src/cmdflag
~/hello_project/src/cmdflag$ ./cmdflag
Usage of ./cmdflag:
  -rating=0: 평점
  -release=false: 개봉 여부
  -runtime=0: 상영 시간
  -title="": 영화 이름
```
