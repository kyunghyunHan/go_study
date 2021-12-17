# 압축

Go 언어는 다양한 압축 알고리즘을 패키지로 제공해줍니다. 보통 압축 알고리즘은 파일, 네트워크 패킷을 압축하여 용량을 줄이고, 압축된 파일, 네트워크 패킷을 해제할 때 사용합니다.

다음은 compress/gzip 패키지에서 제공하는 압축 함수입니다.

- func NewReader(r io.Reader) (*Reader, error): io.Reader 인터페이스로 io.Reader 인터페이스를 따르는 압축 해제 인스턴스 생성
- func NewWriter(w io.Writer) *Writer: io.Writer 인터페이스로 io.Writer 인터페이스를 따르는 압축 인스턴스 생성
다음은 io/ioutil 패키지에서 제공하는 데이터 읽기 함수입니다.

- func ReadAll(r io.Reader) ([]byte, error): io.Reader를 끝(EOF)까지 읽어서 바이트 슬라이스로 리턴
gzip 알고리즘을 사용해서 데이터를 압축한 뒤 파일로 저장해보겠습니다

```
package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile( // 압축할 파일 생성
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := gzip.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 압축 인스턴스 w 생성
	defer w.Close()           // main 함수가 끝나기 직전에 압축 인스턴스를 닫음

	s := "Hello, world!"
	w.Write([]byte(s)) // 압축 인스턴스로 문자열을 압축하여 파일에 저장
	w.Flush()          // 메모리상의 데이터를 파일에 완전히 저장
}
```
gzip.NewWriter 함수에 file 인스턴스를 넣으면 io.Writer 인터페이스를 따르는 압축 인스턴스를 리턴합니다. 그리고 압축 인스턴스의 Write 메서드에 []byte 형식의 데이터를 넣은 뒤 Flush 메서드를 호출하면 자동으로 압축됩니다. 여기서는 문자열을 []byte 형식으로 변환해서 넣었습니다.
Close 메서드로 압축 인스턴스를 반드시 닫아줍니다. 그렇지 않으면 압축이 정상적으로 되지 않습니다.

이번에는 방금 압축한 hello.txt.gz의 압축을 해제해서 읽어보겠습니다.
```
package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("hello.txt.gz") // 압축 파일 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	r, err := gzip.NewReader(file) // file로 io.Reader 인터페이스를 따르는
                                       // 압축 해제 인스턴스 r 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close() // main 함수가 끝나기 직전에 압축 인스턴스를 닫음

	b, err := ioutil.ReadAll(r) // 압축 해제 인스턴스 r을 읽으면
                                    // 압축 해제된 데이터가 리턴됨
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b)) // 문자열로 변환하여 출력
}
```


gzip.NewReader 함수에 file 인스턴스를 넣으면 io.Reader 인터페이스를 따르는 압축 해제 인스턴스를 리턴합니다. 일반 파일 같으면 파일 크기를 구하여 슬라이스를 생성하겠지만 압축 파일은 해제한 뒤의 크기 정보를 구할 수 없습니다. 이때는 ioutil.ReadAll 함수를 사용하면 편리합니다. ioutil.ReadAll 함수에 압축 해제 인스턴스(io.Reader)를 넣으면 자동으로 압축을 해제하여 슬라이스를 리턴합니다.

다른 압축 알고리즘들도 io.Reader, io.Writer 인터페이스를 따르므로 같은 방법으로 압축 및 해제를 할 수 있습니다.

- compress/bzip2: func NewReader(r io.Reader) io.Reader
- compress/zlib:
- - func NewReader(r io.Reader) (io.ReadCloser, error)
- - func NewWriter(w io.Writer) *Writer
- comqpress/flate:
- - func NewReader(r io.Reader) io.ReadCloser
- - func NewWriter(w io.Writer, level int) (*Writer, error)
- compress/lzw:
- - func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
- - func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
