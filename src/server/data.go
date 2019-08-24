package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "gwp_dev_user:gwp@/gwp")
	if err != nil {
		panic(err)
	}
}

func (post Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement, err := Db.Prepare("insert into posts (content, author) values (?, ?)")
	_, err = statement.Exec(post.Content, post.Author)
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&post.Id)
	if err != nil {
		panic(err)
	}

	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}
