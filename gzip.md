# π©π»βπμμΆ

Go μΈμ΄λ λ€μν μμΆ μκ³ λ¦¬μ¦μ ν¨ν€μ§λ‘ μ κ³΅ν΄μ€λλ€. λ³΄ν΅ μμΆ μκ³ λ¦¬μ¦μ νμΌ, λ€νΈμν¬ ν¨ν·μ μμΆνμ¬ μ©λμ μ€μ΄κ³ , μμΆλ νμΌ, λ€νΈμν¬ ν¨ν·μ ν΄μ ν  λ μ¬μ©ν©λλ€.

λ€μμ compress/gzip ν¨ν€μ§μμ μ κ³΅νλ μμΆ ν¨μμλλ€.

- func NewReader(r io.Reader) (*Reader, error): io.Reader μΈν°νμ΄μ€λ‘ io.Reader μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ μμΆ ν΄μ  μΈμ€ν΄μ€ μμ±
- func NewWriter(w io.Writer) *Writer: io.Writer μΈν°νμ΄μ€λ‘ io.Writer μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ μμΆ μΈμ€ν΄μ€ μμ±
λ€μμ io/ioutil ν¨ν€μ§μμ μ κ³΅νλ λ°μ΄ν° μ½κΈ° ν¨μμλλ€.

- func ReadAll(r io.Reader) ([]byte, error): io.Readerλ₯Ό λ(EOF)κΉμ§ μ½μ΄μ λ°μ΄νΈ μ¬λΌμ΄μ€λ‘ λ¦¬ν΄
gzip μκ³ λ¦¬μ¦μ μ¬μ©ν΄μ λ°μ΄ν°λ₯Ό μμΆν λ€ νμΌλ‘ μ μ₯ν΄λ³΄κ² μ΅λλ€

```
package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile( // μμΆν  νμΌ μμ±
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main ν¨μκ° λλκΈ° μ§μ μ νμΌμ λ«μ

	w := gzip.NewWriter(file) // fileλ‘ io.Writer μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ μμΆ μΈμ€ν΄μ€ w μμ±
	defer w.Close()           // main ν¨μκ° λλκΈ° μ§μ μ μμΆ μΈμ€ν΄μ€λ₯Ό λ«μ

	s := "Hello, world!"
	w.Write([]byte(s)) // μμΆ μΈμ€ν΄μ€λ‘ λ¬Έμμ΄μ μμΆνμ¬ νμΌμ μ μ₯
	w.Flush()          // λ©λͺ¨λ¦¬μμ λ°μ΄ν°λ₯Ό νμΌμ μμ ν μ μ₯
}
```
gzip.NewWriter ν¨μμ file μΈμ€ν΄μ€λ₯Ό λ£μΌλ©΄ io.Writer μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ μμΆ μΈμ€ν΄μ€λ₯Ό λ¦¬ν΄ν©λλ€. κ·Έλ¦¬κ³  μμΆ μΈμ€ν΄μ€μ Write λ©μλμ []byte νμμ λ°μ΄ν°λ₯Ό λ£μ λ€ Flush λ©μλλ₯Ό νΈμΆνλ©΄ μλμΌλ‘ μμΆλ©λλ€. μ¬κΈ°μλ λ¬Έμμ΄μ []byte νμμΌλ‘ λ³νν΄μ λ£μμ΅λλ€.
Close λ©μλλ‘ μμΆ μΈμ€ν΄μ€λ₯Ό λ°λμ λ«μμ€λλ€. κ·Έλ μ§ μμΌλ©΄ μμΆμ΄ μ μμ μΌλ‘ λμ§ μμ΅λλ€.

μ΄λ²μλ λ°©κΈ μμΆν hello.txt.gzμ μμΆμ ν΄μ ν΄μ μ½μ΄λ³΄κ² μ΅λλ€.
```
package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("hello.txt.gz") // μμΆ νμΌ μ΄κΈ°
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main ν¨μκ° λλκΈ° μ§μ μ νμΌμ λ«μ

	r, err := gzip.NewReader(file) // fileλ‘ io.Reader μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ
                                       // μμΆ ν΄μ  μΈμ€ν΄μ€ r μμ±
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close() // main ν¨μκ° λλκΈ° μ§μ μ μμΆ μΈμ€ν΄μ€λ₯Ό λ«μ

	b, err := ioutil.ReadAll(r) // μμΆ ν΄μ  μΈμ€ν΄μ€ rμ μ½μΌλ©΄
                                    // μμΆ ν΄μ λ λ°μ΄ν°κ° λ¦¬ν΄λ¨
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b)) // λ¬Έμμ΄λ‘ λ³ννμ¬ μΆλ ₯
}
```


gzip.NewReader ν¨μμ file μΈμ€ν΄μ€λ₯Ό λ£μΌλ©΄ io.Reader μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ μμΆ ν΄μ  μΈμ€ν΄μ€λ₯Ό λ¦¬ν΄ν©λλ€. μΌλ° νμΌ κ°μΌλ©΄ νμΌ ν¬κΈ°λ₯Ό κ΅¬νμ¬ μ¬λΌμ΄μ€λ₯Ό μμ±νκ² μ§λ§ μμΆ νμΌμ ν΄μ ν λ€μ ν¬κΈ° μ λ³΄λ₯Ό κ΅¬ν  μ μμ΅λλ€. μ΄λλ ioutil.ReadAll ν¨μλ₯Ό μ¬μ©νλ©΄ νΈλ¦¬ν©λλ€. ioutil.ReadAll ν¨μμ μμΆ ν΄μ  μΈμ€ν΄μ€(io.Reader)λ₯Ό λ£μΌλ©΄ μλμΌλ‘ μμΆμ ν΄μ νμ¬ μ¬λΌμ΄μ€λ₯Ό λ¦¬ν΄ν©λλ€.

λ€λ₯Έ μμΆ μκ³ λ¦¬μ¦λ€λ io.Reader, io.Writer μΈν°νμ΄μ€λ₯Ό λ°λ₯΄λ―λ‘ κ°μ λ°©λ²μΌλ‘ μμΆ λ° ν΄μ λ₯Ό ν  μ μμ΅λλ€.

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
