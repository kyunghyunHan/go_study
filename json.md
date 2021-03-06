# ๐ฉ๐ปโ๐JSON

encoding/json ํจํค์ง์์ ์ ๊ณตํ๋ JSON ํจ์

- func Marshal(v interface{}) ([]byte, error): Go ์ธ์ด ์๋ฃํ์ JSON ํ์คํธ๋ก ๋ณํ
- func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error): Go ์ธ์ด ์๋ฃํ์ JSON ํ์คํธ๋ก ๋ณํํ๊ณ  ์ฌ๋์ด ๋ณด๊ธฐ ํธํ๋๋ก ๋ค์ฌ์ฐ๊ธฐ๋ฅผ ํด์ค
- func Unmarshal(data []byte, v interface{}) error: JSON ํ์คํธ๋ฅผ Go ์ธ์ด ์๋ฃํ์ผ๋ก ๋ณํ
๋จผ์  JSON ๋ฌธ์๋ฅผ ์ฝ๋ ๋ฐฉ๋ฒ์๋๋ค.

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

	var data map[string]interface{} // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ณต๊ฐ์ ๋งต์ผ๋ก ์ ์ธ

	json.Unmarshal([]byte(doc), &data) // doc๋ฅผ ๋ฐ์ดํธ ์ฌ๋ผ์ด์ค๋ก ๋ณํํ์ฌ ๋ฃ๊ณ ,
                                           // data์ ํฌ์ธํฐ๋ฅผ ๋ฃ์ด์ค

	fmt.Println(data["name"], data["age"]) // maria 10: ๋งต์ ํค๋ฅผ ์ง์ ํ์ฌ ๊ฐ์ ๊ฐ์ ธ์ด
}
```

- JSON ๋ฌธ์๋ฅผ Go ์ธ์ด์์ ์ฝ์ผ๋ ค๋ฉด encoding/json ํจํค์ง์ Unmarshal ํจ์๋ฅผ ์ฌ์ฉํฉ๋๋ค. ์ฌ๊ธฐ์ JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ณต๊ฐ์ด ํ์ํ๋ฐ ๊ฐ๋จํ๊ฒ ๋งต์ผ๋ก ๋ง๋ค ์ ์์ต๋๋ค.
```
var data map[string]interface{} // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ณต๊ฐ์ ๋งต์ผ๋ก ์ ์ธ
```
- JSON ๋ฌธ์์์ ํค๋ ๋ฌธ์์ด์ด๋ฉฐ ๊ฐ์ ๋ฌธ์์ด๊ณผ ์ซ์๋ฅผ ์ฌ์ฉํ๊ณ  ์์ต๋๋ค. ๋ฐ๋ผ์ ๋งต์ ๋ง๋ค ๋ ํค๋ ๋ฌธ์์ด๋ก ์ง์ ํ๊ณ  ๊ฐ์ interface{}๋ก ์ง์ ํ์ฌ ๋ชจ๋  ๊ฐ์ ๋ฃ์ ์ ์๋๋ก ํฉ๋๋ค.

- ์ด์  json.Unmarshal ํจ์์ JSON ๋ฌธ์์ด์ []byte ํ์์ผ๋ก ๋ณํํ์ฌ ๋ฃ๊ณ , ๋งต์ ํฌ์ธํฐ๋ ๋ฃ์ต๋๋ค.

```
json.Unmarshal([]byte(doc), &data) // doc๋ฅผ ๋ฐ์ดํธ ์ฌ๋ผ์ด์ค๋ก ๋ณํํ์ฌ ๋ฃ๊ณ , data์ ํฌ์ธํฐ๋ฅผ ๋ฃ์ด์ค
```
- ์ด์  ๋งต์์ data[โnameโ]์ฒ๋ผ ํค๋ฅผ ์ง์ ํ๋ฉด ๊ฐ์ ๊ฐ์ ธ์ฌ ์ ์์ต๋๋ค.

```
mt.Println(data["name"], data["age"]) // maria 10: ๋งต์ ํค๋ฅผ ์ง์ ํ์ฌ ๊ฐ์ ๊ฐ์ ธ์ด
```
- ์ด๋ฒ์๋ ๋งต ํํ์ ๋ฐ์ดํฐ๋ฅผ JSON ํํ๋ก ๋ณํํด๋ณด๊ฒ ์ต๋๋ค.
```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{}) // ๋ฌธ์์ด์ ํค๋กํ๊ณ  ๋ชจ๋  ์๋ฃํ์ ์ ์ฅํ  ์ ์๋ ๋งต ์์ฑ

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.Marshal(data) // ๋งต์ JSON ๋ฌธ์๋ก ๋ณํ

	fmt.Println(string(doc)) // {"age":10,"name":"maria"}: ๋ฌธ์์ด๋ก ๋ณํํ์ฌ ์ถ๋ ฅ
}
```
- ๋ฌธ์์ด์ ํค๋กํ๊ณ  ๋ชจ๋  ์๋ฃํ์ ๊ฐ์ผ๋ก ์ ์ฅํ  ์ ์๋ ๋งต์ ํ ๋นํฉ๋๋ค. ๊ทธ๋ฆฌ๊ณ  ์ ์ ํ ๋ฐ์ดํฐ๋ฅผ ๋์ํ ๋ค json.Marshal ํจ์๋ฅผ ์ฌ์ฉํ์ฌ JSON ๋ฌธ์๋ก ๋ณํํฉ๋๋ค.

- json.Marshal ํจ์์์ ๋ฆฌํด๋ ๊ฐ์ []byte ํ์์ด๋ฏ๋ก ์ถ๋ ฅํ  ๋๋ string ํ์์ผ๋ก ๋ณํํด์ค๋๋ค.

- JSON ๋ฌธ์๋ก ๋ณํํ์ ๋ ํ ์ค๋ก ๋ถ์ด์ ๋์ค๋ฉด ์ฌ๋์ด ์ฝ๊ธฐ ํ๋ญ๋๋ค. ๋ฐ๋ผ์ ๋ค์๊ณผ ๊ฐ์ด json.MarshalIndent ํจ์๋ฅผ ์ฌ์ฉํ๋ฉด ์ฌ๋์ด ์ฝ๊ฒ ์ฝ์ ์ ์๋๋ก ๋ณํํ  ์ ์์ต๋๋ค.

```
doc, _ := json.MarshalIndent(data, "", "  ")
```
- ์ฒซ ๋ฒ์งธ ๋งค๊ฐ๋ณ์๋ JSON ๋ฌธ์๋ก ๋ง๋ค ๋ฐ์ดํฐ์๋๋ค. ๊ทธ๋ฆฌ๊ณ  ๋ ๋ฒ์งธ ๋งค๊ฐ๋ณ์๋ JSON ๋ฌธ์์ ์ฒซ ์นธ์ ํ์ํ  ๋ฌธ์์ด(Prefix)์ธ๋ฐ ๋ณดํต "" (๋น ๋ฌธ์)๋ฅผ ์ง์ ํฉ๋๋ค. ๋ง์ง๋ง์ผ๋ก ์ธ ๋ฒ์งธ ๋งค๊ฐ๋ณ์๋ ๋ค์ฌ์ฐ๊ธฐํ  ๋ฌธ์์๋๋ค. ๊ณต๋ฐฑ ๋ฌธ์๋ ํญ ๋ฌธ์๋ฅผ ์ฌ์ฉํ  ์ ์์ต๋๋ค.

- Prefix๋ ํ์ํ์ง ์๊ณ , ๋ค์ฌ์ฐ๊ธฐ๋ 2์นธ์ผ๋ก ์ง์ ํ๋ฉด ๋ค์๊ณผ ๊ฐ์ด JSON ๋ฌธ์๋ก ๋ณํ๋ฉ๋๋ค.
```
{
  "age": 10,
  "name": "maria"
}
```

## ๐ฏ๊ตฌ์กฐ์ฒด ํ์ฉ

์์์๋ JSON ๋ฌธ์๊ฐ ํค-๊ฐ 1๋จ๊ณ์ ๋จ์ํ ๊ตฌ์กฐ์์ต๋๋ค. ์ด๋ฒ์๋ ๊ตฌ์กฐ์ฒด๋ฅผ ์ฌ์ฉํ์ฌ ์ข ๋ ๋ณต์กํ ํํ์ธ JSON ๋ฌธ์๋ฅผ ์ฒ๋ฆฌํด๋ณด๊ฒ ์ต๋๋ค(๋ฐฐ์ด, ํค-๋ฐฐ์ด, ํค-๊ฐ ์กฐํฉ).
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
	Author  Author // Author ๊ตฌ์กฐ์ฒด
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author ๊ตฌ์กฐ์ฒด
	Content    string
	Recommends []string  // ๋ฌธ์์ด ๋ฐฐ์ด
	Comments   []Comment // Comment ๊ตฌ์กฐ์ฒด ๋ฐฐ์ด
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

	var data []Article // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์ ์ธ

	json.Unmarshal([]byte(doc), &data) // doc์ ๋ด์ฉ์ ๋ณํํ์ฌ data์ ์ ์ฅ

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (์๋ต)
}
```
```
[{1 Hello, world! {Maria maria@example.com} Hello~ [John Andrew] [{1 {Andrew andrew@hello.com} Hello Maria}]}]
```
- ์ฝ์ด์ฌ JSON ๋ฌธ์์ ๊ตฌ์กฐ์ ๋ง์ถ์ด์ ๊ตฌ์กฐ์ฒด๋ฅผ ์์ฑํฉ๋๋ค.

```
type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author ๊ตฌ์กฐ์ฒด
	Content    string
	Recommends []string  // ๋ฌธ์์ด ๋ฐฐ์ด
	Comments   []Comment // Comment ๊ตฌ์กฐ์ฒด ๋ฐฐ์ด
}
```
- ๊ตฌ์กฐ์ฒด ์์ ํ๋๋ฅผ ์ง์ ํฉ๋๋ค. JSON์ ํ๋ ์์ ๋ค์ ๊ฐ์ฒด(Object)๋ ๋ฐฐ์ด์ด ๋ค์ด๊ฐ๋ ๋ฐฉ์์ด๋ฏ๋ก ์ด์ ๋ง์ถ์ด ๋ค๋ฅธ ๊ตฌ์กฐ์ฒด๋ฅผ ๋ฃ์ด์ค๋๋ค. ํ์ ๊ฐ์ฒด์ด๋ฉด ๊ตฌ์กฐ์ฒด๋ฅผ ๊ทธ๋๋ก ๋ฃ๊ณ  ๋ฐฐ์ด์ด๋ผ๋ฉด ๊ตฌ์กฐ์ฒด๋ฅผ ๋ฐฐ์ด๋ก ๋ง๋ญ๋๋ค.

- ์์ ๋ ๊ฒ์ํ๊ณผ ๋๊ธ์ ๊ตฌ์กฐ์ฒด๋ก ํํํ์ต๋๋ค. ๋ฐ๋ผ์ ๊ฒ์๋ฌผ์ด ์ฌ๋ฌ ๊ฐ์ธ ๋ฐฐ์ด ํํ์ผ ๊ฒ์ด๋ฏ๋ก ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๋ณ์๋ ์ฌ๋ผ์ด์ค ํํ๋ก ์ ์ธํฉ๋๋ค.

```
var data []Article // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์ ์ธ
```
- json.Unmarshal ํจ์๋ฅผ ์ฌ์ฉํ์ฌ ๊ฐ์ ๊ฐ์ ธ์ต๋๋ค.
```
json.Unmarshal([]byte(doc), &data) // doc์ ๋ด์ฉ์ ๋ณํํ์ฌ data์ ์ ์ฅ
```
- ์ด๋ฒ์๋ ๋ฐ์ดํฐ๋ฅผ JSON ํํ๋ก ๋ณํํด๋ณด๊ฒ ์ต๋๋ค.
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
	Author  Author // Author ๊ตฌ์กฐ์ฒด
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author ๊ตฌ์กฐ์ฒด
	Content    string
	Recommends []string  // ๋ฌธ์์ด ๋ฐฐ์ด
	Comments   []Comment // Comment ๊ตฌ์กฐ์ฒด ๋ฐฐ์ด
}

func main() {
	data := make([]Article, 1) // ๊ฐ์ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์์ฑ

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

	doc, _ := json.Marshal(data) // data๋ฅผ JSON ๋ฌธ์๋ก ๋ณํ

	fmt.Println(string(doc)) // [{"Id":1,"Title":"Hello, world!","Au... (์๋ต)
}
```
```
[{"Id":1,"Title":"Hello, world!","Author":{"Name":"Maria","Email":"maria@example.com"},"Content":"Hello~","Recommends":["John","Andrew"],"Comments":[{"Id":1,"Author":{"Name":"Andrew","Email":"andrew@hello.com"},"Content":"Hello Maria"}]}]
```
- ๊ฒ์๋ฌผ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ์ฌ๋ผ์ด์ค๋ฅผ ์ ์ธํ๊ณ  make ํจ์๋ก ๊ณต๊ฐ์ ํ ๋นํฉ๋๋ค. ๊ทธ๋ฆฌ๊ณ  ๋ฐ์ดํฐ๋ฅผ ์ฑ์๋ฃ์ ๋ค json.Marshal ํจ์๋ฅผ ์ฌ์ฉํ์ฌ JSON ๋ฌธ์๋ก ๋ณํํฉ๋๋ค.

```
data := make([]Article, 1) // ๊ฐ์ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์์ฑ

// ... ์๋ต ...

doc, _ := json.Marshal(data) // data๋ฅผ JSON ๋ฌธ์๋ก ๋ณํ
```

- ํจํค์ง์์ ํจ์(๋ณ์, ์์)๋ฅผ ์ธ๋ถ์ ๋ธ์ถํ  ๋ ์ด๋ฆ์ ์ฒซ ๊ธ์๋ฅผ ๋๋ฌธ์๋ก ๋ง๋๋ ๊ฒ์ฒ๋ผ ๊ตฌ์กฐ์ฒด๋ ๊ฐ์ ๊ท์น์ผ๋ก ๋์ํฉ๋๋ค. ์ฆ JSON ๋ฑ์ผ๋ก ์ธ๋ถ์ ๋ธ์ถํ  ๋๋ ๊ตฌ์กฐ์ฒด ํ๋์ ์ฒซ ๊ธ์๋ ๋๋ฌธ์๋ก ๋ง๋ญ๋๋ค. ๋ฐ๋ผ์ ๊ตฌ์กฐ์ฒด ํ๋์ ์ฒซ ๊ธ์๋ฅผ ์๋ฌธ์๋ก ๋ง๋ค๋ฉด JSON ๋ฌธ์์ ํด๋น ํ๋๋ ํฌํจ๋์ง ์์ต๋๋ค.

- ๊ตฌ์กฐ์ฒด ํ๋๊ฐ ๋๋ฌธ์๋ก ์์ํ๋ฉด JSON ๋ฌธ์ ์์ ํค๋ ๋๋ฌธ์๋ก ์์ํ๊ฒ ๋ฉ๋๋ค. ์ฌ๊ธฐ์ JSON ๋ฌธ์ ์์ ํค๋ฅผ ์๋ฌธ์๋ก ์์ํ๊ณ  ์ถ๋ค๋ฉด ๋ค์๊ณผ ๊ฐ์ด ๊ตฌ์กฐ์ฒด ํ๋์ ํ๊ทธ๋ฅผ ์ง์ ํด์ค๋๋ค.

- ํ๋๋ช ์๋ฃํ `json:โํคโ`

```
type ๊ตฌ์กฐ์ฒด๋ช struct {
	ํ๋๋ช ์๋ฃํ `json:"ํค"`
}
```
```
type Author struct {
	Name  string `json:"name"` // ๊ตฌ์กฐ์ฒด ํ๋์ ํ๊ทธ ์ง์ 
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
- ํ๊ทธ๋ ๋ฌธ์์ด ํํ์ด๋ฉฐ ๋ฌธ์์ด ์์ " " (๋ฐ์ดํ)๊ฐ ํฌํจ๋๋ฏ๋ก ` ` (๋ฐฑ์ฟผํธ)๋ก ๊ฐ์ธ์ค๋๋ค. ๊ทธ๋ฆฌ๊ณ  JSON ๋ฌธ์์ด๋ฏ๋ก `json:"ํค"` ํ์์ผ๋ก ํค ์ด๋ฆ์ ์ง์  ์ง์ ํฉ๋๋ค. ์ฌ๊ธฐ์ ํค ์ด๋ฆ์ ๊ตฌ์กฐ์ฒด ํ๋์ ๊ฐ์ ํ์๋ ์์ต๋๋ค.

## ๐ฏJSON ํ์ผ์ฌ์ฉ

- JSON ํ์ผ์ ๋ง๋ค๊ณ , ์ฝ์ด๋ณด๊ฒ ์ต๋๋ค.

- ๋จผ์  ๋ฐ์ดํฐ๋ฅผ JSON ํ์ผ๋ก ์ ์ฅํฉ๋๋ค.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` // ๊ตฌ์กฐ์ฒด ํ๋์ ํ๊ทธ ์ง์ 
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
	data := make([]Article, 1) // ๊ฐ์ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์์ฑ

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

	doc, _ := json.Marshal(data) // data๋ฅผ JSON ๋ฌธ์๋ก ๋ณํ

	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json ํ์ผ์ JSON ๋ฌธ์ ์ ์ฅ
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
- ๊ฒ์๋ฌผ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ์ฌ๋ผ์ด์ค๋ฅผ ์ ์ธํ๊ณ  make ํจ์๋ก ๊ณต๊ฐ์ ํ ๋นํฉ๋๋ค. ๊ทธ๋ฆฌ๊ณ  ๋ฐ์ดํฐ๋ฅผ ์ฑ์๋ฃ์ ๋ค json.Marshal ํจ์๋ฅผ ์ฌ์ฉํ์ฌ JSON ๋ฌธ์๋ก ๋ณํํฉ๋๋ค.

```
data := make([]Article, 1) // ๊ฐ์ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์์ฑ

// ... ์๋ต ...

doc, _ := json.Marshal(data) // data๋ฅผ JSON ๋ฌธ์๋ก ๋ณํ
```
- ioutil.WriteFile ํจ์๋ฅผ ์ฌ์ฉํ์ฌ ๋ณํ๋ JSON ๋ฌธ์๋ฅผ ํ์ผ๋ก ์ ์ฅํฉ๋๋ค.

```
err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json ํ์ผ์ JSON ๋ฌธ์ ์ ์ฅ
```

- ์ธ ๋ฒ์งธ ๋งค๊ฐ๋ณ์์๋ ์ ๋์ค/๋ฆฌ๋์ค ํ์์ ํ์ผ์ ๊ถํ(Permission)์ ์ง์ ํด์ผ ํฉ๋๋ค. ์ฌ๊ธฐ์๋ os.FileMode ํ์์ผ๋ก 0644๋ฅผ ์ง์ ํ์์ต๋๋ค.

- ์์์ JSON ๋ฌธ์๋ฅผ ์ ์ฅํ articles.json ํ์ผ์ ์ฝ์ด๋ณด๊ฒ ์ต๋๋ค.

```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Author struct {
	Name  string `json:"name"` // ๊ฐ์ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์์ฑ
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
	b, err := ioutil.ReadFile("./articles.json") // articles.json ํ์ผ์ ๋ด์ฉ์ ์ฝ์ด์ ๋ฐ์ดํธ ์ฌ๋ผ์ด์ค์ ์ ์ฅ
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์ ์ธ

	json.Unmarshal(b, &data) // JSON ๋ฌธ์์ ๋ด์ฉ์ ๋ณํํ์ฌ data์ ์ ์ฅ

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (์๋ต)
}
```
```
[{1 Hello, world! {Maria maria@example.com} Hello~ [John Andrew] [{1 {Andrew andrew@hello.com} Hello Maria}]}]
```
- ioutil.ReadFile ํจ์๋ก JSON ํ์ผ์ ์ฝ์ด์ต๋๋ค.
```
b, err := ioutil.ReadFile("./articles.json") // articles.json ํ์ผ์ ๋ด์ฉ์ ์ฝ์ด์
                                             // ๋ฐ์ดํธ ์ฌ๋ผ์ด์ค์ ์ ์ฅ
```
- json.Unmarshal ํจ์๋ฅผ ์ฌ์ฉํ์ฌ ๋ฐ์ดํธ ์ฌ๋ผ์ด์ค b์ ์ ์ฅ๋ ๊ฐ์ ๊ฐ์ ธ์ค๋ฉด ๋ฉ๋๋ค.

```
var data []Article // JSON ๋ฌธ์์ ๋ฐ์ดํฐ๋ฅผ ์ ์ฅํ  ๊ตฌ์กฐ์ฒด ์ฌ๋ผ์ด์ค ์ ์ธ

json.Unmarshal(b, &data) // JSON ๋ฌธ์์ ๋ด์ฉ์ ๋ณํํ์ฌ data์ ์ ์ฅ
```
