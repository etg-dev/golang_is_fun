### Concurrency in Go:

Concurrency is a way of structuring a program so that multiple independent tasks can be executed concurrently. In Go, concurrency is achieved through goroutines and channels.

- **Goroutines:**
    - A goroutine is a lightweight thread of execution managed by the Go runtime.
    - You can think of it as a function running concurrently with other goroutines.
    - Goroutines are started with the `go` keyword.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func sayHello() {
      for i := 0; i < 5; i++ {
          fmt.Println("Hello!")
          time.Sleep(time.Millisecond * 500)
      }
  }

  func main() {
      go sayHello()

      // The main function continues to execute independently of the goroutine.
      for i := 0; i < 5; i++ {
          fmt.Println("Main!")
          time.Sleep(time.Millisecond * 300)
      }

      // Ensure the main function doesn't exit before the goroutine completes.
      time.Sleep(time.Second)
  }
  ```

- **Concurrency vs Parallelism:**
    - Concurrency is about structuring a program, and parallelism is about executing multiple tasks simultaneously.
    - Go allows you to write concurrent programs, and if you have a multi-core machine, Go can also provide parallelism.

### Channels:

Channels are a fundamental feature in Go for communication and synchronization between goroutines.

- **Creating Channels:**
    - Channels are created using the `make` function.

  ```go
  ch := make(chan int)
  ```

- **Sending and Receiving:**
    - The `<-` operator is used for sending and receiving data on a channel.

  ```go
  // Sending data
  ch <- 42

  // Receiving data
  data := <-ch
  ```

- **Blocking:**
    - Sending to a channel blocks until there is a receiver.
    - Receiving from a channel blocks until there is a sender.

- **Closing Channels:**
    - It's important to close a channel when you're done with it to avoid deadlocks.
    - Closed channels can still be used to receive remaining values.

  ```go
  close(ch)
  ```

### Unbuffered Channels:

Unbuffered channels provide synchronous communication between goroutines. When a value is sent on an unbuffered channel, it waits until another goroutine is ready to receive that value.

#### Example:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Creating an unbuffered channel
    ch := make(chan int)

    // Sender goroutine
    go func() {
        fmt.Println("Sending data...")
        ch <- 42  // Blocks until a receiver is ready
        fmt.Println("Data sent!")
    }()

    // Receiver goroutine
    go func() {
        time.Sleep(time.Second) // Simulate some processing time
        fmt.Println("Receiving data...")
        data := <-ch  // Blocks until a sender is ready
        fmt.Println("Received data:", data)
    }()

    // Ensure main goroutine doesn't exit prematurely
    select {}
}
```

In this example, the sender and receiver are synchronized. The sender will block until the receiver is ready, and vice versa.

Now, let's move on to buffered channels.

### Buffered Channels:

Buffered channels have a fixed-size buffer that allows them to hold a certain number of elements. Sending to a buffered channel blocks only when the buffer is full, and receiving blocks only when the buffer is empty.

#### Example:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Creating a buffered channel with a capacity of 3
    ch := make(chan int, 3)

    // Sender goroutine
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("Sending data:", i)
            ch <- i  // Non-blocking until the buffer is full
        }
        close(ch)
    }()

    // Receiver goroutine
    go func() {
        for {
            data, ok := <-ch
            if !ok {
                fmt.Println("Channel closed.")
                break
            }
            fmt.Println("Received data:", data)
            time.Sleep(time.Millisecond * 500) // Simulate some processing time
        }
    }()

    // Ensure main goroutine doesn't exit prematurely
    select {}
}
```

In this example, the buffered channel allows the sender to send multiple values without blocking until the buffer is full. The receiver, likewise, can receive values without blocking until the buffer is empty.

Choose between buffered and unbuffered channels based on your specific concurrency requirements. Buffered channels are often useful for scenarios where bursts of data are sent or received.

### sync.WaitGroup:

The `sync.WaitGroup` type is used to wait for a collection of goroutines to finish before proceeding. It provides a way to wait for a specific number of operations to complete.

#### Example:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers done")
}
```

In this example, `wg.Add(1)` increments the counter for each goroutine, and `wg.Done()` decrements it when the goroutine completes. The `wg.Wait()` call blocks until the counter becomes zero, indicating that all goroutines have finished.

### `sync.Mutex` and Locks:

A mutex (`sync.Mutex`) is used to protect shared data from concurrent access. It ensures that only one goroutine can access the shared data at a time.

#### Example:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mutex sync.Mutex

func increment() {
	mutex.Lock()   // Acquire the lock
	defer mutex.Unlock() // Release the lock when the function completes

	counter++
	fmt.Println(counter)
}

func main() {
	for i := 0; i < 5; i++ {
		go increment()
		time.Sleep(time.Millisecond * 100)
	}

	time.Sleep(time.Second)
}
```

In this example, the `mutex.Lock()` ensures that only one goroutine can increment the `counter` at a time. The `defer mutex.Unlock()` ensures that the lock is released even if an error occurs or the function panics.

### Locks in Go:

In addition to `sync.Mutex`, Go provides the `sync.RWMutex` type, which allows multiple readers or a single writer at a time. It's useful when multiple goroutines need to read shared data, and exclusive access is required for writing.

#### Example:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var data = make(map[string]string)
var rwMutex sync.RWMutex

func readData(key string) string {
	rwMutex.RLock() // Acquire a read lock
	defer rwMutex.RUnlock() // Release the read lock when the function completes

	return data[key]
}

func writeData(key, value string) {
	rwMutex.Lock() // Acquire a write lock
	defer rwMutex.Unlock() // Release the write lock when the function completes

	data[key] = value
}

func main() {
	go func() {
		writeData("name", "John")
	}()

	go func() {
		fmt.Println("Name:", readData("name"))
	}()

	time.Sleep(time.Second)
}
```

In this example, the `rwMutex.RLock()` and `rwMutex.Lock()` ensure that multiple goroutines can read the data simultaneously, but only one can write at a time.
