
# ๐ฏGo ๊ณต๋ถ ๋ฐ ์ ๋ฆฌ <img width="100" alt="แแณแแณแแตแซแแฃแบ 2021-11-24 แแฉแแฎ 12 28 54" src="https://user-images.githubusercontent.com/88940298/143169831-99d59153-1735-49a6-8d7c-d3768dfc9400.png">
## ๐ฑ๋ณ์ 
## ๐ฑ์์ const

- go์์๋ ํ์์ ์ ํ์ง ์์ผ๋ฉด go์์ ์์์ ํ์์ ์ง์ ํด์ค๋ค  
- go๊ฐ ๋ค๋ฅธ์ธ์ด๋ ๋ค๋ฅธ ์ ์ ํด๋์ค๊ฐ ์๋ค๋ ๊ฒ์ด๋ค.๋ฉ์๋๋ผ๊ณ  ํํํ๋ ๊ฒ๋ ํํ ๋ฐฉ์์ด ๋ค๋ฅธ ๋ฟ ์ผ๋ฐ ํจ์์ ์ ์ฌํ๊ฒ ํ์๋๋ค  
- go๋ ๋ฆฌํด ๊ฐ์ ์ฌ๋ฌ๊ฐ์ ๊ฐ์ผ๋ก ๋๋ฆด ์ ์๋ค.  
- go&go-routine go ํค์๋๋ ๋ฉํฐ์ค๋ ๋๋ฅผ ์ฌ์ฉํ ๋ ์ ์ฉํ๋ฉฐ ๋์์ฑ์ ๊ฐ์ง๋ค.  
- go get go get์ ํตํ์ฌ ๋ค๋ฅธ ํจํค์ง๋ฅผ ๋ค์ด๋ฐ์ ์ ์๋ค.์ป์ด์จ ํจํค์ง๋ GOPATH๋ก ๋ค์ด๋ก๋ ๋๋ค.GOPATH๋ ์์๊ณต๊ฐ์ด๋ผ๊ณ ๋ ํ๋ฉฐ ์ฌ์ฉ์๋ง๋ค ๋ณ๋๋ก ๊ด๋ฆฌ๋๋ค  
- go mod ์ฌ์ฉ์ ์ ์ ๋ชจ๋์ ๋ง๋ค๋ ์ฌ์ฉํ๋ค. example.comํด๋๋ก ์ง์ํ์ฌ go mod init example.com๋ช๋ น์ด๋ฅผ ํตํด ๋ชจ๋์ ์ ์ํ๋ค  
<img width="355" alt="แแณแแณแแตแซแแฃแบ 2021-10-27 แแฉแแฎ 2 11 44" src="https://user-images.githubusercontent.com/88940298/139003610-95bf570c-fa06-4959-815c-5e2bea14f2c7.png">

## ๐ฑ๋ฐ์ดํธ
byte์๋ ๋ณดํต 16์ง์, ๋ฌธ์ ๊ฐ์ผ๋ก ์ ์ฅํฉ๋๋ค. ์ค๋ฌด์์๋ ๋ฐ์ด๋๋ฆฌ ํ์ผ์์ ๋ฐ์ดํฐ๋ฅผ ์ฝ๊ณ  ์ฐ๊ฑฐ๋, ๋ฐ์ดํฐ๋ฅผ ์ํธํ, ๋ณตํธํํ  ๋ ์ฃผ๋ก ์ฌ์ฉ!
```
var num1 byte = 10   // 10์ง์ ์ ์ฅ
var num2 byte = 0x32 // 16์ง์ ์ ์ฅ
var num3 byte = 'a'  // ๋ฌธ์ ์ ์ฅ
```


## ๐ฑ๋ฃฌ 
rune์ ์ ๋์ฝ๋(UTF-8) ๋ฌธ์ ์ฝ๋๋ฅผ ์ ์ฅํ  ๋ ์ฌ์ฉํฉ๋๋ค. ' ' (์์ ๋ฐ์ดํ)๋ก ๋ฌถ์ด์ฃผ์ด์ผ ํ๋ฉฐ ๋ฌธ์ ๊ทธ๋๋ก ์ ์ฅํ๊ฑฐ๋ \u ๋๋ \U๋ฅผ ์ฌ์ฉํ์ฌ ์ ๋์ฝ๋ ๋ฌธ์ ์ฝ๋๋ก ์ ์ฅํ  ์๋ ์๋ค. ์ฌ๊ธฐ์ \U๋ ๊ฐ์ 16์ง์ 8์๋ฆฌ๋ก ๋ง์ถฐ์ฃผ์ด์ผ ํ๋ค
```
var r1 rune = 'ํ'
var r2 rune = '\ud55c'     // ํ
var r3 rune = '\U0000d55c' // ํ
```


## ๐ฑ์ซ์ ์ฐ์ฐ. 
์ซ์ ์ฐ์ฐ์๋ ๋ง์(+), ๋บ์(-), ๊ณฑ์(*), ๋๋์(/), ๋๋จธ์ง(%), ์ํํธ(<<, >>), ๋นํธ ๋ฐ์ (^) ์ฐ์ฐ์๋ฅผ ์ฌ์ฉํ  ์ ์๋ค
```
var num1 uint8 = 3
var num2 uint8 = 2

fmt.Println(num1 + num2)  // 5
fmt.Println(num1 - num2)  // 1
fmt.Println(num1 * num2)  // 6
fmt.Println(num1 / num2)  // 1
fmt.Println(num1 % num2)  // 1
fmt.Println(num1 << num2) // 12
fmt.Println(num1 >> num2) // 0
fmt.Println(^num1)        // 252: ๋นํธ ๋ฐ์  ์ฐ์ฐ์
```

## ๐ฑ์ค๋ฒํ๋ก์ฐ์ ์ธ๋ํ๋ก์ฐ. 
๊ฐ ์๋ฃํ์์ ์ ์ฅํ  ์ ์๋ ์ต๋ ํฌ๊ธฐ๋ฅผ ๋์ด์๋ฉด ์ค๋ฒํ๋ก์ฐ(Overflow), ์ต์ ํฌ๊ธฐ๋ณด๋ค ์์์ง๋ฉด ์ธ๋ํ๋ก์ฐ(Underflow)  
```
package main

import "fmt"
import "math"

func main() {
	var num1 uint8 = math.MaxUint8 + 1   // ์ค๋ฒํ๋ก์ฐ ์ปดํ์ผ ์๋ฌ ๋ฐ์
	var num2 uint16 = math.MaxUint16 + 1 // ์ค๋ฒํ๋ก์ฐ ์ปดํ์ผ ์๋ฌ ๋ฐ์
	var num3 uint32 = math.MaxUint32 + 1 // ์ค๋ฒํ๋ก์ฐ ์ปดํ์ผ ์๋ฌ ๋ฐ์
	var num4 uint64 = math.MaxUint64 + 1 // ์ค๋ฒํ๋ก์ฐ ์ปดํ์ผ ์๋ฌ ๋ฐ์
}
```
## ๐ฑ๋ฌธ์์ด ๊ธธ์ด. 
```
var s1 string = "ํ๊ธ"
var s2 string = "Hello"

fmt.Println(len(s1)) // 6: UTF-8 ์ธ์ฝ๋ฉ์ ๋ฐ์ดํธ ๊ธธ์ด์ด๋ฏ๋ก 6
fmt.Println(len(s2)) // 5: ์ํ๋ฒณ 5๊ธ์์ด๋ฏ๋ก 5
```

## ๐ฑ์ด๊ฑฐํ. 
## ๐ฑ[ํจํค์ง](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/package.md)
## ๐ฑ[fmt](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/fmt.md)
## ๐ฑ[if๋ฌธ](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/if.md)
## ๐ฑ[for๋ฌธ](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/for.md) 
## ๐ฑ[goto](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/goto.md) 
## ๐ฑ[Switch](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/switch.md) 
## ๐ฑ[Array](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/arr.md)
## ๐ฑ[Slice](https://github.com/kyunghyunHan/go_study/blob/61bf0e44d9f7a92929b2748d7dbd53704b869449/slice.md)
## ๐ฉ๐ปโ๐[Map](https://github.com/kyunghyunHan/go_study/blob/8b7e2994f1f700cd0a1ef59e9e980166b545c5c0/map.md)
## ๐ฉ๐ปโ๐[func](https://github.com/kyunghyunHan/go_study/blob/2ec57d0e5802105739988cedd75e60055ba6600e/func.md)
## ๐ฑ[closure](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/closure.md)
## ๐ฉ๐ปโ๐[defer](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/defer.md)
## ๐ฉ๐ปโ๐[panic](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/panic.md)
## ๐ฑ[Pointers](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/pointer.md)
## ๐ฉ๐ปโ๐[struct](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/struct.md)
## ๐ฉ๐ปโ๐[interface](https://github.com/kyunghyunHan/go_study/blob/a5272a26b9678bea8ee723b0d39b13a252379de2/interface.md)
## ๐ฉ๐ปโ๐[gorutine](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/goroutine.md)
## ๐๐ฅ[chanl](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/channel.md)
## ๐ฉ๐ปโ๐[Mutex](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/mutex.md)
## ๐[reflection](https://github.com/kyunghyunHan/go_study/blob/56bbd6e122d560cfab793f7bb5813e4f46378876/reflection.md)
## ๐ฅ[์ธํฐ๋ท ์์ค์ฌ์ฉ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/iunternetsource.md)
## โ[ํจํค์ง ์์ฑ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/packageadd.md)
## ๐งข[๋ฌธ์ํ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/forder.md)
## โจ[์๋ ฅํจ์ ์ฌ์ฉ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Scan.md)
## ๐ฌ[๋ฌธ์์ด ์์ถ๋ ฅ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/langscan.md)
## ๐ฑ[ํ์ผ ์์ถ๋ ฅ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/filescan.md)
## ๐จโโค๏ธโ๐โ๐จ[์ ๋์ฝ๋](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/UTF-8.md)
## โจ[๋ฌธ์์ด](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/monlang.md)
## ๐[์ ๊ทํํ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/regular.md)
## ๐ผ[ํ์ผ ์ฒ๋ฆฌ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/file.md)
## ๐[์์ถ๋ ฅ ์ธํฐํ์ด์ค](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/ioReader.md)
## ๐ช[JSON](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/json.md)
## ๐[์์ถ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/gzip.md)
## โ๏ธ[์ํธํ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/hash.md)
## ๐ฎ๐ผโโ๏ธ[์ ๋ ฌ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Sort.md)
## โฅ๏ธ[์ปจํ์ด๋](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/container.md)
## โฅ๏ธ[TCPํ๋กํ ์ฝ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/TCP.md)
## ๐ซ[PRCํ๋กํ ์ฝ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/RPC.md)
## ๐[๋ช๋ น์ค](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/ARG.md)
## ๐[์๋ฌ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/Error.md)
## ๐ฏ[๋จ์ํ์คํธ](https://github.com/kyunghyunHan/go_study/blob/3ce5ed0e3541f652e49b8f90248effddb3200dbc/TEST.md)
## ๐์บก์ํ Getter/setter. 
## ๐ฑ[์๋ฒ ์ฐ๊ฒฐ](https://github.com/kyunghyunHan/go_study/blob/f29db2308b47119deb3f708d129a7fa0df3fbeea/golangMysql.md)
## ๐ฑ[Mysql ์ฐ๊ฒฐ]()

