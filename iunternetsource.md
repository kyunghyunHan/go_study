# π©π»βπμΈν°λ· μμ€ μ μ₯μμ ν¨ν€μ§ μ¬μ©νκΈ°

import ν€μλλ λ‘μ»¬μ μλ ν¨ν€μ§λΏλ§ μλλΌ μΈν°λ·μ μμ€ μ μ₯μμ μ¬λΌμ μλ ν¨ν€μ§λ μ¬μ©ν  μ μλ€

Go μΈμ΄μμ μ§μνλ μμ€ μ μ₯μμ λ²μ  κ΄λ¦¬ μμ€νμ λ€μκ³Ό κ°λ€

- GitHub: Git
- BitBucket: Git, Mercurial
- Launchpad: Bazaar
- IBM DevOps Services: Git
λ¨Όμ  μμ€ μ μ₯μμ ν¨ν€μ§λ₯Ό μ¬μ©νκΈ° μ μ μμ€ μ μ₯μμμ μ¬μ©νλ λ²μ  κ΄λ¦¬ μμ€νμ μ€μΉν΄μΌ νλ€. Windowsμ Mac OS Xλ λ€μ μ£Όμμμ μ€μΉ νμΌμ λ°μμ μ€μΉνλ€

- Git: http://git-scm.com/downloads
- Mercurial: http://mercurial.selenic.com/downloads
- Subversion: https://subversion.apache.org/packages.html
- Bazaar: http://wiki.bazaar.canonical.com/Download
λ¦¬λμ€μμλ λ€μ λͺλ ΉμΌλ‘ κ° λ²μ  κ΄λ¦¬ μμ€νμ μ€μΉνλ€
- μ°λΆν¬
```
 sudo apt-get install git
```
- CentOS
```
sudo yum install git
```
```
sudo apt-get install mercurial
```
```
$ sudo yum install mercurial
```
```
 sudo apt-get install subversion
```
```
sudo yum install subversion
```
```
 sudo apt-get install bzr
```
```
$ sudo yum install bzr
```
```
package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("olleh"))
}
```
- μ¬κΈ°μ μ μ₯μ μ£Όμμ κ°μ₯ λ§μ§λ§ λλ ν°λ¦¬κ° ν¨ν€μ§ μ΄λ¦μ΄ λ©λλ€. μ¦, stringutil.Reverseμ²λΌ μ¬μ©νλ©΄ λλ€.

λ€μκ³Ό κ°μ΄ μ μ­ ν¨ν€μ§λ‘ μ¬μ©ν  μ μλ€.


```
import (
	"fmt"
	. "github.com/golang/example/stringutil" // μ μ­ ν¨ν€μ§
)
```
- λ§μ½ λ€λ₯Έ ν¨ν€μ§μ μ΄λ¦μ΄ μ€λ³΅λλ€λ©΄ λ³μΉ­μ μ§μ ν΄λ λλ€.
```
import (
	"fmt"
	strutil "github.com/golang/example/stringutil" // strutilλ‘ λ³μΉ­ μ§μ 
)
```
- μ΄μ  μ½μ(ν°λ―Έλ)μμ GOPATH/src/hello λλ ν°λ¦¬λ‘ μ΄λν λ€ λ€μ λͺλ Ήμ μ€ννλ€(Windowsμμλ λͺλ Ή νλ‘¬ννΈ λλ PowerShellμ μ€ννλ©΄λλ€.). μ λ GOPATHλ₯Ό /home/pyrasis/hello_projectλ‘ μ€μ νμλλ€.

```
~$ cd $GOPATH/src/hello
~/hello_project/src/hello$ go get
```
- μ΄λ κ² go get λͺλ Ήμ μ€ννλ©΄ hello.go νμΌμμ importλ‘ μ€μ ν GitHub μ£Όμμμ ν¨ν€μ§λ₯Ό μλμΌλ‘ λ°μμ¨λ€. GOPATH/src λλ ν°λ¦¬λ₯Ό λ³΄λ©΄ github.com/golang/example λλ ν°λ¦¬κ° λ§λ€μ΄μ Έ μκ³ , ν¨ν€μ§μ μμ€κ° λ€μ΄μλ€.

- λ€μκ³Ό κ°μ΄ go get <μ μ₯μ μ£Όμ> νμμΌλ‘ λͺλ Ήμ μ€ννλ©΄ λ°λ‘ ν¨ν€μ§λ₯Ό λ°μμ¬ μλ μλ€. μ΄λλ ν¨ν€μ§λ₯Ό λ°μμ€κΈ°λ§ νλ€. λ°λΌμ μ΄ ν¨ν€μ§λ₯Ό μ¬μ©νλ €λ©΄ importλ‘ ν¨ν€μ§ μ£Όμ(ν¨ν€μ§ μ΄λ¦)λ₯Ό μ€μ ν΄μΌ νλ€.
```
$ go get github.com/golang/example/stringutil
```
