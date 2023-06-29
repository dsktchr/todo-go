package main

import (
	"context"
	"fmt"
	"time"
	"net/http"
	"strings"
	"strconv"
	"encoding/json"
	"github.com/dsktchr/todo-go/db"
	"github.com/dsktchr/todo-go/todo"
)

var ctx context.Context

type RequestTodo struct {
	Name string `json:"name"`
}

func main() {

	defer db.DB.Close()

	time.Sleep(100 * time.Millisecond)

	if err := db.DB.Ping(); err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	ctx = context.Background()

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println(r)
			todoList := todo.FindAll(ctx)
			fmt.Println(todoList)
		}
		switch r.Method {
		case http.MethodGet:
			fmt.Println(">>GETリクエスト")
			todoList := todo.FindAll(ctx)
			fmt.Println(todoList)
			fmt.Println("<<GETリクエスト")
		case http.MethodPost:
			fmt.Println(">>POSTリクエスト")
			var reqTodo RequestTodo
			err := json.NewDecoder(r.Body).Decode(&reqTodo) 
			defer r.Body.Close()
			if err != nil {
				panic(err)
			}
			fmt.Println(reqTodo)
			todoId := todo.Create(ctx, reqTodo.Name)
			todoItem := todo.FindOne(ctx, todoId)
			fmt.Println(todoItem)
			fmt.Println("<<POSTリクエスト")
		}
	})

	http.HandleFunc("/todo/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			fmt.Println(">>DELETEリクエスト")
			todoPath := strings.Split(r.URL.Path, "/")
			todoId, err := strconv.ParseInt(todoPath[2],10,64)
			if err != nil {
				panic(err)
			}
			todo.Delete(ctx, todoId)
			fmt.Println("<<DELETEリクエスト")
		}
	})

	http.ListenAndServe(":8081", nil)
}
