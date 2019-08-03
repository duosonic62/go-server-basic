package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "gwp_dev_user:gwp@/gwp")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit ?", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	fmt.Println(post)
	statement, err := Db.Prepare("insert into posts (content, author) values (?, ?)")
	_, err = statement.Exec(post.Content, post.Author)
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&post.Id)
	if err != nil {
		panic(err)
	}

	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}

func main() {
	post := Post{Content: "Hello, World!", Author: "Sau"}
	fmt.Println(post)
	err := post.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	readPost.Content = "Bonjour"
	readPost.Author = "Pierre"
	readPost.Update()

	updatePost, _ := GetPost(post.Id)
	fmt.Println(updatePost)

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()

	posts, _ = Posts(10)
	fmt.Println(posts)
}
