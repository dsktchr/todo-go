package main

import (
	"fmt"
	"database/sql"
	"time"
	"net/http"
	"context"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)


var ctx context.Context

func main() {
	db, err := sql.Open("mysql", "todo-user:todopassword@tcp(localhost:3308)/todo")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	time.Sleep(100 * time.Millisecond)

	if err := db.Ping(); err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	createTableStmt := `CREATE TABLE IF NOT EXISTS todos (
id int not null auto_increment primary key,
name varchar(100)
)`
	ctx = context.Background()
	var result sql.Result
	result, err = db.ExecContext(ctx, createTableStmt)

	if err != nil {
		panic(err)
	}

	insertStmt := `INSERT INTO todos (name) VALUES ('お使い')`
	result, err = db.ExecContext(ctx, insertStmt)
	if err != nil {
		panic(err)
	}

	var todoId int64
	todoId, err = result.LastInsertId()

	if err != nil {
		panic(err)
	}
	fmt.Printf("TODO-ID=%v", todoId)

	var (
		id int64
		name string
	)

	err = db.QueryRowContext(ctx, "select id, name from todos where id=?", todoId).Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("id=%v, name=%v", id, name)

	result, err = db.ExecContext(ctx, "update todos set name=? where id=?", name + strconv.FormatInt(id, 10), id)

	if err != nil {
		panic(err)
	}

	var rows int64
	rows, err = result.RowsAffected()
	if rows != 1 {
		panic(err)
	}

}
