package main

import (
	"database/sql"
	"fmt"
	"log"
)


func main() {
	res := selectData("test")
	if res {
		fmt.Println("data exist")
	} else {
		fmt.Println("data not exist")
	}
}

func selectData(name string) bool {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("select name from test where id = ?", 1).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}


