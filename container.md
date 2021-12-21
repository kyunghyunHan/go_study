# 👩🏻‍🎓컨테이너 사용하기


다양한 데이터를 다룰 때 자료구조는 필수적으로 사용하게 됩니다. 자료구조를 일일이 구현하기에는 시간과 노력이 많이 듭니다. Go 언어는 기본 자료구조를 패키지로 제공하므로 편리하게 데이터를 다룰 수 있습니다.

Go 언어에서 제공하는 자료구조는 다음과 같습니다.

- 연결 리스트: 각 노드를 한 줄로 연결한 자료구조입니다.
- 힙: 이진 트리(Binary tree)를 활용한 자료구조입니다. 부모 노드에서 자식 노드가 대소관계를 이루며 정렬이 되는 것이 특징입니다.
- 링: 각 노드가 원형으로 연결된 자료구조입니다.
연결 리스트 사용하기
다음은 container/list 패키지에서 제공하는 연결 리스트 함수입니다.

- func New() *List: 연결 리스트 생성
- func (l *List) PushBack(v interface{}) *Element: 연결 리스트의 맨 뒤에 데이터 추가
- func (l *List) Front() *Element: 연결 리스트의 맨 앞 데이터를 가져옴
- func (l *List) Back() *Element: 연결 리스트의 맨 뒤 데이터를 가져옴
연결 리스트를 생성한 뒤 데이터를 넣고, 각 노드를 순회해보겠습니다.

```
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New() // 연결 리스트 생성
	l.PushBack(10)  // 연결 리스트에 데이터 추가
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value) // 연결 리스트의 맨 앞 데이터를 가져옴
	fmt.Println("Back ", l.Back().Value)   // 연결 리스트의 맨 뒤 데이터를 가져옴

	for e := l.Front(); e != nil; e = e.Next() { // 연결 리스트의 맨 앞부터 끝까지 순회
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
go 언어의 연결 리스트는 이중 연결 리스트(Doubly linked list)입니다. 따라서 앞, 뒤 양쪽 방향으로 순회할 수 있습니다.

list.New 함수로 연결 리스트를 생성한 뒤 PushBack 등의 함수로 데이터를 추가합니다. 다음은 연결 리스트에서 사용할 수 있는 함수입니다.

|함수	|설명|
|----|----|
|PushBack	|맨 뒤에 노드를 추가합니다.|
|PushFront	|맨 앞에 노드를 추가합니다.|
|PushBackList	|맨 뒤에 다른 리스트를 붙입니다.|
|PushFrontList	|맨 앞에 다른 리스트를 붙입니다.|
|InsertAfter	|특정 노드 뒤에 노드를 추가합니다.|
|InsertBefore	|특정 노드 앞에 노드를 추가합니다.|
|MoveAfter	|노드를 특정 노드 뒤로 옮깁니다.|
|MoveBefore	|노드를 특정 노드 앞으로 옮깁니다.|
|MoveToBack	|노드를 맨 뒤로 옮깁니다.|
|MoveToFront|	노드를 맨 앞으로 옮깁니다.|
|Len|	리스트의 노드 개수(길이)를 구합니다.|
|Remove	|특정 노드를 삭제합니다.|


l.Front().Value, l.Back().Value와 같이 리스트에서 맨 앞, 맨 뒤의 노드에서 값을 가져옵니다. 다음과 같이 for 반복문을 사용하면 맨 앞부터 끝까지 리스트를 순회할 수 있습니다.
```
for e := l.Front(); e != nil; e = e.Next() { // 연결 리스트의 맨 앞부터 끝까지 순회
	fmt.Println(e.Value)
}
```

## 💯힙 사용하기
다음은 container/heap 함수에서 제공하는 힙 함수입니다.

- func Init(h Interface): 힙 초기화
- func Push(h Interface, x interface{}): 힙에 데이터 추가
이제 힙을 사용해보겠습니다.

```
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int // 힙을 int 슬라이스로 정의

func (h MinHeap) Len() int {
	return len(h) // 슬라이스의 길이를 구함
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j] // 대소관계 판단
	fmt.Printf("Less %d < %d %t\n", h[i], h[j], r)
	return r
}

func (h MinHeap) Swap(i, j int) {
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i] // 값의 위치를 바꿈
}

func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int)) // 맨 마지막에 값 추가
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // 슬라이스의 맨 마지막 값을 가져옴
	*h = old[0 : n-1] // 맨 마지막 값을 제외한 슬라이스를 다시 저장
	return x
}

func main() {
	data := new(MinHeap) // 힙 생성

	heap.Init(data)    // 힙 초기화
	heap.Push(data, 5) // 힙에 데이터 추가
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "최솟값 : ", (*data)[0])
}
```
힙은 완전 이진 트리(Complete binary tree)를 사용한 자료구조이며 두 가지 종류가 있습니다.

- 최대 힙: 부모 노드의 값이 자식 노드의 값 보다 큰 힙
- 최소 힙: 부모 노드의 값이 자식 노드의 값 보다 작은 힙
예제에서는 최소 힙을 구현해보았습니다. 리스트 같은 자료구조와는 달리 힙은 heap.Interface를 구현해주어야 합니다.

먼저 힙의 데이터 타입을 정의합니다. 여기서는 간단하게 정수형 슬라이스로 정의했습니다

```
type MinHeap []int // 힙을 int 슬라이스로 정의
```
데이터를 넣는 Push 메서드와 데이터를 꺼내는 Pop 메서드를 정의합니다.
```
func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int)) // 맨 마지막에 값 추가
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // 슬라이스의 맨 마지막 값을 가져옴
	*h = old[0 : n-1] // 맨 마지막 값을 제외한 슬라이스를 다시 저장
	return x
}

```
Push 메서드는 슬라이스 마지막에 값을 추가하며 Pop 메서드는 슬라이스의 맨 마지막 값을 가져옵니다. 여기서는 old[n-1]로 맨 마지막 값을 가져온 뒤 *h = old[0 : n-1]처럼 맨 마지막 값을 제외한 슬라이스를 다시 데이터에 저장합니다.

이제 정렬(sort.Interface)을 구현합니다(sort.Interface 인터페이스는 heap.Interface 인터페이스 안에 임베딩되어 있습니다).
```
func (h MinHeap) Len() int {
	return len(h) // 슬라이스의 길이를 구함
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j] // 대소관계 판단
	fmt.Printf("Less %d < %d %t\n", h[i], h[j], r)
	return r
}

func (h MinHeap) Swap(i, j int) {
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i] // 값의 위치를 바꿈
}
```
- Len: 힙의 길이를 구하는 함수입니다. 슬라이스이므로 len 함수를 사용하여 길이를 구합니다.
- Less, Swap: 기존의 값(j)과 새로 들어온 값(i)을 비교하여 새로 들어온 값이 크면 그대로 두고, 새로 들어온 값이 작으면 기존 값과 교체합니다. 즉 부모 노드보다 자식 노드의 값이 크면 그대로 두고, 부모 노드보다 자식 노드가 작으면 서로 값을 바꿉니다(이 식을 반대로 하면 최대 힙이 됩니다). 값을 추가할 때마다 이 작업이 반복됩니다.
결과를 출력해보면 다음과 같이 heap.Push 함수로 값을 넣을 때마다 정렬이 됩니다.
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
&[2 3 7 5] 최솟값 :  2
```

## 💯링 사용하기
다음은 container/ring 패키지에서 제공하는 링 함수입니다.

- func New(n int) *Ring: 링 생성
- func (r *Ring) Do(f func(interface{})): 링의 모든 노드 순회
- func (r *Ring) Move(n int) *Ring: 링을 회전시킴. 매개변수로 양수를 넣으면 시계 방향, 음수를 넣으면 반 시계 방향으로 회전
이제 링(Ring)을 사용해보겠습니다.

```
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(data))       // 노드의 개수를 지정하여 링 생성
	for i := 0; i < r.Len(); i++ { // 링 노드 개수만큼 반복
		r.Value = data[i]      // 링의 노드에 값 넣기
		r = r.Next()           // 다음 노드로  이동
	}

	r.Do(func(x interface{}) { // 링의 모든 노드 순회
		fmt.Println(x)
	})

	fmt.Println("Move forward :")
	r = r.Move(1) // 링을 시계 방향으로 1노드 만큼 회전

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
링은 원형으로 연결된 이중 연결 리스트입니다. 따라서 처음과 끝이 없고, nil을 가리키는 노드는 존재하지 않습니다.
ring.New 함수에 노드의 개수를 지정하여 링을 생성합니다. 그리고 for 반복문을 사용하여 값을 넣습니다. 반드시 뒤로 가든 앞으로 가든 r = r.Next()처럼 현재 위치를 바꿔주어야 합니다.
```
data := []string{"Maria", "John", "Andrew", "James"}

r := ring.New(len(data))       // 노드의 개수를 지정하여 링 생성
for i := 0; i < r.Len(); i++ { // 링 노드 개수만큼 반복
	r.Value = data[i]      // 링의 노드에 값 넣기
	r = r.Next()           // 다음 노드로  이동
}

```
Do 함수는 링의 모든 노드를 순회합니다. 여기서는 모든 노드를 순회하면서 값을 출력합니다.
```
r.Do(func(x interface{}) { // 링의 모든 노드 순회
	fmt.Println(x)
})
```
Move 함수는 링을 회전시킵니다. 양수를 넣으면 앞쪽(시계 방향)으로 회전하고, 음수를 넣으면 뒤쪽(반 시계 방향)으로 회전합니다.
```
r = r.Move(1) // 링을 시계 방향으로 1노드 만큼 회전
```
링도 리스트와 마찬가지로 현재 노드, 이전 노드, 다음 노드의 값을 가져올 수 있습니다. 여기서는 앞쪽 방향으로 1노드 만큼 회전했으므로 현재 값은 John입니다
```
fmt.Println("Curr : ", r.Value)        // Curr :  John
fmt.Println("Prev : ", r.Prev().Value) // Prev :  Maria
fmt.Println("Next : ", r.Next().Value) // Next :  Andrew
```
