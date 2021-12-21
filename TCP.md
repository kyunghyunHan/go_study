# 🌱TCP
Go 언어는 기본 패키지에서 다양한 네트워크 프로토콜을 제공합니다. 그중 TCP 프로토콜은 실시간 처리가 중요한 게임 서버에서 많이 사용됩니다. 또한, HTTP, FTP, SMTP, DHCP 프로토콜 등 인터넷에서 흔히 사용하는 프로토콜도 TCP 프로토콜을 기반으로 구현되어 있습니다.

## 💯서버 작성하기
이번에는 클라이언트에서 보낸 메시지를 다시 응답해주는 TCP 서버를 만들어보겠습니다.

다음은 net 패키지에서 제공하는 TCP 함수입니다.

- func Listen(net, laddr string) (Listener, error): 프로토콜, IP 주소, 포트 번호를 설정하여 네트워크 연결을 대기합니다.
- func (l *TCPListener) Accept() (Conn, error): 클라이언트가 연결되면 TCP 연결(커넥션)을 리턴
- func (l *TCPListener) Close() error: TCP 연결 대기를 닫음
- func (c *TCPConn) Read(b []byte) (int, error): 받은 데이터를 읽음
- func (c *TCPConn) Write(b []byte) (int, error): 데이터를 보냄
- func (c *TCPConn) Close() error: TCP 연결을 닫음
다음 내용을 GOPATH/src/tcpserver/tcpserver.go 파일로 저장합니다. GOPATH를 설정하는 방법은 ‘Unit 3 기본 디렉터리 설정하기’를 참조하기 바랍니다.

```
package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") // TCP 프로토콜에 8000 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음

	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}
}
```

먼저 net.Listen 함수에 tcp와 포트 번호를 설정하여 TCP 연결을 받을 준비를 합니다.

```ln, err := net.Listen("tcp", ":8000") // TCP 프로토콜에 8000 포트로 연결을 받음
if err != nil {
	fmt.Println(err)
	return
}
defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 닫음
```
“:8000”처럼 포트 번호만 설정하면 모든 네트워크 인터페이스(NIC)의 IP 주소에서 연결을 받고, “192.168.0.100:8000”처럼 IP 주소와 함께 설정하면 특정 NIC에서만 TCP 연결을 받습니다. 그리고 TCP 연결 대기 ln은 지연 호출을 사용하여 서버가 끝나면 닫아줍니다.

for 반복문으로 무한 루프를 돌면서 클라이언트의 연결을 기다립니다. 클라이언트와 연결되면 Accept 함수에서 TCP 연결(커넥션) conn이 리턴됩니다. 그 뒤 패킷을 처리할 requestHandler 함수를 고루틴으로 실행해줍니다. 마찬가지로 conn도 지연 호출을 사용하여 서버가 끝나면 닫아줍니다.
```
for {
	conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
	if err != nil {
		fmt.Println(err)
		continue
	}
	defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
}
```
클라이언트에서 받은 패킷을 처리하고, 클라이언트로 패킷을 보내는 requestHandler 함수입니다.
```
func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 가져옴
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
```
for 반복문으로 무한 루프를 돌면서 클라이언트에서 보낸 데이터를 읽어서 다시 클라이언트로 보내는 간단한 구조입니다. TCP 연결 c에서 Read 함수로 클라이언트에서 보낸 데이터를 가져옵니다. 여기서는 4096 크기로 된 []byte 형식의 슬라이스를 사용했고, 바로 문자열(텍스트) 형식으로 변환하여 출력하였습니다. 그리고 c에서 Write 함수로 클라이언트에 데이터를 보냅니다.

보통 TCP 서버에서 클라이언트와 패킷(데이터)을 주고받을 때는 패킷의 최대 크기(길이)를 약속하고, 각 패킷의 크기와 구조를 정의하여 사용합니다. 특히 MMORPG 같은 온라인 게임에서는 패킷의 구조를 독자적으로 만들어서 사용합니다. 패킷 처리를 편리하게 하려면 BSON, Protocol Buffers(protobuf), MessagePack(msgpack) 등의 라이브러리를 사용하면 됩니다.

서버와 클라이언트를 모두 Go 언어로 작성한다면 encoding/gob 패키지가 편리합니다

## 💯클라이언트 작성하기
이제 클라이언트를 작성합니다.

다음은 net 패키지에서 제공하는 TCP 함수입니다.

- func Dial(network, address string) (Conn, error): 프로토콜, IP 주소, 포트 번호를 설정하여 서버에 연결
- func (c *TCPConn) Close() error: TCP 연결을 닫음
- func (c *TCPConn) Read(b []byte) (int, error): 받은 데이터를 읽음
- func (c *TCPConn) Write(b []byte) (int, error): 데이터를 보냄
다음 내용을 GOPATH/src/tcpclient/tcpclient.go 파일로 저장합니다.

```
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000") // TCP 프로토콜, 127.0.0.1:8000 서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	go func(c net.Conn) {
		data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

		for {
			n, err := c.Read(data) // 서버에서 받은 데이터를 읽음
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n])) // 데이터 출력

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "Hello " + strconv.Itoa(i)

			_, err := c.Write([]byte(s)) // 서버로 데이터를 보냄
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
```
예제를 좀 더 명확하게 하고자 이 클라이언트는 서버에 데이터를 보내는 고루틴과 서버에서 받은 데이터를 출력하는 고루틴 두 부분으로 구성되어 있습니다.

먼저 net.Dial 함수에 tcp와 서버의 IP 주소 및 포트 번호를 설정합니다. 서버에 연결되면 net.Conn이 리턴되고 데이터를 받거나 보낼 수 있습니다. 지연 호출을 사용하여 클라이언트가 끝나면 TCP 연결을 닫아줍니다
```
client, err := net.Dial("tcp", "127.0.0.1:8000") // TCP 프로토콜, 127.0.0.1:8000 서버에 연결
if err != nil {
	fmt.Println(err)
	return
}
defer client.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음
```
'
서버에 데이터를 보내는 고루틴입니다. 마찬가지로 TCP 연결 c에 Write 함수로 서버에 데이터를 보냅니다. 여기서는 1초마다 Hello 문자열에 숫자를 1씩 증가시켜서 보냅니다.

```
go func(c net.Conn) {
	i := 0
	for {
		s := "Hello " + strconv.Itoa(i)

		_, err := c.Write([]byte(s)) // 서버로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}

		i++
		time.Sleep(1 * time.Second)
	}
}(client)
```
서버에서 데이터를 받은 뒤 출력하는 고루틴입니다. Read 함수로 서버에서 받은 데이터를 가져옵니다.
```
go func(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

	for {
		n, err := c.Read(data) // 서버에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		time.Sleep(1 * time.Second)
	}
}(client)
```
서버와 클라이언트의 소스 파일을 컴파일한 뒤 실행해봅니다. 먼저 콘솔(터미널)에서 GOPATH/src/tcpserver 디렉터리로 이동한 뒤 서버를 실행합니다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행합니다). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였습니다.
```
~$ cd $GOPATH/src/tcpserver
~/hello_project/src/tcpserver$ ./tcpserver
```
서버를 실행했으면 콘솔에서 GOPATH/src/tcpclient 디렉터리로 이동한 뒤 클라이언트를 실행합니다.
```
~$ cd $GOPATH/src/tcpclient
~/hello_project/src/tcpclient$ ./tcpclient
Hello 0
Hello 1
Hello 2
Hello 3
... (생략)
```
클라이언트를 실행하면 Hello 숫자가 1초마다 출력됩니다. 마찬가지로 서버쪽 콘솔에서도 똑같이 출력됩니다.

```
Hello 0
Hello 1
Hello 2
Hello 3
... (생략)
```
