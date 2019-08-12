package main

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
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

// DB初期化
func init() {
	var err error
	Db, err = sql.Open("mysql", "gwp_dev_user:gwp@/gwp")
	if err != nil {
		panic(err)
	}
}

// コメントを取得
func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("投稿が見つかりません")
		return
	}
	statement, err := Db.Prepare("insert into comments (content, author, post_id) values (?, ?, ?)")
	_, err = statement.Exec(comment.Content, comment.Author, comment.Post.Id)
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&comment.Id)
	if err != nil {
		panic(err)
	}
	return
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
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments where post_id = ?", id)
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		fmt.Println(comment)
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
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
	err := post.Create()
	if err != nil {
		panic(err)
	}
	fmt.Println(post)

	comment := Comment{
		Content: "good",
		Author:  "John",
		Post:    &post,
	}

	comment.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
