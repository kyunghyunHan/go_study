# π©π»βπν΄μ 
Go μΈμ΄μμλ ν΄μ, λμΉ­ν€ μκ³ λ¦¬μ¦, κ³΅κ°ν€ μκ³ λ¦¬μ¦μ λΉλ‘―ν΄ λ€μν μνΈν μκ³ λ¦¬μ¦μ ν¨ν€μ§λ‘ μ κ³΅ν©λλ€. λ³΄ν΅ μνΈν μκ³ λ¦¬μ¦μ νμΌμ μνΈννκ±°λ, λ€νΈμν¬λ‘ λ°μ΄ν°λ₯Ό μ£Όκ³  λ°μ λ λ°μ΄ν°λ₯Ό λ³΄νΈνκΈ° μν΄ μ¬μ©ν©λλ€.

- ν΄μ(Hash) μκ³ λ¦¬μ¦: MD5, SHA1, SHA256, SHA512 λ±μ΄ μμΌλ©° λ°μ΄ν°μμ κ³ μ ν ν΄μ κ°μ μΆμΆν΄λλλ€. μ΄ ν΄μ κ°μΌλ‘ λ°μ΄ν°λ₯Ό λ§λ€μ΄λΌ μλ μμ΅λλ€. λ¨λ°©ν₯ μνΈν μκ³ λ¦¬μ¦μ΄λΌκ³ λ νλ©° ν¨μ€μλλ₯Ό μ μ₯ν  λ μ¬μ©ν©λλ€.
- λμΉ­ν€ μκ³ λ¦¬μ¦(Symmetric-key algorithm): AES, DES, RC4 λ±μ΄ μμΌλ©° μνΈννλ ν€μ λ³΅νΈννλ ν€κ° λμΌν©λλ€.
- κ³΅κ°ν€ μκ³ λ¦¬μ¦(Public key infrastructure, PKI): RSAκ° λνμ μ΄λ©° μνΈννλ ν€μ λ³΅νΈννλ ν€κ° λ€λ¦λλ€. μνΈννλ ν€λ κ³΅κ° ν€λΌκ³  νμ¬ μΈλΆμ κ³΅κ°νκ³ , λ³΅νΈννλ ν€λ λΉλ° ν€(κ°μΈ ν€)λΌκ³ νμ¬ μΈλΆμ λΈμΆνμ§ μμ΅λλ€. κ³΅κ° ν€λ‘λ λΉλ° ν€λ₯Ό λ§λ€μ΄λ΄κΈ° νλ€λλ‘ μ€κ³λμ΄ μμ΅λλ€.
ν΄μ μκ³ λ¦¬μ¦ μ¬μ©νκΈ°
λ€μμ crypto/sha512 ν¨ν€μ§μμ μ κ³΅νλ ν΄μ μκ³ λ¦¬μ¦ ν¨μμλλ€.

- func New() hash.Hash: SHA512 ν΄μ μΈμ€ν΄μ€ μμ±
- func Sum512(data []byte) [Size]byte: SHA512 ν΄μλ₯Ό κ³μ°νμ¬ λ¦¬ν΄
- func (d *digest) Write(p []byte) (nn int, err error): ν΄μ μΈμ€ν΄μ€μ λ°μ΄ν° μΆκ°
- func (d0 *digest) Sum(in []byte) []byte: ν΄μ μΈμ€ν΄μ€μ μ μ₯λ λ°μ΄ν°μ SHA512 ν΄μ κ° μΆμΆ
SHA512 μκ³ λ¦¬μ¦μ μ¬μ©νμ¬ λ°μ΄ν°μμ ν΄μ κ°μ μΆμΆν΄λ³΄κ² μ΅λλ€.

```
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s)) // λ¬Έμμ΄μ SHA512 ν΄μ κ° μΆμΆ
	fmt.Printf("%x\n", h1)

	sha := sha512.New()          // SHA512 ν΄μ μΈμ€ν΄μ€ μμ±
	sha.Write([]byte("Hello, ")) // ν΄μ μΈμ€ν΄μ€μ λ°μ΄ν° μΆκ°
	sha.Write([]byte("world!"))  // ν΄μ μΈμ€ν΄μ€μ λ°μ΄ν° μΆκ°
	h2 := sha.Sum(nil)           // ν΄μ μΈμ€ν΄μ€μ μ μ₯λ λ°μ΄ν°μ SHA512 ν΄μ κ° μΆμΆ
	fmt.Printf("%x\n", h2)
}
```

```
c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421
c1527cd893c124773d811911970c8fe6e857d6df5dc9226bd8a160614c0cd963a4ddea2b94bb7d36021ef9d865d5cea294a82dd49a0bb269f51f6e7a57f79421
```

SHA512 ν΄μ κ°μ μΆμΆνλ λ°©λ²μ κ°λ¨ν©λλ€. sha512.Sum512 ν¨μμ []byte νμμΌλ‘ λ°μ΄ν°λ₯Ό λ£μ΄μ£Όλ©΄ ν΄μ κ°μ΄ λ¦¬ν΄λ©λλ€.

sha512.New ν¨μλ₯Ό μ¬μ©νμ¬ μλ‘μ΄ μΈμ€ν΄μ€λ₯Ό μμ±νμλ€λ©΄ Write ν¨μλ‘ λ°μ΄ν°λ₯Ό λ£μ λ€ Sum ν¨μλ‘ ν΄μ κ°μ λ§λ€λ©΄ λ©λλ€. μ΄λ Sum ν¨μμλ nilμ λ£μ΄μ€λλ€.


## π―AES λμΉ­ν€ μκ³ λ¦¬μ¦ μ¬μ©νκΈ°
λ€μμ crypto/aes ν¨ν€μ§μμ μ κ³΅νλ λμΉ­ν€ μκ³ λ¦¬μ¦ ν¨μμλλ€.

- func NewCipher(key []byte) (cipher.Block, error): λμΉ­ν€ μνΈν λΈλ‘ μμ±
- func (c *aesCipher) Encrypt(dst, src []byte): νλ¬Έμ AES μκ³ λ¦¬μ¦μΌλ‘ μνΈν
- func (c *aesCipher) Decrypt(dst, src []byte): AES μκ³ λ¦¬μ¦μΌλ‘ μνΈνλ λ°μ΄ν°λ₯Ό νλ¬ΈμΌλ‘ λ³΅νΈν
μ΄λ²μλ AES λμΉ­ν€ μκ³ λ¦¬μ¦μ μ¬μ©νμ¬ λ°μ΄ν°λ₯Ό μνΈννκ³  λ³΅νΈνν΄λ³΄κ² μ΅λλ€.
```
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "Hello, Key 12345" // 16λ°μ΄νΈ
	s := "Hello, world! 12"   // 16λ°μ΄νΈ

	block, err := aes.NewCipher([]byte(key)) // AES λμΉ­ν€ μνΈν λΈλ‘ μμ±
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) // νλ¬Έμ AES μκ³ λ¦¬μ¦μΌλ‘ μνΈν
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext) // AES μκ³ λ¦¬μ¦μΌλ‘ μνΈνλ λ°μ΄ν°λ₯Ό νλ¬ΈμΌλ‘ λ³΅νΈν
	fmt.Println(string(plaintext))
}
```
```
a20455c66b97529fa756a0c9a7d2f329
Hello, world! 12
```

AESλ λΈλ‘ μνΈν μκ³ λ¦¬μ¦μ΄λ―λ‘ ν€μ μνΈνν  λ°μ΄ν°μ ν¬κΈ°κ° μΌμ ν΄μΌ ν©λλ€. μ¬κΈ°μλ ν€μ λ°μ΄ν° λͺ¨λ 16λ°μ΄νΈλ‘ λ§λ€μμ΅λλ€.

aes.NewCipher ν¨μμ ν€λ₯Ό λ£μΌλ©΄ μνΈν λΈλ‘(cipher.Block)μ΄ λ¦¬ν΄λ©λλ€. κ·Έλ¦¬κ³  Encrypt ν¨μμ λ°μ΄ν°(s)μ μνΈνλ λ°μ΄ν°μ μ μ₯ν  μ¬λΌμ΄μ€(ciphertext)λ₯Ό λ£μΌλ©΄ μνΈνκ° λ©λλ€. λ§μ°¬κ°μ§λ‘ Decrypt ν¨μμ μνΈνλ λ°μ΄ν°μ λ³΅νΈνλ λ°μ΄ν°λ₯Ό μ μ₯ν  μ¬λΌμ΄μ€λ₯Ό λ£μΌλ©΄ λ³΅νΈνκ° λ©λλ€.


κΈ΄ λ°μ΄ν°λ₯Ό μμ νκ² μνΈννκΈ° μν΄ λμΉ­ν€ μκ³ λ¦¬μ¦μ λ€μν μ΄μ© λ°©μμ μ κ³΅ν©λλ€. κ·Έμ€ CBC(Cipher Block Chaining) λ°©μμ μ¬μ©νμ¬ μνΈνλ₯Ό ν΄λ³΄κ² μ΅λλ€.

λ€μμ crypto/cipher ν¨ν€μ§μμ μ κ³΅νλ μνΈν μ΄μ© λͺ¨λ ν¨μμλλ€.

- func NewCBCEncrypter(b Block, iv []byte) BlockMode: μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ‘ μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
- func (x *cbcEncrypter) CryptBlocks(dst, src []byte): μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘ μνΈν
- func NewCBCDecrypter(b Block, iv []byte) BlockMode: μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ‘ λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
- func (x *cbcDecrypter) CryptBlocks(dst, src []byte): λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘ λ³΅νΈν
λ€μμ io ν¨ν€μ§μμ μ κ³΅νλ μ½κΈ° ν¨μμλλ€.

- func ReadFull(r Reader, buf []byte) (n int, err error): io.Readerμμ bufμ κΈΈμ΄λ§νΌ λ°μ΄ν°λ₯Ό μ½μ

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
	if mod := len(plaintext) % aes.BlockSize; mod != 0 { // λΈλ‘ ν¬κΈ°μ λ°°μκ° λμ΄μΌν¨
		padding := make([]byte, aes.BlockSize-mod)   // λΈλ‘ ν¬κΈ°μμ λͺ¨μλΌλ λΆλΆμ
		plaintext = append(plaintext, padding...)    // μ±μμ€
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // μ΄κΈ°ν λ²‘ν° κ³΅κ°(aes.BlockSize)λ§νΌ λ μμ±
	iv := ciphertext[:aes.BlockSize] // λΆλΆ μ¬λΌμ΄μ€λ‘ μ΄κΈ°ν λ²‘ν° κ³΅κ°μ κ°μ Έμ΄
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // λλ€ κ°μ μ΄κΈ°ν λ²‘ν°μ λ£μ΄μ€
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv) // μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ₯Ό λ£μ΄μ μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘
                                                                // μνΈν

	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 { // λΈλ‘ ν¬κΈ°μ λ°°μκ° μλλ©΄ λ¦¬ν΄
		fmt.Println("μνΈνλ λ°μ΄ν°μ κΈΈμ΄λ λΈλ‘ ν¬κΈ°μ λ°°μκ° λμ΄μΌν©λλ€.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]        // λΆλΆ μ¬λΌμ΄μ€λ‘ μ΄κΈ°ν λ²‘ν° κ³΅κ°μ κ°μ Έμ΄
	ciphertext = ciphertext[aes.BlockSize:] // λΆλΆ μ¬λΌμ΄μ€λ‘ μνΈνλ λ°μ΄ν°λ₯Ό κ°μ Έμ΄

	plaintext := make([]byte, len(ciphertext)) // νλ¬Έ λ°μ΄ν°λ₯Ό μ μ₯ν  κ³΅κ° μμ±
	mode := cipher.NewCBCDecrypter(b, iv)      // μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ₯Ό λ£μ΄μ
                                                   // λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
	mode.CryptBlocks(plaintext, ciphertext)    // λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘ λ³΅νΈν

	return plaintext
}

func main() {
	key := "Hello Key 123456" // 16λ°μ΄νΈ

	s := `λν΄ λ¬Όκ³Ό λ°±λμ°μ΄ λ§λ₯΄κ³  λ³λλ‘
νλλμ΄ λ³΄μ°νμ¬ μ°λ¦¬λλΌ λ§μΈ.
λ¬΄κΆν μΌμ²λ¦¬ νλ €κ°μ°
λν μ¬λ, λνμΌλ‘ κΈΈμ΄ λ³΄μ νμΈ.`

	block, err := aes.NewCipher([]byte(key)) // AES λμΉ­ν€ μνΈν λΈλ‘ μμ±
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s)) // νλ¬Έμ AES μκ³ λ¦¬μ¦μΌλ‘ μνΈν
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext) // AES μκ³ λ¦¬μ¦ μνΈλ¬Έμ νλ¬ΈμΌλ‘ λ³΅νΈν
	fmt.Println(string(plaintext))

}
```
```
6b39496c974990a9890f49ec962f92cd00f3009f724e3e6a7bcdd7b37867... (μλ΅)
λν΄ λ¬Όκ³Ό λ°±λμ°μ΄ λ§λ₯΄κ³  λ³λλ‘... (μλ΅)
```
λ¨Όμ  λΈλ‘ μνΈνλ μνΈνν  λ°μ΄ν°μ κΈΈμ΄κ° λΈλ‘ ν¬κΈ°(aes.BlockSize)μ λ°°μλΌμΌ ν©λλ€. λ°λΌμ λ€μκ³Ό κ°μ΄ λΈλ‘ ν¬κΈ°μ λ°°μκ° λ  μ μλλ‘ λͺ¨μλΌλ λΆλΆμ μ±μμ€λλ€.

```
if mod := len(plaintext) % aes.BlockSize; mod != 0 { // λΈλ‘ ν¬κΈ°μ λ°°μκ° λμ΄μΌν¨
	padding := make([]byte, aes.BlockSize-mod)   // λΈλ‘ ν¬κΈ°μμ λͺ¨μλΌλ λΆλΆμ
	plaintext = append(plaintext, padding...)    // μ±μμ€
}
```
CBC μ΄μ© λ°©μμ μ΄κΈ°ν λ²‘ν°(Initialization Vector, IV)λ₯Ό μ²« λΈλ‘μ λ°°μΉμν€κ³ , κ° λΈλ‘μ μ΄μ  λΈλ‘μ μνΈν κ²°κ³Όμ XORλ©λλ€. λ°λΌμ μ΄κΈ°ν λ²‘ν°λ μνΈνν  λλ§λ€ λ§€λ² λ€λ₯Έ κ°μΌλ‘ μμ±ν΄μΌ ν©λλ€.
```
ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // μ΄κΈ°ν λ²‘ν° κ³΅κ°(aes.BlockSize)λ§νΌ λ μμ±
iv := ciphertext[:aes.BlockSize] // λΆλΆ μ¬λΌμ΄μ€λ‘ μ΄κΈ°ν λ²‘ν° κ³΅κ°μ κ°μ Έμ΄
if _, err := io.ReadFull(rand.Reader, iv); err != nil { // λλ€ κ°μ μ΄κΈ°ν λ²‘ν°μ λ£μ΄μ€
	fmt.Println(err)
	return nil
}
```
μνΈν λ°μ΄ν°λ₯Ό μ μ₯ν  μ¬λΌμ΄μ€λ μ΄κΈ°ν λ²‘ν°κ° λ€μ΄κ° κ³΅κ°(aes.BlockSize)λ§νΌ λ μμ±ν΄μ€λλ€. κ·Έλ¦¬κ³  rand.Readerλ₯Ό io.ReadFull ν¨μλ‘ μ½μΌλ©΄ λλ€ κ°μ μ»μ μ μμ΅λλ€. μ΄ λλ€ κ°μ μ΄κΈ°ν λ²‘ν° ivμ λ£μ΄μ€λλ€.

cipher.NewCBCEncrypter ν¨μμ μνΈν λΈλ‘ bμ μ΄κΈ°ν λ²‘ν° ivλ₯Ό λ£μ΄μ€ λ€ CryptBlocks ν¨μλ‘ μνΈνλ₯Ό νλ©΄ λ©λλ€.
```
mode := cipher.NewCBCEncrypter(b, iv) // μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ₯Ό λ£μ΄μ μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // μνΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘ μνΈν

```
μ΄λ²μλ λ³΅νΈνλ₯Ό μμλ³΄κ² μ΅λλ€. λ³΅νΈνλ₯Ό ν  λ μνΈνλ λ°μ΄ν°μ κΈΈμ΄λ λΈλ‘ ν¬κΈ°μ λ°°μκ° λμ΄μΌ ν©λλ€. λ°λΌμ λ³΅νΈννκΈ° μ μ ν­μ κΈΈμ΄λ₯Ό κ²μ¬ν©λλ€.
```
if len(ciphertext)%aes.BlockSize != 0 { // λΈλ‘ ν¬κΈ°μ λ°°μκ° μλλ©΄ λ¦¬ν΄
	fmt.Println("μνΈνλ λ°μ΄ν°μ κΈΈμ΄λ λΈλ‘ ν¬κΈ°μ λ°°μκ° λμ΄μΌν©λλ€.")
	return nil
}
```
μνΈνλ λ°μ΄ν°(ciphertext)μ μ²« λΈλ‘μμ ivλ₯Ό κΊΌλλλ€. κ·Έλ¦¬κ³  cipher.NewCBCDecrypter ν¨μμ μνΈν λΈλ‘ bμ μ΄κΈ°ν λ²‘ν° ivλ₯Ό λ£μ΄μ€ λ€ CryptBlocks ν¨μλ‘ λ³΅νΈνλ₯Ό νλ©΄ λ©λλ€.
```
iv := ciphertext[:aes.BlockSize]        // λΆλΆ μ¬λΌμ΄μ€λ‘ μ΄κΈ°ν λ²‘ν° κ³΅κ°μ κ°μ Έμ΄
ciphertext = ciphertext[aes.BlockSize:] // λΆλΆ μ¬λΌμ΄μ€λ‘ μνΈνλ λ°μ΄ν°λ₯Ό κ°μ Έμ΄

plaintext := make([]byte, len(ciphertext)) // νλ¬Έ λ°μ΄ν°λ₯Ό μ μ₯ν  κ³΅κ° μμ±
mode := cipher.NewCBCDecrypter(b, iv)      // μνΈν λΈλ‘κ³Ό μ΄κΈ°ν λ²‘ν°λ₯Ό λ£μ΄μ
                                           // λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€ μμ±
mode.CryptBlocks(plaintext, ciphertext)    // λ³΅νΈν λΈλ‘ λͺ¨λ μΈμ€ν΄μ€λ‘ λ³΅νΈν
```
## π―RSA κ³΅κ°ν€ μκ³ λ¦¬μ¦ μ¬μ©νκΈ°
λμΉ­ν€ μκ³ λ¦¬μ¦μ μνΈ ν€κ° μ μΆλλ©΄ μνΈνλ λ°μ΄ν°λ₯Ό λͺ¨λ ν μ μμ΅λλ€. νΉν λ€νΈμν¬λ‘ μνΈ ν€λ₯Ό μ£Όκ³  λ°μΌλ©΄ λΈμΆλ  μνμ΄ μ»€μ§λλ€. λ°λΌμ λ€νΈμν¬λ‘ λ°μ΄ν°λ₯Ό μ£Όκ³  λ°μλλ κ³΅κ°ν€ μνΈν μκ³ λ¦¬μ¦μ λ§μ΄ μ¬μ©ν©λλ€. λ¨ κ³΅κ°ν€ μκ³ λ¦¬μ¦μ λμΉ­ν€ μκ³ λ¦¬μ¦μ λΉν΄ μλκ° λλ¦¬λ―λ‘ λμΉ­ν€ μκ³ λ¦¬μ¦μ μνΈ ν€λ§ κ³΅κ°ν€ μκ³ λ¦¬μ¦μΌλ‘ μνΈννμ¬ λ€νΈμν¬λ‘ μ μ‘νκΈ°λ ν©λλ€.

λ€μμ crypto/rsa ν¨ν€μ§μμ μ κ³΅νλ κ³΅κ°ν€ μκ³ λ¦¬μ¦ ν¨μμλλ€.

- func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error): κ°μΈ ν€μ κ³΅κ°ν€ μμ±
- func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error): νλ¬Έμ κ³΅κ° ν€λ‘ μνΈν
- func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error): μνΈνλ λ°μ΄ν°λ₯Ό κ°μΈ ν€λ‘ λ³΅νΈν
RSA κ³΅κ°ν€ μκ³ λ¦¬μ¦μ μ¬μ©νμ¬ μνΈνμ λ³΅νΈνλ₯Ό ν΄λ³΄κ² μ΅λλ€.
```
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // κ°μΈ ν€μ κ³΅κ°ν€ μμ±
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // κ°μΈ ν€ λ³μ μμ κ³΅κ° ν€κ° λ€μ΄μμ

	s := `λν΄ λ¬Όκ³Ό λ°±λμ°μ΄ λ§λ₯΄κ³  λ³λλ‘
νλλμ΄ λ³΄μ°νμ¬ μ°λ¦¬λλΌ λ§μΈ.
λ¬΄κΆν μΌμ²λ¦¬ νλ €κ°μ°
λν μ¬λ, λνμΌλ‘ κΈΈμ΄ λ³΄μ νμΈ.`

	ciphertext, err := rsa.EncryptPKCS1v15( // νλ¬Έμ κ³΅κ° ν€λ‘ μνΈν
		rand.Reader,
		publicKey, // κ³΅κ°ν€
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15( // μνΈνλ λ°μ΄ν°λ₯Ό κ°μΈ ν€λ‘ λ³΅νΈν
		rand.Reader,
		privateKey, // κ°μΈν€
		ciphertext,
	)

	fmt.Println(string(plaintext))
}
```
rsa.GenerateKey ν¨μμ rand.Readerμ ν€μ κΈΈμ΄λ₯Ό λ£μΌλ©΄ κ°μΈ ν€μ κ³΅κ°ν€κ° μμ±λ©λλ€. ν€μ κΈΈμ΄λ λΉνΈ λ¨μμ΄λ©° 2μ μ κ³±μΌλ‘ μλ ₯ν©λλ€.
rsa.EncryptPKCS1v15 ν¨μμ rand.Reader, publicKey, μνΈνν  λ°μ΄ν°λ₯Ό λ£μΌλ©΄ μνΈνκ° λ©λλ€. κ·Έλ¦¬κ³  rsa.DecryptPKCS1v15 ν¨μμ rand.Reader, privateKey, μνΈνλ λ°μ΄ν°λ₯Ό λ£μΌλ©΄ λ³΅νΈνκ° λ©λλ€.

μ΄μ²λΌ κ³΅κ°ν€ μκ³ λ¦¬μ¦μ μνΈνν  λλ κ³΅κ° ν€λ₯Ό μ¬μ©νκ³ , λ³΅νΈνν  λλ κ°μΈ ν€λ₯Ό μ¬μ©ν©λλ€. κ³΅κ° ν€λ₯Ό μΈλΆμ κ³΅κ°νμ¬ λ€λ₯Έ μ¬λνν μ λ¬ν λ€ μνΈνλ₯Ό νκ³ , μνΈνλ λ°μ΄ν°λ₯Ό λ°λ μ¬λμ μκΈ°κ° κ°μ§ κ°μΈν€λ‘ λ³΅νΈννλ λ°©μμλλ€. κ°μΈ ν€λ μΈλΆμ λΈμΆν  μΌμ΄ μμΌλ©° κ³΅κ° ν€λ‘λ κ°μΈ ν€λ₯Ό μΆμΆνκΈ° μ΄λ ΅κΈ° λλ¬Έμ μμ νκ² μνΈνλ λ°μ΄ν°λ₯Ό μ£Όκ³  λ°μ μ μμ΅λλ€.

μ΄λ²μλ RSA μκ³ λ¦¬μ¦μ μ¬μ©νμ¬ λ©μμ§λ₯Ό μλͺ(Signing)νκ³  μΈμ¦(Verification)νλ λ°©λ²μ μμλ³΄κ² μ΅λλ€.

λ€μμ crypto/rsa ν¨ν€μ§μμ μ κ³΅νλ μλͺ, μΈμ¦ ν¨μμλλ€.

- func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error): κ°μΈ ν€λ‘ μλͺ
- func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (err error): κ³΅κ°ν€λ‘ μλͺ κ²μ¦
λ€μμ crypto/md5 ν¨ν€μ§μμ μ κ³΅νλ ν΄μ ν¨μμλλ€.

- func New() hash.Hash: MD5 ν΄μ μΈμ€ν΄μ€ μμ±
- func (d *digest) Write(p []byte) (nn int, err error): ν΄μ μΈμ€ν΄μ€μ λ°μ΄ν° μΆκ°
- func (d0 *digest) Sum(in []byte) []byte: ν΄μ μΈμ€ν΄μ€μ μ μ₯λ λ°μ΄ν°μ MD5 ν΄μ κ° μΆμΆ

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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // κ°μΈ ν€μ κ³΅κ° ν€ μμ±
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // κ°μΈ ν€ λ³μ μμ κ³΅κ° ν€κ° λ€μ΄μμ

	message := "μλνμΈμ. Go μΈμ΄"
	hash := md5.New()           // ν΄μ μΈμ€ν΄μ€ μμ±
	hash.Write([]byte(message)) // ν΄μ μΈμ€ν΄μ€μ λ¬Έμμ΄ μΆκ°
	digest := hash.Sum(nil)     // λ¬Έμμ΄μ MD5 ν΄μ κ° μΆμΆ

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15( // κ°μΈ ν€λ‘ μλͺ
		rand.Reader,
		privateKey, // κ°μΈ ν€
		h1,
		digest,     // MD5 ν΄μ κ°
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15( // κ³΅κ° ν€λ‘ μλͺ κ²μ¦
		publicKey, // κ³΅κ° ν€
		h2,
		digest,    // MD5 ν΄μ κ°
		signature, // μλͺ κ°
	)

	if err != nil {
		fmt.Println("κ²μ¦ μ€ν¨")
	} else {
		fmt.Println("κ²μ¦ μ±κ³΅")
	}
}
```
λ¨Όμ  μλͺν  λ°μ΄ν°μ ν΄μ κ°μ κ΅¬ν©λλ€. μ¬κΈ°μλ md5λ‘ ν΄μ κ°μ κ΅¬ν©λλ€.
```
message := "μλνμΈμ. Go μΈμ΄"
hash := md5.New()           // ν΄μ μΈμ€ν΄μ€ μμ±
hash.Write([]byte(message)) // ν΄μ μΈμ€ν΄μ€μ λ¬Έμμ΄ μΆκ°
digest := hash.Sum(nil)     // λ¬Έμμ΄μ MD5 ν΄μ κ° μΆμΆ
```
μ΄μ  βμλνμΈμ. Go μΈμ΄βμ ν΄μ κ°μ κ°μΈ ν€λ‘ μλͺν©λλ€. μ¬κΈ°μ μλͺμ΄λ λ΄κ° βμλνμΈμ. Go μΈμ΄βλΌλ λ©μμ§λ₯Ό λ³΄λλ€κ³  λ³΄μ¦νλ κ²μλλ€.
```
var h1 crypto.Hash
signature, err := rsa.SignPKCS1v15( // κ°μΈ ν€λ‘ μλͺ
	rand.Reader,
	privateKey, // κ°μΈ ν€
	h1,
	digest,     // MD5 ν΄μ κ°
)
```
rsa.SignPKCS1v15 ν¨μμ rand.Reader, crypto.Hash, κ°μΈ ν€, ν΄μ κ°μ λ£μΌλ©΄ μλͺμ΄ λ©λλ€. μ΄λ κ² ν΄μ λ©μμ§(message), λ©μμ§μ λν ν΄μ κ°(digest), μλͺ(signature)μ΄ λͺ¨λ μ€λΉλμμ΅λλ€.

λ©μμ§λ₯Ό λ°λ μ¬λμ μ λ§ μ  μ¬λ(κ°μΈ ν€λ₯Ό κ°μ§ μ¬λ)μ΄ βμλνμΈμ. Go μΈμ΄βλΌλ λ§μ νλμ§ κ²μ¦(μΈμ¦)μ ν΄μΌλ©λλ€.
```
var h2 crypto.Hash
err = rsa.VerifyPKCS1v15( // κ³΅κ° ν€λ‘ μλͺ κ²μ¦
	publicKey, // κ³΅κ° ν€
	h2,
	digest,    // MD5 ν΄μ κ°
	signature, // μλͺ κ°
)
```
rsa.VerifyPKCS1v15 ν¨μμ κ³΅κ° ν€, crypto.Hash, ν΄μ κ°, μλͺμ λ£μΌλ©΄ κ²μ¦μ΄ λ©λλ€. err κ°μ΄ nilμ΄λ©΄ μ μμ μΈ λ©μμ§μλλ€.

λ©μμ§, λ©μμ§μ λν ν΄μ κ°, μλͺ, κ³΅κ° ν€λ λͺ¨λ κ³΅κ°λ μ λ³΄μλλ€. μ΄ κ°λ€μ μ΄μ©ν΄μ λ©μμ§κ° λ³μ‘°λμ§ μκ³ , μ¬λ°λ₯Έ μ¬λννμ μλμ§ κ²μ¦νλ κ²μλλ€.

- λ©μμ§μ λν ν΄μ κ°μΌλ‘ λ©μμ§κ° λ³μ‘°λμλμ§ νμΈν  μ μμ΅λλ€.
- μλͺ κ°μΌλ‘ λ©μμ§κ° μ¬λ°λ₯Έ μ¬λννμ μλμ§ νμΈν  μ μμ΅λλ€. μ¦ μλͺμ ν  μ μλ μ¬λμ κ°μΈ ν€λ₯Ό κ°μ§ μ¬λλ°μ μκΈ° λλ¬Έμλλ€.
κ°μΈ ν€λ₯Ό κ°μ§κ³  μλͺμ νλ μ£Όμ²΄λ μ¬λμ΄ λ  μλ μκ³  μ»΄ν¨ν°κ° λ  μλ μμ΅λλ€. μ¬κΈ°μλ νΈμμ μ¬λμ΄λΌ μΉ­νμ΅λλ€. λ³΄ν΅ κΈμ΅κ±°λμμ λ§μ΄ μ°λ κ³΅μΈμΈμ¦μκ° μλͺκ³Ό μΈμ¦ λ°©μμ RSA μκ³ λ¦¬μ¦μ μ¬μ©ν©λλ€. λ€λ₯Έ μ¬λ‘λ‘λ μΈν°λ·μμ νμΌμ λ°μ λ ν΄μ κ°κ³Ό μλͺ νμΌμ΄ ν¨κ» μλ κ³³μ΄ μμ΅λλ€. μ΄λλ ν΄μ κ°μΌλ‘ νμΌ λ³μ‘° μ¬λΆλ₯Ό νμΈνκ³ , μλͺ νμΌλ‘ κ³΅μ κ°λ°μ(νμ¬, λ¨μ²΄)κ° λ°°ν¬ν κ²μΈμ§ κ²μ¦ν  μ μμ΅λλ€.
