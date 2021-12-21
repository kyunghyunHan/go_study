# 👩🏻‍🎓입출력 인터페이스

- 입출력 인터페이스를 받는 여러 패키지가 있으므로 화면에 출력하는 부분부터 파일, 문자열까지 같은 인터페이스로 처리할 수 있습니다.

io.Reader와 io.Writer의 정의는 다음과 같습니다.

```
type Reader interface {
        Read(p []byte) (n int, err error)
}

type Writer interface {
        Write(p []byte) (n int, err error)
}
```
어떤 구조체든 매개변수로 바이트 슬라이스를 받고, 정수와 에러 값을 리턴하는 Read 함수를 가지고 있으면 io.Reader 인터페이스를 따른다고 할 수 있습니다. 마찬가지로 매개변수로 바이트 슬라이스를 받고 정수와 에러 값을 리턴하는 Write 함수를 가지고 있으면 io.Writer 인터페이스를 따른다고 할 수 있습니다.

파일 처리하기
다음은 bufio 패키지에서 제공하는 입출력 인터페이스 함수입니다.

- func NewReader(rd io.Reader) *Reader: io.Reader 인터페이스로 io.Reader 인터페이스 따르는 읽기 인스턴스 생성
- func NewWriter(w io.Writer) *Writer: io.Writer 인터페이스로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 생성
- func (b *Writer) WriteString(s string) (int, error): 문자열을 버퍼에 저장
- func (b *Writer) Flush() error: 버퍼의 데이터를 파일에 저장
bufio는 Buffered I/O를 뜻하며 io.Reader, ioWriter 인터페이스를 받습니다. 이제 bufio를 사용하여 파일을 읽고 써보겠습니다.

```
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성,
                                                  // 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644),                // 파일 권한은 644
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := bufio.NewWriter(file)     // io.Writer 인터페이스를 따르는 file로
                                       // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	w.WriteString("Hello, world!") // 쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
	w.Flush()                      // 버퍼의 내용을 파일에 저장

	r := bufio.NewReader(file)   // io.Reader 인터페이스를 따르는 file로
                                     // io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
	fi, _ := file.Stat()         // 파일 정보 구하기
	b := make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET) // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	r.Read(b)                 // 읽기 인스턴스로 파일의 내용을 읽어서 b에 저장

	fmt.Println(string(b)) // 문자열로 변환하여 b의 내용 출력
}
```
- bufio.NewWriter 함수에 file 인스턴스를 넣으면 io.Write 인스턴스를 따르는 쓰기 인스턴스를 리턴합니다(물론 file 인스턴스도 io.Writer 인터페이스를 따릅니다). 그리고 WriteString과 같은 메서드로 파일에 문자열을 저장할 수 있습니다. 단, bufio이므로 파일에 바로 저장하지 않고 임시 공간(버퍼)에 쌓아둡니다. 따라서 버퍼의 내용을 파일에 완전히 저장하려면 Flush 메서드를 사용합니다.
- 마찬가지로 bufio.NewReader 함수에 file 인스턴스를 넣으면 io.Reader 인터페이스를 따르는 읽기 인스턴스를 리턴합니다(file 인스턴스도 io.Reader 인터페이스도 따릅니다). 그리고 읽기 인스턴스로는 읽기 작업을 할 수 있습니다. 여기서 파일을 읽기 전에 file.Stat 함수로 파일 크기를 구해온 뒤 슬라이스를 생성합니다.

- 앞에서 먼저 WriteString 함수로 파일을 조작했으므로 파일 읽기/쓰기 위치(offset)가 이동하였습니다. 그러므로 file.Seek 함수로 offset을 초기화해 준 뒤 Read 메서드로 파일의 내용을 읽습니다.

- 매우 큰 파일을 읽을 때는 읽을 크기만큼의 슬라이스를 할당 한 뒤 file.Seek 함수로 읽을 위치를 옮겨가면서 일부만 읽으면 됩니다.

## 💯문자열을 파일로 저장하기

다음은 strings 패키지에서 제공하는 입출력 인터페이스 함수입니다.

- func NewReader(s string) *Reader: 문자열로 io.Reader 인터페이스를 따르는 읽기 인스턴스를 생성합니다.
다음은 bufio 패키지에서 제공하는 입출력 인터페이스 함수입니다.

- func (b *Writer) ReadFrom(r io.Reader) (n int64, err error): io.Reader의 데이터를 읽어서 io.Writer에 저장
이번에는 문자열을 io.Reader 인터페이스로 만든 뒤 파일에 저장해보겠습니다.

```
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
                                  // 읽기 인스턴스 r 생성

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 file로
                                   // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	w.ReadFrom(r)              // 읽기 인스턴스 r에서 데이터를 읽어서 w의 버퍼에 저장
	w.Flush()                  // 버퍼의 내용을 파일에 저장
}
```
- strings.NewReader 함수를 사용하면 문자열로 io.Reader 인터페이스를 따르는 읽기 인스턴스를 만들 수 있습니다. 그리고 w.ReadFrom(r)처럼 io.Reader를 io.Writer로 가져올 수 있습니다. 데이터를 가져온 뒤에는 Flush 메서드를 사용하여 버퍼의 내용을 파일에 저장합니다.

- 다음과 같이 io.Copy 함수를 사용하면 io.Reader의 데이터를 io.Writer로 복사할 수 있습니다. 앞의 예제와 마찬가지로 문자열을 파일에 저장합니다.

- func Copy(dst Writer, src Reader) (written int64, err error): io.Reader를 io.Writer로 복사

```
w := bufio.NewWriter(file)
io.Copy(w, r)
w.Flush()
```
## 💯문자열을 파일로 출력하기

다음과 같이 문자열 io.Reader를 그대로 화면에 출력할 수도 있습니다.

```
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
                                  // 읽기 인스턴스 r 생성

	io.Copy(os.Stdout, r) // os.Stdout에 io.Reader를 복사하면 화면에 그대로 출력됨
}
```
```
Hello, world!
```
- strings.NewReader 함수를 사용하면 문자열을 io.Reader로 만들 수 있습니다. os.Stdout도 io.Writer 인터페이스를 따르기 때문에 io.Copy 함수로 io.Reader를 복사해주면 콘솔(터미널)에 그대로 출력됩니다.

## 💯기본 입출력 함수 사용하기
이번에는 기본 입출력 함수와 입출력 인터페이스를 함께 사용해보겠습니다.

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는
                                  // 읽기 인스턴스 r 생성

	var s1, s2 string
	n, _ := fmt.Fscanf(r, "%s %s", &s1, &s2) // 형식을 지정하여 읽기 인스턴스 r에서
                                                 // 문자열 읽기

	fmt.Println("입력 개수:", n)             // 입력 개수: 2
	fmt.Println(s1)                          // Hello,
	fmt.Println(s2)                          // world!
}
```
- fmt.Fscanf 함수는 file 인스턴스뿐만 아니라 io.Reader 인터페이스를 따르는 모든 인스턴스를 사용할 수 있습니다. 따라서 문자열로 만든 io.Reader도 사용할 수 있습니다. 여기서는 strings.NewReader 함수로 만든 io.Reader에서 fmt.Fscanf 함수로 형식에 맞게 입력을 받습니다.

- 다음은 bufio와 기본 입출력 함수를 사용하여 파일에 값을 저장합니다.
```
package main

import (
	"bufio"
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

	w := bufio.NewWriter(file)                  // file로 io.Writer 인터페이스를 따르는
                                                    // 쓰기 인스턴스 w 생성
	fmt.Fprintf(w, "%d,%f,%s", 1, 1.1, "Hello") // 형식을 지정하여 쓰기 인스턴스 w의 버퍼에
                                                    // 1, 1.1, Hello 저장
	w.Flush()                                   // 버퍼의 내용을 파일에 저장
}
```
bufio.NewWriter 함수로 만든 쓰기 인스턴스에 fmt.Fprintf 함수로 형식을 설정하여 값을 출력하고, Flush 함수로 버퍼의 내용을 파일에 저장합니다. 이렇게 하면 hello.txt 파일에 1,1.100000,Hello가 저장됩니다.
이처럼 io.Reader, io.Writer 인터페이스만 맞으면 어디든 활용할 수 있습니다. Go 언어로 프로그래밍할 때는 이러한 인터페이스 사용법에 익숙해지는 것이 좋습니다.

## 🌱읽기, 쓰기 인터페이스를 함께 사용하기
지금까지 io.Reader, io.Writer를 각각 따로 사용했는데 이번에는 io.ReadWriter 인터페이스로 읽기/쓰기를 처리해보겠습니다.

io.ReadWriter의 정의는 다음과 같습니다.

```
type ReadWriter interface {
        Reader
        Writer
}
```
- 즉, io.ReadWriter 인터페이스는 io.Reader, io.Writer 인터페이스를 포함(임베딩)하고 있습니다.

- 다음은 bufio 패키지에서 제공하는 입출력 인터페이스 함수입니다.

- func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error): io.Reader에서 문자열 한 줄을 읽어서 바이트 슬라이스로 리턴
이제 io.ReadWriter 인터페이스를 사용하여 파일에 문자열을 쓴 뒤 다시 읽어보겠습니다.
```
package main

import (
	"bufio"
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

	r := bufio.NewReader(file) // file로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	rw := bufio.NewReadWriter(r, w) // r, w를 사용하여 io.ReadWriter 인터페이스를 따르는
                                        // 읽기/쓰기 인스턴스 생성
	rw.WriteString("Hello, world!") // 읽기/쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
	rw.Flush()                      // 버퍼의 내용을 파일에 저장

	file.Seek(0, os.SEEK_SET)   // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	data, _, _ := rw.ReadLine() // 파일에서 문자열 한 줄을 읽어서 data에 저장
	fmt.Println(string(data))   // 문자열로 변환하여 data의 내용 출력
}
```
- bufio.NewReader, bufio.NewWriter 함수를 사용하여 io.Reader, io.Writer 인터페이스를 따르는 읽기, 쓰기 인스턴스를 생성합니다. 그리고 bufio.NewReadWriter 함수에 읽기, 쓰기 인스턴스를 넣으면 io.ReadWriter 인터페이스를 따르는 읽기/쓰기 인스턴스가 생성됩니다.

- 읽기/쓰기 인스턴스의 WriteString으로 버퍼에 문자열을 쓴 뒤 버퍼의 내용을 파일에 저장합니다. 그 다음에는 파일의 읽기 위치를 맨 앞으로 이동하고, 읽기/쓰기 인스턴스의 ReadLine으로 파일에서 문자열 한 줄을 읽습니다.

- 이처럼 io.ReadWriter 인터페이스를 사용하면 읽기 및 쓰기 동작을 인스턴스 한 개로 처리할 수 있습니다
