package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo represents a todo item with a description and due date.
type Todo struct {
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

// Todos represents a collection of todo items.
type Todos []Todo

func main() {
	fmt.Println("Welcome to the TodoCLI App!")

	// Load existing todos from file
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice (1-3): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTodo(&todos)
		case 2:
			listTodos(todos)
		case 3:
			saveAndExit(todos)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
		}
	}
}

func addTodo(todos *Todos) {
	fmt.Println("\nAdding a new todo:")
	var description string
	fmt.Print("Enter todo description: ")
	fmt.Scanln(&description)

	var dueDateString string
	fmt.Print("Enter due date (YYYY-MM-DD): ")
	fmt.Scanln(&dueDateString)

	// Parse due date string to time.Time
	dueDate, err := time.Parse("2006-01-02", dueDateString)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Add new todo to the list
	newTodo := Todo{
		Description: description,
		DueDate:     dueDate,
	}
	*todos = append(*todos, newTodo)

	fmt.Println("Todo added successfully!")
}

func listTodos(todos Todos) {
	fmt.Println("\nListing todos:")
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	for i, todo := range todos {
		fmt.Printf("%d. %s - Due: %s\n", i+1, todo.Description, todo.DueDate.Format("2006-01-02"))
	}
}

func saveAndExit(todos Todos) {
	// Save todos to file
	err := saveTodos(todos)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		return
	}

	fmt.Println("Todos saved. Exiting the TodoCLI App. Goodbye!")
	os.Exit(0)
}

func loadTodos() (Todos, error) {
	// Read from todos.json file
	fileContent, err := os.ReadFile("todos.json")
	if err != nil {
		return nil, err
	}

	var todos Todos
	err = json.Unmarshal(fileContent, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func saveTodos(todos Todos) error {
	// Convert todos to JSON format
	fileContent, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	// Write to todos.json file
	err = os.WriteFile("todos.json", fileContent, 0644)
	if err != nil {
		return err
	}

	// what is the meaning of the 0644
	// the first number determined the type of file
	// which is in our case it just a regular file
	// 644 refer to linux permissions
	// r (read): 4
	//w (write): 2
	// x (execute): 1
	// so the 644 means -> ( read + write ) , ( write ) , ( write )
	//                          owner       ,   group   ,   other
	// mean owner of file and read and write to file , other users in that group can
	// write to file and the other user also able to write to file

	return nil
}
