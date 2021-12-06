# Map
Go 언어는 기본 자료형으로 맵을 지원합니다. 맵은 해시 테이블, 딕셔너리라고도 하며 키-값 형태로 자료를 저장합니다. 또한, 슬라이스와 마찬가지로 레퍼런스 타입입니다. 이제 var 키워드로 맵을 생성해보겠습니다.

```
var a map[string]int // 키는 string, 값은 int인 맵 선언
```

[ ] (대괄호)안에는 키의 자료형을 지정하고 그 뒤에 값의 자료형을 지정합니다.

맵은 make 함수를 사용하여 공간을 할당해야 값을 넣을 수 있습니다. 여기서 맵 선언과 동시에 make 함수를 사용하면 map 키워드와 자료형을 생략할 수 있습니다. 또한, 일반 변수와 마찬가지로 :=를 사용하여 var를 생략할 수 있습니다

var a map[string]int = make(map[string]int) // make 함수로 키는 string 값은 int인 맵에 공간 할당
var b = make(map[string]int)                // 맵을 선언할 때 map 키워드와 자료형 생략
c := make(map[string]int)                   // 맵을 선언할 때 var, map 키워드와 자료형 생략

맵을 선언할 때와 동일하게 map 키워드로 키, 값 자료형을 지정합니다. 이렇게 하면 문자열 키에 정수형 값을 저장합니다.

맵을 생성하면서 키와 값을 초기화하려면 { } (중괄호)를 사용합니다. 중괄호에서 키와 값은 한 줄로 나열해도 되고 여러 줄로 나열해도 됩니다(여러 줄로 나열할 때는 마지막 키와 값에도 콤마를 붙여줍니다).
```
a := map[string]int{"Hello": 10, "world": 20}

b := map[string]int{
	"Hello": 10,
	"world": 20, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
}

fmt.Println(a["Hello"]) // 10
fmt.Println(b["world"]) // 10

```
```
10
20
```

## 맵에 데이터 저장하고 조회하기

맵에 데이터를 넣으려면 [ ] (대괄호)안에 키를 지정하고 값을 대입합니다. 마찬가지로 값을 조회할 때도 대괄호 안에 키를 지정하면 됩니다. 다음은 태양계 행성들의 공전 주기를 맵에 저장한 것입니다.

```
solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

solarSystem["Mercury"] = 87.969 // 맵[키] = 값
solarSystem["Venus"] = 224.70069
solarSystem["Earth"] = 365.25641
solarSystem["Mars"] = 686.9600
solarSystem["Jupiter"] = 4333.2867
solarSystem["Saturn"] = 10756.1995
solarSystem["Uranus"] = 30707.4896
solarSystem["Neptune"] = 60223.3528

fmt.Println(solarSystem["Earth"]) // 365.25641
```
```
365.2564
```

맵에서 존재하지 않는 키를 조회했을 때는 빈 값이 출력됩니다. 여기서는 값의 자료형이 float32이므로 0이 출력됩니다. 자료형이 문자열이라면 빈 문자열(““)이 출력됩니다

```
fmt.Println(solarSystem["Pluto"]) // 0: 존재하지 않는 키를 조회(명왕성은 행성 지위 상실...)
```
```
value, ok := solarSystem["Pluto"] // 맵에 키가 있는지 검사할 때는 리턴값을 두 개 활용
fmt.Println(value, ok)            // 0 false: 맵에 키가 없으면 두 번째 리턴값으로 false가 리턴됨

if value, ok := solarSystem["Saturn"]; ok {
	fmt.Println(value) // 10756.1995
}
```
```
0 false
10756.199
```

## 맵 순회
맵의 모든 데이터를 출력해보겠습니다. 배열, 슬라이스와 마찬가지로 맵도 for 반복문에서 range 키워드를 사용합니다.

- for 키, 값 := range 맵 { }
```
for key, value := range solarSystem { // 반복문이 실행될 때마다 키와 값이 자동으로 변수에 들어감
	fmt.Println(key, value)
}
```
range 키워드를 사용하면 반복문이 실행될 때 마다 맵의 키와 값이 자동으로 변수에 들어갑니다.

range의 리턴값에서 키 변수를 사용하고 싶지 않다면 _ (밑줄 문자)를 사용합니다.
```
for _, value := range solarSystem { // 키 변수를 사용하고 싶지 않다면 _ 사용
	fmt.Println(value)
}
```

## 맵에서 데이터 삭제
맵에서 값을 삭제하려면 delete 함수를 사용합니다.
- delete(맵, 삭제할_키)

```
a := map[string]int{"Hello": 10, "world": 20}

delete(a, "world") // 맵 a에서 world 키 삭제

fmt.Println(a) // map[Hello:10]
```
## 맵안에 맵 생성
맵의 값 안에는 일반 자료형뿐만 아니라 맵 자체도 들어갈 수 있습니다.

- map[키_자료형]map[키_자료형]값_자료형
다음은 지구형 행성의 반지름, 질량, 공전주기를 맵으로 표현한 예제입니다.

```
terrestrialPlanet := map[string]map[string]float32{
	"Mercury": map[string]float32{
		"meanRadius":    2439.7,
		"mass":          3.3022E+23,
		"orbitalPeriod": 87.969,
	},
	"Venus": map[string]float32{
		"meanRadius":    6051.8,
		"mass":          4.8676E+24,
		"orbitalPeriod": 224.70069,
	},
	"Earth": map[string]float32{
		"meanRadius":    6371.0,
		"mass":          5.97219E+24,
		"orbitalPeriod": 365.25641,
	},
	"Mars": map[string]float32{
		"meanRadius":    3389.5,
		"mass":          6.4185E+23,
		"orbitalPeriod": 686.9600,
	},
}

fmt.Println(terrestrialPlanet["Mars"]["mass"]) // 6.4185E+23
```
```
6.4185e+23
```
