package test02

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	PostId  int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "myuser:mypass@(127.0.0.1:3306)/gwp?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	_, err = Db.Exec("insert into comments (content, author, post_id)"+
		" values (?, ?, ?);", comment.Content, comment.Author, comment.Post.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = "+
		"?", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, post_id, content, author from comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Content, &comment.Author)
		if err != nil {
			return
		}

		if comment.PostId == post.Id {
			post.Comments = append(post.Comments, comment)
		}
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	result, err := Db.Exec("insert into posts (content, author) values (?, ?);",
		post.Content, post.Author)

	a, _ := result.LastInsertId()
	post.Id = int(a)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	err := comment.Create()
	if err != nil {
		fmt.Println("err = " + err.Error())
	}
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
	////
}
