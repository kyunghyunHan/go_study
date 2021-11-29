
# GoLang Mysql 

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
