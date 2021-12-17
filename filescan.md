# 파일 입출력

값을 파일로 저장하거나 파일에서 변수로 값을 가져올 수 있습니다.

다음은 os 패키지에서 제공하는 파일 처리 함수입니다.

- func Create(name string) (file *File, err error): 기존 파일을 열거나 새 파일을 생성
- func Open(name string) (file *File, err error): 기존 파일을 열기
- func (f *File) Close() error: 열린 파일을 닫음
다음은 fmt 패키지에서 제공하는 입출력 함수이며 함수를 사용하기 전에 os 패키지의 파일 처리 함수로 파일을 생성하거나 열어야 합니다.

- func Fprint(w io.Writer, a ...interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 파일에 저장
- func Fprintln(w io.Writer, a ...interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장
- func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error): 형식을 지정하여 파일에 저장
- func Fscan(r io.Reader, a ...interface{}) (n int, err error): 파일을 읽은 뒤 공백, 개행 문자로 구분된 문자열에서 입력을 받음
- func Fscanln(r io.Reader, a ...interface{}) (n int, err error): 파일을 읽은 뒤 공백으로 구분된 문자열에서 입력을 받음
- func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error): 파일을 읽은 뒤 문자열에서 형식을 지정하여 입력을 받음
먼저 값을 파일로 저장해보겠습니다.

```
package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello1.txt")        // hello1.txt 파일 생성
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprint(file1, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 파일에 저장

	file2, _ := os.Create("hello2.txt")          // hello2.txt 파일 생성
	defer file2.Close()                          // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintln(file2, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤
	                                             // 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장

	file3, _ := os.Create("hello3.txt")                     // hello3.txt 파일 생성
	defer file3.Close()                                     // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, world!") // 형식을 지정하여 파일에 저장
}
```

지금까지 사용했던 함수와 사용 방법이 비슷합니다. fmt.Fprint 함수는 주어진 값을 그대로 파일에 저장합니다. 마찬가지로 fmt.Fprintln 함수도 값을 파일에 저장하지만 개행 문자(\n)을 넣어줍니다. 첫 번째 매개변수에 파일 인스턴스를 넣어주고, 두 번째 매개변수부터는 저장할 값 또는 변수를 나열합니다.

fmt.Fprintf 함수는 첫 번째 매개변수에 파일 인스턴스를 넣어주고 두 번째 매개변수에 출력할 형태(포맷)을 지정해줍니다. 세 번째 매개변수부터는 저장할 값 또는 변수를 나열합니다. 여기서 % (형식 지정자)로 값을 저장할 위치와 타입을 정하며 %의 개수와 변수 또는 값의 개수가 같아야 합니다.

파일에 대한 자세한 설명은 ‘Unit 49 파일 처리하기’를 참조하기 바랍니다.

반대로 파일에서 입력을 받아보겠습니다. 파일은 방금 만든 파일(hello1.txt, hello2.txt, hello3.txt)을 그대로 사용합니다.
```
package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")          // hello1.txt 파일 열기
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	n, _ := fmt.Fscan(file1, &num1, &num2, &s) // 파일을 읽은 뒤 공백, 개행 문자로
	                                           // 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)                // 입력 개수: 3
	fmt.Println(num1, num2, s)                 // 1 1.1 Hello

	file2, _ := os.Open("hello2.txt")    // hello2.txt 파일 열기
	defer file2.Close()                  // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanln(file2, &num1, &num2, &s) // 파일을 읽은 뒤 공백으로
	                                     // 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)          // 입력 개수: 3
	fmt.Println(num1, num2, s)           // 1 1.1 Hello

	file3, _ := os.Open("hello3.txt")               // hello3.txt 파일 열기
	defer file3.Close()                             // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanf(file3, "%d,%f,%s", &num1, &num2, &s) // 파일을 읽은 뒤 문자열에서
	                                                // 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)                     // 입력 개수: 3
	fmt.Println(num1, num2, s)                      // 1 1.1 Hello
}
```
fmt.Fscan 함수는 공백이나 개행 문자(\n)로 구분된 파일에서 입력을 받고, fmt.Fscanln 함수는 공백으로 구분된 파일에서 입력을 받습니다. 첫 번째 매개변수에 파일 인스턴스를 넣고, 두 번째 매개변수부터는 입력받을 변수를 레퍼런스 형식으로 넣습니다.

fmt.Fscanf 함수는 설정한 포맷대로 파일에서 입력을 받습니다. 첫 번째 매개변수에 파일 인스턴스를 넣고, 두 번째 매개변수에 포맷을 설정합니다. 그리고 세 번째 매개변수부터는 입력받을 변수를 레퍼런스 형식으로 넣습니다. 여기서는 포맷을 %d,%f,%s로 설정했으므로 , (콤마)로 구분된 파일의 내용을 변수에 저장합니다.
