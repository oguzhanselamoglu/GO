package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	firstName    string
	lastName     string
	email        string
	roleId       string
	enterpriseId string
	deleted      bool
	passive      bool
}

func Database() {
	fmt.Println("Database")
	var err error
	connStr := "postgres://veboni:Xidok4096H!@localhost/veboni-ae7949a8-7aa3-4038-b085-a98139bbe329?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	//	addUser("oguzhan", "selamoglu", "oguzhan@kod.com.tr")

	getUser()

	fmt.Println(userList)
	//defer rows.Close()

	CheckError(err)

}

var userList []User

func getUser() {
	rows, err := db.Query(`select firstname,lastname,email,role_id,enterprises_id,deleted,passive from "users"`)
	if err == nil {
		for rows.Next() {
			var firstname, lastname, email, roleId, enterpriseId string
			var deleted, passive bool

			err = rows.Scan(&firstname, &lastname, &email, &roleId, &enterpriseId, &deleted, &passive)
			CheckError(err)

			userList = append(userList, User{
				firstName:    firstname,
				lastName:     lastname,
				email:        email,
				roleId:       roleId,
				enterpriseId: "",
				deleted:      deleted,
				passive:      passive,
			})

		}
		rows.Close()
	}
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func addUser(firstName, lastName, email string) {
	stmt, err := db.Prepare(`insert into "users"("firstname","lastname","email","role_id","enterprises_id","authorizeenterprise","deleted","passive") values ($1, $2, $3, '46955c7a-a650-41f9-bad4-d2d8e130cb90','ae7949a8-7aa3-4038-b085-a98139bbe329',true,false,false) `)
	CheckError(err)
	res, err := stmt.Exec(firstName, lastName, email)
	CheckError(err)
	//	id, err := res.LastInsertId()
	//	CheckError(err)
	fmt.Println("Kayıtlı id", res)

}
