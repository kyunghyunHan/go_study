
# ๐ฉ๐ปโ๐GoLang Mysql 

1.Bash ์ธ์ ์์

2.ํ ๋๋ ํฐ๋ฆฌ์์ ํ๋ก์ ํธ ํด๋์์ฑ (์: mkdir -p ~/go/src/mysqlgo/).

3.๋๋ ํฐ๋ฆฌ๋ฅผ ํด๋๋ก ๋ณ๊ฒฝ(์: cd ~/go/src/mysqlgo/).

4.ํ์ฌ ํ ๋๋ ํฐ๋ฆฌ์ go ํด๋์ ๊ฐ์ด ์ ํจํ ์์ค ๋๋ ํฐ๋ฆฌ๋ฅผ ๊ฐ๋ฆฌํค๋๋ก GOPATH ํ๊ฒฝ ๋ณ์๋ฅผ ์ค์  Bash ์ธ์์ export GOPATH=~/go๋ฅผ ์คํํ์ฌ go ๋๋ ํฐ๋ฆฌ๋ฅผ ํ์ฌ ์ธ ์ธ์์ ๋ํ GOPATH๋ก ์ถ๊ฐ

5.go get github.com/go-sql-driver/mysql ๋ช๋ น์ ์คํํ์ฌ MySQL์ฉ go-sql-driver(์๋ฌธ)๋ฅผ ์ค์น
```
 mkdir -p ~/go/src/mysqlgo/
cd ~/go/src/mysqlgo/
export GOPATH=~/go/
go get github.com/go-sql-driver/mysql
```
์์ฝ!!!, Go ์ค์น ํ ๋ค์ bash ๋ช๋ น์ ์คํํฉ๋๋ค.์ฐ๊ฒฐ



### ์ฐ๊ฒฐ์ ๋ณด ๊ฐ์ ธ์ค๊ธฐ

MySQL์ฉ Azure Database์ ์ฐ๊ฒฐํ๋ ๋ฐ ํ์ํ ์ฐ๊ฒฐ ์ ๋ณด ํ์ธ  ์ ๊ทํ๋ ์๋ฒ ์ด๋ฆ ๋ฐ ๋ก๊ทธ์ธ ์๊ฒฉ ์ฆ๋ช์ด ํ์!!

1.Azure Portal์ ๋ก๊ทธ์ธ
2.Azure Portal์ ์ผ์ชฝ ๋ฉ๋ด์์ ๋ชจ๋  ๋ฆฌ์์ค ๋ฅผ ํด๋ฆญํ ๋ค์, ๋ฐฉ๊ธ ๋ง๋  ์๋ฒ๋ฅผ ๊ฒ์(์: mydemoserver).
3.์๋ฒ ์ด๋ฆ์ ํด๋ฆญ
4.์๋ฒ์ ๊ฐ์ ํจ๋์ ์๋ ์๋ฒ ์ด๋ฆ ๊ณผ ์๋ฒ ๊ด๋ฆฌ์ ๋ก๊ทธ์ธ ์ด๋ฆ ์ ๊ธฐ๋ก ํ๊ธฐ  ์ํธ๋ฅผ ์์ด๋ฒ๋ฆฌ๋ฉด ์ด ํจ๋์์ ์ํธ๋ฅผ ์ฌ์ค์  ๊ฐ๋ฅ


### Go ์ฝ๋ ์์ฑ ๋ฐ ์คํ

Golang ์ฝ๋๋ฅผ ์์ฑํ๋ ค๋ฉด Microsoft Windows์ ๋ฉ๋ชจ์ฅ, Ubuntu์ vi ๋๋ Nano, macOS์ TextEdit๊ณผ ๊ฐ์ ๊ฐ๋จํ ํ์คํธ ํธ์ง๊ธฐ๋ฅผ ์ฌ์ฉํ  ์ ์์ต๋๋ค. ๋ณด๋ค ํ๋ถํ IDE(๋ํํ ๊ฐ๋ฐ ํ๊ฒฝ)๋ฅผ ์ ํธํ๋ ๊ฒฝ์ฐ Jetbrains์ Gogland, Microsoft์ Visual Studio Code ๋๋ Atom์ ์ฌ์ฉํด ๋ณด์ธ์.
์๋ ์น์์์ Go ์ฝ๋๋ฅผ ํ์คํธ ํ์ผ์ ๋ถ์ฌ๋ฃ๊ณ  %USERPROFILE%\go\src\mysqlgo\createtable.go(Windows ๊ฒฝ๋ก) ๋๋ ~/go/src/mysqlgo/createtable.go(Linux ๊ฒฝ๋ก)์ ๊ฐ์ด *.go ํ์ผ ํ์ฅ๋ช์ด ํฌํจ๋ ํ๋ก์ ํธ ํด๋์ ์ ์ฅํฉ๋๋ค.
์ฝ๋์์ HOST, DATABASE, USER ๋ฐ PASSWORD ์์๋ฅผ ์ฐพ์ ์์  ๊ฐ์ ์ฌ์ฉ์ ๊ณ ์ ์ ๊ฐ์ผ๋ก ๋ฐ๊ฟ๋๋ค.
๋ช๋ น ํ๋กฌํํธ ๋๋ Bash ์ธ์ ์์ํฉ๋๋ค. ๋๋ ํฐ๋ฆฌ๋ฅผ ํ๋ก์ ํธ ํด๋๋ก ๋ณ๊ฒฝํฉ๋๋ค. ์๋ฅผ ๋ค์ด Windows์์๋ cd %USERPROFILE%\go\src\mysqlgo\์ด๊ณ , Linux์์๋ cd ~/go/src/mysqlgo/์๋๋ค. ์ธ๊ธ๋ ์ผ๋ถ IDE ํธ์ง๊ธฐ์์๋ ์ธ ๋ช๋ น ์์ด ๋๋ฒ๊ทธ ๋ฐ ๋ฐํ์ ๊ธฐ๋ฅ์ ์ ๊ณตํฉ๋๋ค.
์ ํ๋ฆฌ์ผ์ด์์ ์ปดํ์ผํ๊ณ  ์คํํ๋ ค๋ฉด go run createtable.go ๋ช๋ น์ ์๋ ฅํ์ฌ ์ฝ๋๋ฅผ ์คํํฉ๋๋ค.
๋๋ go build createtable.go ๋ค์ดํฐ๋ธ ์ ํ๋ฆฌ์ผ์ด์์ ์ฝ๋๋ฅผ ๋น๋ํ๋ ค๋ฉด createtable.exe๋ฅผ ์์ํ์ฌ ์ ํ๋ฆฌ์ผ์ด์์ ์คํํฉ๋๋ค.

### ํ์ด๋ธ ์ฐ๊ฒฐ, ์์ฑ ๋ฐ ๋ฐ์ดํฐ ์ฝ์

๋ค์ ์ฝ๋๋ฅผ ์ฌ์ฉํ์ฌ ์๋ฒ์ ์ฐ๊ฒฐํ๊ณ , ํ์ด๋ธ์ ๋ง๋ค๊ณ , INSERT SQL ๋ฌธ์ ํตํด ๋ฐ์ดํฐ๋ฅผ ๋ก๋ํฉ๋๋ค.

์ด ์ฝ๋๋ ์ธ ๊ฐ์ ํจํค์ง, ์ฆ sql ํจํค์ง, MySQL์ฉ Azure Database์ ํต์ ํ  ๋๋ผ์ด๋ฒ๋ก ์ฌ์ฉ๋๋ go sql driver for mysql, ๋ช๋ น์ค์ ์ถ๋ ฅ๋๋ ์์ถ๋ ฅ์ ์ํ fmt ํจํค์ง๋ฅผ ๊ฐ์ ธ์ต๋๋ค.

sql.Open() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ MySQL์ฉ Azure Database์ ์ฐ๊ฒฐํ๊ณ  db.Ping() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ ์ฐ๊ฒฐ์ ํ์ธํฉ๋๋ค. ๋ฐ์ดํฐ๋ฒ ์ด์ค ํธ๋ค์ ์ด๋ฌํ ๊ณผ์  ๋ด๋ด ์ฌ์ฉ๋๋ฉฐ ๋ฐ์ดํฐ๋ฒ ์ด์ค ์๋ฒ์ ๋ํ ์ฐ๊ฒฐ ํ์ ๋ณด์ ํฉ๋๋ค. Exec() ๋ฉ์๋๋ฅผ ์ฌ๋ฌ ๋ฒ ํธ์ถํ์ฌ ์ฌ๋ฌ DDL ๋ช๋ น์ ์คํํฉ๋๋ค. ๋ํ Prepare() ๋ฐ Exec()๋ฅผ ์ฌ์ฉํ์ฌ ๋ค๋ฅธ ๋งค๊ฐ ๋ณ์๋ก ์ค๋น๋ ๋ฌธ์ ์คํํ์ฌ 3๊ฐ ํ์ ์ฝ์ํฉ๋๋ค. ๋งค๋ฒ ์ฌ์ฉ์ ์ง์  checkError() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ, ์ค๋ฅ๊ฐ ๋ฐ์ํ๋์ง์ ์๋๋ฌ ์ข๋ฃํ๋์ง ํ์ธํฉ๋๋ค.

host, database, user ๋ฐ password ์์๋ ์ํ๋ ๊ฐ์ผ๋ก ๋ฐ๊ฟ๋๋ค

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

### ๋ฐ์ดํฐ ์ฝ๊ธฐ

SELECT SQL ๋ฌธ์ ์ฌ์ฉํ์ฌ ๋ฐ์ดํฐ๋ฅผ ์ฐ๊ฒฐํ๊ณ  ์ฝ์ผ๋ ค๋ฉด ๋ค์ ์ฝ๋๋ฅผ ์ฌ์ฉํ์ธ์.

์ด ์ฝ๋๋ ์ธ ๊ฐ์ ํจํค์ง, ์ฆ sql ํจํค์ง, MySQL์ฉ Azure Database์ ํต์ ํ  ๋๋ผ์ด๋ฒ๋ก ์ฌ์ฉ๋๋ go sql driver for mysql, ๋ช๋ น์ค์ ์ถ๋ ฅ๋๋ ์์ถ๋ ฅ์ ์ํ fmt ํจํค์ง๋ฅผ ๊ฐ์ ธ์ต๋๋ค.

sql.Open() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ MySQL์ฉ Azure Database์ ์ฐ๊ฒฐํ๊ณ  db.Ping() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ ์ฐ๊ฒฐ์ ํ์ธํฉ๋๋ค. ๋ฐ์ดํฐ๋ฒ ์ด์ค ํธ๋ค์ ์ด๋ฌํ ๊ณผ์  ๋ด๋ด ์ฌ์ฉ๋๋ฉฐ ๋ฐ์ดํฐ๋ฒ ์ด์ค ์๋ฒ์ ๋ํ ์ฐ๊ฒฐ ํ์ ๋ณด์ ํฉ๋๋ค. Query() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ select ๋ช๋ น์ ์คํํฉ๋๋ค. ๊ทธ๋ฐ ๋ค์ Next()๋ฅผ ์คํํ์ฌ ๊ฒฐ๊ณผ ์งํฉ์ ๋ฐ๋ณตํ๊ณ  Scan()์ ์ฌ์ฉํ์ฌ ์ด ๊ฐ์ ๊ตฌ๋ฌธ ๋ถ์ํ์ฌ ๋ณ์์ ๊ฐ์ ์ ์ฅํฉ๋๋ค. ๋งค๋ฒ ์ฌ์ฉ์ ์ง์  checkError() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ, ์ค๋ฅ๊ฐ ๋ฐ์ํ์ฌ ์๋๋ฌ ์ข๋ฃํ๋์ง ํ์ธํฉ๋๋ค.

host, database, user ๋ฐ password ์์๋ ์ํ๋ ๊ฐ์ผ๋ก ๋ฐ๊ฟ๋๋ค.

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

### ๋ฐ์ดํฐ ์๋ฐ์ดํธ

UPDATE SQL ๋ฌธ์ ์ฌ์ฉํ์ฌ ๋ฐ์ดํฐ๋ฅผ ์ฐ๊ฒฐํ๊ณ  ์๋ฐ์ดํธํ๋ ค๋ฉด ๋ค์ ์ฝ๋๋ฅผ ์ฌ์ฉํ์ธ์.

์ด ์ฝ๋๋ ์ธ ๊ฐ์ ํจํค์ง, ์ฆ sql ํจํค์ง, MySQL์ฉ Azure Database์ ํต์ ํ  ๋๋ผ์ด๋ฒ๋ก ์ฌ์ฉ๋๋ go sql driver for mysql, ๋ช๋ น์ค์ ์ถ๋ ฅ๋๋ ์์ถ๋ ฅ์ ์ํ fmt ํจํค์ง๋ฅผ ๊ฐ์ ธ์ต๋๋ค.

sql.Open() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ MySQL์ฉ Azure Database์ ์ฐ๊ฒฐํ๊ณ  db.Ping() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ ์ฐ๊ฒฐ์ ํ์ธํฉ๋๋ค. ๋ฐ์ดํฐ๋ฒ ์ด์ค ํธ๋ค์ ์ด๋ฌํ ๊ณผ์  ๋ด๋ด ์ฌ์ฉ๋๋ฉฐ ๋ฐ์ดํฐ๋ฒ ์ด์ค ์๋ฒ์ ๋ํ ์ฐ๊ฒฐ ํ์ ๋ณด์ ํฉ๋๋ค. Exec() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ update ๋ช๋ น์ ์คํํฉ๋๋ค. ๋งค๋ฒ ์ฌ์ฉ์ ์ง์  checkError() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ, ์ค๋ฅ๊ฐ ๋ฐ์ํ์ฌ ์๋๋ฌ ์ข๋ฃํ๋์ง ํ์ธํฉ๋๋ค.

host, database, user ๋ฐ password ์์๋ ์ํ๋ ๊ฐ์ผ๋ก ๋ฐ๊ฟ๋๋ค.

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

### ๋ฐ์ดํฐ ์ญ์ 

๋ค์ ์ฝ๋๋ฅผ ์ฌ์ฉํ์ฌ ๋ฐ์ดํฐ๋ฅผ ์ฐ๊ฒฐํ๊ณ  DELETE SQL ๋ฌธ์ ํตํด ๋ฐ์ดํฐ๋ฅผ ์ ๊ฑฐํฉ๋๋ค.

์ด ์ฝ๋๋ ์ธ ๊ฐ์ ํจํค์ง, ์ฆ sql ํจํค์ง, MySQL์ฉ Azure Database์ ํต์ ํ  ๋๋ผ์ด๋ฒ๋ก ์ฌ์ฉ๋๋ go sql driver for mysql, ๋ช๋ น์ค์ ์ถ๋ ฅ๋๋ ์์ถ๋ ฅ์ ์ํ fmt ํจํค์ง๋ฅผ ๊ฐ์ ธ์ต๋๋ค.

sql.Open() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ MySQL์ฉ Azure Database์ ์ฐ๊ฒฐํ๊ณ  db.Ping() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ ์ฐ๊ฒฐ์ ํ์ธํฉ๋๋ค. ๋ฐ์ดํฐ๋ฒ ์ด์ค ํธ๋ค์ ์ด๋ฌํ ๊ณผ์  ๋ด๋ด ์ฌ์ฉ๋๋ฉฐ ๋ฐ์ดํฐ๋ฒ ์ด์ค ์๋ฒ์ ๋ํ ์ฐ๊ฒฐ ํ์ ๋ณด์ ํฉ๋๋ค. Exec() ๋ฉ์๋๋ฅผ ํธ์ถํ์ฌ delete ๋ช๋ น์ ์คํํฉ๋๋ค. ๋งค๋ฒ ์ฌ์ฉ์ ์ง์  checkError() ๋ฉ์๋๋ฅผ ์ฌ์ฉํ์ฌ, ์ค๋ฅ๊ฐ ๋ฐ์ํ์ฌ ์๋๋ฌ ์ข๋ฃํ๋์ง ํ์ธํฉ๋๋ค.

host, database, user ๋ฐ password ์์๋ ์ํ๋ ๊ฐ์ผ๋ก ๋ฐ๊ฟ๋๋ค.

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

### ๋ฆฌ์์ค ์ ๋ฆฌ
```
az group delete \
    --name $AZ_RESOURCE_GROUP \
    --yes
```


#์๋ฒ์ฐ๊ฒฐ

<img width="509" alt="แแณแแณแแตแซแแฃแบ 2021-10-26 แแฉแแฅแซ 12 36 59" src="https://user-images.githubusercontent.com/88940298/138727331-671b1969-66a1-4be2-acb9-177d1152a4b3.png">

3000 ์๋ฒ  ์ฐ๊ฒฐ
