## warm up with go

create a file with name hw.go and copy and past this file

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

In order to compile Go code and create a binary executable file, you need to use
the go build command. What go build does is create an executable file for you to
distribute and execute manually. This means that the go build command requires
an additional step for running your code.

The generated executable is automatically named after the source code filename
without the .go file extension. Therefore, because of the hw.go source file, the
executable will be called hw.

run this command in your terminal

```bash
$ go build hw.go
$ ./hw
Hello World!
```

instead of this command you can use this `go run fileName.go`

It's more or less the equivalent of running go build X.go -o /tmp/random-tmp-folder/exe && /tmp/random-tmp-folder/exe

Now that we know how to compile Go code

you should know
Go programs are organized in packages—even the smallest Go program should
be delivered as a package. The package keyword helps you define the name of a
new package, which can be anything you want with just one exception: if you are
creating an executable application and not just a package that will be shared by other
applications or packages, you should name your package main. we explore about packages later

Important formatting and coding rules
You should know that Go comes with some strict formatting and coding rules that
help the developer avoid beginner mistakes and bugs—once you learn these few
rules and Go idiosyncrasies as well as the implications they have for your code, you
will be free to concentrate on the actual functionality of your code. Additionally, the
Go compiler is here to help you follow these rules with its expressive error messages
and warnings. Last, Go offers standard tooling (gofmt) that can format your code for
you so you never have to think about it.
The following is a list of important Go rules that will help you while reading this
chapter:
* Go code is delivered in packages and you are free to use the functionality
found in existing packages. However, if you are going to import a package,
you should use some of this functionality—there are some exceptions to
this rule that mainly have to with initializing connections, but they are not
important for now.

* You either use a variable or you do not declare it at all. This rule helps you
avoid errors such as misspelling an existing variable or function name.
* There is only one way to format curly braces in Go.

* Coding blocks in Go are embedded in curly braces even if they contain just a
single statement or no statements at all.

Go functions can return multiple values.
You cannot automatically convert between different data types, even if
they are of the same kind. As an example, you cannot implicitly convert
an integer to a floating point.

### Defining and using variables
Imagine that you wanted to perform some basic mathematical calculations with Go.
In that case, you need to define variables to keep your input and your results.

Go provides multiple ways to declare new variables in order to make the variable
declaration process more natural and convenient. You can declare a new variable
using the var keyword followed by the variable name, followed by the desired data
type

**If there
is an initial value given, you can omit the data type and the compiler will guess it for
you.**

There is also the := notation, which can be used instead of a var declaration. :=
defines a new variable by inferring the data of the value that follows it. The official
name for := is **short assignment statement** and it is very frequently used in Go,
especially for getting the return values from functions and for loops with the range
keyword.

You rarely see the use of var in Go; the var keyword is mostly used for
declaring global or local variables without an initial value. The reason for the former
is that every statement that exists outside of the code of a function must begin with a
keyword such as func or var.

**This means that the short assignment statement cannot be used outside of a function
environment**

**Therefore, although you can declare local variables using either var or :=, only
const (when the value of a variable is not going to change)**

Programs tend to display information, which means that they need to print data or
send it somewhere for other software to store or process it. For printing data on the
screen, Go uses the functionality of the fmt package. If you want Go to take care of
the printing, then you might want to use the fmt.Println() function.

```go
package main

import (
    "fmt"
    "math"
)

var Global int = 1234
var AnotherGlobal = -5678

func main() {
    var j int
    i := Global + AnotherGlobal
    fmt.Println("Initial j value:", j)
    j = Global

    // math.Abs() requires a float64 parameter, 
    // so we type cast it appropriately
    k := math.Abs(float64(AnotherGlobal))
    fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
```
In Go, each variable is given a default value when it is declared, even if you don't explicitly assign a value to it. This default value is known as the "zero value." The zero value depends on the type of the variable.

### Numeric Types:

1. **Integer Types (int, int8, int16, int32, int64):**
   - Zero Value: 0

2. **Unsigned Integer Types (uint, uint8, uint16, uint32, uint64):**
   - Zero Value: 0

3. **Floating-Point Types (float32, float64):**
   - Zero Value: 0.0

### Textual Types:

1. **String:**
   - Zero Value: ""

2. **Rune (alias for int32, represents a Unicode code point):**
   - Zero Value: 0

### Boolean Type:

1. **bool:**
   - Zero Value: false

### Composite Types:

1. **Array:**
   - Zero Value: An array of the respective type where all elements are set to their zero values. For example, if you have an array of integers, all elements will be set to 0.

2. **Slice:**
   - Zero Value: nil

3. **Map:**
   - Zero Value: nil

4. **Struct:**
   - Zero Value: A struct with all fields set to their zero values. For numeric types, it's 0; for strings, it's ""; for bool, it's false, and so on.

### Pointer Types:

1. **Pointer to any type:**
   - Zero Value: nil

### Function Types:

1. **Function:**
   - Zero Value: nil

### Interface Types:

1. **Interface:**
   - Zero Value: nil

### Channel Type:

1. **Channel:**
   - Zero Value: nil

### Custom Types:

1. **Custom Types (user-defined types):**
   - The zero value is determined by the zero value of the underlying type.

### Example:

```go
package main

import "fmt"

func main() {
    var intVar int
    var floatVar float64
    var stringVar string
    var boolVar bool

    fmt.Printf("intVar: %v\n", intVar)               // 0
    fmt.Printf("floatVar: %v\n", floatVar)           // 0.0
    fmt.Printf("stringVar: %v\n", stringVar)         // ""
    fmt.Printf("boolVar: %v\n", boolVar)             // false
}
```
