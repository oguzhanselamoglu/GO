package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Database() {
	fmt.Println("Database")

	connStr := "postgres://veboni:Xidok4096H!@localhost/veboni-ae7949a8-7aa3-4038-b085-a98139bbe329?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT "code" FROM "boards"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var code string

		err = rows.Scan(&code)
		CheckError(err)

		fmt.Println(code)
	}

	CheckError(err)

}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
