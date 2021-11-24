# go_study 
go 



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

제로갑:배열만 선언해놓고, 값을 할당하지 않은 경우 각 타입에 맞는 제로값이 설정된다
인덱스 제로베이스 : 0 부터 시작 length-1까지 해당
슬라이스:크기가 정해지지 않은 배열이며 길이가 동적으로 변할수 있다. go의 슬라이스는 배열의 뷰(view)일뿐 다른 의미는 가지지 않는다 . make()빌트인 함수 사용하여 생성
또한 슬라이스를 변경하면 원본 배열도 바뀌며 원본 배열을 바꾸면 슬라이스도 변경된다.
append()슬라이스에 값을 추가할떄 사용 

<img width="461" alt="스크린샷 2021-10-27 오후 2 35 41" src="https://user-images.githubusercontent.com/88940298/139005924-44e7384d-59ac-48a1-a11d-bdac7f2309c1.png">

맵은 자바스크립트의 객체 리터럴과 비슷하게 키와 값을 가질 수 있는 자료구도이다 맵을 사용하면 배열과 슬라이스를 사용했을떄 거처야할 일부 번거로운 연산과정을 거칠 필요가 없이 키와 값이 매핑되어 있기 떄문에 접근이나 값을 수정하는 것이 수월하다. 데이터가 순서대로 정렬되지 않는 특징이 있어 정려로딘 데이터를 원한다면 별도의 정렬 함수를 사용해야한다.
delete()함수를 사용하면 맵에서 특정 키에 해당하는 값을 제거 할수 있다
for~range를 사용하여 순회할수 있으며, 순서가 지정되있지 않기 떄문에 무엇이 먼저 나올수 있을지 모른다는것점이 있다.  
<img width="430" alt="스크린샷 2021-10-27 오후 2 53 48" src="https://user-images.githubusercontent.com/88940298/139007738-0376e5af-079c-44c8-ae20-228d1847a9f3.png">
<img width="467" alt="스크린샷 2021-10-27 오후 3 09 01" src="https://user-images.githubusercontent.com/88940298/139009338-e0c7544e-8e2a-4d2b-831b-f5cbd68e4c7d.png">  


## Struct

go 에는 클래스가 없다.객체라는 단어는 거의 사용하지 않으며 c언어처럼 구조체라는 존재가 있다.

#
구조체는 클래스처럼 객체를 찍어내기위한 판 이라 접근하면 좋다.struct라는 키워드를 사용하며 person이라는구조체를 선언하고 main에서
구조체를 타입으로 갖는 변수를 선언하고 값을 할당하는 모습을 보여주고 있다.
메서드는 일반적은 메서드파라메터외에 리시버파라메터라는것을 사용한다.이는 go의 특징이다.
func다음에 파라메타 같은 형태로 나타낸 것을 볼수 있는데 리시버파라메타이다.this키워드를 대신하며 this에비해 타입을 명시해줄수 있다.
또한 리시버파라메타에 대해 포인터 자료형으로 주었는데 포인터로 주지 않으면 값복사가 되어 실제로 원본이 변화하지 않게 된다.따라서 포인터 리시버파라메타를 주어 원본도 바뀔수 있게 한것이다

캡슐화 Getter/setter
Main패키지에서 접근 가능한것은 대문자로 되어있는 메서드이며 setter를 만들떄 항상 Set*의 형태로한다
또한 New로 시작하는 것은 생성자라 생각하면 좋다




## GoLang Mysql 

1.Bash 셸을 시작

2.홈 디렉터리에서 프로젝트 폴더생성 (예: mkdir -p ~/go/src/mysqlgo/).

3.디렉터리를 폴더로 변경(예: cd ~/go/src/mysqlgo/).

4.현재 홈 디렉터리의 go 폴더와 같이 유효한 소스 디렉터리를 가리키도록 GOPATH 환경 변수를 설정 Bash 셸에서 export GOPATH=~/go를 실행하여 go 디렉터리를 현재 셸 세션에 대한 GOPATH로 추가

5.go get github.com/go-sql-driver/mysql 명령을 실행하여 MySQL용 go-sql-driver(영문)를 설치
```
 mkdir -p ~/go/src/mysqlgo/
cd ~/go/src/mysqlgo/
export GOPATH=~/go/
go get github.com/go-sql-driver/mysql
```
요약!!!, Go 설치 후 다음 bash 명령을 실행합니다.연결



### 연결정보 가져오기

MySQL용 Azure Database에 연결하는 데 필요한 연결 정보 확인  정규화된 서버 이름 및 로그인 자격 증명이 필요!!

1.Azure Portal에 로그인
2.Azure Portal의 왼쪽 메뉴에서 모든 리소스 를 클릭한 다음, 방금 만든 서버를 검색(예: mydemoserver).
3.서버 이름을 클릭
4.서버의 개요 패널에 있는 서버 이름 과 서버 관리자 로그인 이름 을 기록 하기  암호를 잊어버리면 이 패널에서 암호를 재설정 가능


### Go 코드 작성 및 실행

Golang 코드를 작성하려면 Microsoft Windows의 메모장, Ubuntu의 vi 또는 Nano, macOS의 TextEdit과 같은 간단한 텍스트 편집기를 사용할 수 있습니다. 보다 풍부한 IDE(대화형 개발 환경)를 선호하는 경우 Jetbrains의 Gogland, Microsoft의 Visual Studio Code 또는 Atom을 사용해 보세요.
아래 섹션에서 Go 코드를 텍스트 파일에 붙여넣고 %USERPROFILE%\go\src\mysqlgo\createtable.go(Windows 경로) 또는 ~/go/src/mysqlgo/createtable.go(Linux 경로)와 같이 *.go 파일 확장명이 포함된 프로젝트 폴더에 저장합니다.
코드에서 HOST, DATABASE, USER 및 PASSWORD 상수를 찾아 예제 값을 사용자 고유의 값으로 바꿉니다.
명령 프롬프트 또는 Bash 셸을 시작합니다. 디렉터리를 프로젝트 폴더로 변경합니다. 예를 들어 Windows에서는 cd %USERPROFILE%\go\src\mysqlgo\이고, Linux에서는 cd ~/go/src/mysqlgo/입니다. 언급된 일부 IDE 편집기에서는 셸 명령 없이 디버그 및 런타임 기능을 제공합니다.
애플리케이션을 컴파일하고 실행하려면 go run createtable.go 명령을 입력하여 코드를 실행합니다.
또는 go build createtable.go 네이티브 애플리케이션에 코드를 빌드하려면 createtable.exe를 시작하여 애플리케이션을 실행합니다.

### 테이블 연결, 생성 및 데이터 삽입

다음 코드를 사용하여 서버에 연결하고, 테이블을 만들고, INSERT SQL 문을 통해 데이터를 로드합니다.

이 코드는 세 개의 패키지, 즉 sql 패키지, MySQL용 Azure Database와 통신할 드라이버로 사용되는 go sql driver for mysql, 명령줄에 출력되는 입출력을 위한 fmt 패키지를 가져옵니다.

sql.Open() 메서드를 호출하여 MySQL용 Azure Database에 연결하고 db.Ping() 메서드를 사용하여 연결을 확인합니다. 데이터베이스 핸들은 이러한 과정 내내 사용되며 데이터베이스 서버에 대한 연결 풀을 보유합니다. Exec() 메서드를 여러 번 호출하여 여러 DDL 명령을 실행합니다. 또한 Prepare() 및 Exec()를 사용하여 다른 매개 변수로 준비된 문을 실행하여 3개 행을 삽입합니다. 매번 사용자 지정 checkError() 메서드를 사용하여, 오류가 발생했는지와 서둘러 종료했는지 확인합니다.

host, database, user 및 password 상수는 원하는 값으로 바꿉니다

```
package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

const (
    host     = "mydemoserver.mysql.database.azure.com"
    database = "quickstartdb"
    user     = "myadmin@mydemoserver"
    password = "yourpassword"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    // Initialize connection string.
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    // Initialize connection object.
    db, err := sql.Open("mysql", connectionString)
    checkError(err)
    defer db.Close()

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database.")

    // Drop previous table of same name if one exists.
    _, err = db.Exec("DROP TABLE IF EXISTS inventory;")
    checkError(err)
    fmt.Println("Finished dropping table (if existed).")

    // Create table.
    _, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
    checkError(err)
    fmt.Println("Finished creating table.")

    // Insert some data into table.
    sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
    res, err := sqlStatement.Exec("banana", 150)
    checkError(err)
    rowCount, err := res.RowsAffected()
    fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

    res, err = sqlStatement.Exec("orange", 154)
    checkError(err)
    rowCount, err = res.RowsAffected()
    fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

    res, err = sqlStatement.Exec("apple", 100)
    checkError(err)
    rowCount, err = res.RowsAffected()
    fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
    fmt.Println("Done.")
}
```

### 데이터 읽기

SELECT SQL 문을 사용하여 데이터를 연결하고 읽으려면 다음 코드를 사용하세요.

이 코드는 세 개의 패키지, 즉 sql 패키지, MySQL용 Azure Database와 통신할 드라이버로 사용되는 go sql driver for mysql, 명령줄에 출력되는 입출력을 위한 fmt 패키지를 가져옵니다.

sql.Open() 메서드를 호출하여 MySQL용 Azure Database에 연결하고 db.Ping() 메서드를 사용하여 연결을 확인합니다. 데이터베이스 핸들은 이러한 과정 내내 사용되며 데이터베이스 서버에 대한 연결 풀을 보유합니다. Query() 메서드를 호출하여 select 명령을 실행합니다. 그런 다음 Next()를 실행하여 결과 집합을 반복하고 Scan()을 사용하여 열 값을 구문 분석하여 변수에 값을 저장합니다. 매번 사용자 지정 checkError() 메서드를 사용하여, 오류가 발생하여 서둘러 종료했는지 확인합니다.

host, database, user 및 password 상수는 원하는 값으로 바꿉니다.

```
package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

const (
    host     = "mydemoserver.mysql.database.azure.com"
    database = "quickstartdb"
    user     = "myadmin@mydemoserver"
    password = "yourpassword"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    // Initialize connection string.
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    // Initialize connection object.
    db, err := sql.Open("mysql", connectionString)
    checkError(err)
    defer db.Close()

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database.")

    // Variables for printing column data when scanned.
    var (
        id       int
        name     string
        quantity int
    )

    // Read some data from the table.
    rows, err := db.Query("SELECT id, name, quantity from inventory;")
    checkError(err)
    defer rows.Close()
    fmt.Println("Reading data:")
    for rows.Next() {
        err := rows.Scan(&id, &name, &quantity)
        checkError(err)
        fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
    }
    err = rows.Err()
    checkError(err)
    fmt.Println("Done.")
}
```

### 데이터 업데이트

UPDATE SQL 문을 사용하여 데이터를 연결하고 업데이트하려면 다음 코드를 사용하세요.

이 코드는 세 개의 패키지, 즉 sql 패키지, MySQL용 Azure Database와 통신할 드라이버로 사용되는 go sql driver for mysql, 명령줄에 출력되는 입출력을 위한 fmt 패키지를 가져옵니다.

sql.Open() 메서드를 호출하여 MySQL용 Azure Database에 연결하고 db.Ping() 메서드를 사용하여 연결을 확인합니다. 데이터베이스 핸들은 이러한 과정 내내 사용되며 데이터베이스 서버에 대한 연결 풀을 보유합니다. Exec() 메서드를 호출하여 update 명령을 실행합니다. 매번 사용자 지정 checkError() 메서드를 사용하여, 오류가 발생하여 서둘러 종료했는지 확인합니다.

host, database, user 및 password 상수는 원하는 값으로 바꿉니다.

```
package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

const (
    host     = "mydemoserver.mysql.database.azure.com"
    database = "quickstartdb"
    user     = "myadmin@mydemoserver"
    password = "yourpassword"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    // Initialize connection string.
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    // Initialize connection object.
    db, err := sql.Open("mysql", connectionString)
    checkError(err)
    defer db.Close()

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database.")

    // Modify some data in table.
    rows, err := db.Exec("UPDATE inventory SET quantity = ? WHERE name = ?", 200, "banana")
    checkError(err)
    rowCount, err := rows.RowsAffected()
    fmt.Printf("Updated %d row(s) of data.\n", rowCount)
    fmt.Println("Done.")
}
```

### 데이터 삭제

다음 코드를 사용하여 데이터를 연결하고 DELETE SQL 문을 통해 데이터를 제거합니다.

이 코드는 세 개의 패키지, 즉 sql 패키지, MySQL용 Azure Database와 통신할 드라이버로 사용되는 go sql driver for mysql, 명령줄에 출력되는 입출력을 위한 fmt 패키지를 가져옵니다.

sql.Open() 메서드를 호출하여 MySQL용 Azure Database에 연결하고 db.Ping() 메서드를 사용하여 연결을 확인합니다. 데이터베이스 핸들은 이러한 과정 내내 사용되며 데이터베이스 서버에 대한 연결 풀을 보유합니다. Exec() 메서드를 호출하여 delete 명령을 실행합니다. 매번 사용자 지정 checkError() 메서드를 사용하여, 오류가 발생하여 서둘러 종료했는지 확인합니다.

host, database, user 및 password 상수는 원하는 값으로 바꿉니다.

```
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

const (
    host     = "mydemoserver.mysql.database.azure.com"
    database = "quickstartdb"
    user     = "myadmin@mydemoserver"
    password = "yourpassword"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    // Initialize connection string.
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

    // Initialize connection object.
    db, err := sql.Open("mysql", connectionString)
    checkError(err)
    defer db.Close()

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database.")

    // Modify some data in table.
    rows, err := db.Exec("DELETE FROM inventory WHERE name = ?", "orange")
    checkError(err)
    rowCount, err := rows.RowsAffected()
    fmt.Printf("Deleted %d row(s) of data.\n", rowCount)
    fmt.Println("Done.")
}
```

### 리소스 정리
```
az group delete \
    --name $AZ_RESOURCE_GROUP \
    --yes
```


#서버연결

<img width="509" alt="스크린샷 2021-10-26 오전 12 36 59" src="https://user-images.githubusercontent.com/88940298/138727331-671b1969-66a1-4be2-acb9-177d1152a4b3.png">

3000 서버  연결
