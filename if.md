# ❓if문

```
if 조건식 {
	// 여기에 조건식이 true일 때 실행할 코드를 작성합니다.
}
```
이제 간단하게 숫자가 5 이상인지 판별해보겠습니다.
```
i := 10

if i >= 5 {
	fmt.Println("5 이상")
}
```
if 조건문을 작성할 때 ( ) (괄호)를 사용하지 않으며 조건문을 시작하는 줄에서 { (여는 중괄호)가 시작됩니다. * 조건식: 조건식의 결과가 참(true)이면 중괄호 블록을 실행합니다. 따라서 조건식의 결과는 반드시 불형이되어야 합니다. C 언어와 달리 조건식의 결과가 정수형, 실수형, 문자열 값 등이 되면 컴파일 에러가 발생합니다.

다음과 같이 여는 중괄호를 다음 줄에 작성하면 컴파일 에러가 발생합니다.

```
i := 10

if i >= 5 // 컴파일 에러
{
	fmt.Println("5 이상")
}
```
if 조건식이 거짓일 때는 else 키워드로 코드를 실행할 수 있습니다. else 키워드는 } else {처럼 닫는 중괄호와 여는 중괄호가 같은 줄에 있어야 합니다.

```
i := 4

if i >= 5 {
	fmt.Println("5 이상")
} else {
	fmt.Println("5 미만")
}
```

if의 조건식이 거짓일 때 다른 조건식을 검사하도록 else if 키워드를 사용할 수도 있습니다. else 키워드와 마찬가지로 닫는 중괄호와 여는 중괄호가 같은 줄에 있어야 합니다

```
i := 6

if i >= 10 {
	fmt.Println("10 이상")
} else if i >= 5 && i < 10 {
	fmt.Println("5 이상, 10 미만")
} else if i >= 0 && i < 5 {
	fmt.Println("0 이상, 5 미만")
}
```

```
var b []byte
var err error

b, err = ioutil.ReadFile("./hello.txt")

if err == nil {
	fmt.Printf("%s", b)
}

결과: Hello, world!
```
이 코드를 다음과 같이 if 조건문 안에서 함수를 실행한 뒤 조건을 판단하는 방식으로 바꿀 수 있습니다.

```
if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
	fmt.Printf("%s", b)
}
```


if 조건문 안에서 함수를 실행하고 ; (세미콜론)으로 구분한 뒤 조건식을 작성합니다. ioutil.ReadFile 함수로 파일을 열면 b에는 파일의 내용이 들어가고, err에는 에러 값이 들어갑니다. 이때 에러가 없으면 fmt.Printf 함수를 실행하여 파일 내용을 출력합니다. 이렇게 if 안에서 함수를 실행하고 에러 값을 확인하면 코드양을 줄일 수 있습니다.

if 조건문 안에서 변수를 생성했을 때 변수는 if 조건식이 참일 때 중괄호 블록뿐만 아니라 else, else if 중괄호 블록에서도 사용할 수 있습니다. 하지만 if 조건문 바깥에서는 변수를 사용할 수 없습니다.

```
if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
	fmt.Printf("%s", b) // 변수 b를 사용할 수 있음
} else {
	fmt.Println(err) // 변수 err을 사용할 수 있음
}

fmt.Println(b)   // 변수 b를 사용할 수 없음. 컴파일 에러
fmt.Println(err) // 변수 err을 사용할 수 없음. 컴파일 에러
```


