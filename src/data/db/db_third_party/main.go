package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db: "author"`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("mysql", "gwp_dev_user:gwp@/gwp")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	fmt.Println(id)
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.AuthorName)
	fmt.Println(post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	state, err := Db.Prepare("insert into posts (content, author) values (?, ?)")
	_, err = state.Exec(post.Content, post.AuthorName)
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

func main() {
	post := Post{Content: "Hello", AuthorName: "jmoiron"}
	post.Create()
	fmt.Println(post)
	readPost := Post{}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)

}
