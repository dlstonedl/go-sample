package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
	sex      string
	age      int
}

func main() {
	db, err := sql.Open("mysql", "root:@/mybatis_analysis")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from user where id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(1)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.id, &user.username, &user.sex, &user.age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", user)
	}
}
