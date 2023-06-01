package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("TODOリストアプリ")

	for {
		fmt.Println("\n[1] TODOの追加")
		fmt.Println("[2] TODOの表示")
		fmt.Println("[3] 終了")
		fmt.Print("選択してください: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTodo()
		case "2":
			displayTodos()
		case "3":
			fmt.Println("アプリを終了します。")
			return
		default:
			fmt.Println("無効な選択です。もう一度選択してください。")
		}
	}
}

func addTodo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nTODOの追加")
	fmt.Print("タスクを入力してください: ")

	scanner.Scan()
	task := scanner.Text()

	todo := Todo{
		ID:   len(todos) + 1,
		Task: task,
	}

	todos = append(todos, todo)

	fmt.Println("TODOが追加されました。")
}

func displayTodos() {
	fmt.Println("\nTODOリスト")

	if len(todos) == 0 {
		fmt.Println("TODOはありません。")
		return
	}

	for _, todo := range todos {
		fmt.Printf("%d. %s\n", todo.ID, todo.Task)
	}
}