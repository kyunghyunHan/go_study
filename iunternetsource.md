# 👩🏻‍🎓인터넷 소스 저장소의 패키지 사용하기

import 키워드는 로컬에 있는 패키지뿐만 아니라 인터넷의 소스 저장소에 올라와 있는 패키지도 사용할 수 있다

Go 언어에서 지원하는 소스 저장소와 버전 관리 시스템은 다음과 같다

- GitHub: Git
- BitBucket: Git, Mercurial
- Launchpad: Bazaar
- IBM DevOps Services: Git
먼저 소스 저장소의 패키지를 사용하기 전에 소스 저장소에서 사용하는 버전 관리 시스템을 설치해야 한다. Windows와 Mac OS X는 다음 주소에서 설치 파일을 받아서 설치한다

- Git: http://git-scm.com/downloads
- Mercurial: http://mercurial.selenic.com/downloads
- Subversion: https://subversion.apache.org/packages.html
- Bazaar: http://wiki.bazaar.canonical.com/Download
리눅스에서는 다음 명령으로 각 버전 관리 시스템을 설치한다
- 우분투
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
- 여기서 저장소 주소의 가장 마지막 디렉터리가 패키지 이름이 됩니다. 즉, stringutil.Reverse처럼 사용하면 된다.

다음과 같이 전역 패키지로 사용할 수 있다.


```
import (
	"fmt"
	. "github.com/golang/example/stringutil" // 전역 패키지
)
```
- 만약 다른 패키지와 이름이 중복된다면 별칭을 지정해도 된다.
```
import (
	"fmt"
	strutil "github.com/golang/example/stringutil" // strutil로 별칭 지정
)
```
- 이제 콘솔(터미널)에서 GOPATH/src/hello 디렉터리로 이동한 뒤 다음 명령을 실행한다(Windows에서는 명령 프롬프트 또는 PowerShell을 실행하면된다.). 저는 GOPATH를 /home/pyrasis/hello_project로 설정하였된다.

```
~$ cd $GOPATH/src/hello
~/hello_project/src/hello$ go get
```
- 이렇게 go get 명령을 실행하면 hello.go 파일에서 import로 설정한 GitHub 주소에서 패키지를 자동으로 받아온다. GOPATH/src 디렉터리를 보면 github.com/golang/example 디렉터리가 만들어져 있고, 패키지의 소스가 들어있다.

- 다음과 같이 go get <저장소 주소> 형식으로 명령을 실행하면 바로 패키지를 받아올 수도 있다. 이때는 패키지를 받아오기만 한다. 따라서 이 패키지를 사용하려면 import로 패키지 주소(패키지 이름)를 설정해야 한다.
```
$ go get github.com/golang/example/stringutil
```
