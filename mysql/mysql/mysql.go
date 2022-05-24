package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err1 := sql.Open("mysql", "root:Ovopark#2021@(47.97.255.245:31117)/turbo?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err1 != nil {
		fmt.Println("Open:", err1.Error())
	}
	err2 := db.Ping()
	if err2 != nil {
		fmt.Println("Ping:", err2.Error())
	}
	//create_table(db)
	//insert_(db)
	//query_single_row(db)
	query_all(db)
}

func query_all(db *sql.DB) {
	query_sql := `SELECT * FROM users`
	rows, _ := db.Query(query_sql)
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
		users = append(users, u)
	}
	bt, _ := json.Marshal(users)
	fmt.Println(string(bt))
}

func query_single_row(db *sql.DB) {
	var user User
	query_sql := `SELECT * FROM users WHERE id = ?`
	row := db.QueryRow(query_sql, 1).Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if row != nil {
		fmt.Println("Query error:", row.Error())
	}
	fmt.Println(user.CreatedAt)
}

func insert_(db *sql.DB) {
	username := "lisi"
	password := "123456"
	create_at := time.Now()
	fmt.Println(create_at)
	insert_sql := `INSERT INTO users (username,password,created_at) VALUES (?,?,?)`
	result, err := db.Exec(insert_sql, username, password, create_at)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 获取档次插入的主键自增的id
	id, err2 := result.LastInsertId()
	if err2 == nil {
		fmt.Println("插入ID：", id)
	}
}

const create_table_sql = `CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
);`

func create_table(db *sql.DB) {
	_, err := db.Exec(create_table_sql)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type User struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
}
