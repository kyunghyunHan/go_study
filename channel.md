# π©π»βπμ±λ

- μ±λ(channel)μ κ³ λ£¨ν΄λΌλ¦¬ λ°μ΄ν°λ₯Ό μ£Όκ³  λ°κ³ , μ€ν νλ¦μ μ μ΄νλ κΈ°λ₯μ΄λ€.
- λͺ¨λ  νμμ μ±λλ‘ μ¬μ©ν  μ μλ€. κ·Έλ¦¬κ³  μ±λ μμ²΄λ κ°μ΄ μλ λ νΌλ°μ€ νμμ΄λ€
```
κ³ λ£¨ν΄  <-> μ±λ <->κ³ λ£¨ν΄
```
- μ±λμ κ°μ λ³΄λ΄κ³  κΊΌλ

- μ±λμ λ€μκ³Ό κ°μ νμμΌλ‘ μ¬μ©

- make(chan μλ£ν)
κ³ λ£¨ν΄κ³Ό μ±λμ μ¬μ©νμ¬ λ μ μ κ°μ λνκΈ°
```
package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b          // μ±λμ aμ bμ ν©μ λ³΄λ
}

func main() {
	c := make(chan int) // intν μ±λ μμ±

	go sum(1, 2, c)     // sumμ κ³ λ£¨ν΄μΌλ‘ μ€νν λ€ μ±λμ λ§€κ°λ³μλ‘ λκ²¨μ€

	n := <-c            // μ±λμμ κ°μ κΊΌλΈ λ€ nμ λμ
	fmt.Println(n)      // 3
}
```
- μ±λμ μ¬μ©νκΈ° μ μλ λ°λμ make ν¨μλ‘ κ³΅κ°μ ν λΉν΄μΌ νλ€. κ·Έλ¦¬κ³  μ΄λ κ² μμ±νλ©΄ λκΈ° μ±λ(synchronous channel)μ΄ μμ±λλ€.

- λ€μκ³Ό κ°μ΄ :=λ₯Ό μ¬μ©νμ§ μκ³ , var ν€μλλ‘ μ±λμ μ μΈνκ³  ν λΉν  μλ μλ€

```
var c chan int // chan intν λ³μ μ μΈ
c = make(chan int)
```
- sum ν¨μλ₯Ό κ³ λ£¨ν΄μΌλ‘ μ€ννλ©΄μ λν  κ°κ³Ό μ±λ λ³μλ₯Ό λ£λλ€. μ±λμ λ§€κ°λ³μλ‘ λ°λ ν¨μλ λ°λμ go ν€μλλ₯Ό μ¬μ©νμ¬ κ³ λ£¨ν΄μΌλ‘ μ€νν΄μΌ νλ€

- ν¨μμμ μ±λμ λ§€κ°λ³μλ‘ λ°μ λλ λ€μκ³Ό κ°μ νμμΌλ‘ μ¬μ©νλ€.

- λ³μλͺ chan μλ£ν
```
//                         β intν μ±λμ λ§€κ°λ³μλ‘ λ°μ
func sum(a int, b int, c chan int) {
	c <- a + b
//        β μ±λμ κ°μ λ³΄λ
}
```
- μ±λμ κ°μ λ³΄λΌ λλ λ€μκ³Ό κ°μ νμμΌλ‘ μ¬μ©νλ€.

- μ±λ <- κ°
- μ±λ λ³μμλ =λ‘ κ°μ λμνμ§ μκ³  <- μ°μ°μλ₯Ό μ¬μ©νλ€. μ¬κΈ°μλ sum ν¨μ μμμ aμ bλ₯Ό λν κ°μ μ±λ cλ‘ λ³΄λΈλ€

- μ΄μ  main ν¨μμμλ μ±λμμ κ°μ κ°μ Έμ¨λ€

- <- μ±λ
```
n := <-c
```
- μ΄λ²μλ <- μ°μ°μλ₯Ό μ¬μ©νμ§λ§ μμκ° λ°λλ‘ λμ΄μλ€. μ¦ μ±λ cμμ κ°μ κ°μ Έμ¨ λ€ λ³μ nμ μμ±νμ¬ λμνλ€(fmt.Println(<-c)μ²λΌ λ³μμ λμνμ§ μκ³  λ°λ‘ μ¬μ©ν  μλ μλ€).

- <-cλ μ±λμμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ λκΈ°νλ€. κ·Έλ¦¬κ³  μ±λμ κ°μ΄ λ€μ΄μ€λ©΄ λκΈ°λ₯Ό λλ΄κ³  λ€μ μ½λλ₯Ό μ€ννλ€. λ°λΌμ μ±λμ κ°μ μ£Όκ³  λ°λ λμμ λκΈ°ν μ­ν κΉμ§ μννλ€.
μμ½νμλ©΄ λ€μκ³Ό κ°λ€.

- μ±λ <-: μ±λμ κ°μ λ³΄λλλ€.
- <- μ±λ: μ±λμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ κΈ°λ€λ¦° λ€ κ°μ΄ λ€μ΄μ€λ©΄ κ°μ κ°μ Έμ΅λλ€.
- κ°μ Έμ¨ κ°μ :=, =λ₯Ό μ¬μ©νμ¬ λ³μμ λμν  μ μλ€.
- κ°μ κ°μ Έμ€λ©΄ λ€μ μ½λλ₯Ό μ€ννλ€.

## π―λκΈ°μ±λ

- λκΈ° μ±λμ νΉμ±
- λ€μμ κ³ λ£¨ν΄κ³Ό λ©μΈ ν¨μλ₯Ό λ²κ°μκ°λ©΄μ μ€ν

```
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // λκΈ° μ±λ μμ±
	count := 3              // λ°λ³΅ν  νμ

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // κ³ λ£¨ν΄μ trueλ₯Ό λ³΄λ, κ°μ κΊΌλΌ λκΉμ§ λκΈ°
			fmt.Println("κ³ λ£¨ν΄ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
			time.Sleep(1 * time.Second) // 1μ΄ λκΈ°
		}
	}()

	for i := 0; i < count; i++ {
		<-done                         // μ±λμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ λκΈ°, κ°μ κΊΌλ
		fmt.Println("λ©μΈ ν¨μ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
	}
}
```
- make(chan bool)μ²λΌ μ±λμ μμ±ν©λλ€. μ¬κΈ°μλ μ±λλ‘ κ°μ μ£Όκ³  λ°μλ μ€μ λ‘ μ¬μ©νμ§ μμΌλ―λ‘ μλ£νμ ν° μλ―Έκ° μλ€. make ν¨μμ λ§€κ° λ³μλ₯Ό νλλ§ μ§μ νμΌλ―λ‘ λκΈ° μ±λμ μμ±νλ€. 

- λ¨Όμ  κ³ λ£¨ν΄μ μμ±νκ³ , λ°λ³΅λ¬Έμ μ€νν  λλ§λ€ μ±λ doneμ true κ°μ λ³΄λΈ λ€ 1μ΄λ₯Ό κΈ°λ€λ¦°λ€

```
go func() {
	for i := 0; i < count; i++ {
		done <- true                // κ³ λ£¨ν΄μ trueλ₯Ό λ³΄λ, κ°μ κΊΌλΌ λκΉμ§ λκΈ°
		fmt.Println("κ³ λ£¨ν΄ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
		time.Sleep(1 * time.Second) // 1μ΄ λκΈ°
	}
}()
```
- λκΈ° μ±λμ΄λ―λ‘ doneμ κ°μ λ³΄λ΄λ©΄ λ€λ₯Έ μͺ½μμ κ°μ κΊΌλΌ λκΉμ§ λκΈ°νλ€. λ°λΌμ λ°λ³΅λ¬Έλ μ€νλμ§ μμΌλ―λ‘ βκ³ λ£¨ν΄ : μ«μβκ° κ³μ μΆλ ₯λμ§ μλλ€.

- μ΄μ  λ©μΈ ν¨μμμλ λ°λ³΅λ¬Έμ μ€νν  λλ§λ€ μ±λ doneμμ κ°μ κΊΌλΈλ€
```
for i := 0; i < count; i++ {
	<-done                         // μ±λμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ λκΈ°, κ°μ κΊΌλ
	fmt.Println("λ©μΈ ν¨μ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
}
```
- <-done λΆλΆμμ μ±λμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ λκΈ°νλ€ λ¨Όμ  μμ κ³ λ£¨ν΄μμ doneμ κ°μ λ³΄λκΈ° λλ¬Έμ κ°μ κΊΌλΈ λ€ λ€μ μ½λλ‘ μ§ννλ€. κ·Έλ¦¬κ³  κ³ λ£¨ν΄ μͺ½μ λκΈ°λ μ’λ£λλ©΄μ λ€μ λ°λ³΅λ¬Έμ΄ μ€νλ λ€ μ±λμ κ°μ λ³΄λΈλ€. κ·Έλ¦¬κ³  λ©μΈ ν¨μλ μ±λμμ κ°μ κΊΌλ΄κ³ , λ€μ κ³ λ£¨ν΄λ μ±λμ κ°μ λ³΄λΈλ€. λ°λΌμ λ€μκ³Ό κ°μ΄ κ³ λ£¨ν΄ β λ©μΈ ν¨μ β κ³ λ£¨ν΄ β λ©μΈ ν¨μ μμλ‘ μ€νλλ€.
```
κ³ λ£¨ν΄ :  0
λ©μΈ ν¨μ :  0
κ³ λ£¨ν΄ :  1
λ©μΈ ν¨μ :  1
κ³ λ£¨ν΄ :  2
λ©μΈ ν¨μ :  2
```
- λκΈ° μ±λμ λ³΄λ΄λ μͺ½μμλ κ°μ λ°μ λκΉμ§ λκΈ°νκ³ , λ°λ μͺ½μμλ μ±λμ κ°μ΄ λ€μ΄μ¬ λκΉμ§ λκΈ°νλ€. λ°λΌμ, λκΈ° μ±λμ νμ©νλ©΄ κ³ λ£¨ν΄μ μ½λ μ€ν μμλ₯Ό μ μ΄ν  μ μλ€.

## π―μ±λλ²νΌλ§

- λ€μμ μ±λμ λ²νΌκ° κ°λμ°¨λ©΄ κ°μ κΊΌλ΄μ μΆλ ₯νλ€

- make(chan μλ£ν, λ²νΌ_κ°μ)

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // λ²νΌκ° 2κ°μΈ λΉλκΈ° μ±λ μμ±
	count := 4                 // λ°λ³΅ν  νμ

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // μ±λμ trueλ₯Ό λ³΄λ, λ²νΌκ° κ°λμ°¨λ©΄ λκΈ°
			fmt.Println("κ³ λ£¨ν΄ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
		}
	}()

	for i := 0; i < count; i++ {
		<-done                         // λ²νΌμ κ°μ΄ μμΌλ©΄ λκΈ°, κ°μ κΊΌλ
		fmt.Println("λ©μΈ ν¨μ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
	}
}

```
- μ±λμ λ²νΌλ₯Ό 1κ° μ΄μ μ€μ νλ©΄ λΉλκΈ° μ±λ(asynchronous channel)μ΄ μμ±λλ€. λΉλκΈ° μ±λμ λ³΄λ΄λ μͺ½μμ λ²νΌκ° κ°λ μ°¨λ©΄ μ€νμ λ©μΆκ³  λκΈ°νλ©° λ°λ μͺ½μμλ λ²νΌμ κ°μ΄ μμΌλ©΄ λκΈ°νλ€.

- κ³ λ£¨ν΄μ μμ±νκ³ , λ°λ³΅λ¬Έμ μ€νν  λλ§λ€ μ±λ doneμ true κ°μ λ³΄λΈλ€.
```
go func() {
	for i := 0; i < count; i++ {
		done <- true                // μ±λμ trueλ₯Ό λ³΄λ, λ²νΌκ° κ°λμ°¨λ©΄ λκΈ°
		fmt.Println("κ³ λ£¨ν΄ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
	}
}()
```
- λΉλκΈ° μ±λμ΄λ―λ‘ λ²νΌκ° κ°λ μ°°λκΉμ§ κ°μ κ³μ λ³΄λΈλ€. μ¬κΈ°μλ μ±λμ λ²νΌλ₯Ό 2κ°λ‘ μ€μ νμΌλ―λ‘ doneμ trueλ₯Ό 2λ² λ³΄λΈ λ€ κ·Έ λ€μ λ£¨νμμ λκΈ°νλ€(λ²νΌκ° κ°λ μ°¨λ©΄ λκΈ°νλ―λ‘ iκ° 0, 1μΌ λ μ±λμ κ°μ λ³΄λΈ λ€ iκ° 2μΌλ done <- trueμμ λκΈ°νλ€).

- μ΄μ  λ©μΈ ν¨μμμλ λ°λ³΅λ¬Έμ μ€νν  λλ§λ€ μ±λ doneμμ κ°μ κΊΌλΈλ€.
```
for i := 0; i < count; i++ {
	<-done                         // λ²νΌμ κ°μ΄ μμΌλ©΄ λκΈ°, κ°μ κΊΌλ
	fmt.Println("λ©μΈ ν¨μ : ", i) // λ°λ³΅λ¬Έμ λ³μ μΆλ ₯
}
```
- λΉλκΈ° μ±λμ λ²νΌκ° 2κ°μ΄λ―λ‘ doneμλ μ΄λ―Έ κ°μ΄ 2κ° λ€μ΄μλ€. λ°λΌμ λ£¨νλ₯Ό λ λ² λ°λ³΅νλ©΄μ <-doneμμ κ°μ κΊΌλΈλ€. κ·Έ λ€μμλ μ±λμ΄ λΉμμΌλ―λ‘ μ€νμ λ©μΆκ³  λκΈ°νλ€. κ·Έλ¦¬κ³  λ€μ κ³ λ£¨ν΄ μͺ½μμ κ°μ λ λ² λ³΄λ΄κ³ , λ©μΈ ν¨μμμ λ λ² κΊΌλΈλ€. μ¦ κ³ λ£¨ν΄ β κ³ λ£¨ν΄ β λ©μΈ ν¨μ β λ©μΈ ν¨μ μμλ‘ μ€νλλ€.
```
κ³ λ£¨ν΄ :  0
κ³ λ£¨ν΄ :  1
λ©μΈ ν¨μ :  0
λ©μΈ ν¨μ :  1
κ³ λ£¨ν΄ :  2
κ³ λ£¨ν΄ :  3
λ©μΈ ν¨μ :  2
λ©μΈ ν¨μ :  3
```

## π©π»βπrange,close μ¬μ©
- λ€μμ 0λΆν° 4κΉμ§ μ±λμ κ°μ λ³΄λ΄κ³ , λ€μ μ±λμμ κ°μ κΊΌλ΄μ μΆλ ₯νλ€.
```
package main

import "fmt"

func main() {
	c := make(chan int) // intν μ±λμ μμ±

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // μ±λμ κ°μ λ³΄λ
		}
		close(c)       // μ±λμ λ«μ
	}()

	for i := range c { // rangeλ₯Ό μ¬μ©νμ¬ μ±λμμ κ°μ κΊΌλ
		fmt.Println(i) // κΊΌλΈ κ°μ μΆλ ₯
	}
}
```
```
0
1
2
3
4
```
- for λ°λ³΅λ¬Έ μμμ range ν€μλλ₯Ό μ¬μ©νλ©΄ μ±λμ΄ λ«ν λκΉμ§ λ°λ³΅νλ©΄μ κ°μ κΊΌλΈλ€. μ¬κΈ°μλ λμμ κ³ λ£¨ν΄ μμμ μ±λ cμ 0λΆν° 4κΉμ§ κ°μ λ³΄λΈ λ€ closeλ₯Ό μ¬μ©νμ¬ μ±λμ λ«λλ€. μ΄λ κ² νλ©΄ rangeλ‘ 0λΆν° 4κΉμ§ κΊΌλ΄κ³ , κ°μ μΆλ ₯ν λ€ λ°λ³΅λ¬Έμ΄ μ’λ£λλ€.

 rangeμ close ν¨μμ νΉμ§

- μ΄λ―Έ λ«ν μ±λμ κ°μ λ³΄λ΄λ©΄ ν¨λμ΄ λ°μνλ€
- μ±λμ λ«μΌλ©΄ range λ£¨νκ° μ’λ£λλ€.
- μ±λμ΄ μ΄λ €μκ³ , κ°μ΄ λ€μ΄μ€μ§ μλλ€λ©΄ rangeλ μ€νλμ§ μκ³  κ³μ λκΈ°νλ€ λ§μ½ λ€λ₯Έ κ³³μμ μ±λμ κ°μ λ³΄λλ€λ©΄(μ±λμ κ°μ΄ λ€μ΄μ€λ©΄) κ·ΈλλΆν° rangeκ° κ³μ λ°λ³΅λλ€
- μ λ¦¬νμλ©΄ rangeλ μ±λμ κ°μ΄ λͺ κ°λ λ€μ΄μ¬μ§ λͺ¨λ₯΄κΈ° λλ¬Έμ κ°μ΄ λ€μ΄μ¬ λλ§λ€ κ³μ κΊΌλ΄κΈ° μν΄ μ¬μ©νλ€

- μ±λμ κ°μ Έμ¨ λ€ λ λ²μ§Έ λ¦¬ν΄κ°μΌλ‘ μ±λμ΄ λ«νλμ§ νμΈν  μ μ

```
c := make(chan int, 1)

go func() {
	c <- 1
}()

a, ok := <-c       // μ±λμ΄ λ«νλμ§ νμΈ
fmt.Println(a, ok) // 1 true: μ±λμ μ΄λ € μκ³  1μ κΊΌλ

close(c)           // μ±λμ λ«μ
a, ok = <-c        // μ±λμ΄ λ«νλμ§ νμΈ
fmt.Println(a, ok) // 0 false: μ±λμ΄ λ«νμ
```
## π―λ³΄λ΄κΈ° μ μ© λ° λ°κΈ° μ μ© μ±λ

- λ³΄λ΄κΈ° μ μ© μ±λκ³Ό λ°κΈ° μ μ© μ±λμ κ°μ νλ¦μ΄ ν λ°©ν₯μΌλ‘ κ³ μ λ μ±λμ΄λ€

- λ€μμ 0λΆν° 4κΉμ§ μ±λμ κ°μ λ³΄λ΄κ³ , λ€μ μ±λμμ κ°μ κΊΌλ΄μ μΆλ ₯νλ€. κ·Έλ¦¬κ³  λ°λ³΅λ¬Έμ΄ λλ λ€ μ±λμ 100μ λ³΄λΈ λ€ λ€μ κΊΌλ΄μ μΆλ ₯νλ€
```
package main

import "fmt"

func producer(c chan<- int) { // λ³΄λ΄κΈ° μ μ© μ±λ
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100           // μ±λμ κ°μ λ³΄λ

	//fmt.Println(<-c) // μ±λμμ κ°μ κΊΌλ΄λ©΄ μ»΄νμΌ μλ¬
}

func consumer(c <-chan int) { // λ°κΈ° μ μ© μ±λ
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c) // μ±λμ κ°μ κΊΌλ

	// c <- 1        // μ±λμ κ°μ λ³΄λ΄λ©΄ μ»΄νμΌ μλ¬
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
```
```
0
1
2
3
4
100
```
- λ³΄λ΄κΈ° μ μ© λ° λ°κΈ° μ μ© μ±λμ μ±λ μ λ€λ‘ <- μ°μ°μλ₯Ό λΆμ¬μ λ§λ­λλ€. λ³΄ν΅ ν¨μμ λ§€κ°λ³μλ‘ μ¬μ©νκ±°λ, κ΅¬μ‘°μ²΄μ νλλ‘ μ¬μ©νλ€.

- λ³΄λ΄κΈ° μ μ©(send-only): chan<- μλ£ν νμμ΄λ€. c chan<- intλ intν λ³΄λ΄κΈ° μ μ© μ±λ cλ₯Ό λ»νλ€. λ³΄λ΄κΈ° μ μ© μ±λμ κ°μ λ³΄λΌ μλ§ μμΌλ©° κ°μ κ°μ Έμ€λ €κ³  νλ©΄ μ»΄νμΌ μλ¬κ° λ°μνλ€.
- λ°κΈ° μ μ©(receive-only): <-chan μλ£ν νμμ΄λ€. c <-chan intλ intν λ°κΈ° μ μ© μ±λ cλ₯Ό λ»νλ€. λ°κΈ° μ μ© μ±λμ range ν€μλ λλ <- μ±λ νμμΌλ‘ κ°μ κΊΌλΌ μλ§ μμΌλ©° κ°μ λ³΄λ΄λ €κ³  νλ©΄ μ»΄νμΌ μλ¬κ° λ°μνλ€.
- chan ν€μλλ₯Ό κΈ°μ€μΌλ‘ <- (νμ΄ν)κ° λΆμ λ°©ν₯μ λ³΄λ©΄ λ³΄λ΄κΈ° μ μ©μΈμ§ λ°κΈ° μ μ©μΈμ§ μ μ μλ€. μ¦ chan<-μ chan ν€μλλ‘ <-κ° λ€μ΄κ°λ―λ‘ λ³΄λ΄κΈ° μ μ©, <-chanμ chan ν€μλμμ <-κ° λμ€κ³  μμΌλ―λ‘ λ°κΈ° μ μ© μ±λμ΄λ€.

- μ¬κΈ°μλ producer ν¨μλ λ§€κ°λ³μλ‘ λ³΄λ΄κΈ° μ μ© μ±λμ μ¬μ©νκ³ , consumer ν¨μλ λ§€κ°λ³μλ‘ λ°κΈ° μ μ© μ±λμ μ¬μ©νλ€. λ°λΌμ producer ν¨μλ κ°μ λ³΄λ΄κΈ°λ§ νκ³ , consumer ν¨μλ κ°μ κΊΌλ΄κΈ°λ§ νλ€.

- μ΄λ²μλ μ±λμ λ¦¬ν΄κ°μΌλ‘ μ¬μ©νμ¬, λ€μμ λ μλ₯Ό λν λ€ μ±λλ‘ λ¦¬ν΄νλ€.

```
package main

import "fmt"

//                    β ν¨μμ λ¦¬ν΄ κ°μ int ν λ°κΈ° μ μ© μ±λ
func sum(a, b int) <-chan int {
	out := make(chan int) // μ±λ μμ±
	go func() {
		out <- a + b // μ±λμ aμ bμ ν©μ λ³΄λ
	}()
	return out           // μ±λ λ³μ μμ²΄λ₯Ό λ¦¬ν΄
}

func main() {
	c := sum(1, 2)   // μ±λμ λ¦¬ν΄κ°μΌλ‘ λ°μμ cμ λμ

	fmt.Println(<-c) // 3: μ±λμμ κ°μ κΊΌλ
}
```
- sum ν¨μλ λ°κΈ° μ μ© μ±λμ λ¦¬ν΄νλλ‘ λ§λ€μκ³ . μ±λμ λ¦¬ν΄νλ €λ©΄ λ¨Όμ  make ν¨μλ‘ μ±λμ μμ±νλ€ κ·Έλ¦¬κ³  κ³ λ£¨ν΄ μμμ μ±λμ κ°μ λ³΄λΈ λ€ κ³ λ£¨ν΄ λ°κΉ₯μμ μ±λμ λ¦¬ν΄νλ€

- sum ν¨μλ₯Ό μ¬μ©νμ¬ μ±λμ λ¦¬ν΄κ°μΌλ‘ λ°μμΌλ©΄ <-cμ²λΌ κ°μ κΊΌλ΄λ©΄ λλ€.

- μ΄λ²μλ μ±λλ§ μ¬μ©νμ¬ κ°μ λν΄λ³΄κ² λ€.
```
package main

import "fmt"

//                    β ν¨μμ λ¦¬ν΄ κ°μ int ν λ°κΈ° μ μ© μ±λ
func num(a, b int) <-chan int {
	out := make(chan int) // intν μ±λ μμ±
	go func() {
		out <- a   // μ±λμ aμ κ°μ λ³΄λ
		out <- b   // μ±λμ bμ κ°μ λ³΄λ
		close(out) // μ±λμ λ«μ
	}()
	return out // μ±λ λ³μ μμ²΄λ₯Ό λ¦¬ν΄
}

//            β ν¨μμ λ§€κ°λ³μλ intν λ°κΈ° μ μ© μ±λ
func sum(c <-chan int) <-chan int {
//                        β ν¨μμ λ¦¬ν΄ κ°μ intν λ°κΈ° μ μ© μ±λ
	out := make(chan int) // intν μ±λ μμ±
	go func() {
		r := 0
		for i := range c { // rangeλ₯Ό μ¬μ©νμ¬ μ±λμ΄ λ«ν λκΉμ§ κ°μ κΊΌλ
			r = r + i  // κΊΌλΈ κ°μ λͺ¨λ λν¨
		}
		out <- r           // λν κ²°κ³Όλ₯Ό μ±λμ λ³΄λ
	}()
	return out // μ±λ λ³μ μμ²΄λ₯Ό λ¦¬ν΄
}

func main() {
	c := num(1, 2) // 1κ³Ό 2κ° λ€μ΄μλ μ±λμ΄ λ¦¬ν΄λ¨
	out := sum(c)  // μ±λ cλ₯Ό λ§€κ°λ³μμ λκ²¨μ λͺ¨λ λν¨, λν κ°μ΄ λ€μ΄μλ out μ±λμ λ¦¬ν΄

	fmt.Println(<-out) // 3: out μ±λμμ κ°μ κΊΌλ
}
```
```
func num(a, b int) <-chan int {
	out := make(chan int) // intν μ±λ μμ±
	go func() {
		out <- a   // μ±λμ aμ κ°μ λ³΄λ
		out <- b   // μ±λμ bμ κ°μ λ³΄λ
		close(out) // μ±λμ λ«μ
	}()
	return out // μ±λ λ³μ μμ²΄λ₯Ό λ¦¬ν΄
}
```
- μ΄μ  μ±λμλ μ«μ λ κ°κ° μ μ₯λμ΄ μλ€. κ·Έλ¦¬κ³  close ν¨μλ‘ μ±λμ λ«μμ range ν€μλμ λ°λ³΅μ΄ λλλλ‘ νλ€.
- λ€μκ³Ό κ°μ΄ sum ν¨μμμλ range ν€μλλ‘ μ±λμμ κ°μ λ κ° κΊΌλ΄μ λͺ¨λ λνλ€. κ·Έλ¦¬κ³  λν κ°μ λ¦¬ν΄μ© μ±λμ λ³΄λΈλ€.
```
func sum(c <-chan int) <-chan int {
	out := make(chan int) // intν μ±λ μμ±
	go func() {
		r := 0
		for i := range c { // rangeλ₯Ό μ¬μ©νμ¬ μ±λμ΄ λ«ν λκΉμ§ κ°μ κΊΌλ
			r = r + i  // κΊΌλΈ κ°μ λͺ¨λ λν¨
		}
		out <- r           // λν κ²°κ³Όλ₯Ό μ±λμ λ³΄λ
	}()
	return out // μ±λ λ³μ μμ²΄λ₯Ό λ¦¬ν΄
}
```
- num ν¨μκ° λ¦¬ν΄ν μ±λμλ 1κ³Ό 2κ° λ€μ΄μλ€. κ·Έλ¦¬κ³  μ΄ μ±λμ sum ν¨μμ λ£μ΄μ£Όλ©΄ κ°μ΄ λͺ¨λ λν΄μ§λ€. λ§μ§λ§μΌλ‘ λν κ°μ sum ν¨μκ° λ¦¬ν΄ν μ±λμμ κΊΌλ΄λ©΄ λλ€
```
c := num(1, 2) // 1κ³Ό 2κ° λ€μ΄μλ μ±λμ΄ λ¦¬ν΄λ¨
out := sum(c)  // μ±λ cλ₯Ό λ§€κ°λ³μμ λκ²¨μ λͺ¨λ λν¨, λν κ°μ΄ λ€μ΄μλ out μ±λμ λ¦¬ν΄

fmt.Println(<-out) // 3: out μ±λμμ κ°μ κΊΌλ
```
## π―μλ νΈ μ¬μ©
- Go μΈμ΄λ μ¬λ¬ μ±λμ μμ½κ² μ¬μ©ν  μ μλλ‘ select λΆκΈ°λ¬Έμ μ κ³΅νλ€.
- select { case <- μ±λ: μ½λ }

```
select {
case <-μ±λ1:
	// μ±λ1μ κ°μ΄ λ€μ΄μμ λ μ€νν  μ½λλ₯Ό μμ±ν©λλ€.
case <-μ±λ2:
	// μ±λ2μ κ°μ΄ λ€μ΄μμ λ μ€νν  μ½λλ₯Ό μμ±ν©λλ€.
default:
	// λͺ¨λ  caseμ μ±λμ κ°μ΄ λ€μ΄μ€μ§ μμμ λ μ€νν  μ½λλ₯Ό μμ±ν©λλ€.
}
```
- select λΆκΈ°λ¬Έμ switch λΆκΈ°λ¬Έκ³Ό λΉμ·νμ§λ§ select ν€μλ λ€μ κ²μ¬ν  λ³μλ₯Ό λ°λ‘ μ§μ νμ§ μμΌλ©° κ° μ±λμ κ°μ΄ λ€μ΄μ€λ©΄ ν΄λΉ caseκ° μ€νλλ€.(close ν¨μλ‘ μ±λμ λ«μμ λλ caseκ° μ€νλλ€.). κ·Έλ¦¬κ³  λ³΄ν΅ selectλ₯Ό κ³μ μ²λ¦¬ν  μ μλλ‘ forλ‘ λ°λ³΅ν΄μ€λ€(λ°λ³΅νμ§ μμΌλ©΄ ν λ²λ§ μ€ννκ³  λλΈλ€).

- switch λΆκΈ°λ¬Έκ³Ό λ§μ°¬κ°μ§λ‘ select λΆκΈ°λ¬Έλ default μΌμ΄μ€λ₯Ό μ§μ ν  μ μμΌλ©° caseμ μ§μ λ μ±λμ κ°μ΄ λ€μ΄μ€μ§ μμμ λ μ¦μ μ€νλλ€. λ¨, defaultμ μ μ ν μ²λ¦¬λ₯Ό νμ§ μμΌλ©΄ CPU μ½μ΄λ₯Ό λͺ¨λ μ μ νλ―λ‘ μ£Όμνλ€
λ€μμ μ±λ 2κ°λ₯Ό μμ±νκ³  100λ°λ¦¬μ΄, 500λ°λ¦¬μ΄ κ°κ²©μΌλ‘ μ«μμ λ¬Έμμ΄μ λ³΄λΈ λ€ κΊΌλ΄μ μΆλ ₯νλ€
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // intν μ±λ μμ±
	c2 := make(chan string) // string μ±λ μμ±

	go func() {
		for {
			c1 <- 10                           // μ±λ c1μ 10μ λ³΄λΈ λ€
			time.Sleep(100 * time.Millisecond) // 100 λ°λ¦¬μ΄ λκΈ°
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // μ±λ c2μ Hello, world!λ₯Ό λ³΄λΈ λ€
			time.Sleep(500 * time.Millisecond) // 500 λ°λ¦¬μ΄ λκΈ°
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:                // μ±λ c1μ κ°μ΄ λ€μ΄μλ€λ©΄ κ°μ κΊΌλ΄μ iμ λμ
				fmt.Println("c1 :", i) // i κ°μ μΆλ ₯
			case s := <-c2:                // μ±λ c2μ κ°μ΄ λ€μ΄μλ€λ©΄ κ°μ κΊΌλ΄μ sμ λμ
				fmt.Println("c2 :", s) // s κ°μ μΆλ ₯
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10μ΄ λμ νλ‘κ·Έλ¨ μ€ν
}
```
```
c2 : Hello, world!
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c2 : Hello, world!
c1 : 10
... (μλ΅)
```

- μ€νμ ν΄λ³΄λ©΄ selectμμ λ²κ°μκ°λ©΄μ 10κ³Ό Hello, world!κ° μΆλ ₯λλ€  μ±λ c2μ Hello, world!λ₯Ό λ³΄λΈ μͺ½μμλ 500 λ°λ¦¬μ΄ λκΈ°νκ³ , μ±λ c1μ 10μ λ³΄λΈ μͺ½μμλ 100 λ°λ¦¬μ΄ λκΈ°νλ―λ‘ 10μ΄ λ λ§μ΄ μΆλ ₯λλ€
- caseμμλ case i := <-c1:μ²λΌ μ±λμμ κ°μ κΊΌλΈ λ€ λ³μμ λ°λ‘ μ μ₯ν  μ μλ€. λ§μ½ κΊΌλΈ κ°μ μ¬μ©νμ§ μλλ€λ©΄ case <-c1:μ²λΌ λ³μλ₯Ό μλ΅ν΄λ λλ€.

```
select {
case i := <-c1:                // μ±λ c1μ κ°μ΄ λ€μ΄μλ€λ©΄ κ°μ κΊΌλ΄μ iμ λμ
	fmt.Println("c1 :", i) // i κ°μ μΆλ ₯
case s := <-c2:                // μ±λ c2μ κ°μ΄ λ€μ΄μλ€λ©΄ κ°μ κΊΌλ΄μ sμ λμ
	fmt.Println("c2 :", s) // s κ°μ μΆλ ₯
}
```
- time.After ν¨μλ₯Ό μ¬μ©νλ©΄ μκ° μ ν μ²λ¦¬λ₯Ό ν  μ μλ€. time.Afterλ νΉμ  μκ°μ΄ μ§λλ©΄ νμ¬ μκ°μ μ±λλ‘ λ³΄λΈλ€.
```
select {
case i := <-c1:
	fmt.Println("c1 : ", i)
case s := <-c2:
	fmt.Println("c2 : ", s)
case <-time.After(50 * time.Millisecond): // 50 λ°λ¦¬μ΄ ν νμ¬ μκ°μ΄ λ΄κΈ΄ μ±λμ΄ λ¦¬ν΄λ¨
	fmt.Println("timeout")
}
```
- μ΄μ²λΌ caseμμλ time.Afterμ κ°μ΄ λ°κΈ° μ μ© μ±λμ λ¦¬ν΄νλ ν¨μλ₯Ό μ¬μ©ν  μ μλ€.

- select λΆκΈ°λ¬Έμ μ±λμ κ°μ λ³΄λΌ μλ μλ€.

- case μ±λ <- κ°: μ½λ
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // intν μ±λ μμ±
	c2 := make(chan string) // string μ±λ μμ±

	go func() {
		for {
			i := <-c1                          // μ±λ c1μμ κ°μ κΊΌλΈ λ€ iμ λμ
			fmt.Println("c1 :", i)             // i κ°μ μΆλ ₯
			time.Sleep(100 * time.Millisecond) // 100 λ°λ¦¬μ΄ λκΈ°
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // μ±λ c2μ Hello, world!λ₯Ό λ³΄λ
			time.Sleep(500 * time.Millisecond) // 500 λ°λ¦¬μ΄ λκΈ°
		}
	}()

	go func() {
		for { // λ¬΄ν λ£¨ν
			select {
			case c1 <- 10:                 // λ§€λ² μ±λ c1μ 10μ λ³΄λ
			case s := <-c2:                // c2μ κ°μ΄ λ€μ΄μμ λλ κ°μ κΊΌλΈ λ€ sμ λμ
				fmt.Println("c2 :", s) // s κ°μ μΆλ ₯
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10μ΄ λμ νλ‘κ·Έλ¨ μ€ν
}
```
```
c2 : Hello, world!
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c2 : Hello, world!
c1 : 10
... (μλ΅)
```
- select λΆκΈ°λ¬Έμμ μ±λμ κ°μ λ³΄λ΄λ caseκ° μλ€λ©΄ ν­μ κ°μ λ³΄λΈλ€. νμ§λ§ μ±λμ κ°μ΄ λ€μ΄μμ λλ κ°μ λ°λ caseκ° μ€νλλ€.
- μ¬κΈ°μλ selectμμ λ§€λ² μ±λ c1μ κ°μ λ³΄λ΄μ§λ§ μ±λ c2μ κ°μ΄ λ€μ΄μ€λ©΄ c2μμ κ°μ κΊΌλ΄μ μΆλ ₯νλ€.
```
for { // λ¬΄ν λ£¨ν
	select {
	case c1 <- 10:             // λ§€λ² μ±λ c1μ 10μ λ³΄λ
	case s := <-c2:            // c2μ κ°μ΄ λ€μ΄μμ λλ κ°μ κΊΌλΈ λ€ sμ λμ
		fmt.Println("c2 :", s) // s κ°μ μΆλ ₯
	}
}
```
- μμ μμλ λ³΄λ΄λ μ±λκ³Ό λ°λ μ±λμ λ κ° μ¬μ©νμ§λ§ λ€μκ³Ό κ°μ΄ μ±λ c1 ν κ°λ‘ selectμμ κ°μ λ³΄λ΄κ±°λ λ°μ μλ μλ€.
```
μμ΅λλ€.

c1 := make(chan int) // intν μ±λ μμ±

go func() {
	for {
		i := <-c1                          // μ±λ c1μμ κ°μ κΊΌλΈ λ€ iμ λμ
		fmt.Println("c1 :", i)             // i κ°μ μΆλ ₯
		time.Sleep(100 * time.Millisecond) // 100 λ°λ¦¬μ΄ λκΈ°
	}
}()

go func() {
	for {
		c1 <- 20                           // μ±λ c1μ 20μ λ³΄λ
		time.Sleep(500 * time.Millisecond) // 100 λ°λ¦¬μ΄ λκΈ°
	}
}()

go func() {
	for { // λ¬΄ν λ£¨ν
		select {                       // μ±λ c1 ν κ°λ‘ κ°μ λ³΄λ΄κ±°λ λ°μ
		case c1 <- 10:                 // λ§€λ² μ±λ c1μ 10μ λ³΄λ
		case i := <-c1:                // c1μ κ°μ΄ λ€μ΄μμ λλ κ°μ κΊΌλΈ λ€ iμ λμ
			fmt.Println("c1 :", i) // i κ°μ μΆλ ₯
		}
	}
}()

time.Sleep(10 * time.Second) // 10μ΄ λμ νλ‘κ·Έλ¨ μ€ν
```
```
c1 : 20
c1 : 10
c1 : 10
c1 : 10
c1 : 10
c1 : 20
c1 : 10
... (μλ΅)
```
- μ¬κΈ°μλ λ§€λ² μ±λμ κ°μ λ³΄λ΄μ§λ§, select λΆκΈ°λ¬Έμ΄ μλ λ€λ₯Έ μͺ½μμ μ±λμ κ°μ λ³΄λ΄μ κ°μ΄ λ€μ΄μλ€λ©΄ κ°μ λ°λ caseκ° μ€νλλ€.

