# go_study


<img width="509" alt="스크린샷 2021-10-26 오전 12 36 59" src="https://user-images.githubusercontent.com/88940298/138727331-671b1969-66a1-4be2-acb9-177d1152a4b3.png">

3000 서버  연결

<img width="403" alt="스크린샷 2021-10-26 오전 12 39 35" src="https://user-images.githubusercontent.com/88940298/138727383-95f932d6-ffde-467d-bd05-5e3806c4906f.png">

go 에서는 ;(세미클론이 필요없다) 선택적으로 사용할 수는 있음
go에서는 패키지 개념 패키지 내부에서는 변수,상수,함수,구조체,메서드가 선언되어 있다.
모든 go파일은 패키지 내부에속해야한다 main패키지가 아닌 다른코드를 실행시킬수 없음

import"fmt" 다른 패키지를 포함하기 위해 사용 go 의패키지는 src에 모여있으며 pkg폴더에는 그러한 패키지가 컴파일된 소스코드가 모여있다
변수 선언 var 사용 함수안에서는 단축변수선언 사용가능 :=

변수를 여러개를 선언하고 초기화할떄에는 반드시 선언하는 변수와 초기화하는 값의 개수가 같아야함 타입은 달라도 됨

go에서는 패키지,타입과 똑같은 이름의 변수를 사용하여도 에러가 나지 않는다

상수 const

go에서는 타입을 정하지 않으면 go에서 알아서 타입을 지정해준다
 
go가 다른언어랑 다른 점은 클래스가 없다는 것이다.메서드라고 표현하는 것도 표현 방식이 다른 뿐 일반 함수와 유사하게 표시된다
go는 리턴 값을 여러개의 값으로 돌릴 수 있다.
defer키워드는 go에서 지원하는 특이한 키워드중 하나 이다. defer 을 통해  먼저 예약이 가능 하며 , defer은 나중에 설정된 defer부터 실행된다.

go&go-routine go 키워드는 멀티스레드를 사용할때 유용하며 동시성을 가진다.
 
go get go get을 통하여 다른 패키지를 다운받을 수 있다.얻어온 패키지는 GOPATH로 다운로드 된다.GOPATH는 작업공간이라고도 하며 사용자마다 별도로 관리된다

go mod 사용자 정의 모듈을 만들떄 사용한다. example.com폴더로 진입하여 go mod init example.com명령어를 통해 모듈을 정의한다


<img width="355" alt="스크린샷 2021-10-27 오후 2 11 44" src="https://user-images.githubusercontent.com/88940298/139003610-95bf570c-fa06-4959-815c-5e2bea14f2c7.png">


## if문 
조건문을 사용할떄 변수를 함께 선언할 수있다.

swtich 조건문과 마찬가지로 사용법은 switch,case와 비슷하지만 break문이 필요없다 .새로운 키워드인 fallthrough를 선택했다.fallthrough를사용해야 다음으로 넘어 갈 수있다.또한 다른 언어와 다르게 조건을 생략하고 case를 사용할 수 잇다.

<img width="300" alt="스크린샷 2021-10-26 오전 12 40 06" src="https://user-images.githubusercontent.com/88940298/138727421-15f9ef34-9bfb-4202-b1fd-4ef9ad1e5b04.png">
## for문 

반복문을 사용할 떄는 for키워드를 이용한다. go 언어에는 while문이 없기 떄문에 for문에서 모든 반복을 해결한다. 무한루프의 경우 아무것도 안주면 되어 타언어에 비해 심플하다 배열이나 슬라이드,문을 순회할떄 range키워드를 사용하여 순회 할 수 있다. 키 인덱스가 들어가며 두번쨰에는 값이 들어간다.
반복문을 제어하는 방법으로는 break,continue가 존재한다.break를 사용하면 반복문에서 즉시 탈출하며,continue를사용하면 그 이후 표현식은 무시하고 다시 조건문으로 돌아간다

<img width="171" alt="스크린샷 2021-10-27 오후 2 20 51" src="https://user-images.githubusercontent.com/88940298/139004540-59a4b021-cb73-4120-aa5c-1f8cf0c42b13.png">

## Pointers
& :메모리의 주소값
* : 주소에 담긴 값을 살펴볼수 잇음

<img width="340" alt="스크린샷 2021-10-27 오후 2 26 28" src="https://user-images.githubusercontent.com/88940298/139005018-256a8d69-4597-40b3-aa49-0cd0cdc570c3.png">

##Array ana Slices

