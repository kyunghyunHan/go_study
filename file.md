# 👩🏻‍🎓파일 처리

## 💯파일쓰기 
os 패키지에서 제공하는 파일 함수와 파일 쓰기 함수

- func Create(name string) (file *File, err error): 기존 파일을 열거나 새 파일을 생성
- func (f *File) Close() error: 열린 파일을 닫음
- func (f *File) Write(b []byte) (n int, err error): 파일에 값을 씀. 파일에 쓴 데이터의 길이와 에러 값을 리턴
hello.txt 파일에 Hello, world! 문자열을 저장

```
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt") // hello.txt 파일 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello, world!"

	n, err := file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
}
```
```
13 바이트 저장 완료
```
- os.Create 함수로 hello.txt 파일을 생성,(hello.txt 파일이 이미 있으면 파일을 열고 값을 덮어씁니다). 여기서 파일 생성에 성공하면 file 인스턴스가 리턴됩니다. 그리고 파일을 열었으면 항상 file.Close 함수로 닫아줍니다(파일을 닫을 때 지연 호출 defer를 사용하면 편리합니다. 코드상에서는 파일을 연 즉시 파일을 닫는 코드가 나오므로 파일 닫기를 까먹는 일이 줄어듭니다).

file 인스턴스에서 Write 함수를 사용하면 []byte 슬라이스의 내용을 파일에 저장할 수 있습니다. 여기서는 문자열 변수 s를 []byte 형식으로 변환하여 Hello, world!를 파일에 저장했습니다.

## 💯파일 읽기

 os 패키지에서 제공하는 파일 열기, 파일 정보, 파일 읽기 함수

- func Open(name string) (file *File, err error): 파일 열기
- func (f *File) Stat() (fi FileInfo, err error): 파일의 정보를 얻어옴
- func (f *File) Read(b []byte) (n int, err error): 파일에서 값을 읽음. 파일에서 읽은 데이터의 길이와 에러 값을 리턴
```
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt") // hello.txt 파일을 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	fi, err := file.Stat() // 파일 정보를 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	n, err := file.Read(data) // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
```
```
13 바이트 읽기 완료
Hello, world!
```
- os.Open 함수로 hello.txt 파일을 엽니다(hello.txt 파일이 없으면 에러가 발생합니다). 여기서 파일 열기에 성공하면 file 인스턴스가 리턴됩니다. 단, os.Open 함수로 파일을 열었을 때는 읽기만 할 수 있습니다.

- 파일을 읽기 전에 파일 크기를 구해와야 하므로 file.Stat 함수로 파일의 정보를 얻어옵니다. 그리고 fi.Size 함수로 파일 크기만큼 슬라이스를 생성합니다. 마지막으로 file.Read 함수에 슬라이스를 넣으면 파일의 내용을 읽어올 수 있습니다.

## 💯파일 읽기 쓰기
os 패키지에서 제공하는 파일 열기 함수와 파일 포인터 설정 함수

- func OpenFile(name string, flag int, perm FileMode) (file *File, err error): 파일 플래그, 파일 모드를 지정하여 파일 열기
- func (f *File) Seek(offset int64, whence int) (ret int64, err error): 파일을 읽거나 쓸 위치로 이동
```
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
                                                  // 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644))                // 파일 권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat() // 파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET) // Write 함수로 인해 파일 읽기/쓰기 위치가
                                  // 이동했으므로 file.Seek 함수로 읽기/쓰기 위치를
                                  // 파일의 맨 처음(0)으로 이동
	n, err = file.Read(data)  // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
```
```
15 바이트 저장 완료
15 바이트 읽기 완료
안녕하세요
```
- os.OpenFile 함수는 파일 이름, 플래그, 파일 모드를 받습니다. 
- 플래그의 종류

|플래그	|설명|
|-----|-------|
|os.O_RDONLY	|읽기 전용으로 파일을 엽니다.|
|os.O_WRONLY	|쓰기 전용으로 파일을 엽니다.|
|os.O_RDWR	|읽기, 쓰기를 모두 할 수 있도록 파일을 엽니다.|
|os.O_APPEND|	파일을 저장했을 때 끝부분에 내용을 추가합니다.|
|os.O_CREATE|	파일이 없으면 파일을 새로 생성합니다. 파일이 있으면 파일을 해당 파일을 엽니다.|
|os.O_EXCL	|파일이 없을 때만 새로 생성하며 O_CREATE와 함께 사용합니다.|
|os.O_SYNC	|동기(Synchronous) I/O를 사용합니다.|
|os.O_TRUNC	|파일이 있다면 파일을 연 뒤 내용을 삭제합니다.|


- 플래그는 OR(|) 연산자를 사용하여 조합할 수 있습니다. 여기서는 os.O_RDWR로 읽기, 쓰기를 하고, os.O_CREATE로 파일이 없으면 파일을 생성하고, os.O_TRUNC로 파일이 있을 때는 파일을 연 뒤 내용을 삭제합니다. 그리고 파일 모드는 유닉스/리눅스 형식의 권한(Permission)을 지정합니다. 여기서는 os.FileMode 타입으로 0644를 지정했습니다.

- file.Write 함수로 파일에 문자열을 저장합니다. 그리고 다시 file.Stat 함수로 파일 크기를 구해와서 슬라이스를 생성합니다.

- 파일을 읽기 전에 file.Read, file.Write 함수로 파일을 읽거나 쓰면 해당 크기만큼 읽기/쓰기 위치(Offset)가 이동합니다. 따라서 앞에서 file.Write 함수로 파일에 내용을 저장했으므로 Offset이 15바이트만큼 앞으로 이동했습니다. 파일을 처음부터 읽으려면 file.Seek(0, os.SEEK_SET)처럼 파일 읽기/쓰기 위치를 초기화해줍니다.

file.Seek 함수에서 첫 번째 매개변수는 offset(바이트)이며 두 번째 매개변수는 기준점입니다. 다음은 설정할 수 있는 기준점입니다.

- os.SEEK_SET: 파일의 맨 처음을 기준으로 합니다.
- os.SEEK_CUR: 현재 읽기/쓰기 offset 값을 기준으로 합니다.
- os.SEEK_END: 파일의 맨 끝을 기준으로 합니다.
file.Seek(0, os.SEEK_SET)이면 파일의 맨 처음부터 0바이트이므로 파일의 맨 앞이 됩니다. 그리고 file.Read 함수로 파일의 내용을 읽습니다.
```
유닉스/리눅스 파일 권한은 세 부분으로 구성되어 있습니다.

파일 소유자(User)
파일 그룹(Group)
파일에 권한이 없는 소유자(Others)
여기서 읽기, 쓰기, 실행 권한을 설정할 수 있습니다.

읽기(Read, r): 4
쓰기(Write, w): 2
실행(Execute, x): 1
```
파일 권한이 0644라면 소유자는 읽고 쓸 수 있으며(6=4+2) 그룹과 권한이 없는 소유자는 읽기(4)만 할 수 있습니다. 처음 0은 Sticky bit, SetUID, SetGID 중 어떤 것도 사용하지 않는 설정입니다.


## 💯ioutil 패키지 사용하기
ioutil 패키지를 사용하면 좀 더 간단하게 파일을 읽고 쓸 수 있습니다.

ioutil 패키지에서 제공하는 파일 함수

- func WriteFile(filename string, data []byte, perm os.FileMode) error: 파일 쓰기. 에러 값을 리턴
- func ReadFile(filename string) ([]byte, error): 파일 읽기. 읽은 데이터(바이트 슬라이스)와 에러 값을 리턴

```
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := "Hello, world!"

	err := ioutil.WriteFile("hello.txt", []byte(s), os.FileMode(644))
                                                  // s를 []byte 바이트 슬라이스로 변환, s를 hello.txt 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile("hello.txt") // hello.txt의 내용을 읽어서 바이트 슬라이스 리턴
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
```
ioutil.WriteFile 함수는 파일명, 저장할 데이터, 파일 모드만 지정해주면 간단하게 파일에 데이터를 저장할 수 있습니다(이미 파일이 있으면 데이터를 덮어씁니다). 마찬가지로 ioutil.ReadFile 함수는 파일명만 지정해주면 간단하게 파일을 읽어올 수 있습니다.
