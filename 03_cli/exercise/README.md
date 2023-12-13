# TodoCLI Golang Exercise

## Objective

Create a simple Command Line Interface (CLI) application in Golang for managing todo items with due dates. The application should allow users to add new todos with due dates, list existing todos, and store them in a JSON file for persistence.

## Steps

1. **Setup Project:**
    - Create a new directory for your project, e.g., `todo-cli`.
    - Inside the project directory, initialize a Go module using `go mod init todo-cli`.

2. **Define Todo Struct:**
    - Create a `main.go` file in your project.
    - Define a `Todo` struct with fields for description (`string`) and due date (`time.Time`).

3. **Define Todos Slice:**
    - Define a `Todos` type as a slice of `Todo` structs.
    - This will represent the collection of todo items.

4. **Implement Main Function:**
    - Write the `main` function that serves as the entry point of the program.
    - Print a welcome message to the console.

5. **Load Existing Todos:**
    - Create a function `loadTodos` to read existing todos from a JSON file (`todos.json`) and return a `Todos` slice.

6. **Add Todo Functionality:**
    - Implement a function `addTodo` that takes user input for todo description and due date.
    - Parse the due date input to a `time.Time` type.
    - Add the new todo to the `Todos` slice.

7. **List Todos Functionality:**
    - Implement a function `listTodos` to display existing todos.
    - Print each todo with its index, description, and due date.

8. **Save Todos to File:**
    - Create a function `saveTodos` that takes a `Todos` slice and writes it to a JSON file (`todos.json`).

9. **Menu and User Interaction:**
    - Inside the `main` function, implement a simple menu system:
        - Add Todo
        - List Todos
        - Exit

10. **Error Handling:**
    - Add error handling for user input, date parsing, and file operations.
    - Display appropriate error messages to the user.

11. **Usage Instructions:**
    - Provide instructions on how to clone the repository, run the program, and interact with the TodoCLI in a `README.md` file.

12. **Testing:**
    - Test the application by adding todos, listing them, and ensuring that they are persisted to the `todos.json` file.

13. **Bonus (Optional):**
    - Add additional features, such as marking todos as completed, removing todos, or sorting todos by due date.

14. **Documentation and License:**
    - Add comments to your code for better readability.
    - Include a license file (e.g., `LICENSE`) with the project.

15. **Publish to GitHub (Optional):**
    - If you're comfortable, consider publishing your project on GitHub and share it with others.

## Resources

- Golang Documentation: https://golang.org/doc/
- Golang Time Package: https://pkg.go.dev/time
- JSON Package: https://pkg.go.dev/encoding/json

## Example Output

```plaintext
Welcome to the TodoCLI App!

Choose an option:
1. Add Todo
2. List Todos
3. Exit

Enter your choice (1-3): 1
Adding a new todo:
Enter todo description: Finish project
Enter due date (YYYY-MM-DD): 2023-12-31
Todo added successfully!

Choose an option:
1. Add Todo
2. List Todos
3. Exit

Enter your choice (1-3): 2
Listing todos:
1. Finish project - Due: 2023-12-31

Choose an option:
1. Add Todo
2. List Todos
3. Exit

Enter your choice (1-3): 3
Todos saved. Exiting the TodoCLI App. Goodbye!
