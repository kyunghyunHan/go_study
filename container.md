# ๐ฉ๐ปโ๐์ปจํ์ด๋ ์ฌ์ฉํ๊ธฐ


๋ค์ํ ๋ฐ์ดํฐ๋ฅผ ๋ค๋ฃฐ ๋ ์๋ฃ๊ตฌ์กฐ๋ ํ์์ ์ผ๋ก ์ฌ์ฉํ๊ฒ ๋ฉ๋๋ค. ์๋ฃ๊ตฌ์กฐ๋ฅผ ์ผ์ผ์ด ๊ตฌํํ๊ธฐ์๋ ์๊ฐ๊ณผ ๋ธ๋ ฅ์ด ๋ง์ด ๋ญ๋๋ค. Go ์ธ์ด๋ ๊ธฐ๋ณธ ์๋ฃ๊ตฌ์กฐ๋ฅผ ํจํค์ง๋ก ์ ๊ณตํ๋ฏ๋ก ํธ๋ฆฌํ๊ฒ ๋ฐ์ดํฐ๋ฅผ ๋ค๋ฃฐ ์ ์์ต๋๋ค.

Go ์ธ์ด์์ ์ ๊ณตํ๋ ์๋ฃ๊ตฌ์กฐ๋ ๋ค์๊ณผ ๊ฐ์ต๋๋ค.

- ์ฐ๊ฒฐ ๋ฆฌ์คํธ: ๊ฐ ๋ธ๋๋ฅผ ํ ์ค๋ก ์ฐ๊ฒฐํ ์๋ฃ๊ตฌ์กฐ์๋๋ค.
- ํ: ์ด์ง ํธ๋ฆฌ(Binary tree)๋ฅผ ํ์ฉํ ์๋ฃ๊ตฌ์กฐ์๋๋ค. ๋ถ๋ชจ ๋ธ๋์์ ์์ ๋ธ๋๊ฐ ๋์๊ด๊ณ๋ฅผ ์ด๋ฃจ๋ฉฐ ์ ๋ ฌ์ด ๋๋ ๊ฒ์ด ํน์ง์๋๋ค.
- ๋ง: ๊ฐ ๋ธ๋๊ฐ ์ํ์ผ๋ก ์ฐ๊ฒฐ๋ ์๋ฃ๊ตฌ์กฐ์๋๋ค.
์ฐ๊ฒฐ ๋ฆฌ์คํธ ์ฌ์ฉํ๊ธฐ
๋ค์์ container/list ํจํค์ง์์ ์ ๊ณตํ๋ ์ฐ๊ฒฐ ๋ฆฌ์คํธ ํจ์์๋๋ค.

- func New() *List: ์ฐ๊ฒฐ ๋ฆฌ์คํธ ์์ฑ
- func (l *List) PushBack(v interface{}) *Element: ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ๋ค์ ๋ฐ์ดํฐ ์ถ๊ฐ
- func (l *List) Front() *Element: ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ์ ๋ฐ์ดํฐ๋ฅผ ๊ฐ์ ธ์ด
- func (l *List) Back() *Element: ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ๋ค ๋ฐ์ดํฐ๋ฅผ ๊ฐ์ ธ์ด
์ฐ๊ฒฐ ๋ฆฌ์คํธ๋ฅผ ์์ฑํ ๋ค ๋ฐ์ดํฐ๋ฅผ ๋ฃ๊ณ , ๊ฐ ๋ธ๋๋ฅผ ์ํํด๋ณด๊ฒ ์ต๋๋ค.

```
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New() // ์ฐ๊ฒฐ ๋ฆฌ์คํธ ์์ฑ
	l.PushBack(10)  // ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋ฐ์ดํฐ ์ถ๊ฐ
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value) // ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ์ ๋ฐ์ดํฐ๋ฅผ ๊ฐ์ ธ์ด
	fmt.Println("Back ", l.Back().Value)   // ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ๋ค ๋ฐ์ดํฐ๋ฅผ ๊ฐ์ ธ์ด

	for e := l.Front(); e != nil; e = e.Next() { // ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ์๋ถํฐ ๋๊น์ง ์ํ
		fmt.Println(e.Value)
	}
}
```
```
Front  10
Back  30
10
20
30
```
go ์ธ์ด์ ์ฐ๊ฒฐ ๋ฆฌ์คํธ๋ ์ด์ค ์ฐ๊ฒฐ ๋ฆฌ์คํธ(Doubly linked list)์๋๋ค. ๋ฐ๋ผ์ ์, ๋ค ์์ชฝ ๋ฐฉํฅ์ผ๋ก ์ํํ  ์ ์์ต๋๋ค.

list.New ํจ์๋ก ์ฐ๊ฒฐ ๋ฆฌ์คํธ๋ฅผ ์์ฑํ ๋ค PushBack ๋ฑ์ ํจ์๋ก ๋ฐ์ดํฐ๋ฅผ ์ถ๊ฐํฉ๋๋ค. ๋ค์์ ์ฐ๊ฒฐ ๋ฆฌ์คํธ์์ ์ฌ์ฉํ  ์ ์๋ ํจ์์๋๋ค.

|ํจ์	|์ค๋ช|
|----|----|
|PushBack	|๋งจ ๋ค์ ๋ธ๋๋ฅผ ์ถ๊ฐํฉ๋๋ค.|
|PushFront	|๋งจ ์์ ๋ธ๋๋ฅผ ์ถ๊ฐํฉ๋๋ค.|
|PushBackList	|๋งจ ๋ค์ ๋ค๋ฅธ ๋ฆฌ์คํธ๋ฅผ ๋ถ์๋๋ค.|
|PushFrontList	|๋งจ ์์ ๋ค๋ฅธ ๋ฆฌ์คํธ๋ฅผ ๋ถ์๋๋ค.|
|InsertAfter	|ํน์  ๋ธ๋ ๋ค์ ๋ธ๋๋ฅผ ์ถ๊ฐํฉ๋๋ค.|
|InsertBefore	|ํน์  ๋ธ๋ ์์ ๋ธ๋๋ฅผ ์ถ๊ฐํฉ๋๋ค.|
|MoveAfter	|๋ธ๋๋ฅผ ํน์  ๋ธ๋ ๋ค๋ก ์ฎ๊น๋๋ค.|
|MoveBefore	|๋ธ๋๋ฅผ ํน์  ๋ธ๋ ์์ผ๋ก ์ฎ๊น๋๋ค.|
|MoveToBack	|๋ธ๋๋ฅผ ๋งจ ๋ค๋ก ์ฎ๊น๋๋ค.|
|MoveToFront|	๋ธ๋๋ฅผ ๋งจ ์์ผ๋ก ์ฎ๊น๋๋ค.|
|Len|	๋ฆฌ์คํธ์ ๋ธ๋ ๊ฐ์(๊ธธ์ด)๋ฅผ ๊ตฌํฉ๋๋ค.|
|Remove	|ํน์  ๋ธ๋๋ฅผ ์ญ์ ํฉ๋๋ค.|


l.Front().Value, l.Back().Value์ ๊ฐ์ด ๋ฆฌ์คํธ์์ ๋งจ ์, ๋งจ ๋ค์ ๋ธ๋์์ ๊ฐ์ ๊ฐ์ ธ์ต๋๋ค. ๋ค์๊ณผ ๊ฐ์ด for ๋ฐ๋ณต๋ฌธ์ ์ฌ์ฉํ๋ฉด ๋งจ ์๋ถํฐ ๋๊น์ง ๋ฆฌ์คํธ๋ฅผ ์ํํ  ์ ์์ต๋๋ค.
```
for e := l.Front(); e != nil; e = e.Next() { // ์ฐ๊ฒฐ ๋ฆฌ์คํธ์ ๋งจ ์๋ถํฐ ๋๊น์ง ์ํ
	fmt.Println(e.Value)
}
```

## ๐ฏํ ์ฌ์ฉํ๊ธฐ
๋ค์์ container/heap ํจ์์์ ์ ๊ณตํ๋ ํ ํจ์์๋๋ค.

- func Init(h Interface): ํ ์ด๊ธฐํ
- func Push(h Interface, x interface{}): ํ์ ๋ฐ์ดํฐ ์ถ๊ฐ
์ด์  ํ์ ์ฌ์ฉํด๋ณด๊ฒ ์ต๋๋ค.

```
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int // ํ์ int ์ฌ๋ผ์ด์ค๋ก ์ ์

func (h MinHeap) Len() int {
	return len(h) // ์ฌ๋ผ์ด์ค์ ๊ธธ์ด๋ฅผ ๊ตฌํจ
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j] // ๋์๊ด๊ณ ํ๋จ
	fmt.Printf("Less %d < %d %t\n", h[i], h[j], r)
	return r
}

func (h MinHeap) Swap(i, j int) {
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i] // ๊ฐ์ ์์น๋ฅผ ๋ฐ๊ฟ
}

func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int)) // ๋งจ ๋ง์ง๋ง์ ๊ฐ ์ถ๊ฐ
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // ์ฌ๋ผ์ด์ค์ ๋งจ ๋ง์ง๋ง ๊ฐ์ ๊ฐ์ ธ์ด
	*h = old[0 : n-1] // ๋งจ ๋ง์ง๋ง ๊ฐ์ ์ ์ธํ ์ฌ๋ผ์ด์ค๋ฅผ ๋ค์ ์ ์ฅ
	return x
}

func main() {
	data := new(MinHeap) // ํ ์์ฑ

	heap.Init(data)    // ํ ์ด๊ธฐํ
	heap.Push(data, 5) // ํ์ ๋ฐ์ดํฐ ์ถ๊ฐ
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "์ต์๊ฐ : ", (*data)[0])
}
```
ํ์ ์์  ์ด์ง ํธ๋ฆฌ(Complete binary tree)๋ฅผ ์ฌ์ฉํ ์๋ฃ๊ตฌ์กฐ์ด๋ฉฐ ๋ ๊ฐ์ง ์ข๋ฅ๊ฐ ์์ต๋๋ค.

- ์ต๋ ํ: ๋ถ๋ชจ ๋ธ๋์ ๊ฐ์ด ์์ ๋ธ๋์ ๊ฐ ๋ณด๋ค ํฐ ํ
- ์ต์ ํ: ๋ถ๋ชจ ๋ธ๋์ ๊ฐ์ด ์์ ๋ธ๋์ ๊ฐ ๋ณด๋ค ์์ ํ
์์ ์์๋ ์ต์ ํ์ ๊ตฌํํด๋ณด์์ต๋๋ค. ๋ฆฌ์คํธ ๊ฐ์ ์๋ฃ๊ตฌ์กฐ์๋ ๋ฌ๋ฆฌ ํ์ heap.Interface๋ฅผ ๊ตฌํํด์ฃผ์ด์ผ ํฉ๋๋ค.

๋จผ์  ํ์ ๋ฐ์ดํฐ ํ์์ ์ ์ํฉ๋๋ค. ์ฌ๊ธฐ์๋ ๊ฐ๋จํ๊ฒ ์ ์ํ ์ฌ๋ผ์ด์ค๋ก ์ ์ํ์ต๋๋ค

```
type MinHeap []int // ํ์ int ์ฌ๋ผ์ด์ค๋ก ์ ์
```
๋ฐ์ดํฐ๋ฅผ ๋ฃ๋ Push ๋ฉ์๋์ ๋ฐ์ดํฐ๋ฅผ ๊บผ๋ด๋ Pop ๋ฉ์๋๋ฅผ ์ ์ํฉ๋๋ค.
```
func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int)) // ๋งจ ๋ง์ง๋ง์ ๊ฐ ์ถ๊ฐ
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // ์ฌ๋ผ์ด์ค์ ๋งจ ๋ง์ง๋ง ๊ฐ์ ๊ฐ์ ธ์ด
	*h = old[0 : n-1] // ๋งจ ๋ง์ง๋ง ๊ฐ์ ์ ์ธํ ์ฌ๋ผ์ด์ค๋ฅผ ๋ค์ ์ ์ฅ
	return x
}

```
Push ๋ฉ์๋๋ ์ฌ๋ผ์ด์ค ๋ง์ง๋ง์ ๊ฐ์ ์ถ๊ฐํ๋ฉฐ Pop ๋ฉ์๋๋ ์ฌ๋ผ์ด์ค์ ๋งจ ๋ง์ง๋ง ๊ฐ์ ๊ฐ์ ธ์ต๋๋ค. ์ฌ๊ธฐ์๋ old[n-1]๋ก ๋งจ ๋ง์ง๋ง ๊ฐ์ ๊ฐ์ ธ์จ ๋ค *h = old[0 : n-1]์ฒ๋ผ ๋งจ ๋ง์ง๋ง ๊ฐ์ ์ ์ธํ ์ฌ๋ผ์ด์ค๋ฅผ ๋ค์ ๋ฐ์ดํฐ์ ์ ์ฅํฉ๋๋ค.

์ด์  ์ ๋ ฌ(sort.Interface)์ ๊ตฌํํฉ๋๋ค(sort.Interface ์ธํฐํ์ด์ค๋ heap.Interface ์ธํฐํ์ด์ค ์์ ์๋ฒ ๋ฉ๋์ด ์์ต๋๋ค).
```
func (h MinHeap) Len() int {
	return len(h) // ์ฌ๋ผ์ด์ค์ ๊ธธ์ด๋ฅผ ๊ตฌํจ
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j] // ๋์๊ด๊ณ ํ๋จ
	fmt.Printf("Less %d < %d %t\n", h[i], h[j], r)
	return r
}

func (h MinHeap) Swap(i, j int) {
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i] // ๊ฐ์ ์์น๋ฅผ ๋ฐ๊ฟ
}
```
- Len: ํ์ ๊ธธ์ด๋ฅผ ๊ตฌํ๋ ํจ์์๋๋ค. ์ฌ๋ผ์ด์ค์ด๋ฏ๋ก len ํจ์๋ฅผ ์ฌ์ฉํ์ฌ ๊ธธ์ด๋ฅผ ๊ตฌํฉ๋๋ค.
- Less, Swap: ๊ธฐ์กด์ ๊ฐ(j)๊ณผ ์๋ก ๋ค์ด์จ ๊ฐ(i)์ ๋น๊ตํ์ฌ ์๋ก ๋ค์ด์จ ๊ฐ์ด ํฌ๋ฉด ๊ทธ๋๋ก ๋๊ณ , ์๋ก ๋ค์ด์จ ๊ฐ์ด ์์ผ๋ฉด ๊ธฐ์กด ๊ฐ๊ณผ ๊ต์ฒดํฉ๋๋ค. ์ฆ ๋ถ๋ชจ ๋ธ๋๋ณด๋ค ์์ ๋ธ๋์ ๊ฐ์ด ํฌ๋ฉด ๊ทธ๋๋ก ๋๊ณ , ๋ถ๋ชจ ๋ธ๋๋ณด๋ค ์์ ๋ธ๋๊ฐ ์์ผ๋ฉด ์๋ก ๊ฐ์ ๋ฐ๊ฟ๋๋ค(์ด ์์ ๋ฐ๋๋ก ํ๋ฉด ์ต๋ ํ์ด ๋ฉ๋๋ค). ๊ฐ์ ์ถ๊ฐํ  ๋๋ง๋ค ์ด ์์์ด ๋ฐ๋ณต๋ฉ๋๋ค.
๊ฒฐ๊ณผ๋ฅผ ์ถ๋ ฅํด๋ณด๋ฉด ๋ค์๊ณผ ๊ฐ์ด heap.Push ํจ์๋ก ๊ฐ์ ๋ฃ์ ๋๋ง๋ค ์ ๋ ฌ์ด ๋ฉ๋๋ค.
```
Push 5
Push 2
Less 2 < 5 true
Swap 5 2
Push 7
Less 7 < 2 false
Push 3
Less 3 < 5 true
Swap 5 3
Less 3 < 2 false
&[2 3 7 5] ์ต์๊ฐ :  2
```

## ๐ฏ๋ง ์ฌ์ฉํ๊ธฐ
๋ค์์ container/ring ํจํค์ง์์ ์ ๊ณตํ๋ ๋ง ํจ์์๋๋ค.

- func New(n int) *Ring: ๋ง ์์ฑ
- func (r *Ring) Do(f func(interface{})): ๋ง์ ๋ชจ๋  ๋ธ๋ ์ํ
- func (r *Ring) Move(n int) *Ring: ๋ง์ ํ์ ์ํด. ๋งค๊ฐ๋ณ์๋ก ์์๋ฅผ ๋ฃ์ผ๋ฉด ์๊ณ ๋ฐฉํฅ, ์์๋ฅผ ๋ฃ์ผ๋ฉด ๋ฐ ์๊ณ ๋ฐฉํฅ์ผ๋ก ํ์ 
์ด์  ๋ง(Ring)์ ์ฌ์ฉํด๋ณด๊ฒ ์ต๋๋ค.

```
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(data))       // ๋ธ๋์ ๊ฐ์๋ฅผ ์ง์ ํ์ฌ ๋ง ์์ฑ
	for i := 0; i < r.Len(); i++ { // ๋ง ๋ธ๋ ๊ฐ์๋งํผ ๋ฐ๋ณต
		r.Value = data[i]      // ๋ง์ ๋ธ๋์ ๊ฐ ๋ฃ๊ธฐ
		r = r.Next()           // ๋ค์ ๋ธ๋๋ก  ์ด๋
	}

	r.Do(func(x interface{}) { // ๋ง์ ๋ชจ๋  ๋ธ๋ ์ํ
		fmt.Println(x)
	})

	fmt.Println("Move forward :")
	r = r.Move(1) // ๋ง์ ์๊ณ ๋ฐฉํฅ์ผ๋ก 1๋ธ๋ ๋งํผ ํ์ 

	fmt.Println("Curr : ", r.Value)        // Curr :  John
	fmt.Println("Prev : ", r.Prev().Value) // Prev :  Maria
	fmt.Println("Next : ", r.Next().Value) // Next :  Andrew
}
```
```
Maria
John
Andrew
James
Move forward :
Curr :  John
Prev :  Maria
Next :  Andrew
```
๋ง์ ์ํ์ผ๋ก ์ฐ๊ฒฐ๋ ์ด์ค ์ฐ๊ฒฐ ๋ฆฌ์คํธ์๋๋ค. ๋ฐ๋ผ์ ์ฒ์๊ณผ ๋์ด ์๊ณ , nil์ ๊ฐ๋ฆฌํค๋ ๋ธ๋๋ ์กด์ฌํ์ง ์์ต๋๋ค.
ring.New ํจ์์ ๋ธ๋์ ๊ฐ์๋ฅผ ์ง์ ํ์ฌ ๋ง์ ์์ฑํฉ๋๋ค. ๊ทธ๋ฆฌ๊ณ  for ๋ฐ๋ณต๋ฌธ์ ์ฌ์ฉํ์ฌ ๊ฐ์ ๋ฃ์ต๋๋ค. ๋ฐ๋์ ๋ค๋ก ๊ฐ๋  ์์ผ๋ก ๊ฐ๋  r = r.Next()์ฒ๋ผ ํ์ฌ ์์น๋ฅผ ๋ฐ๊ฟ์ฃผ์ด์ผ ํฉ๋๋ค.
```
data := []string{"Maria", "John", "Andrew", "James"}

r := ring.New(len(data))       // ๋ธ๋์ ๊ฐ์๋ฅผ ์ง์ ํ์ฌ ๋ง ์์ฑ
for i := 0; i < r.Len(); i++ { // ๋ง ๋ธ๋ ๊ฐ์๋งํผ ๋ฐ๋ณต
	r.Value = data[i]      // ๋ง์ ๋ธ๋์ ๊ฐ ๋ฃ๊ธฐ
	r = r.Next()           // ๋ค์ ๋ธ๋๋ก  ์ด๋
}

```
Do ํจ์๋ ๋ง์ ๋ชจ๋  ๋ธ๋๋ฅผ ์ํํฉ๋๋ค. ์ฌ๊ธฐ์๋ ๋ชจ๋  ๋ธ๋๋ฅผ ์ํํ๋ฉด์ ๊ฐ์ ์ถ๋ ฅํฉ๋๋ค.
```
r.Do(func(x interface{}) { // ๋ง์ ๋ชจ๋  ๋ธ๋ ์ํ
	fmt.Println(x)
})
```
Move ํจ์๋ ๋ง์ ํ์ ์ํต๋๋ค. ์์๋ฅผ ๋ฃ์ผ๋ฉด ์์ชฝ(์๊ณ ๋ฐฉํฅ)์ผ๋ก ํ์ ํ๊ณ , ์์๋ฅผ ๋ฃ์ผ๋ฉด ๋ค์ชฝ(๋ฐ ์๊ณ ๋ฐฉํฅ)์ผ๋ก ํ์ ํฉ๋๋ค.
```
r = r.Move(1) // ๋ง์ ์๊ณ ๋ฐฉํฅ์ผ๋ก 1๋ธ๋ ๋งํผ ํ์ 
```
๋ง๋ ๋ฆฌ์คํธ์ ๋ง์ฐฌ๊ฐ์ง๋ก ํ์ฌ ๋ธ๋, ์ด์  ๋ธ๋, ๋ค์ ๋ธ๋์ ๊ฐ์ ๊ฐ์ ธ์ฌ ์ ์์ต๋๋ค. ์ฌ๊ธฐ์๋ ์์ชฝ ๋ฐฉํฅ์ผ๋ก 1๋ธ๋ ๋งํผ ํ์ ํ์ผ๋ฏ๋ก ํ์ฌ ๊ฐ์ John์๋๋ค
```
fmt.Println("Curr : ", r.Value)        // Curr :  John
fmt.Println("Prev : ", r.Prev().Value) // Prev :  Maria
fmt.Println("Next : ", r.Next().Value) // Next :  Andrew
```
