# π©π»βπλ¦¬νλ μ

- λ¦¬νλ μμ μ€ν μμ (Runtime, λ°νμ)μ μΈν°νμ΄μ€λ κ΅¬μ‘°μ²΄ λ±μ νμ μ λ³΄λ₯Ό μ»μ΄λ΄κ±°λ κ²°μ νλ κΈ°λ₯μ΄λ€. λ¦¬νλ μμ Java, C#μ²λΌ κ°μ λ¨Έμ  μμμ μ€νλλ μΈμ΄λ Python, Ruby λ±μ μ€ν¬λ¦½νΈ μΈμ΄μμ μ£Όλ‘ μ¬μ©νμλ€. λ§μ°¬κ°μ§λ‘ Go μΈμ΄λ κΈ°λ³Έ ν¨ν€μ§μμ λ¦¬νλ μμ μ κ³΅νλ€

κ°λ¨νκ² λ³μμ κ΅¬μ‘°μ²΄μ νμμ νμν΄λ³΄κ² μ΅λλ€.

```
package main

import (
	"fmt"
	"reflect"
)

type Data struct { // κ΅¬μ‘°μ²΄ μ μ
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // int: reflect.TypeOfλ‘ μλ£ν μ΄λ¦ μΆλ ₯

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s)) // string: reflect.TypeOfλ‘ μλ£ν μ΄λ¦ μΆλ ₯

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f)) // float32: reflect.TypeOfλ‘ μλ£ν μ΄λ¦ μΆλ ₯

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // main.Data: reflect.TypeOfλ‘ κ΅¬μ‘°μ²΄ μ΄λ¦ μΆλ ₯
}
```
```
int
string
float32
main.Data
```
- reflect.TypeOf ν¨μλ₯Ό μ¬μ©νλ©΄ μΌλ° μλ£νμ΄λ κ΅¬μ‘°μ²΄μ νμμ μ μ μλ€. μ¬κΈ°μλ int, string, float32 ν λ³μμ μλ£νμ΄ μΆλ ₯λλ€. λ§μ°¬κ°μ§λ‘ κ΅¬μ‘°μ²΄λ νμμ μμλΌ μ μλλ° Data κ΅¬μ‘°μ²΄λ main ν¨ν€μ§ μμ μν΄μκΈ° λλ¬Έμ main.Dataλ‘ λμ¨λ€.

- λ¦¬νλ μμΌλ‘ λ³μμ νμλΏλ§ μλλΌ κ°μ λν μμΈν μ λ³΄λ μ»μ΄μ¬ μ μλ€.
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)  // fμ νμ μ λ³΄λ₯Ό tμ μ μ₯
	v := reflect.ValueOf(f) // fμ κ° μ λ³΄λ₯Ό vμ μ μ₯

	fmt.Println(t.Name())                    // float64: μλ£ν μ΄λ¦ μΆλ ₯
	fmt.Println(t.Size())                    // 8: μλ£ν ν¬κΈ° μΆλ ₯
	fmt.Println(t.Kind() == reflect.Float64) // true: μλ£ν μ’λ₯λ₯Ό μμλ΄μ΄ 
                                                 // reflect.Float64μ λΉκ΅
	fmt.Println(t.Kind() == reflect.Int64)   // false: μλ£ν μ’λ₯λ₯Ό μμλ΄μ΄ reflect.Int64μ λΉκ΅

	fmt.Println(v.Type())                    // float64: κ°μ΄ λ΄κΈ΄ λ³μμ μλ£ν μ΄λ¦ μΆλ ₯
	fmt.Println(v.Kind() == reflect.Float64) // true: κ°μ΄ λ΄κΈ΄ λ³μμ μλ£ν μ’λ₯λ₯Ό 
                                                 // μμλ΄μ΄ reflect.Float64μ λΉκ΅
	fmt.Println(v.Kind() == reflect.Int64)   // false: κ°μ΄ λ΄κΈ΄ λ³μμ μλ£ν μ’λ₯λ₯Ό 
                                                 // μμλ΄μ΄ reflect.Int64μ λΉκ΅
	fmt.Println(v.Float())                   // 1.3: κ°μ μ€μνμΌλ‘ μΆλ ₯
}
```
```
float64
8
true
false
float64
true
false
1.3
```
- reflect.TypeOf ν¨μλ‘ float64 λ³μμ νμ μ λ³΄ reflect.Typeλ₯Ό μ»μ΄μλ€. νμ μ λ³΄λ‘λ νμμ μ΄λ¦, λ³μ(νμ)μ ν¬κΈ° λ±μ μ μ μμΌλ©° Kind ν¨μλ₯Ό μ¬μ©νλ©΄ μμ νμμΌλ‘λ νμ μ’λ₯λ μ μ μλ€.

- reflect.ValueOf ν¨μλ‘ float64 λ³μμ κ° μ λ³΄ reflect.Valueλ₯Ό μ»μ΄μ€λ©΄ νμ μ λ³΄, νμ μ’λ₯, λ³μμ μ μ₯λ κ°μ μ μ μλ€. μ¬κΈ°μ λ³μκ° float64λΌλ©΄ v.Float(), intλΌλ©΄ v.Int(), stringμ΄λΌλ©΄ v.String()κ³Ό κ°μ΄ κ° νμμ λ§λ ν¨μλ₯Ό μ¬μ©νλ©΄ λ³μμ μ μ₯λ κ°μ κ°μ Έμ¬ μ μλ€.

## π―κ΅¬μ‘°μ²΄ νκ·Έ κ°μ Έμ€κΈ° 
```
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"μ΄λ¦" tag2:"Name"` // κ΅¬μ‘°μ²΄μ νκ·Έ μ€μ 
	age  int    `tag1:"λμ΄" tag2:"Age"`  // κ΅¬μ‘°μ²΄μ νκ·Έ μ€μ 
}

func main() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2")) // true μ΄λ¦ Name

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2")) // true λμ΄ Age
}
```
```
true μ΄λ¦ Name
true λμ΄ Age
```
- κ΅¬μ‘°μ²΄ νλμ νκ·Έλ `νκ·Έλͺ:"λ΄μ©"` νμμΌλ‘ μ§μ νλ€. νκ·Έλ λ¬Έμμ΄ ννμ΄λ©° λ¬Έμμ΄ μμ " " (λ°μ΄ν)κ° ν¬ν¨λλ―λ‘ ` ` (λ°±μΏΌνΈ)λ‘ κ°μΈμ€λ€. μ¬λ¬ κ°λ₯Ό μ§μ ν  λλ κ³΅λ°±μΌλ‘ κ΅¬λΆν΄μ€λ€.

- reflect.TypeOf ν¨μμ κ΅¬μ‘°μ²΄ μΈμ€ν΄μ€λ₯Ό λ£μΌλ©΄ reflect.Typeμ΄ λ¦¬ν΄λλ€. μ¬κΈ°μ FieldByName ν¨μμ κ΅¬μ‘°μ²΄μ νλ μ΄λ¦μ μ§μ νλ©΄ νλ μ λ³΄λ₯Ό μ»μ μ μλ€. λ λ²μ§Έ λ¦¬ν΄κ°μ ν΄λΉ μ΄λ¦μΌλ‘ νλκ° μ‘΄μ¬νλμ§ μ¬λΆμ΄λ€. νλ μ λ³΄λ₯Ό μ»μ λ€μλ name.Tag.Get(βtag1β)κ³Ό κ°μ΄ νκ·Έλ₯Ό κ°μ Έμ¬ μ μλ€.

## π―ν¬μΈν°μ μΈν°νμ΄μ€μ κ° κ°μ Έμ€κΈ° 
```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))               // *int
	fmt.Println(reflect.ValueOf(a))              // <*int Value>
	fmt.Println(reflect.ValueOf(a).Int())        // λ°νμ μλ¬
	fmt.Println(reflect.ValueOf(a).Elem())       // <int Value>
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 1

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))         // int
	fmt.Println(reflect.ValueOf(b))        // <int Value>
	fmt.Println(reflect.ValueOf(b).Int())  // 1
	fmt.Println(reflect.ValueOf(b).Elem()) // λ°νμ μλ¬
}
```
- ν¬μΈν°λ μΌλ° λ³μμλ λ€λ₯΄κ² κ°μ κ°μ Έμ€λ €λ©΄ reflect.ValueOf ν¨μλ‘ κ° μ λ³΄ reflect.Valueλ₯Ό μ»μ΄μ¨ λ€ λ€μ Elem ν¨μλ‘ κ° μ λ³΄λ₯Ό κ°μ ΈμμΌ νλ€. κ·Έλ¦¬κ³  λ³μμ νμμ λ§λ Int, Float, String λ±μ ν¨μλ‘ κ°μ κ°μ Έμ¨λ€.

- μ¬κΈ°μλ int ν¬μΈν° aμ κ° μ λ³΄μμ λ°λ‘ Int ν¨μλ‘ κ°μ κ°μ Έμ€λ €λ©΄ λ°νμ μλ¬κ° λ°μνλ€. λ°λΌμ Elem ν¨μλ‘ ν¬μΈν°μ λ©λͺ¨λ¦¬μ μ μ₯λ μ€μ  κ° μ λ³΄λ₯Ό κ°μ ΈμμΌ νλ€.

- λΉ μΈν°νμ΄μ€ bμ 1μ λμνλ©΄ νμ μ λ³΄λ intμ΄κ³  κ° μ λ³΄λ intμ΄λ€. λ°λΌμ μΈν°νμ΄μ€μ κ°μ κ°μ Έμ€λ €λ©΄ λ³μ νμμ λ§λ Int, Float, String λ±μ ν¨μλ₯Ό μ¬μ©νλ©΄ λλ€.

## π―λμ μΌλ‘ ν¨μ μμ±

- λ¦¬νλ μμ μ¬μ©νμ¬ λμ μΌλ‘ ν¨μλ₯Ό λ§λ€μ΄λ΄λ λ°©λ²

- λ¨Όμ  reflect.MakeFunc ν¨μλ₯Ό μ¬μ©νλ λ°©λ²μ΄λ€.
```
package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value { // λ§€κ°λ³μμ λ¦¬ν΄κ°μ λ°λμ 
                                               // []reflect.Valueλ₯Ό μ¬μ©
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	var hello func() // ν¨μλ₯Ό λ΄μ λ³μ μ μΈ

	fn := reflect.ValueOf(&hello).Elem() // helloμ μ£Όμλ₯Ό λκΈ΄ λ€ ElemμΌλ‘ κ° μ λ³΄λ₯Ό κ°μ Έμ΄

	v := reflect.MakeFunc(fn.Type(), h) // hμ ν¨μ μ λ³΄λ₯Ό μμ±

	fn.Set(v) // helloμ κ° μ λ³΄μΈ fnμ hμ ν¨μ μ λ³΄ vλ₯Ό μ€μ νμ¬ ν¨μλ₯Ό μ°κ²°

	hello()
}
```
```
Hello, world!
```
- λ¨Όμ  λ§€κ°λ³μμ λ¦¬ν΄κ°μΌλ‘ []reflect.Valueλ₯Ό μ¬μ©νλ ν¨μ hλ₯Ό μ μνλ€. μ΄ ν¨μμμ μ€μ  κΈ°λ₯μ κ΅¬ννλ©° μ¬κΈ°μλ Hello, world!λ₯Ό μΆλ ₯νλ€. λ¦¬νλ μμΌλ‘ μμ±νλ ν¨μλ λ°λμ []reflect.Valueλ₯Ό λ§€κ°λ³μλ‘ λ°κ³  λ¦¬ν΄κ°μΌλ‘ μ¬μ©ν΄μΌ νλ€.

- μ΄μ  ν¨μλ₯Ό λ΄μ μ μλ λ³μλ₯Ό helloλ₯Ό μ μΈνλ€. κ·Έλ¦¬κ³  reflect.ValueOf ν¨μμ helloμ μ£Όμ(λ νΌλ°μ€)λ₯Ό λκΈ΄ λ€ Elem ν¨μλ‘ μ€μ  κ° μ λ³΄λ₯Ό κ°μ Έμ¨λ€. κ·Έ λ€μμλ reflect.MakeFunc(fn.Type(), h)κ³Ό κ°μ΄ ν¨μ hλ₯Ό λ£μ΄μ ν¨μ μ λ³΄λ₯Ό μμ±νλ€. λ§μ§λ§μΌλ‘ fn.Set(v)μ κ°μ΄ helloμ κ° μ λ³΄μΈ fnμ ν¨μ μ λ³΄ vλ₯Ό μ€μ νμ¬ ν¨μλ₯Ό μ°κ²°ν΄μ€λ€.
λͺ¨λ  κ³Όμ μ λ§μΉ λ€ hello ν¨μλ₯Ό νΈμΆνλ©΄ Hello, world!κ° μΆλ ₯λλ€.. μ¦ hello ν¨μλ h ν¨μλ₯Ό μ΄μ©νμ¬ λμ μΌλ‘ μμ±λ ν¨μμ΄λ€. νμ§λ§ Hello, world!λ₯Ό μΆλ ₯νκΈ°μλ λ³΅μ‘νκΈ°λ§ νλ€,

- λμ  ν¨μ μμ±μ μ’λ μ λλ‘ νμ©νκΈ° μν΄ ν¨μ νλλ‘ μ μ, μ€μ, λ¬Έμμ΄ λνκΈ°λ₯Ό λͺ¨λ μ²λ¦¬νλ ν¨μλ₯Ό μμ±

- λ¨Όμ  λ κ°μ λνλ sum ν¨μλ₯Ό κ΅¬ν
```
...

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("νμμ΄ λ€λ¦λλ€.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

...
```

- λ¦¬νλ μμΌλ‘ μμ±λ  ν¨μμ΄λ―λ‘ λ§€κ°λ³μμ λ¦¬ν΄κ°μ νμμ []reflect.Valueμ΄λ€. κ·Έλ¬λ―λ‘ λ κ°μ λν  λ μ²« λ²μ§Έ λ§€κ°λ³μλ args[0], λ λ²μ§Έ λ§€κ°λ³μλ args[1]μ΄ λλ€.μ¬κΈ°μλ λ κ°μ νμ μ’λ₯λ₯Ό λΉκ΅ν΄μ λ€λ₯΄λ©΄ nilμ λ¦¬ν΄μ νλλ‘ κ΅¬ννλ€. νμ§λ§ κ΅¬ννκΈ°μ λ°λΌμ float32μ int32λ₯Ό λνμ¬ float64λ₯Ό λ§λ€ μλ μκ³ , stringκ³Ό int32λ₯Ό λνμ¬ stringμ λ§λ€ μλ μλ€.

- switch λΆκΈ°λ¬ΈμΌλ‘ νμ μ’λ₯λ₯Ό κ΅¬λΆνκ³ , κ°μ μ’λ₯μ νμμ λ¬Άμ΄μ μ²λ¦¬νλ€.. return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}μ κ°μ΄ νμ μ’λ₯μ λ§κ² λ³μμμ κ°μ κΊΌλ΄μ λνκ³ , []reflect.Value μ¬λΌμ΄μ€λ‘ λ§λ€μ΄μ λ¦¬ν΄νλ€.

- μ΄μ  λ§€κ°λ³μλ‘ λ°μ λ³μμ ν¨μλ₯Ό μ°κ²°ν΄μ£Όλ ν¨μλ₯Ό κ΅¬ν
```
...

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

...
```
- λΉ μΈν°νμ΄μ€λ‘ λ§€κ°λ³μλ₯Ό λ°μ λ€ reflect.ValueOf ν¨μμ Elem ν¨μλ₯Ό μ¬μ©νμ¬ μ€μ  κ° μ λ³΄λ₯Ό κ°μ Έμ¨λ€. μ΄λ λ§€κ°λ³μλ‘ λ€μ΄μ¬ κ°μ ν¨μλ₯Ό μ μ₯ν  μ μλ λ³μμ΄λ€. κ·Έ λ€μμλ reflect.MakeFunc(fn.Type(), sum)κ³Ό κ°μ΄ ν¨μ sumμ λ£μ΄μ ν¨μ μ λ³΄λ₯Ό μμ±, λ§μ§λ§μΌλ‘ fn.Set(v)μ κ°μ΄ fptrμ κ° μ λ³΄μΈ fnμ ν¨μ μ λ³΄ vλ₯Ό μ€μ νμ¬ ν¨μλ₯Ό μ°κ²°ν΄μ€λ€

μ΄λ κ² νλ©΄ λ§€κ°λ³μ, λ¦¬ν΄κ° μλ£νμ ννκ° λ€μν ν¨μλ₯Ό λμ μΌλ‘ μ°κ²°μν¬ μ μλ€.

-  main ν¨μμμ ν¨μλ₯Ό μ μ₯ν  μ μλ λ³μλ₯Ό μ μΈ, ν¨μ λͺ¨μμ νκ³  μμ§λ§ μ€μ  ν¨μκ° μ μλμ΄ μμ§ μκΈ° λλ¬Έμ νΈμΆμ ν  μ μλ μνμ΄λ€.
```
...

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

...
}

```
- λ κ°μ λνλ―λ‘ λ§€κ°λ³μλ λ κ°μ΄λ©° κ° νμλ³λ‘ μ μΈνλ€. μ¬κΈ°μ λ¦¬ν΄κ°μ ν¬κΈ°κ° κ°μ₯ ν° νμμ μ¬μ©νλ€. μ¦ μ μλ int64, μ€μλ float64λ₯Ό μ¬μ©ν΄μΌ νλ€.

- μμμ μ μΈν λ³μμ λ νΌλ°μ€(ν¬μΈν°)λ₯Ό makeSum ν¨μμ λ£μ΄μ£Όλ©΄ μμ ν ννμ ν¨μκ° μμ±λλ€.
```
func main() {
...

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
```
```
3
5.599999904632568
Hello, world!
```
- intSum ν¨μλ μ μλ₯Ό λνκ³ , floatSum ν¨μλ μ€μλ₯Ό λνκ³ , stringSum ν¨μλ λ¬Έμμ΄μ λΆμΈλ€. ν¨μλ μΈ μ’λ₯μ§λ§ μ€μ λ‘λ μμμ κ΅¬νν sum ν¨μ νλλ‘ μ²λ¦¬λλ€. μ΄μ²λΌ λ¦¬νλ μμ λμ  ν¨μ μμ± κΈ°λ₯μ νμ©νλ©΄ νμλ³λ‘ ν¨μλ₯Ό μ¬λ¬ λ² κ΅¬ννμ§ μμλ λλ€.

-  μ μ²΄ μμ€ μ½λ
```
package main

import (
	"fmt"
	"reflect"
)

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("νμμ΄ λ€λ¦λλ€.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
```
