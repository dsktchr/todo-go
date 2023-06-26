package main

import (
	"context"
	"fmt"
	"time"
	"net/http"
	"github.com/dsktchr/todo-go/db"
	"github.com/dsktchr/todo-go/todo"
)

var ctx context.Context

func main() {

	defer db.DB.Close()

	time.Sleep(100 * time.Millisecond)

	if err := db.DB.Ping(); err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	ctx = context.Background()

	//newTodoId := todo.Create(ctx)
	//fmt.Printf("TodoId=%v のアイテムを作成しました\n", newTodoId)

	//todoList := todo.FindAll(ctx)
	//fmt.Println(todoList)

	//todo.Update(ctx, newTodoId, "お花見")

	//updatedTodo := todo.FindOne(ctx, newTodoId)
	//fmt.Println(updatedTodo)

	//todo.Delete(ctx, newTodoId)

	//todoList = todo.FindAll(ctx)
	//fmt.Println(todoList)

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		todoList := todo.FindAll(ctx)
		fmt.Println(todoList)
	})

	http.ListenAndServe(":8081", nil)
}
