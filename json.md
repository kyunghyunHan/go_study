# 👩🏻‍🎓JSON

encoding/json 패키지에서 제공하는 JSON 함수

- func Marshal(v interface{}) ([]byte, error): Go 언어 자료형을 JSON 텍스트로 변환
- func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error): Go 언어 자료형을 JSON 텍스트로 변환하고 사람이 보기 편하도록 들여쓰기를 해줌
- func Unmarshal(data []byte, v interface{}) error: JSON 텍스트를 Go 언어 자료형으로 변환
먼저 JSON 문서를 읽는 방법입니다.

```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`

	var data map[string]interface{} // JSON 문서의 데이터를 저장할 공간을 맵으로 선언

	json.Unmarshal([]byte(doc), &data) // doc를 바이트 슬라이스로 변환하여 넣고,
                                           // data의 포인터를 넣어줌

	fmt.Println(data["name"], data["age"]) // maria 10: 맵에 키를 지정하여 값을 가져옴
}
```

- JSON 문서를 Go 언어에서 읽으려면 encoding/json 패키지의 Unmarshal 함수를 사용합니다. 여기서 JSON 문서의 데이터를 저장할 공간이 필요한데 간단하게 맵으로 만들 수 있습니다.
```
var data map[string]interface{} // JSON 문서의 데이터를 저장할 공간을 맵으로 선언
```
- JSON 문서에서 키는 문자열이며 값은 문자열과 숫자를 사용하고 있습니다. 따라서 맵을 만들 때 키는 문자열로 지정하고 값은 interface{}로 지정하여 모든 값을 넣을 수 있도록 합니다.

- 이제 json.Unmarshal 함수에 JSON 문자열을 []byte 형식으로 변환하여 넣고, 맵의 포인터도 넣습니다.

```
json.Unmarshal([]byte(doc), &data) // doc를 바이트 슬라이스로 변환하여 넣고, data의 포인터를 넣어줌
```
- 이제 맵에서 data[“name”]처럼 키를 지정하면 값을 가져올 수 있습니다.

```
mt.Println(data["name"], data["age"]) // maria 10: 맵에 키를 지정하여 값을 가져옴
```
- 이번에는 맵 형태의 데이터를 JSON 형태로 변환해보겠습니다.
```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.Marshal(data) // 맵을 JSON 문서로 변환

	fmt.Println(string(doc)) // {"age":10,"name":"maria"}: 문자열로 변환하여 출력
}
```
- 문자열을 키로하고 모든 자료형을 값으로 저장할 수 있는 맵을 할당합니다. 그리고 적절히 데이터를 대입한 뒤 json.Marshal 함수를 사용하여 JSON 문서로 변환합니다.

- json.Marshal 함수에서 리턴된 값은 []byte 형식이므로 출력할 때는 string 형식으로 변환해줍니다.

- JSON 문서로 변환했을 때 한 줄로 붙어서 나오면 사람이 읽기 힘듭니다. 따라서 다음과 같이 json.MarshalIndent 함수를 사용하면 사람이 쉽게 읽을 수 있도록 변환할 수 있습니다.

```
doc, _ := json.MarshalIndent(data, "", "  ")
```
- 첫 번째 매개변수는 JSON 문서로 만들 데이터입니다. 그리고 두 번째 매개변수는 JSON 문서의 첫 칸에 표시할 문자열(Prefix)인데 보통 "" (빈 문자)를 지정합니다. 마지막으로 세 번째 매개변수는 들여쓰기할 문자입니다. 공백 문자나 탭 문자를 사용할 수 있습니다.

- Prefix는 표시하지 않고, 들여쓰기는 2칸으로 지정하면 다음과 같이 JSON 문서로 변환됩니다.
```
{
  "age": 10,
  "name": "maria"
}
```

## 💯구조체 활용

앞에서는 JSON 문서가 키-값 1단계의 단순한 구조였습니다. 이번에는 구조체를 사용하여 좀 더 복잡한 형태인 JSON 문서를 처리해보겠습니다(배열, 키-배열, 키-값 조합).
```
package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author // Author 구조체
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}

func main() {
	doc := `
	[{
		"Id": 1,
		"Title": "Hello, world!",
		"Author": {
			"Name": "Maria",
			"Email": "maria@example.com"
		},
		"Content": "Hello~",
		"Recommends": [
			"John",
			"Andrew"
		],
		"Comments": [{
			"id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "andrew@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal([]byte(doc), &data) // doc의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
}
```
```
[{1 Hello, world! {Maria maria@example.com} Hello~ [John Andrew] [{1 {Andrew andrew@hello.com} Hello Maria}]}]
```
- 읽어올 JSON 문서의 구조에 맞추어서 구조체를 작성합니다.

```
type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}
```
- 구조체 안에 필드를 지정합니다. JSON은 필드 안에 다시 객체(Object)나 배열이 들어가는 방식이므로 이에 맞추어 다른 구조체를 넣어줍니다. 하위 객체이면 구조체를 그대로 넣고 배열이라면 구조체를 배열로 만듭니다.

- 예제는 게시판과 댓글을 구조체로 표현했습니다. 따라서 게시물이 여러 개인 배열 형태일 것이므로 데이터를 저장할 변수도 슬라이스 형태로 선언합니다.

```
var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언
```
- json.Unmarshal 함수를 사용하여 값을 가져옵니다.
```
json.Unmarshal([]byte(doc), &data) // doc의 내용을 변환하여 data에 저장
```
- 이번에는 데이터를 JSON 형태로 변환해보겠습니다.
```
package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author // Author 구조체
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	fmt.Println(string(doc)) // [{"Id":1,"Title":"Hello, world!","Au... (생략)
}
```
```
[{"Id":1,"Title":"Hello, world!","Author":{"Name":"Maria","Email":"maria@example.com"},"Content":"Hello~","Recommends":["John","Andrew"],"Comments":[{"Id":1,"Author":{"Name":"Andrew","Email":"andrew@hello.com"},"Content":"Hello Maria"}]}]
```
- 게시물 데이터를 저장할 슬라이스를 선언하고 make 함수로 공간을 할당합니다. 그리고 데이터를 채워넣은 뒤 json.Marshal 함수를 사용하여 JSON 문서로 변환합니다.

```
data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

// ... 생략 ...

doc, _ := json.Marshal(data) // data를 JSON 문서로 변환
```

- 패키지에서 함수(변수, 상수)를 외부에 노출할 때 이름의 첫 글자를 대문자로 만드는 것처럼 구조체도 같은 규칙으로 동작합니다. 즉 JSON 등으로 외부에 노출할 때는 구조체 필드의 첫 글자는 대문자로 만듭니다. 따라서 구조체 필드의 첫 글자를 소문자로 만들면 JSON 문서에 해당 필드는 포함되지 않습니다.

- 구조체 필드가 대문자로 시작하면 JSON 문서 안의 키도 대문자로 시작하게 됩니다. 여기서 JSON 문서 안의 키를 소문자로 시작하고 싶다면 다음과 같이 구조체 필드에 태그를 지정해줍니다.

- 필드명 자료형 `json:”키”`

```
type 구조체명 struct {
	필드명 자료형 `json:"키"`
}
```
```
type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}
```
- 태그는 문자열 형태이며 문자열 안에 " " (따옴표)가 포함되므로 ` ` (백쿼트)로 감싸줍니다. 그리고 JSON 문서이므로 `json:"키"` 형식으로 키 이름을 직접 지정합니다. 여기서 키 이름은 구조체 필드와 같을 필요는 없습니다.

## 💯JSON 파일사용

- JSON 파일을 만들고, 읽어보겠습니다.

- 먼저 데이터를 JSON 파일로 저장합니다.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
- 게시물 데이터를 저장할 슬라이스를 선언하고 make 함수로 공간을 할당합니다. 그리고 데이터를 채워넣은 뒤 json.Marshal 함수를 사용하여 JSON 문서로 변환합니다.

```
data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

// ... 생략 ...

doc, _ := json.Marshal(data) // data를 JSON 문서로 변환
```
- ioutil.WriteFile 함수를 사용하여 변환된 JSON 문서를 파일로 저장합니다.

```
err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
```

- 세 번째 매개변수에는 유닉스/리눅스 형식의 파일의 권한(Permission)을 지정해야 합니다. 여기서는 os.FileMode 타입으로 0644를 지정하였습니다.

- 앞에서 JSON 문서를 저장한 articles.json 파일을 읽어보겠습니다.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Author struct {
	Name  string `json:"name"` // 값을 저장할 구조체 슬라이스 생성
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
}
```
```
[{1 Hello, world! {Maria maria@example.com} Hello~ [John Andrew] [{1 {Andrew andrew@hello.com} Hello Maria}]}]
```
- ioutil.ReadFile 함수로 JSON 파일을 읽어옵니다.
```
b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서
                                             // 바이트 슬라이스에 저장
```
- json.Unmarshal 함수를 사용하여 바이트 슬라이스 b에 저장된 값을 가져오면 됩니다.

```
var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장
```
