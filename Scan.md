# π©π»βπμλ ₯ν¨μ

λ€μμ fmt ν¨ν€μ§μμ μ κ³΅νλ νμ€ μλ ₯ ν¨μμ΄λ€.

- func Scan(a ...interface{}) (n int, err error): μ½μμμ κ³΅λ°±, μ μ€λ‘ κ΅¬λΆνμ¬ μλ ₯μ λ°μ
- func Scanln(a ...interface{}) (n int, err error): μ½μμμ κ³΅λ°±μΌλ‘ κ΅¬λΆνμ¬ μλ ₯μ λ°μ
- func Scanf(format string, a ...interface{}) (n int, err error): μ½μμμ νμμ μ§μ νμ¬ μλ ₯μ λ°μ

```
package main

import "fmt"

func main() {
	var s1, s2 string
	n, _ := fmt.Scan(&s1, &s2) // fmt.Scan ν¨μμ λ λ²μ§Έ λ¦¬ν΄κ°μ μλ΅
	fmt.Println("μλ ₯ κ°μ:", n)
	fmt.Println(s1, s2)
}
```

- fmt.Scan ν¨μλ νμ€ μλ ₯μμ μλ ₯λ°μ λ¬Έμ, μ«μλ₯Ό λ³μμ μ μ₯νλ€. μλ ₯λ°μ κ°μλ§νΌ λ§€κ°λ³μμ λ³μλ₯Ό λ£μ΄μ£Όλ©΄ λλ©° λ°λμ λ νΌλ°μ€(ν¬μΈν°) ννλ‘ λ£λλ€. κ·Έλ¦¬κ³  μλ ₯μ΄ μ±κ³΅ν κ°μλ₯Ό λ¦¬ν΄νλ€.

- μμ€ νμΌμ μ»΄νμΌνμ¬ μ½μ(ν°λ―Έλ)μμ μ€ννλ€. κ·Έλ¦¬κ³  λ€μκ³Ό κ°μ΄ λ¬Έμμ΄μ κ³΅λ°±μΌλ‘ κ΅¬λΆνμ¬ μλ ₯νλ©΄ λλ€.

```
$ go build
$ ./scan
Hello, world! (μλ ₯)
μλ ₯ κ°μ: 2
Hello, world!
```
- fmt.Scan ν¨μλ κ³΅λ°±λΏλ§ μλλΌ μ μ€(Newline)μΌλ‘λ κ°μ κ΅¬λΆν  μ μμΌλ―λ‘ λ€μκ³Ό κ°μ΄ λ¬Έμμ΄μ λ μ€λ‘ μλ ₯ν΄λ λλ€. λ§μ½ fmt.Scan ν¨μμ λ§€κ°λ³μλ₯Ό 4κ° λ£μλ€λ©΄ 4λ² μλ ₯μ λ°μ κ²μ΄λ€.
```
$ ./scan
Hello, (μλ ₯)
world! (μλ ₯)
μλ ₯ κ°μ: 2
Hello, world!
```
- fmt.Scanlnμ ν μ€μμ κ³΅λ°±μΌλ‘λ§ κ°μ κ΅¬λΆνλ©° μν°ν€λ₯Ό μλ ₯νμ¬ μ μ€λ‘ λμ΄κ°λ©΄ μλ ₯μ μ’λ£λλ€.
```
var s1, s2 string
n, _ := fmt.Scanln(&s1, &s2)
fmt.Println("μλ ₯ κ°μ:", n)
fmt.Println(s1, s2)
```
```
$ go build
$ ./scan
Hello, world! (μλ ₯)
μλ ₯ κ°μ: 2
Hello, world!
```
```
package main

import "fmt"

func main() {
	var num1, num2 int
	n, _ := fmt.Scanf("%d,%d", &num1, &num2) // μ μνμΌλ‘ νμμ μ§μ νμ¬ μλ ₯μ λ°μ
	fmt.Println("μλ ₯ κ°μ:", n)
	fmt.Println(num1, num2)
}
```
```
fmt.Scanf ν¨μμ μ²« λ²μ§Έ λ§€κ°λ³μμ νμ μ§μ μλ₯Ό μ΄μ©νμ¬ μλ ₯λ°μ νμμ μ€μ ν©λλ€. κ·Έλ¦¬κ³  λ λ²μ§Έ λ§€κ°λ³μλΆν°λ μλ ₯λ°μ λ³μλ₯Ό λ νΌλ°μ€ νμμΌλ‘ λ£μ΅λλ€. μ¬κΈ°μ % (νμ μ§μ μ)κ°μμ λ³μμ κ°μκ° κ°μμΌ ν©λλ€.

μμ€ νμΌμ μ»΄νμΌνμ¬ μ½μ(ν°λ―Έλ)μμ μ€νν λ€ 1,2μ²λΌ μ«μλ₯Ό , (μ½€λ§)λ‘ κ΅¬λΆνμ¬ μλ ₯ν©λλ€.
```
```
$ go build
$ ./scan
1,2 (μλ ₯)
μλ ₯ κ°μ: 2
1 2
```
