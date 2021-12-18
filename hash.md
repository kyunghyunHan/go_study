# 해시 
Go 언어에서는 해시, 대칭키 알고리즘, 공개키 알고리즘을 비롯해 다양한 암호화 알고리즘을 패키지로 제공합니다. 보통 암호화 알고리즘은 파일을 암호화하거나, 네트워크로 데이터를 주고 받을 때 데이터를 보호하기 위해 사용합니다.

- 해시(Hash) 알고리즘: MD5, SHA1, SHA256, SHA512 등이 있으며 데이터에서 고유한 해시 값을 추출해냅니다. 이 해시 값으로 데이터를 만들어낼 수는 없습니다. 단방향 암호화 알고리즘이라고도 하며 패스워드를 저장할 때 사용합니다.
- 대칭키 알고리즘(Symmetric-key algorithm): AES, DES, RC4 등이 있으며 암호화하는 키와 복호화하는 키가 동일합니다.
- 공개키 알고리즘(Public key infrastructure, PKI): RSA가 대표적이며 암호화하는 키와 복호화하는 키가 다릅니다. 암호화하는 키는 공개 키라고 하여 외부에 공개하고, 복호화하는 키는 비밀 키(개인 키)라고하여 외부에 노출하지 않습니다. 공개 키로는 비밀 키를 만들어내기 힘들도록 설계되어 있습니다.
해시 알고리즘 사용하기
다음은 crypto/sha512 패키지에서 제공하는 해시 알고리즘 함수입니다.

- func New() hash.Hash: SHA512 해시 인스턴스 생성
- func Sum512(data []byte) [Size]byte: SHA512 해시를 계산하여 리턴
- func (d *digest) Write(p []byte) (nn int, err error): 해시 인스턴스에 데이터 추가
- func (d0 *digest) Sum(in []byte) []byte: 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출
SHA512 알고리즘을 사용하여 데이터에서 해시 값을 추출해보겠습니다.

```
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s)) // 문자열의 SHA512 해시 값 추출
	fmt.Printf("%x\n", h1)

	sha := sha512.New()          // SHA512 해시 인스턴스 생성
	sha.Write([]byte("Hello, ")) // 해시 인스턴스에 데이터 추가
	sha.Write([]byte("world!"))  // 해시 인스턴스에 데이터 추가
	h2 := sha.Sum(nil)           // 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출
	fmt.Printf("%x\n", h2)
}
```

```
c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421
c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421
```

SHA512 해시 값을 추출하는 방법은 간단합니다. sha512.Sum512 함수에 []byte 형식으로 데이터를 넣어주면 해시 값이 리턴됩니다.

sha512.New 함수를 사용하여 새로운 인스턴스를 생성하였다면 Write 함수로 데이터를 넣은 뒤 Sum 함수로 해시 값을 만들면 됩니다. 이때 Sum 함수에는 nil을 넣어줍니다.


## AES 대칭키 알고리즘 사용하기
다음은 crypto/aes 패키지에서 제공하는 대칭키 알고리즘 함수입니다.

- func NewCipher(key []byte) (cipher.Block, error): 대칭키 암호화 블록 생성
- func (c *aesCipher) Encrypt(dst, src []byte): 평문을 AES 알고리즘으로 암호화
- func (c *aesCipher) Decrypt(dst, src []byte): AES 알고리즘으로 암호화된 데이터를 평문으로 복호화
이번에는 AES 대칭키 알고리즘을 사용하여 데이터를 암호화하고 복호화해보겠습니다.
```
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "Hello, Key 12345" // 16바이트
	s := "Hello, world! 12"   // 16바이트

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) // 평문을 AES 알고리즘으로 암호화
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext) // AES 알고리즘으로 암호화된 데이터를 평문으로 복호화
	fmt.Println(string(plaintext))
}
```
```
a20455c66b97529fa756a0c9a7d2f329
Hello, world! 12
```

AES는 블록 암호화 알고리즘이므로 키와 암호화할 데이터의 크기가 일정해야 합니다. 여기서는 키와 데이터 모두 16바이트로 만들었습니다.

aes.NewCipher 함수에 키를 넣으면 암호화 블록(cipher.Block)이 리턴됩니다. 그리고 Encrypt 함수에 데이터(s)와 암호화된 데이터을 저장할 슬라이스(ciphertext)를 넣으면 암호화가 됩니다. 마찬가지로 Decrypt 함수에 암호화된 데이터와 복호화된 데이터를 저장할 슬라이스를 넣으면 복호화가 됩니다.


긴 데이터를 안전하게 암호화하기 위해 대칭키 알고리즘은 다양한 운용 방식을 제공합니다. 그중 CBC(Cipher Block Chaining) 방식을 사용하여 암호화를 해보겠습니다.

다음은 crypto/cipher 패키지에서 제공하는 암호화 운용 모드 함수입니다.

- func NewCBCEncrypter(b Block, iv []byte) BlockMode: 암호화 블록과 초기화 벡터로 암호화 블록 모드 인스턴스 생성
- func (x *cbcEncrypter) CryptBlocks(dst, src []byte): 암호화 블록 모드 인스턴스로 암호화
- func NewCBCDecrypter(b Block, iv []byte) BlockMode: 암호화 블록과 초기화 벡터로 복호화 블록 모드 인스턴스 생성
- func (x *cbcDecrypter) CryptBlocks(dst, src []byte): 복호화 블록 모드 인스턴스로 복호화
다음은 io 패키지에서 제공하는 읽기 함수입니다.

- func ReadFull(r Reader, buf []byte) (n int, err error): io.Reader에서 buf의 길이만큼 데이터를 읽음

```
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 { // 블록 크기의 배수가 되어야함
		padding := make([]byte, aes.BlockSize-mod)   // 블록 크기에서 모자라는 부분을
		plaintext = append(plaintext, padding...)    // 채워줌
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
	iv := ciphertext[:aes.BlockSize] // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // 랜덤 값을 초기화 벡터에 넣어줌
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv) // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // 암호화 블록 모드 인스턴스로
                                                                // 암호화

	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 { // 블록 크기의 배수가 아니면 리턴
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]        // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	ciphertext = ciphertext[aes.BlockSize:] // 부분 슬라이스로 암호화된 데이터를 가져옴

	plaintext := make([]byte, len(ciphertext)) // 평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(b, iv)      // 암호화 블록과 초기화 벡터를 넣어서
                                                   // 복호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(plaintext, ciphertext)    // 복호화 블록 모드 인스턴스로 복호화

	return plaintext
}

func main() {
	key := "Hello Key 123456" // 16바이트

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s)) // 평문을 AES 알고리즘으로 암호화
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext) // AES 알고리즘 암호문을 평문으로 복호화
	fmt.Println(string(plaintext))

}
```
```
6b39496c974990a9890f49ec962f92cd00f3009f724e3e6a7bcdd7b37867... (생략)
동해 물과 백두산이 마르고 닳도록... (생략)
```
먼저 블록 암호화는 암호화할 데이터의 길이가 블록 크기(aes.BlockSize)의 배수라야 합니다. 따라서 다음과 같이 블록 크기의 배수가 될 수 있도록 모자라는 부분을 채워줍니다.

```
if mod := len(plaintext) % aes.BlockSize; mod != 0 { // 블록 크기의 배수가 되어야함
	padding := make([]byte, aes.BlockSize-mod)   // 블록 크기에서 모자라는 부분을
	plaintext = append(plaintext, padding...)    // 채워줌
}
```
CBC 운용 방식은 초기화 벡터(Initialization Vector, IV)를 첫 블록에 배치시키고, 각 블록은 이전 블록의 암호화 결과와 XOR됩니다. 따라서 초기화 벡터는 암호화할 때마다 매번 다른 값으로 생성해야 합니다.
```
ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
iv := ciphertext[:aes.BlockSize] // 부분 슬라이스로 초기화 벡터 공간을 가져옴
if _, err := io.ReadFull(rand.Reader, iv); err != nil { // 랜덤 값을 초기화 벡터에 넣어줌
	fmt.Println(err)
	return nil
}
```
암호화 데이터를 저장할 슬라이스는 초기화 벡터가 들어갈 공간(aes.BlockSize)만큼 더 생성해줍니다. 그리고 rand.Reader를 io.ReadFull 함수로 읽으면 랜덤 값을 얻을 수 있습니다. 이 랜덤 값을 초기화 벡터 iv에 넣어줍니다.

cipher.NewCBCEncrypter 함수에 암호화 블록 b와 초기화 벡터 iv를 넣어준 뒤 CryptBlocks 함수로 암호화를 하면 됩니다.
```
mode := cipher.NewCBCEncrypter(b, iv) // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // 암호화 블록 모드 인스턴스로 암호화

```
이번에는 복호화를 알아보겠습니다. 복호화를 할 때 암호화된 데이터의 길이는 블록 크기의 배수가 되어야 합니다. 따라서 복호화하기 전에 항상 길이를 검사합니다.
```
if len(ciphertext)%aes.BlockSize != 0 { // 블록 크기의 배수가 아니면 리턴
	fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
	return nil
}
```
암호화된 데이터(ciphertext)의 첫 블록에서 iv를 꺼냅니다. 그리고 cipher.NewCBCDecrypter 함수에 암호화 블록 b와 초기화 벡터 iv를 넣어준 뒤 CryptBlocks 함수로 복호화를 하면 됩니다.
```
iv := ciphertext[:aes.BlockSize]        // 부분 슬라이스로 초기화 벡터 공간을 가져옴
ciphertext = ciphertext[aes.BlockSize:] // 부분 슬라이스로 암호화된 데이터를 가져옴

plaintext := make([]byte, len(ciphertext)) // 평문 데이터를 저장할 공간 생성
mode := cipher.NewCBCDecrypter(b, iv)      // 암호화 블록과 초기화 벡터를 넣어서
                                           // 복호화 블록 모드 인스턴스 생성
mode.CryptBlocks(plaintext, ciphertext)    // 복호화 블록 모드 인스턴스로 복호화
```
## RSA 공개키 알고리즘 사용하기
대칭키 알고리즘은 암호 키가 유출되면 암호화된 데이터를 모두 풀 수 있습니다. 특히 네트워크로 암호 키를 주고 받으면 노출될 위험이 커집니다. 따라서 네트워크로 데이터를 주고 받을때는 공개키 암호화 알고리즘을 많이 사용합니다. 단 공개키 알고리즘은 대칭키 알고리즘에 비해 속도가 느리므로 대칭키 알고리즘의 암호 키만 공개키 알고리즘으로 암호화하여 네트워크로 전송하기도 합니다.

다음은 crypto/rsa 패키지에서 제공하는 공개키 알고리즘 함수입니다.

- func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error): 개인 키와 공개키 생성
- func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error): 평문을 공개 키로 암호화
- func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error): 암호화된 데이터를 개인 키로 복호화
RSA 공개키 알고리즘을 사용하여 암호화와 복호화를 해보겠습니다.
```
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개키 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	ciphertext, err := rsa.EncryptPKCS1v15( // 평문을 공개 키로 암호화
		rand.Reader,
		publicKey, // 공개키
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15( // 암호화된 데이터를 개인 키로 복호화
		rand.Reader,
		privateKey, // 개인키
		ciphertext,
	)

	fmt.Println(string(plaintext))
}
```
rsa.GenerateKey 함수에 rand.Reader와 키의 길이를 넣으면 개인 키와 공개키가 생성됩니다. 키의 길이는 비트 단위이며 2의 제곱으로 입력합니다.
rsa.EncryptPKCS1v15 함수에 rand.Reader, publicKey, 암호화할 데이터를 넣으면 암호화가 됩니다. 그리고 rsa.DecryptPKCS1v15 함수에 rand.Reader, privateKey, 암호화된 데이터를 넣으면 복호화가 됩니다.

이처럼 공개키 알고리즘은 암호화할 때는 공개 키를 사용하고, 복호화할 때는 개인 키를 사용합니다. 공개 키를 외부에 공개하여 다른 사람한테 전달한 뒤 암호화를 하고, 암호화된 데이터를 받는 사람은 자기가 가진 개인키로 복호화하는 방식입니다. 개인 키는 외부에 노출할 일이 없으며 공개 키로는 개인 키를 추출하기 어렵기 때문에 안전하게 암호화된 데이터를 주고 받을 수 있습니다.

이번에는 RSA 알고리즘을 사용하여 메시지를 서명(Signing)하고 인증(Verification)하는 방법을 알아보겠습니다.

다음은 crypto/rsa 패키지에서 제공하는 서명, 인증 함수입니다.

- func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error): 개인 키로 서명
- func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (err error): 공개키로 서명 검증
다음은 crypto/md5 패키지에서 제공하는 해시 함수입니다.

- func New() hash.Hash: MD5 해시 인스턴스 생성
- func (d *digest) Write(p []byte) (nn int, err error): 해시 인스턴스에 데이터 추가
- func (d0 *digest) Sum(in []byte) []byte: 해시 인스턴스에 저장된 데이터의 MD5 해시 값 추출

```
package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개 키 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음

	message := "안녕하세요. Go 언어"
	hash := md5.New()           // 해시 인스턴스 생성
	hash.Write([]byte(message)) // 해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil)     // 문자열의 MD5 해시 값 추출

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15( // 개인 키로 서명
		rand.Reader,
		privateKey, // 개인 키
		h1,
		digest,     // MD5 해시 값
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15( // 공개 키로 서명 검증
		publicKey, // 공개 키
		h2,
		digest,    // MD5 해시 값
		signature, // 서명 값
	)

	if err != nil {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}
}
```
먼저 서명할 데이터의 해시 값을 구합니다. 여기서는 md5로 해시 값을 구합니다.
```
message := "안녕하세요. Go 언어"
hash := md5.New()           // 해시 인스턴스 생성
hash.Write([]byte(message)) // 해시 인스턴스에 문자열 추가
digest := hash.Sum(nil)     // 문자열의 MD5 해시 값 추출
```
이제 “안녕하세요. Go 언어”의 해시 값을 개인 키로 서명합니다. 여기서 서명이란 내가 “안녕하세요. Go 언어”라는 메시지를 보냈다고 보증하는 것입니다.
```
var h1 crypto.Hash
signature, err := rsa.SignPKCS1v15( // 개인 키로 서명
	rand.Reader,
	privateKey, // 개인 키
	h1,
	digest,     // MD5 해시 값
)
```
rsa.SignPKCS1v15 함수에 rand.Reader, crypto.Hash, 개인 키, 해시 값을 넣으면 서명이 됩니다. 이렇게 해서 메시지(message), 메시지에 대한 해시 값(digest), 서명(signature)이 모두 준비되었습니다.

메시지를 받는 사람은 정말 저 사람(개인 키를 가진 사람)이 “안녕하세요. Go 언어”라는 말을 했는지 검증(인증)을 해야됩니다.
```
var h2 crypto.Hash
err = rsa.VerifyPKCS1v15( // 공개 키로 서명 검증
	publicKey, // 공개 키
	h2,
	digest,    // MD5 해시 값
	signature, // 서명 값
)
```
rsa.VerifyPKCS1v15 함수에 공개 키, crypto.Hash, 해시 값, 서명을 넣으면 검증이 됩니다. err 값이 nil이면 정상적인 메시지입니다.

메시지, 메시지에 대한 해시 값, 서명, 공개 키는 모두 공개된 정보입니다. 이 값들을 이용해서 메시지가 변조되지 않고, 올바른 사람한테서 왔는지 검증하는 것입니다.

- 메시지에 대한 해시 값으로 메시지가 변조되었는지 확인할 수 있습니다.
- 서명 값으로 메시지가 올바른 사람한테서 왔는지 확인할 수 있습니다. 즉 서명을 할 수 있는 사람은 개인 키를 가진 사람밖에 없기 때문입니다.
개인 키를 가지고 서명을 하는 주체는 사람이 될 수도 있고 컴퓨터가 될 수도 있습니다. 여기서는 편의상 사람이라 칭했습니다. 보통 금융거래에서 많이 쓰는 공인인증서가 서명과 인증 방식에 RSA 알고리즘을 사용합니다. 다른 사례로는 인터넷에서 파일을 받을 때 해시 값과 서명 파일이 함께 있는 곳이 있습니다. 이때는 해시 값으로 파일 변조 여부를 확인하고, 서명 파일로 공식 개발자(회사, 단체)가 배포한 것인지 검증할 수 있습니다.
