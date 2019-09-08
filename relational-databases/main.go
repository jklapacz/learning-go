package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
)

type Profile struct{
	ProfileId int
	Name string
	Age int
}

func main() {
	db, err := sql.Open("sqlite3", "../godb.db")
	checkError(err)

	statement, err := db.Prepare("insert into Profile (name, age) values (?, ?)")
	checkError(err)

	statement.Exec("Jessica", 30)

	update, err := db.Prepare("update Profile set age = ? where name = ?")
	checkError(err)

	update.Exec(40, "Jessica")

	var profile Profile
	rows, err := db.Query("select id, age, name from Profile")
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&profile.ProfileId, &profile.Age, &profile.Name)
		checkError(err)
		fmt.Println(profile)
	}
	rows.Close()
	db.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
