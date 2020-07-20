package main

import "database/sql"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "myuser:mypass@(127.0.0.1:3306)/gwp?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
}
