# ๐ฉ๐ปโ๐์ง์ฐํธ์ถํจ์

- defer ํจ์๋ช()
- defer ํจ์๋ช(๋งค๊ฐ๋ณ์)

```
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
        //โ ํจ์์ ํธ์ถ์ด ์ง์ฐ๋จ.
	defer world() // ํ์ฌ ํจ์(main())๊ฐ ๋๋๊ธฐ ์ง์ ์ ํธ์ถ
	hello()
	hello()
	hello()
}
```
```
Hello
Hello
Hello
world
```
- ์ฌ๊ธฐ์ hello ํจ์๋ณด๋ค world ํจ์๋ฅผ ๋จผ์  ํธ์ถํ๋๋ผ๋ ์ถ๋ ฅ์ ํด๋ณด๋ฉด world ํจ์๊ฐ ๋์ค์ ํธ์ถ๋๋. ์ฆ defer๋ฅผ ์ฌ์ฉํ ํจ์๋ ํ์ฌ ํจ์(main())๊ฐ ๋๋๊ธฐ ์ง์ ์ ํธ์ถ๋๋ฏ๋ก ๋ค๋ฅธ ํจ์๋ค ๋ณด๋ค ๋งจ ๋ง์ง๋ง์ ํธ์ถ๋๋ค.

```
package main

import (
	"fmt"
	"os"
)

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() // ์ง์ฐ ํธ์ถํ file.Close()๊ฐ ๋งจ ๋ง์ง๋ง์ ํธ์ถ๋จ

	if err != nil {
		fmt.Println(err)
		return // file.Close() ํธ์ถ
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() ํธ์ถ
	}

	fmt.Println(string(buf))

	// file.Close() ํธ์ถ
}

func main() {
	ReadHello()
}
```
os.Open์ผ๋ก ํ์ผ์ ์ฐ ๋ค ์ง์ฐ ํธ์ถ๋ก file.Close๋ฅผ ํธ์ถํ๋ฉด ํจ์๊ฐ ๋๋  ๋ ๋ฌด์กฐ๊ฑด file.Close๋ก ํ์ผ์ ๋ซ๋๋ค. ํ์ผ ์ด๊ธฐ๋ ์ฑ๊ณตํ์ง๋ง ์ค๊ฐ์ file.Read์์ ์คํจํ๋๋ผ๋ ์ง์ฐ ํธ์ถ์ ์ฌ์ฉํ์ผ๋ฏ๋ก ํจ์๊ฐ ์ข๋ฃ๋  ๋ file.Close๋ก ํ์ผ์ ๋ซ๋๋ค ํนํ ์ด๋ฐ ๋ฐฉ์์ ํ๋ก๊ทธ๋จ ํ๋ฆ์ ๋ถ๊ธฐ๊ฐ ๋ง์ ์๋ฌ ์ฒ๋ฆฌ๊ฐ ๋ณต์กํด์ง๋ ์ ์ฉํ๋ค
