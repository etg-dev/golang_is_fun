Controlling program flow

So far, we have seen Go variables but how do we change the flow of a Go program
based on the value of a variable or some other condition? Go supports the **if/
else and switch** control structures. Both control structures can be found in most
modern programming languages, so if you have already programmed in another
programming language,

Certainly! In Go, conditionals are expressed primarily through `if` statements and the `switch` statement. Here's a comprehensive guide on conditions in Go:

### `if` Statements:

The basic syntax of an `if` statement in Go is as follows:

```go
if condition {
    // code to be executed if the condition is true
} else {
    // code to be executed if the condition is false
}
```

You can also have an `if` statement without an else block:

```go
if condition {
    // code to be executed if the condition is true
}
```

**Example:**

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is not greater than 5")
    }
}
```

### `switch` Statement:

The `switch` statement in Go is more flexible than in some other languages. It can be used with different types, not just constants. Here's the basic syntax:

```go
switch expression {
case value1:
    // code to be executed if expression == value1
case value2:
    // code to be executed if expression == value2
// more cases...
default:
    // code to be executed if none of the cases match
}
```

**Example:**

```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Friday":
        fmt.Println("It's almost the weekend!")
    default:
        fmt.Println("It's a regular day.")
    }
}
```

In Go, a `switch` statement can also be used without an expression, which is equivalent to `switch true`. 
This allows you to express complex conditions in each case:

```go
package main

import "fmt"

func main() {
    age := 25

    switch {
    case age < 18:
        fmt.Println("You are a minor.")
    case age >= 18 && age < 65:
        fmt.Println("You are an adult.")
    default:
        fmt.Println("You are a senior.")
    }
}
```

### Type Switch:

```go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Tuesday":
        fmt.Println("It's the second day of the week.")
    case "Wednesday":
        fmt.Println("It's the middle of the week.")
    case "Thursday":
        fmt.Println("It's almost the weekend!")
    case "Friday":
        fmt.Println("It's the end of the workweek.")
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("It's an unknown day.")
    }
}
```

In this example, the `switch` statement is used to check the value of the `day` variable. Each case corresponds to a specific day of the week. If the value of `day` matches any of the cases, the corresponding message is printed. If there is no match, the `default` case is executed.

Note that you can also use multiple values in a single case, as shown with "Saturday" and "Sunday" sharing the same case.

Understanding how to use `if` and `switch` statements effectively is fundamental for writing clear and concise Go code.

let's use a very common pattern in Go that is used
almost everywhere. This pattern says that if the value of an error variable as
returned from a function is nil, then everything is OK with the function execution.
Otherwise, there is an error condition somewhere that needs special care. This
pattern is usually implemented as follows:
```go
err := anyFunctionCall()
if err != nil {
// Do something if there is an error
}
```
err is the variable that holds the error value as returned from a function and !=
means that the value of the err variable is not nil. You will see similar code multiple
times in Go programs.

### Iterating with for loops and range

This section is all about iterating in Go. Go supports for loops as well as the range
keyword for iterating over all the elements of arrays, slices, and (as you will see in
Chapter 3, Composite Data Types) maps. An example of Go simplicity is the fact that
Go provides support for the for keyword only, instead of including direct support
for while loops. However, depending on how you write a for loop, it can function as
a while loop or an infinite loop. Moreover, for loops can implement the functionality
of JavaScript's forEach function when combined with the range keyword.

### `for` Loop:

The `for` loop in Go is quite flexible and can be used for various types of iterations.

#### Basic `for` Loop:

```go
package main

import "fmt"

func main() {
    // Simple for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

In this example, the loop initializes `i` to 0, continues as long as `i` is less than 5, and increments `i` in each iteration.

#### Infinite Loop:

```go
package main

import "fmt"

func main() {
    // Infinite loop
    for {
        fmt.Println("This is an infinite loop.")
        // To break out of the loop, use a break statement
        break
    }
}
```

### `range` Keyword:

The `range` keyword is used to iterate over various data structures like arrays, slices, maps, strings, and channels.

#### Arrays and Slices:

```go
package main

import "fmt"

func main() {
    // Iterating over an array
    numbers := [3]int{1, 2, 3}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Iterating over a slice
    fruits := []string{"apple", "banana", "cherry"}
    for index, value := range fruits {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }
}
```

#### Maps:

```go
package main

import "fmt"

func main() {
    // Iterating over a map
    colors := map[string]string{
        "red":    "#FF0000",
        "green":  "#00FF00",
        "blue":   "#0000FF",
    }

    for key, value := range colors {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }
}
```

#### Strings:

```go
package main

import "fmt"

func main() {
    // Iterating over a string
    message := "Hello, Go!"
    for index, char := range message {
        fmt.Printf("Index: %d, Character: %c\n", index, char)
    }
}
```

### `goto` Statement:

The `goto` statement in Go is available but is not recommended for regular use because it can make code less readable and harder to maintain. It is used to transfer control to a labeled statement.

#### Example:

```go
package main

import "fmt"

func main() {
    i := 0

loop:
    fmt.Println(i)
    i++

    if i < 5 {
        goto loop
    }
}
```

In this example, the `goto` statement is used to create a labeled loop. However, it's crucial to note that using `goto` can lead to less maintainable and less readable code, and it's generally considered a best practice to use `for` loops and `range` instead.

While `goto` can be useful in certain situations, it's essential to be cautious and use it sparingly to avoid introducing unnecessary complexity to your code.

Remember, in most cases, `for` loops and `range` provide cleaner and more readable alternatives for iteration in Go.