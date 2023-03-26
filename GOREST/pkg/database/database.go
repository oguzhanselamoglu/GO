package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	recId        string
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

	//updateUser("3c67cf5e-c0d3-11ed-89d6-0242ac120003", "veli", "ali", "aaaa@aaa.com")
	//deleteUser("4df3c7fa-c0d3-11ed-baf2-0242ac120003")
	getUser()

	fmt.Println(UserList)
	//defer rows.Close()

	CheckError(err)

}

func deleteUser(recId string) {
	stmt, err := db.Prepare(`delete from "users" where recid =$1 `)
	CheckError(err)
	_, err = stmt.Exec(recId)
}
func updateUser(recId, firstname, lastname, email string) {
	stmt, err := db.Prepare(`update "users" set firstname =$1, lastname = $2, email = $3 where recid = $4`)
	CheckError(err)
	_, err = stmt.Exec(firstname, lastname, email, recId)
	CheckError(err)

	fmt.Println("Kayıt silindi")

}

var UserList []User

func getUser() {
	rows, err := db.Query(`select recid,firstname,lastname,email,role_id,enterprises_id,deleted,passive from "users"`)
	if err == nil {
		for rows.Next() {
			var firstname, lastname, email, roleId, enterpriseId, recid string
			var deleted, passive bool

			err = rows.Scan(&recid, &firstname, &lastname, &email, &roleId, &enterpriseId, &deleted, &passive)
			CheckError(err)

			UserList = append(UserList, User{
				recId:        recid,
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
