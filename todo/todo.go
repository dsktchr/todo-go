package todo


import (
	"log"
	"context"
	"github.com/dsktchr/todo-go/db"
)

type Todo struct {
	Id   int64
	Name string
}

func FindAll(ctx context.Context) []Todo {
	rows, err := db.DB.QueryContext(ctx, "SELECT id, name FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	todoList := make([]Todo, 0)
	
	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.Id, &todo.Name); err != nil {
			log.Fatal(err)
		}

		todoList = append(todoList, todo)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todoList
}

func FindOne(ctx context.Context, todoId int64) Todo {
	row := db.DB.QueryRowContext(ctx, "SELECT id, name from todos WHERE id=?", todoId)
	todo := Todo{}
	if err := row.Scan(&todo.Id, &todo.Name); err != nil {
		log.Fatal(err)
	}
	
	return todo
}

func Create(ctx context.Context, name string) int64 {
	result, err := db.DB.ExecContext(ctx, "INSERT INTO todos (name) VALUES (?)", name)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func Update(ctx context.Context, todoId int64, name string) {
	result, err := db.DB.ExecContext(ctx, "UPDATE todos SET name=? WHERE id=?", name, todoId)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatal("1アイテム以上追加されています. rows=", rows) 
	}
}


func Delete(ctx context.Context, todoId int64) {
	result, err := db.DB.ExecContext(ctx, "DELETE FROM todos WHERE id=?", todoId)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatal("1アイテムの削除を期待しています. 実際に削除されたRow=?", rows)
	}
}
