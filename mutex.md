# 👩🏻‍🎓동기화 객체
- Go 언어에서는 채널 이외에도 고루틴의 실행 흐름을 제어하는 동기화 객체를 제공합니다.

대표적인 동기화(synchronization) 객체는 다음과 같다

- Mutex: 뮤텍스입니다. 상호 배제(mutual exclusion)라고도 하며 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용
- RWMutex: 읽기/쓰기 뮤텍스입니다. 읽기와 쓰기 동작을 나누어서 잠금(락)을 걸 수 있다
- Cond: 조건 변수(condition variable). 대기하고 있는 하나의 객체를 깨울 수도 있고 여러 개를 동시에 깨울 수도 있다
- Once: 특정 함수를 딱 한 번만 실행할 때 사용
- Pool: 멀티 스레드(고루틴)에서 사용할 수 있는 객체 풀이다 자주 사용하는 객체를 풀에 보관했다가 다시 사용
- WaitGroup: 고루틴이 모두 끝날 때까지 기다리는 기능
- Atomic: 원자적 연산이라고도 하며 더 이상 쪼갤 수 없는 연산이라는 뜻이다 멀티 스레드(고루틴), 멀티코어 환경에서 안전하게 값을 연산하는 기능이다
## 💯뮤텍스 사용하기
- 뮤텍스는 여러 고루틴이 공유하는 데이터를 보호할 때 사용하며 sync 패키지에서 제공하는 뮤텍스 구조체와 함수는 다음과 같다

- sync.Mutex
- func (m *Mutex) Lock(): 뮤텍스 잠금
- func (m *Mutex) Unlock(): 뮤텍스 잠금 해제
다음은 고루틴 두 개에서 각각 1,000번씩 슬라이스에 값을 추가
```
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{} // int형 슬라이스 생성

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			data = append(data, 1)  // data 슬라이스에 1을 추가

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			data = append(data, 1)  // data 슬라이스에 1을 추가

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)      // 2초 대기

	fmt.Println(len(data))           // data 슬라이스의 길이 출력
}
```
```
1883 (매번 달라질 수 있음)
```
- 실행을 해보면 대략 1800~1990 사이의 값이 나온다. data 슬라이스에 1을 2,000번 추가했으므로 data의 길이가 2000이 되어야 하는데 그렇지가 않다. 두 고루틴이 경합을 벌이면서 동시에 data에 접근했기 때문에 append 함수가 정확하게 처리되지 않은 상황이다 이러한 상황을 경쟁 조건(Race condition)이라고 한다

- runtime.Gosched 함수는 다른 고루틴이 CPU를 사용할 수 있도록 양보(yield)한다. 지금까지 time.Sleep 함수를 사용했지만 runtime.Gosched 함수가 좀 더 명확하다

이제 data 슬라이스를 뮤텍스로 보호한다
```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			mutex.Lock()            // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1)  // data 슬라이스에 1을 추가
			mutex.Unlock()          // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			mutex.Lock()            // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1)  // data 슬라이스에 1을 추가
			mutex.Unlock()          // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기

	fmt.Println(len(data)) // data 슬라이스의 길이 출력
}
```
- 뮤텍스는 sync.Mutex를 할당한 뒤에 고루틴에서 Lock, Unlock 함수로 사용한다. 보호를 시작할 부분에서 Lock 함수를 사용하고, 보호를 끝낼 부분에서 Unlock 함수를 사용한다 Lock, Unlock 함수는 반드시 짝을 맞추어주어야 하며 짝이 맞지 않으면 데드락(deadlock, 교착 상태)이 발생하므로 주의해야한다.

여기서는 data 슬라이스를 보호할 것이므로 두 고루틴 모두 data = append(data, 1) 부분 위 아래로 Lock, Unlock 함수를 사용한다. 이제 실행을 해보면 정확히 2000이 출력이된다.

## 💯읽기, 쓰기 뮤텍스 사용하기

읽기, 쓰기 뮤텍스는 읽기 동작과 쓰기 동작을 나누어 잠금(락)을 걸 수 있다

- 읽기 락(Read Lock): 읽기 락끼리는 서로를 막지 않는다. 하지만 읽기 시도 중에 값이 바뀌면 안 되므로 쓰기 락은 막는다
- 쓰기 락(Write Lock): 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안 되고, 다른 곳에서 값을 바꾸면 안 되므로 읽기, 쓰기 락 모두 막는다.
sync 패키지에서 제공하는 읽기, 쓰기 뮤텍스 구조체와 함수는 다음과 같다.

- sync.RWMutex
- func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock(): 쓰기 뮤텍스 잠금, 잠금 해제
- func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock(): 읽기 뮤텍스 잠금 및 잠금 해제
 읽기 쓰기 뮤텍스를 사용하지 않고 고루틴에서 값을 출력

```
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int = 0

	go func() {                                       // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			data += 1                         // data에 값 쓰기
			fmt.Println("write   : ", data)   // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
		}
	}()

	go func() {                                       // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 1 : ", data)    // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)       // 1초 대기
		}
	}()

	go func() {                                       // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 2 : ", data)    // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)       // 2초 대기
		}
	}()

	time.Sleep(10 * time.Second)              // 10초 동안 프로그램 실행
}
```

- 값을 쓰는 고루틴을 1개 생성하고, 값을 읽는 고루틴을 2개 생성!, 값을 쓰는 고루틴은 10 밀리초마다 반복하고, 값을 읽는 고루틴은 1초, 2초마다 반복한다. 실행해보면 매번 불규칙적인 모양으로 출력될 것이고, 특별한 순서를 찾기 어렵다
-  변수에 1만 더하는 간단한 상황이라 큰 문제가 없지만, 실제 상황에서는 중요한 데이터를 쓰고 있는데 다른 곳에서 이전 데이터를 읽는다던지, 읽기 시도 중에 값이 바뀐다던지 하는 상황이 생길 수 있다.

```
read 1 :  1
read 2 :  1
write  :  1
write  :  2
write  :  3
read 1 :  3
read 1 :  3
read 2 :  3
read 2 :  3
```
```
]package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int = 0
	var rwMutex = new(sync.RWMutex) // 읽기, 쓰기 뮤텍스 생성

	go func() {                                       // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                    // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1                         // data에 값 쓰기
			fmt.Println("write  : ", data)    // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
			rwMutex.Unlock()                  // 쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() {                                       // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                   // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 1 : ", data)    // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)       // 1초 대기
			rwMutex.RUnlock()                 // 읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	go func() {                                       // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                   // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 2 : ", data)    // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)       // 2초 대기
			rwMutex.RUnlock()                 // 읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	time.Sleep(10 * time.Second)                      // 10초 동안 프로그램 실행
}
```
- 읽기, 쓰기 뮤텍스는 sync.RWMutex를 할당한 뒤에 고루틴에서 RLock, RUnlock, Lock, Unlock 함수로 사용한다. 읽기 동작을 시작할 부분에서 RLock 함수를 사용하고, 읽기 동작을 끝낼 부분에서 RUnlock 함수를 사용한다. 그리고 쓰기 동작을 시작할 부분에서 Lock 함수를 사용하고, 쓰기 동작을 끝낼 부분에서 Unlock 함수를 사용한다.
- RLock, RUnlock, Lock, Unlock 함수는 반드시 짝을 맞춰야 하며 짝이 맞지 않으면 데드락(deadlock, 교착 상태)이 발생하므로 주의해야한다.

- 실행해보면 각각의 순서는 매번 달라지지만 실행되는 모양은 규칙적이다. 즉 read 1, read 2 읽기 동작이 모두 끝나야 write 쓰기 동작이 시작된다. 마찬가지로 쓰기 동작이 끝나야 읽기 동작이 시작된다. 읽기 동작끼리는 서로를 막지 않으므로 항상 동시에 실행된다.

```
read 1 :  0
read 2 :  0  ← read 1, read 2 읽기 동작이 모두 끝나야
write  :  1  ← write 쓰기 동작이 시작됨
read 1 :  1
read 2 :  1
write  :  2
read 2 :  2
read 1 :  2
write  :  3
```
- 읽기, 쓰기 뮤텍스는 중요한 쓰기 작업을 할 때 다른 곳에서 이전 데이터를 읽지 못하도록 방지하거나, 읽기 작업을 할 때 데이터가 바뀌는 상황을 방지할 때 사용한다. 특히 읽기, 쓰기 뮤텍스는 쓰기 동작보다 읽기 동작이 많을 때 유리하다.

## 💯조건변수

조건 변수는 대기하고 있는 객체 하나만 깨우거나 여러 개를 동시에 깨울 때 사용한다.

sync 패키지에서 제공하는 조건 변수의 함수는 다음과 같다.

- sync.Cond
- func NewCond(l Locker) *Cond: 조건 변수 생성
- func (c *Cond) Wait(): 고루틴 실행을 멈추고 대기
- func (c *Cond) Signal(): 대기하고 있는 고루틴 하나만 깨움
- func (c *Cond) Broadcast(): 대기하고 있는 모든 고루틴을 깨움
먼저 대기하고 있는 고루틴을 하나씩 깨워우기
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) {                        // 고루틴 3개 생성
			mutex.Lock()                    // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true                       // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait()                     // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock()                  // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()                // 뮤텍스 잠금, cond.Signal() 보호 시작
		fmt.Println("signal : ", i)
		cond.Signal()               // 대기하고 있는 고루틴을 하나씩 깨움
		mutex.Unlock()              // 뮤텍스 잠금 해제, cond.Signal() 보고 종료
	}

	fmt.Scanln()
}
```
- 조건 변수는 뮤텍스를 먼저 생성한 뒤 sync.NewCond 함수로 생성한다. 대기할 때는 고루틴 안에서 Wait 함수를 사용하며, 대기하는 고루틴을 깨울 때는 Signal 함수를 사용한다. 여기서 Wait 함수 부분은 뮤텍스의 Lock, Unlock 함수로 보호해야 한다.
- 여기서는 채널을 사용하여 모든 고루틴이 생성될 때까지 기다린 뒤 차례대로 Signal 함수를 사용하여 대기하고 있는 고루틴을 깨운다.
```
wait begin :  0
wait begin :  1
wait begin :  2
signal :  0
signal :  1
signal :  2
wait end :  0
wait end :  1
wait end :  2
```
- 이번에는 대기하고 있는 모든 고루틴을 꺠운다
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) {                        // 고루틴 3개 생성
			mutex.Lock()                    // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true                       // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait()                     // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock()                  // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	mutex.Lock()             // 뮤텍스 잠금, cond.Broadcast() 보호 시작
	fmt.Println("broadcast")
	cond.Broadcast()         // 대기하고 있는 모든 고루틴을 깨움
	mutex.Unlock()           // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	fmt.Scanln()
}
```
- 마찬가지로 대기할 때는 고루틴에서 Wait 함수를 사용한다. 대기하고 있는 모든 고루틴을 깨울 때는 Broadcast 함수를 사용하면 된다.

- 여기서는 Broadcast 함수를 사용한 즉시 대기하고 있던 고루틴 3개가 깨어난다.
```
wait begin :  1
wait begin :  2
wait begin :  0
broadcast
wait end :  2
wait end :  1
wait end :  0
```

## 👩🏻‍🎓함수한번만 사용
Once를 사용하면 함수를 한 번만 실행할 수 있다. 다음은 고루틴 안에서 Hello, world!를 출력한다.

sync 패키지에서 제공하는 Once의 구조체와 함수는 다음과 같다
- sync.Once
- func (*Once) Do(f func()): 함수를 한 번만 실행
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ {
		go func(n int) {   // 고루틴 3개 생성
			fmt.Println("goroutine : ", n)

			once.Do(hello) // 고루틴은 3개지만 hello 함수를 한 번만 실행
		}(i)
	}

	fmt.Scanln()
}
```
- Once는 sync.Once를 할당한 뒤에 Do 함수로 사용한다. Do 함수에는 실행할 함수 이름을 지정하거나, 클로저 형태로 함수를 지정할 수 있다. Once는 어떤 상황이든 상관없이 지정된 함수를 딱 한 번만 실행시킨다

- 여기서는 고루틴을 3개 실행시키고 각각 Once의 Do 함수로 hello 함수를 실행시켰지만 실제로는 한 번만 실행되었다.
```
goroutine :  1
Hello, world!
goroutine :  0
goroutine :  2
```
- 이처럼 Once는 한 번만 실행되는 특성이 있으므로 복잡한 반복문 안에서 각종 초기화를 할 때 유용한다.

## 💯풀사용
풀은 객체(메모리)를 사용한 후 보관해두었다가 다시 사용하게 해주는 기능이다. 객체를 반복해서 할당하면 메모리 사용량도 늘어나고, 메모리를 해제해야 하는 가비지 컬렉터에게도 부담이 된다. 즉, 풀은 일종의 캐시라고 할 수 있으며 메모리 할당과 해제 횟수를 줄여 성능을 높이고자 할 때 사용한다. 그리고 풀은 여러 고루틴에서 동시에 사용할 수 있다.

sync 패키지에서 제공하는 풀의 구조체와 함수는 다음과 같다.

- sync.Pool
- func (p *Pool) Get() interface{}: 풀에 보관된 객체를 가져옴
- func (p *Pool) Put(x interface{}): 풀에 객체를 보관
풀을 사용하여 정수 10개짜리 슬라이스를 공유, 첫 번째 고루틴 그룹에서는 슬라이스에 랜덤한 숫자를 10개를 저장한 뒤 출력하고, 두 번째 고루틴 그룹에서는 짝수 10개를 저장한 뒤 출력한다.

```
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { // Data 구조체 정의
	tag    string  // 풀 태그
	buffer []int   // 데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{                            // 풀 할당
		New: func() interface{} {             // Get 함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	for i := 0; i < 10; i++ {
		go func() {                                          // 고루틴 10개 생성
			data := pool.Get().(*Data)                   // 풀에서 *Data 타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)  // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data)                            // data 내용 출력
			data.tag = "used"                            // 객체가 사용되었다는 태그 설정
			pool.Put(data)                               // 풀에 객체를 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {                               // 고루틴 10개 생성
			data := pool.Get().(*Data)        // 풀에서 *Data 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n    // 슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data)                 // data 내용 출력
			data.tag = "used"                 // 객체가 사용되었다는 태그 설정
			pool.Put(data)                    // 풀에 객체 보관
		}()
	}

	fmt.Scanln()
}
```
- 풀은 sync.Pool을 할당한 뒤에 Get, Put 함수로 사용한다. 먼저 다음과 같이 sync.Pool을 할당 한 뒤 New 필드에 초기화 함수를 만들어준다.
```
pool := sync.Pool{                            // 풀 할당
	New: func() interface{} {             // Get 함수를 사용했을 때 호출될 함수 정의
		data := new(Data)             // 새 메모리 할당
		data.tag = "new"              // 태그 설정
		data.buffer = make([]int, 10) // 슬라이스 공간 할당
		return data                   // 할당한 메모리(객체) 리턴
	},
}
```
- New 필드에 정의된 함수는 Get 함수를 사용했을 때 호출된다. 단, 풀에 객체가 없을 때만 호출되므로 객체를 생성하고, 메모리를 할당하는 코드를 작성한다. 풀에 객체가 들어있다면 New 필드의 함수는 호출되지 않고, 보관된 객체가 리턴한다. 여기서는 객체를 새로 할당했다는 의미에서 tag 필드에 new를 대입한다.(tag 필드는 풀의 객체 사용 상황을 알아보기 위한 예제이며 필수 요소는 아니다)
- 풀에서 Get 함수로 객체를 꺼낸 뒤에는 반드시 Type assertion을 해주어야 한다. 여기서는 New 필드의 함수에서 new(Data)로 메모리를 할당했으므로 포인터 형인 (*Data)로 변환된다.
```
data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴
```
- 이제 슬라이스에 적절히 데이터를 채워넣은 뒤 화면에 출력한다. 그리고 객체 사용이 끝났다는 의미에서 tag 필드에 used를 대입한다. 객체를 사용이 끝났으므로 다시 Put 함수를 사용하여 객체를 풀에 보관한다.
```
for index := range data.buffer {
	data.buffer[index] = rand.Intn(100)  // 슬라이스에 랜덤 값 저장
}
fmt.Println(data)                            // data 내용 출력
data.tag = "used"                            // 객체가 사용되었다는 태그 설정
pool.Put(data)
```
- 출력 결과를 보면 객체가 새로 할당되는 횟수는 얼마 안되고 대부분 풀에 보관된 객체를 사용하는 것을 알 수 있다.
```
&{new [81 87 47 59 81 18 25 40 56 0]} ← 객체가 새로 할당된 경우
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [41 8 87 31 29 56 37 31 85 26]}
&{new [94 11 62 89 28 74 11 45 37 6]}
&{used [13 90 94 63 33 47 78 24 59 53]} ← 풀에 보관된 객체를 사용
&{used [57 21 89 99 0 5 88 38 3 55]}
&{used [51 10 5 56 66 28 61 2 83 46]}
&{used [63 76 2 18 47 94 77 63 96 20]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [23 53 37 33 41 59 33 43 91 2]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{used [0 2 4 6 8 10 12 14 16 18]}
&{new [95 66 28 58 47 47 87 88 90 15]}
&{new [78 36 46 7 40 3 52 43 5 98]}
```
## 💯대기그룹

대기 그룹은 고루틴이 모두 끝날 때까지 기다릴 때 사용합니다. 앞에서는 time.Sleep, fmt.Scanln 함수를 사용하여 고루틴이 끝날 때까지 임시로 대기했습니다. 이번에는 대기 그룹을 사용하여 고루틴이 끝날 때까지 기다려보겠습니다.

sync 패키지에서 제공하는 대기 그룹의 구조체와 함수는 다음과 같습니다.

- sync.WaitGroup
- func (wg *WaitGroup) Add(delta int): 대기 그룹에 고루틴 개수 추가
- func (wg *WaitGroup) Done(): 고루틴이 끝났다는 것을 알려줄 때 사용
- func (wg *WaitGroup) Wait(): 모든 고루틴이 끝날 때까지 기다림
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1)            // 반복할 때마다 wg.Add 함수로 1씩 추가
		go func(n int) {     // 고루틴 10개 생성
			fmt.Println(n)
			wg.Done()    // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
	fmt.Println("the end")
}
```
- 대기 그룹은 sync.WaitGroup을 할당 한 뒤에 Add, Done, Wait 함수로 사용한다. 고루틴을 생성할 때 Add 함수로 고루틴 개수를 추가해준다 그리고 고루틴 안에서 Done 함수를 사용하여 고루틴이 끝났다는 것을 알려준다. 마지막으로 Wait 함수를 사용하여 모든 고루틴이 끝날 때까지 기다린다.

- Add 함수에 설정한 값과 Done 함수가 호출되는 횟수는 같아야 한다. 즉 Add(3)으로 설정했다면 Done 함수는 3번 호출되야 합니다. 이 횟수가 맞지 않으면 패닉이 발생하므로 주의한다.

- Done 함수는 다음과 같이 defer와 함께 사용해서 지연 호출로도 사용할 수 있다.
```
for i := 0; i < 10; i++ {
	wg.Add(1)
	go func(n int) {
		defer wg.Done() // 고루틴이 끝나기 직전에 wg.Done 함수 호출
		fmt.Println(n)
	}(i)
}
```

## 💯원차적 연산

- 원자적 연산은 더 이상 쪼갤 수 없는 연산이라는 뜻이다. 따라서 여러 스레드(고루틴), CPU 코어에서 같은 변수(메모리)를 수정할 때 서로 영향을 받지 않고 안전하게 연산할 수 있다. 보통 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있디.

- 고루틴을 사용하여 정수형 변수를 2,000번은 더하고, 1,000번은 빼보겠습니다
```
]package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {       // 고루틴 2,000개 생성
			data += 1 // 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {       // 고루틴 1,000개 생성
			data -= 1 // 1씩 뺌
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
```
```
948 (매번 달라질 수 있음)
```
- 실행해보면 0 + 2000 - 1000은 1000이 되어야하는데 그렇지가 않다(실행할 때마다, 시스템마다 실행 결과는 달라질 수 있습니다). 여러 변수에 고루틴이 동시에 접근하면서 정확하게 연산이 되지 않았기 때문이다

- 이번에는 원자적 연산을 사용하여 계산해보겟다.
```
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {                       // 고루틴 2,000개 생성
			atomic.AddInt32(&data, 1) // 원자적 연산으로 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {                        // 고루틴 1,000개 생성
			atomic.AddInt32(&data, -1) // 원자적 연산으로 1씩 뺌
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
```
```
1000
```
- 이제 정확하게 1000이 출력된다.

- 원자적 연산에는 메모리 주소와 수정할 값을 넣습니다. 따라서 atomic.AddInt32(&data, 1)처럼 & (참조)를 사용하여 data 변수의 메모리 주소를 대입한다.

다음은 sync/atomic 패키지에서 제공하는 원자적 연산의 종류입니다.

- Add 계열: 변수에 값을 더하고 결과를 리턴합니다.
- CompareAndSwap 계열: 변수 A와 B를 비교하여 같으면 C를 대입합니다. 그리고 A와 B가 같으면 true, 다르면 false를 리턴합니다.
- Load 계열: 변수에서 값을 가져옵니다.
- Store 계열: 변수에 값을 저장합니다.
- Swap 계열: 변수에 새 값을 대입하고, 이전 값을 리턴합니다.
