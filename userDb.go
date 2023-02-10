package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpwd = "ROBrob528%@*"
const dbname = "life_care"

func GetUsers() []User {

	db, err := sql.Open("mysql", dbuser+":"+dbpwd+"@tcp(127.0.0.1:3306)/"+dbname)

	// If there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("ERR:", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("select * from User")

	if err != nil {
		fmt.Println("ERR", err.Error())
		return nil
	}
	users := []User{}

	for results.Next() {
		var iterator User

		err = results.Scan(&iterator.Id, &iterator.Name, &iterator.Password, &iterator.Email)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, iterator)

		// fmt.Println(iterator)
	}

	return users
}

func AddUser(user User) {
	db, err := sql.Open("mysql", dbuser+":"+dbpwd+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished executing
	defer db.Close()
	fmt.Println(user)
	insert, err := db.Query("insert into user (id, name, email, password) values (?, ?, ?, ?)", user.Id, user.Name, user.Email, user.Password)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
