# ๐ฏinterface

- ๋ฉ์๋์ ์งํฉ์ฒด 
```
type hello interface { // ์ธํฐํ์ด์ค ์ ์
}

func main() {
	var h hello    // ์ธํฐํ์ด์ค ์ ์ธ
	fmt.Println(h) // <nil>: ๋น ์ธํฐํ์ด์ค์ด๋ฏ๋ก nil์ด ์ถ๋ ฅ๋จ
}
```
- ์ธํฐํ์ด์ค๋ฅผ ์ ์ธํ๋ ๋ฐฉ๋ฒ์  var ๋ณ์๋ช ์ธํฐํ์ด์ค ํ์์ผ๋ก ์ ์ธํ๋ค ๊ทธ๋ฆฌ๊ณ  ์ธํฐํ์ด์ค๋ ๋ค๋ฅธ ์๋ฃํ๊ณผ ๋์ผํ๊ฒ ํจ์์ ๋งค๊ฐ๋ณ์, ๋ฆฌํด๊ฐ, ๊ตฌ์กฐ์ฒด์ ํ๋๋ก ์ฌ์ฉํ  ์ ์๋ค.

- ์ฌ๊ธฐ์๋ ๋น ์ธํฐํ์ด์ค๋ฅผ ์ ์ํ๊ณ  ์ ์ธ! , ๋น ์ธํฐํ์ด์ค๊ธฐ ๋๋ฌธ์ ํ  ์ ์๋ ๊ฒ์ด ์์ผ๋ฉฐ fmt.Println์ผ๋ก ์ถ๋ ฅํด๋ด๋ nil์ด ๋์จ๋ค

- ๋ฉ์๋๋ฅผ ๊ฐ์ง๋ ์ธํฐํ์ด์ค ์ ์

- type ์ธํฐํ์ด์ค๋ช interface { ๋ฉ์๋ }

```
type ์ธํฐํ์ด์ค๋ช interface {
	๋ฉ์๋1() ๋ฆฌํด๊ฐ_์๋ฃํ
	๋ฉ์๋2()                  // ๋ฆฌํด๊ฐ์ด ์์ ๋
}
```

- { } (์ค๊ดํธ) ๋ธ๋ก์์ ๋ฉ์๋ ์ด๋ฆ, ๋งค๊ฐ๋ณ์ ์๋ฃํ, ๋ฆฌํด๊ฐ ์๋ฃํ์ ์ง์ ํ์ฌ ํ ์ค ์ฉ ๋์ดํ๋ค. ์ฌ๊ธฐ์๋ , (์ฝค๋ง)๋ก ๊ตฌ๋ถํ์ง ์๋๋ค

- ๋ค์์ int ์๋ฃํ์ ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐํ๊ณ , ์ธํฐํ์ด์ค๋ก ํด๋น ๋ฉ์๋๋ฅผ ํธ์ถํ๋ค
```
package main

import "fmt"

type MyInt int // intํ์ MyInt๋ก ์ ์

func (i MyInt) Print() { // MyInt์ Print ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐ
	fmt.Println(i)
}

type Printer interface { // Print ๋ฉ์๋๋ฅผ ๊ฐ์ง๋ ์ธํฐํ์ด์ค ์ ์
	Print()
}

func main() {
	var i MyInt = 5

	var p Printer // ์ธํฐํ์ด์ค ์ ์ธ

	p = i     // i๋ฅผ ์ธํฐํ์ด์ค p์ ๋์
	p.Print() // 5: ์ธํฐํ์ด์ค๋ฅผ ํตํ์ฌ MyInt์ Print ๋ฉ์๋ ํธ์ถ
}
```
- ๋จผ์  type ์_์๋ฃํ ์๋ฃํ ํ์์ผ๋ก ๊ธฐ์กด ์๋ฃํ์ ์ ์๋ฃํ์ผ๋ก ์ ์ํ  ์ ์๋ค. ์ฌ๊ธฐ์๋ ๊ธฐ๋ณธ ์๋ฃํ์ธ int์ ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐํ๊ธฐ ์ํด MyInt๋ฅผ ์๋ก ์ ์ํ๋ค. ๊ทธ ๋ค์ MyInt์ Print ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐํ๊ณ , ํ์ฌ ๋ณ์์ ๊ฐ์ ์ถ๋ ฅํ๋๋ก ๊ตฌํํ์ฟ๋ค.
- ์ธํฐํ์ด์ค ์ ์(๋ณดํต ์ธํฐํ์ด์ค์ ์ด๋ฆ์ ~er ํํ(์: Reader, Writer))

```
type Printer interface {
	Print()
}
```
- main ํจ์ ์์์ ์ธํฐํ์ด์ค๋ฅผ ์ ์ธํ ๋ค ๋ณ์๋ฅผ ๋์,  ์ธํฐํ์ด์ค์ . (์ )์ ์ฌ์ฉํ์ฌ ๋ฉ์๋๋ฅผ ํธ์ถํ๋ค.์ฌ๊ธฐ์๋ p.Print()์ฒ๋ผ ์ธํฐํ์ด์ค์ Print ๋ฉ์๋๋ฅผ ํธ์ถํ์ง๋ง ์ค์ ๋ก๋ MyInt์ Print ๋ฉ์๋๊ฐ ํธ์ถ๋๋ค. ์ฆ ์ธํฐํ์ด์ค์ ๋ด๊ธด ์ค์  ํ์(์๋ฃํ, ๊ตฌ์กฐ์ฒด)์ ๋ฉ์๋๊ฐ ํธ์ถ๋๋ค.
- ๋ค์์ int ์๋ฃํ๊ณผ ์ฌ๊ฐํ ๊ตฌ์กฐ์ฒด์ ๋ด์ฉ์ ์ถ๋ ฅํ๊ณ , int ์๋ฃํ๊ณผ ์ฌ๊ฐํ ๊ตฌ์กฐ์ฒด์ ์ธ์คํด์ค๋ฅผ ๋ด์ ์ ์๋ ์ธํฐํ์ด์ค๋ฅผ ์ ์ํ ์์ 
```
package main

import "fmt"

type MyInt int // int ํ์ MyInt๋ก ์ ์

func (i MyInt) Print() { // MyInt์ Print ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐ
	fmt.Println(i)
}

type Rectangle struct { // ์ฌ๊ฐํ ๊ตฌ์กฐ์ฒด ์ ์
	width, height int
}

func (r Rectangle) Print() { // Rectangle์ Print ๋ฉ์๋๋ฅผ ์ฐ๊ฒฐ
	fmt.Println(r.width, r.height)
}

type Printer interface { // Print ๋ฉ์๋๋ฅผ ๊ฐ์ง๋ ์ธํฐํ์ด์ค ์ ์
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer // ์ธํฐํ์ด์ค ์ ์ธ

	p = i     // i๋ฅผ ์ธํฐํ์ด์ค p์ ๋์
	p.Print() // 5: ์ธํฐํ์ด์ค p๋ฅผ ํตํ์ฌ MyInt์ Print ๋ฉ์๋ ํธ์ถ

	p = r     // r์ ์ธํฐํ์ด์ค p์ ๋์
	p.Print() // 10 20: ์ธํฐํ์ด์ค p๋ฅผ ํตํ์ฌ Rectangle์ Print ๋ฉ์๋ ํธ์ถ
}
```
- MyInt๋ฅผ ์ ์ํ๊ณ , ๋๋น์ ๋์ด๋ฅผ ๊ฐ์ง๋ ์ฌ๊ฐํ ๊ตฌ์กฐ์ฒด Rectangle์ ์ ์!!, ๊ทธ๋ฆฌ๊ณ  MyInt, Rectangle ๋ชจ๋ ์์ ์ ๋ด์ฉ์ ์ถ๋ ฅํ๋ Print ํจ์๋ฅผ ๊ตฌํ!! ์ด์  ๋ ํ์ ๋ชจ๋ ๋๊ฐ์ Print ํจ์๋ฅผ ๊ฐ์ง๊ณ  ์๋ค(์ฌ๊ธฐ์ ๋๊ฐ์ ํจ์๋ผ๋ ๊ฒ์ ํจ์ ์ด๋ฆ, ๋งค๊ฐ๋ณ์ ์๋ฃํ, ๋ฆฌํด๊ฐ ์๋ฃํ์ด ๋ชจ๋ ๊ฐ์ ์ํ๋ฅผ ๋ปํจ).

- MyInt ์๋ฃํ, Ractangle ๊ตฌ์กฐ์ฒด, Printer ์ธํฐํ์ด์ค๋ฅผ ๊ทธ๋ฆผ์ผ๋ก ํํํ๋ฉด ๋ค์๊ณผ ๊ฐ๋ค
```
func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer // ์ธํฐํ์ด์ค ์ ์ธ

	p = i     // i๋ฅผ ์ธํฐํ์ด์ค p์ ๋์
	p.Print() // 5: ์ธํฐํ์ด์ค p๋ฅผ ํตํ์ฌ MyInt์ Print ๋ฉ์๋ ํธ์ถ

	p = r     // r์ ์ธํฐํ์ด์ค p์ ๋์
	p.Print() // 10 20: ์ธํฐํ์ด์ค p๋ฅผ ํตํ์ฌ Rectangle์ Print ๋ฉ์๋ ํธ์ถ
}
```
- Printer ์ธํฐํ์ด์ค p์ MyInt ํ ๋ณ์ i๋ฅผ ๋์!!, ๋ง์ฐฌ๊ฐ์ง๋ก ๋ค์ p์ Rectangle ์ธ์คํด์ค๋ฅผ ๋์ํ์๋ค. ์ด๋ ๊ฒ ์ ํ ๋ค๋ฅธ ํ์์ ์ธํฐํ์ด์ค์ ๋์ํ  ์ ์๋ค. ์ฆ ์ธํฐํ์ด์ค๋ ์๋ฃํ์ด๋  ๊ตฌ์กฐ์ฒด๋  ํ์์ ์๊ด์์ด ๋ฉ์๋ ์งํฉ๋ง ๊ฐ์ผ๋ฉด ๋์ผํ ํ์์ผ๋ก ๋ณธ๋ค
- ์ธํฐํ์ด์ค์ ๋ฉ์๋๋ฅผ ํธ์ถํ๋ฉด ์ธํฐํ์ด์ค์ ๋ด๊ธด ์ค์  ํ์์ ๋ฉ์๋๊ฐ ํธ์ถ๋๋ฏ๋ก MyInt ํ ๋ณ์๋ฅผ ๋์ํ ๋ค Print ํจ์๋ฅผ ํธ์ถํ๋ฉด 5๊ฐ ์ถ๋ ฅ๋๊ณ , Rectangle ์ธ์คํด์ค๋ฅผ ๋์ํ๋ฉด 10๊ณผ 20์ด ์ถ๋ ฅ๋๋ค.

- ์ธํฐํ์ด์ค๋ฅผ ์ ์ธํ๋ฉด์ ์ด๊ธฐํํ๋ ค๋ฉด ๋ค์๊ณผ ๊ฐ์ด :=๋ฅผ ์ฌ์ฉํ๋ฉด ๋๋ค. ์ธํฐํ์ด์ค์๋ ( ) (๊ดํธ)๋ฅผ ์ฌ์ฉํ์ฌ ๋ณ์๋ ์ธ์คํด์ค๋ฅผ ๋ฃ์ด์ค๋ค

```
var i MyInt = 5
r := Rectangle{10, 20}

p1 := Printer(i) // ์ธํฐํ์ด์ค๋ฅผ ์ ์ธํ๋ฉด์ i๋ก ์ด๊ธฐํ
p1.Print()       // 5

p2 := Printer(r) // ์ธํฐํ์ด์ค๋ฅผ ์ ์ธํ๋ฉด์ r๋ก ์ด๊ธฐํ
p2.Print()       // 10 20
```
- ๋ค์๊ณผ ๊ฐ์ด ๋ฐฐ์ด(์ฌ๋ผ์ด์ค) ํํ๋ก๋ ์ธํฐํ์ด์ค๋ฅผ ์ด๊ธฐํ ํ  ์ ์์ต๋๋ค.
```
var i MyInt = 5
r := Rectangle{10, 20}

pArr := []Printer{i, r} // ์ฌ๋ผ์ด์ค ํํ๋ก ์ธํฐํ์ด์ค ์ด๊ธฐํ
for index, _ := range pArr {
	pArr[index].Print() // ์ฌ๋ผ์ด์ค๋ฅผ ์ํํ๋ฉด์ Print ๋ฉ์๋ ํธ์ถ
}

for _, value := range pArr {
	value.Print() // ์ฌ๋ผ์ด์ค๋ฅผ ์ํํ๋ฉด์ Print ๋ฉ์๋ ํธ์ถ
}
```
## ๐ฏ๋ํ์ดํ
- ๋ํ์ดํ: ๊ฐ ๊ฐ์ด๋ ์ธ์คํด์ค์ ์ค์  ํ์์ ์๊ดํ์ง ์๊ณ , ๊ตฌํ๋ ๋ฉ์๋๋ก๋ง ํ์์ ํ๋จํ๋ ๋ฐฉ์,์ด ์ฉ์ด๋ ๋ค์๊ณผ ๊ฐ์ ๋ ํ์คํธ(์ค๋ฆฌ ํ์คํธ)์์ ์ ๋
- โ๋ง์ฝ ์ด๋ค ์๊ฐ ์ค๋ฆฌ์ฒ๋ผ ๊ฑท๊ณ , ํค์์น๊ณ , ๊ฝฅ๊ฝฅ๊ฑฐ๋ฆฌ๋ ์๋ฆฌ๋ฅผ ๋ธ๋ค๋ฉด ๋๋ ๊ทธ ์๋ฅผ ์ค๋ฆฌ๋ผ ๋ถ๋ฅด๊ฒ ๋ค.โ

Go ์ธ์ด๋ก๋ ๋ค์๊ณผ ๊ฐ์ด ๋ ํ์ดํ์ ๊ตฌํํ  ์ ์๋ค

```
package main

import "fmt"

type Duck struct { // ์ค๋ฆฌ(Duck) ๊ตฌ์กฐ์ฒด ์ ์
}

func (d Duck) quack() {     // ์ค๋ฆฌ์ quack ๋ฉ์๋ ์ ์
	fmt.Println("๊ฝฅ~!") // ์ค๋ฆฌ ์ธ์ ์๋ฆฌ
}

func (d Duck) feathers() { // ์ค๋ฆฌ์ feathers ๋ฉ์๋ ์ ์
	fmt.Println("์ค๋ฆฌ๋ ํฐ์๊ณผ ํ์ ํธ์ ๊ฐ์ง๊ณ  ์์ต๋๋ค.")
}

type Person struct { // ์ฌ๋(Person) ๊ตฌ์กฐ์ฒด ์ ์
}

func (p Person) quack() {                           // ์ฌ๋์ quack ๋ฉ์๋ ์ ์
	fmt.Println("์ฌ๋์ ์ค๋ฆฌ๋ฅผ ํ๋ด๋๋๋ค. ๊ฝฅ~!") // ์ฌ๋์ด ์ค๋ฆฌ ์๋ฆฌ๋ฅผ ํ๋ด๋
}

func (p Person) feathers() { // ์ฌ๋์ feathers ๋ฉ์๋ ์ ์
	fmt.Println("์ฌ๋์ ๋์์ ๊นํธ์ ์ฃผ์์ ๋ณด์ฌ์ค๋๋ค.")
}

type Quacker interface { // quack, feathers ๋ฉ์๋๋ฅผ ๊ฐ์ง๋ Quacker ์ธํฐํ์ด์ค ์ ์
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker ์ธํฐํ์ด์ค๋ก quack ๋ฉ์๋ ํธ์ถ
	q.feathers() // Quacker ์ธํฐํ์ด์ค๋ก feathers ๋ฉ์๋ ํธ์ถ
}

func main() {
	var donald Duck // ์ค๋ฆฌ ์ธ์คํด์ค ์์ฑ
	var john Person // ์ฌ๋ ์ธ์คํด์ค ์์ฑ

	inTheForest(donald) // ์ธํฐํ์ด์ค๋ฅผ ํตํ์ฌ ์ค๋ฆฌ์ quack, feather ๋ฉ์๋ ํธ์ถ
	inTheForest(john)   // ์ธํฐํ์ด์ค๋ฅผ ํตํ์ฌ ์ฌ๋์ quack, feather ๋ฉ์๋ ํธ์ถ
}
```
- Quacker ์ธํฐํ์ด์ค๋ quack, feathers ํจ์๋ฅผ ์ ์ํ๊ณ  ์๋ค. ๊ทธ๋ฆฌ๊ณ  ์ค๋ฆฌ(Duck)์ ์ฌ๋(Person) ๋ชจ๋ quack, feathers ํจ์๋ฅผ ๊ตฌํํ์๋ค

- ์ค์ ๋ก ์ฌ์ฉํ  ๋๋ ๋ค์๊ณผ ๊ฐ์ด inTheForest ํจ์์์ Quacker ์ธํฐํ์ด์ค๋ฅผ ๋งค๊ฐ๋ณ์๋ก ๋ฐ๋๋ค. ์ฌ๊ธฐ์ ์ค๋ฆฌ๋  ์ฌ๋์ด๋  inTheForest ํจ์์ ๋ฃ์ ์ ์์ผ๋ฉฐ quack, feathers ๋ฉ์๋๋ฅผ ํธ์ถํ๋ค. ์ค๋ฆฌ๋ ์ง์ง โ๊ฝฅ~โํ๊ณ  ์ธ๊ฒ์ด๊ณ , ์ฌ๋์ด๋ผ๋ฉด ์ค๋ฆฌ๋ฅผ ํ๋ด๋ด์ด ์๋ฆฌ๋ฅผ ๋ผ ๊ฒ์ด๋ค

```
func inTheForest(q Quacker) {
	q.quack()    // Quacker ์ธํฐํ์ด์ค๋ก quack ๋ฉ์๋ ํธ์ถ
	q.feathers() // Quacker ์ธํฐํ์ด์ค๋ก feathers ๋ฉ์๋ ํธ์ถ
}

func main() {
	var donald Duck // ์ค๋ฆฌ ์ธ์คํด์ค ์์ฑ
	var john Person // ์ฌ๋ ์ธ์คํด์ค ์์ฑ

	inTheForest(donald) // ์ธํฐํ์ด์ค๋ฅผ ํตํ์ฌ ์ค๋ฆฌ์ quack, feather ๋ฉ์๋ ํธ์ถ
	inTheForest(john)   // ์ธํฐํ์ด์ค๋ฅผ ํตํ์ฌ ์ฌ๋์ quack, feather ๋ฉ์๋ ํธ์ถ
}
```
```
๊ฝฅ~!
์ค๋ฆฌ๋ ํฐ์๊ณผ ํ์ ํธ์ ๊ฐ์ง๊ณ  ์์ต๋๋ค.
์ฌ๋์ ์ค๋ฆฌ๋ฅผ ํ๋ด๋๋๋ค. ๊ฝฅ~!
์ฌ๋์ ๋์์ ๊นํธ์ ์ฃผ์์ ๋ณด์ฌ์ค๋๋ค.
```
- ์ค๋ฆฌ๋  ์ฌ๋์ด๋  ์๊ด์์ด ๊ฝฅ(quack)๊ณผ ๊นํธ(feathers) ๋ฉ์๋๋ง ๊ฐ์ก๋ค๋ฉด ๋ชจ๋ ๊ฐ์ ์ธํฐํ์ด์ค๋ก ์ฒ๋ฆฌํ  ์ ์๋ค

- ํ์์ด ํน์  ์ธํฐํ์ด์ค๋ฅผ ๊ตฌํํ๋์ง ๊ฒ์ฌํ๋ ค๋ฉด ๋ค์๊ณผ ๊ฐ์ด ์ฌ์ฉํ๋ค.

- interface{}(์ธ์คํด์ค).(์ธํฐํ์ด์ค)

```
var donald Duck

if v, ok := interface{}(donald).(Quacker); ok {
	fmt.Println(v, ok)
}
```
```
{} true
```
Duck ํ์์ ์ธ์คํด์ค donald๋ฅผ ๋น ์ธํฐํ์ด์ค์ ๋ฃ์ ๋ค Quacker ์ธํฐํ์ด์ค์ ๊ฐ์์ง ํ์ธํ๋ค. ์ฒซ ๋ฒ์งธ ๋ฆฌํด๊ฐ์ ๊ฒ์ฌํ๋ ์ธ์คํด์ค์ด๋ฉฐ ๋ ๋ฒ์งธ ๋ฆฌํด๊ฐ์ ์ธ์คํด์ค๊ฐ ํด๋น ์ธํฐํ์ด์ค๋ฅผ ๊ตฌํํ๊ณ  ์๋์ง ์ฌ๋ถ, ๊ตฌํํ๊ณ  ์๋ค๋ฉด true ๊ทธ๋ ์ง ์์ผ๋ฉด false!! 
## ๐ฏ๋น ์ธํฐํ์ด์ค ์ฌ์ฉ

- ์ธํฐํ์ด์ค์ ์๋ฌด ๋ฉ์๋๋ ์ ์๋์ด ์์ง ์์ผ๋ฉด ๋ชจ๋  ํ์์ ์ ์ฅ ๊ฐ๋ฅ! 
```
func f1(arg interface{}) { // ๋ชจ๋  ํ์์ ์ ์ฅํ  ์ ์์
}
```
```
type Any interface{} // ์ธํฐํ์ด์ค์ ๋ฉ์๋๋ฅผ ์ง์ ํ์ง ์์

func f2(arg Any) {   // ๋ชจ๋  ํ์์ ์ ์ฅํ  ์ ์์
}
```
- ๋น ์ธํฐํ์ด์ค ํ์์ ํจ์์ ๋งค๊ฐ๋ณ์, ๋ฆฌํด๊ฐ, ๊ตฌ์กฐ์ฒด์ ํ๋๋ก ์ฌ์ฉํ  ์ ์๋ค

- ๋ชจ๋  ํ์์ ๋ฐ์์ ์ถ๋ ฅํ๋ ํจ์

```
v
package main

import (
	"fmt"
	"strconv"
)

//                      โ ๋น ์ธํฐํ์ด์ค๋ฅผ ์ฌ์ฉํ์ฌ ๋ชจ๋  ํ์์ ๋ฐ์
func formatString(arg interface{}) string {
	//       โ ์ธํฐํ์ด์ค์ ์ ์ฅ๋ ํ์์ ๋ฐ๋ผ case ์คํ
	switch arg.(type) {
	case int:                      // arg๊ฐ intํ์ด๋ผ๋ฉด
		i := arg.(int)         // intํ์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return strconv.Itoa(i) // strconv.Itoa ํจ์๋ฅผ ์ฌ์ฉํ์ฌ i๋ฅผ ๋ฌธ์์ด๋ก ๋ณํ
	case float32:                  // arg๊ฐ float32ํ์ด๋ผ๋ฉด
		f := arg.(float32)     // float32ํ์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
                                       // strconv.FormatFloat ํจ์๋ฅผ ์ฌ์ฉํ์ฌ f๋ฅผ ๋ฌธ์์ด๋ก ๋ณํ
	case float64:                  // arg๊ฐ float64ํ์ด๋ผ๋ฉด
		f := arg.(float64)     // float64ํ์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return strconv.FormatFloat(f, 'f', -1, 64)
                                       // strconv.FormatFloat ํจ์๋ฅผ ์ฌ์ฉํ์ฌ f๋ฅผ ๋ฌธ์์ด๋ก ๋ณํ
	case string:                   // arg๊ฐ string์ด๋ผ๋ฉด
		s := arg.(string)      // string์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return s               // string์ด๋ฏ๋ก ๊ทธ๋๋ก ๋ฆฌํด
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))
}
```
```
1
2.5
Hello, world!
```
- formatString ํจ์์ ๋งค๊ฐ๋ณ์๋ฅผ ๋ณด๋ฉด ํ์์ interface{}๋ก ์ง์ ํ์ฟ๋ค. ์ด๋ ๊ฒ ํ๋ฉด ๋ชจ๋  ํ์์ ์ฒ๋ฆฌํ  ์ ์๋ค
- ์ธํฐํ์ด์ค์ ์ ์ฅ๋ ํ์์ ์์๋ด๋ ค๋ฉด switch ๋ถ๊ธฐ๋ฌธ ์์์ arg.(type)์ฒ๋ผ ์ธํฐํ์ด์ค ๋ณ์์ .(type)์ ์ฌ์ฉํ๋ค ๋จ ์ด ๋ฐฉ๋ฒ์ switch ๋ถ๊ธฐ๋ฌธ ์์์๋ง ์ฌ์ฉํ  ์ ์๊ณ , ์ผ๋ฐ์ ์ธ ๋ฐฉ๋ฒ์ผ๋ก๋ ์ฌ์ฉํ  ์ ์๋ค.
- ํ์์ ๋ฐ๋ผ case๋ก ๋๋๋ค. 
- ๋น ์ธํฐํ์ด์ค ๋ณ์๋ ๊ทธ๋๋ก ์ฌ์ฉํ  ์ ์์ผ๋ฏ๋ก arg.(int), arg.(float32), arg.(float64), arg.(string)์ฒ๋ผ ํ์์ ์ง์ ํ์ฌ ๊ฐ์ ๊ฐ์ ธ์ค๋ฉฐ, ์ด๋ ๊ฒ ํ์์ ์ํ๋ ํํ๋ก ๋ฐ๊พธ๋ ์์์ Type assertion์ด๋ผ๊ณ  ํ๋ค
- ๊ฐ ํ์์ ๋ง๊ฒ strconv.Itoa, strconv.FormatFloat ํจ์๋ฅผ ์ฌ์ฉํ์ฌ ๊ฐ์ ๋ฌธ์์ด๋ก ๋ง๋ค๋ฉฐ, string์ ๊ฐ์ ๋ฌธ์์ด์ด๋ฏ๋ก ๋ณํํ์ง ์๊ณ  ๊ทธ๋๋ก ๋ฆฌํดํ๋ค.

- ์ผ๋ฐ ์๋ฃํ๋ฟ๋ง ์๋๋ผ ๊ตฌ์กฐ์ฒด ์ธ์คํด์ค ๋ฐ ํฌ์ธํฐ๋ ๋น ์ธํฐํ์ด์ค๋ก ๋๊ธธ ์ ์๋ค.
```
package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Person ๊ตฌ์กฐ์ฒด ์ ์
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:               // arg์ ํ์์ด Person์ด๋ผ๋ฉด
		p := arg.(Person)  // Person ํ์์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return p.name + " " + strconv.Itoa(p.age) // ๊ฐ ํ๋๋ฅผ ํฉ์ณ์ ๋ฆฌํด
	case *Person:              // arg์ ํ์์ด Person ํฌ์ธํฐ๋ผ๋ฉด
		p := arg.(*Person) // *Person ํ์์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์ด
		return p.name + " " + strconv.Itoa(p.age) // ๊ฐ ํ๋๋ฅผ ํฉ์ณ์ ๋ฆฌํด
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(&Person{"John", 12}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString(andrew))
}
```
```
Maria 20
John 12
Andrew 35
```
- switch arg.(type) { }์ผ๋ก ์ธํฐํ์ด์ค์ ์ ์ฅ๋ ํ์์ ์์๋ธ ๋ค ๊ฐ ๊ตฌ์กฐ์ฒด ํ์๋ณ๋ก case๋ฅผ ๋ง๋ค์ด ์ฒ๋ฆฌํ๋ฉฐ, ์ฌ๊ธฐ์ ๊ตฌ์กฐ์ฒด๋ฅผ ๊ทธ๋๋ก ๋๊ฒจ์คฌ๋ค๋ฉด case Person:์ผ๋ก ์ฒ๋ฆฌํ๊ณ , ๊ตฌ์กฐ์ฒด์ ํฌ์ธํฐ๋ฅผ ๋๊ฒจ์คฌ๋ค๋ฉด case *Person:์ผ๋ก ์ฒ๋ฆฌํ๋ค ๋ง์ฐฌ๊ฐ์ง๋ก ๊ตฌ์กฐ์ฒด์ผ ๋๋ arg.(Person), ํฌ์ธํฐ์ผ ๋๋ arg.(*Person)์ผ๋ก ๊ฐ์ ๊ฐ์ ธ์จ๋ค

์ธํฐํ์ด์ค์ ์ ์ฅ๋ ํ์์ด ํน์  ํ์์ธ์ง ๊ฒ์ฌํ๋ ค๋ฉด ๋ค์๊ณผ ๊ฐ์ด ์ฌ์ฉ
```
var t interface{}
t = Person{"Maria", 20}

if v, ok := t.(Person); ok {
	fmt.Println(v, ok)
}
```
```
{Maria 20} true
```
์ธํฐํ์ด์ค.(ํ์) ํ์์ด๋ฉฐ, ์ฒซ ๋ฒ์งธ ๋ฆฌํด๊ฐ์ ํด๋น ํ์์ผ๋ก ๋ ๊ฐ์ด๋ฉฐ ๋ ๋ฒ์งธ ๋ฆฌํด๊ฐ์ ํ์์ด ๋ง๋์ง ์ฌ๋ถ๋ฅผํ์ธ . ํ์์ด ์ผ์นํ๋ฉด true ๊ทธ๋ ์ง ์์ผ๋ฉด false์ด๋ค
