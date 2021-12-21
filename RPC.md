# 👩🏻‍🎓RPC

RPC(Remote Procedure Call, 원격 프로시저 호출)는 원격에서 함수를 실행하는 기술입니다. 서버에 있는 함수를 로컬에 있는 함수처럼 사용할 수 있기 때문에 매우 편리합니다. RPC는 Go 언어에서 기본 패키지로 지원하므로 간단하게 구현할 수 있습니다.

## 서버 작성하기
이번에는 원격에서 덧셈 함수를 호출하여 두 수를 더해보겠습니다.

다음은 net/rpc 패키지에서 제공하는 RPC 함수입니다.

- func Register(rcvr interface{}) error: RPC로 사용할 함수 등록
다음 내용을 GOPATH/src/rpcserver/rpcserver.go 파일로 저장합니다. GOPATH를 설정하는 방법은 ‘Unit 3 기본 디렉터리 설정하기’를 참조하기 바랍니다.
```
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의

type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B // 두 값을 더하여 리턴값 구조체에 넣어줌
	return nil
}

func main() {
	rpc.Register(new(Calc)) // Calc 타입의 인스턴스를 생성하여 RPC 서버에 등록
	ln, err := net.Listen("tcp", ":6000") // TCP 프로토콜에 6000번 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // main 함수가 종료되기 직전에 연결 대기를 닫음

	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go rpc.ServeConn(conn) // RPC를 처리하는 함수를 고루틴으로 실행
	}
}
```
- RPC 서버에 함수를 등록하려면 함수만으로는 안되고, 구조체나 일반 자료형과 같은 타입에 메서드 형태로 구성되어 있어야 합니다. 여기서는 Calc 타입을 int 형으로 정의했는데, 다른 자료형이나 빈 구조체로 정의해도 됩니다.
```
type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의
```
- 덧셈 함수의 매개변수 및 리턴값을 정의합니다.
```
type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B // 두 값을 더하여 리턴값 구조체에 넣어줌
	return nil
}
```
- RPC로 호출되는 함수는 매개변수 두 개와 error 형 리턴값 형식입니다. 따라서 매개변수를 위한 구조체 Args와 리턴값을 위한 구조체(Reply)를 정의해줍니다. 여기서 구조체 필드의 이름에서 첫글자는 항상 대문자로 시작해야 합니다.

- 덧셈 함수 Sum을 Calc 타입의 메서드로 정의해줍니다. 그리고 reply.C = args.A + args.B처럼 매개변수 구조체에서 받은 값을 더해서 리턴값 구조체에 넣어줍니다. 매개변수와 리턴값이 구조체 형식이므로 다양한 자료형을 사용할 수 있고, 리턴값의 개수도 제한이 없습니다.

- RPC 서버에서 함수가 처리될 수 있도록 등록을 해주고, TCP 연결을 받을 준비를 합니다.
```
rpc.Register(new(Calc)) // Calc 타입의 인스턴스를 생성하여 RPC 서버에 등록
ln, err := net.Listen("tcp", ":6000") // TCP 프로토콜에 6000번 포트로 연결을 받음
if err != nil {
	fmt.Println(err)
	return
}
defer ln.Close() // main 함수가 종료되기 직전에 연결 대기를 닫음
```
- 우리는 Calc 타입의 메서드를 사용할 것이므로 Calc 타입의 메모리를 할당하여 rpc.Register 함수에 넣어줍니다. 그리고 net.Listen 함수에 tcp와 포트 번호를 설정합니다. “:6000”처럼 포트 번호만 설정하면 모든 네트워크 인터페이스(NIC)의 IP 주소에서 연결을 받고, “192.168.0.100:6000”처럼 IP 주소와 함께 설정하면 특정 NIC에서만 TCP 연결을 받습니다. 그리고 TCP 연결 대기 ln은 지연 호출을 사용하여 서버가 끝나면 닫아줍니다.
- for 반복문으로 무한 루프를 돌면서 클라이언트의 연결을 기다립니다. 클라이언트와 연결되면 Accept 함수에서 TCP 연결(커넥션) conn이 리턴됩니다. 그 뒤 RPC를 처리할 rpc.ServeConn 함수를 고루틴으로 실행해줍니다. 마찬가지로 conn도 지연 호출을 사용하여 서버가 끝나면 닫아줍니다.
```
for {
	conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
	if err != nil {
		continue
	}
	defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	go rpc.ServeConn(conn) // RPC를 처리하는 함수를 고루틴으로 실행
}
```
이제 RPC 요청이 오면 알아서 적절한 함수를 실행해줍니다.

## 💯클라이언트 작성하기
이제 클라이언트를 작성합니다.

다음은 net/rpc 패키지에서 제공하는 RPC 함수입니다.

- func Dial(network, address string) (*Client, error): 프로토콜, IP 주소, 포트 번호를 설정하여 RPC 서버에 연결
- func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error: RPC 서버의 함수를 호출(동기)
- func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call: RPC 서버의 함수를 고루틴으로 호출(비동기)
다음 내용을 GOPATH/src/rpcclient/rpcclient.go 파일로 저장합니다.

```
package main

import (
	"fmt"
	"net/rpc"
)

type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6000") // RPC 서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기 직전에 RPC 연결을 닫음

	// 동기 호출
	args := &Args{1, 2}
	reply := new(Reply)
	err = client.Call("Calc.Sum", args, reply) // Calc.Sum 함수 호출
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	// 비동기 호출
	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil) // Calc.Sum 함수를 고루틴으로 호출
	<-sumCall.Done // 함수가 끝날 때까지 대기
	fmt.Println(reply.C)
}
```
RPC 서버의 함수를 호출할 때 사용할 매개변수와 리턴값 구조체를 정의합니다. 이 부분은 서버와 동일하게 작성해야 하며 서버와 공유하는 패키지로 만들면 편리합니다.
```
type Args struct { // 매개변수
	A, B int
}

type Reply struct { // 리턴값
	C int
}
```
먼저 net.Dial 함수에 tcp와 서버의 IP 주소 및 포트 번호를 설정합니다. 서버에 연결되면 net.Conn이 리턴되고 데이터를 받거나 보낼 수 있습니다. 지연 호출을 사용하여 클라이언트가 끝나면 TCP 연결을 닫아줍니다.

```
client, err := rpc.Dial("tcp", "127.0.0.1:6000") // RPC 서버에 연결
if err != nil {
	fmt.Println(err)
	return
}
defer client.Close() // main 함수가 끝나기 직전에 RPC 연결을 닫음
```
이제 서버의 Sum 함수를 실행합니다.
```
// 동기 호출
args := &Args{1, 2}
reply := new(Reply)
err = client.Call("Calc.Sum", args, reply) // Calc.Sum 함수 호출
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(reply.C)
```
매개변수 구조체 args를 채워넣고, 리턴값을 담을 reply도 메모리를 할당합니다. 그리고 client.Call 함수에 실행할 함수 이름, args, reply를 넣어줍니다. 서버에서 Sum 함수를 작성할 때 Calc 타입에 Sum 메서드로 만들었으므로 함수 이름은 Calc.Sum이 됩니다. client.Call 함수가 성공하면 reply에 결괏값이 저장됩니다.

이번에는 비동기로 서버의 Sum 함수를 실행하는 방법입니다.
```
// 비동기 호출
args.A = 4
args.B = 9
sumCall := client.Go("Calc.Sum", args, reply, nil) // Calc.Sum 함수를 고루틴으로 호출
<-sumCall.Done // 함수가 끝날 때까지 대기
fmt.Println(reply.C)
```
- client.Go 함수에 실행할 함수 이름, args, reply를 넣어주면 rpc.Call 인스턴스가 리턴됩니다. 마지막 매개변수는 함수 실행이 끝났는지 확인하기 위한 채널입니다. 여기에 nil을 넣으면 채널이 새로 할당되어 리턴됩니다
- 마지막으로 <-sumCall.Done처럼 Done 채널에 값이 들어올 때까지 기다리면 됩니다.

- 이제 서버와 클라이언트의 소스 파일을 컴파일한 뒤 실행해봅니다. 먼저 콘솔(터미널)에서 GOPATH/src/rpcserver 디렉터리로 이동한 뒤 서버를 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였습니다.

```
~$ cd $GOPATH/src/rpcserver
~/hello_project/src/rpcserver$ ./rpcserver
```
서버를 실행했으면 콘솔에서 GOPATH/src/rpcclient 디렉터리로 이동한 뒤 클라이언트를 실행합니다.
```
~$ cd $GOPATH/src/rpcclient
~/hello_project/src/rpcclient$ ./rpcclient
3
13
```
클라이언트를 실행하면 1과 2를 더한 값 3 그리고 4와 9를 더한 값 13이 출력됩니다.
