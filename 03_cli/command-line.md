## Working with command-line arguments

Although typing user input when needed might look like a nice idea, this is not
usually how real software works. Usually, user input is given in the form of
command-line arguments to the executable file. By default, command-line arguments
in Go are stored in the os.Args slice. Go also offers the flag package for parsing
command-line arguments, but there are better and more powerful alternatives.

It is important to know that the
os.Args slice is properly initialized by Go and is available to the program when
referenced.

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if there are enough command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <your_name>")
		return
	}

	// Extract the name from command-line arguments
	name := strings.Join(os.Args[1:], " ")

	// Greeting the user
	fmt.Printf("Hello, %s! Nice to meet you.\n", name)
}
```

In this version:

1. **Command-Line Arguments (`os.Args`):** We use `os.Args` to access the command-line arguments. `os.Args` is a slice that holds the command-line arguments, where the first element is the name of the executable itself.

2. **Check Argument Count (`if len(os.Args) < 2 { fmt.Println("Usage: go run main.go <your_name>") return }`):** We check if there are enough command-line arguments. If not, we print a usage message and exit the program.

3. **Extracting Name (`name := strings.Join(os.Args[1:], " ")`):** We use `strings.Join` to concatenate the command-line arguments (excluding the first element, which is the program name) into a single string representing the user's name.

4. **Greeting the User (`fmt.Printf("Hello, %s! Nice to meet you.\n", name)`):** We print a greeting message to the console, addressing the user by the name extracted from the command-line arguments.

To run this program, you would execute it from the command line like this:

```bash
go run main.go YourName
```

Another example of CLI application

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the Golang CLI Calculator!")

	// Get the first number from the user
	fmt.Print("Enter the first number: ")
	var input1 string
	fmt.Scanln(&input1)

	// Convert the first input to an integer
	num1, err1 := strconv.Atoi(input1)
	if err1 != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Get the second number from the user
	fmt.Print("Enter the second number: ")
	var input2 string
	fmt.Scanln(&input2)

	// Convert the second input to an integer
	num2, err2 := strconv.Atoi(input2)
	if err2 != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Perform the calculation
	sum := num1 + num2

	// Display the result
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
}
```

In this version:

1. **Print Welcome Message (`fmt.Println("Welcome to the Golang CLI Calculator!")`):** We print a welcome message to the console.

2. **Get First Number (`fmt.Print("Enter the first number: ")`, `fmt.Scanln(&input1)`):** We prompt the user to enter the first number and use `fmt.Scanln` to read the input.

3. **Convert First Input (`num1, err1 := strconv.Atoi(input1)`):** We convert the first input to an integer using `strconv.Atoi` and handle any conversion errors.

4. **Get Second Number (`fmt.Print("Enter the second number: ")`, `fmt.Scanln(&input2)`):** We prompt the user to enter the second number and use `fmt.Scanln` to read the input.

5. **Convert Second Input (`num2, err2 := strconv.Atoi(input2)`):** We convert the second input to an integer using `strconv.Atoi` and handle any conversion errors.

6. **Perform Calculation (`sum := num1 + num2`):** We calculate the sum of the two numbers.

7. **Display Result (`fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)`):** We print the result, including the numbers and their sum.

To run this program, save the code in a file with a ".go" extension and execute it:

```bash
go run main.go
```
