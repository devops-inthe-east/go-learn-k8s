# Day 1: Learning Go & System Design with Kubernetes

## Fundamental Concepts & Q&A

### 1. Library Discovery & Navigation
**Q: When working with libraries, how do we know what functions are available and what parameters they require?**
**A:** Use your IDE's IntelliSense (hovering over types or typing `package.`), use the "Go to Definition" feature (F12) to read the actual source code, or read the official documentation at `pkg.go.dev`.

### 2. Exported vs. Unexported (Visibility)
**Q: Why are words like `Until` capitalized, but `period` is lowercase?**
**A:** Go does not use `public` or `private` keywords. Instead, visibility is determined by capitalization:
- **Capital Letter (`Until`):** Exported (Public). It can be imported and used by other packages.
- **Lowercase Letter (`period`):** Unexported (Private). It can only be used inside its own package.

### 3. Channels and Concurrency
**Q: What is happening with `<-chan struct{}`?**
**A:** This is a receive-only Channel. A channel (`chan`) is a pipe that allows concurrent functions (goroutines) to send messages to each other. The `<-` means the function can only listen/read from it. `struct{}` is an empty data structure taking 0 bytes. Together, this pipe acts as a highly efficient signal to tell the function to stop looping.

### 4. Function Wrappers
**Q: Is `JitterUntil` a sub-function of `Until`?**
**A:** No, it is just another normal public function in the same package. `Until` acts as a **Wrapper**. It provides a simple API for developers by calling `JitterUntil` underneath and automatically filling in annoying default values (like `0.0` for jitter).

### 5. Function Types & Primitives
**Q: What do the types in `JitterUntil(f func(), period time.Duration, 0.0, true, stopCh)` mean?**
**A:**
- `func()`: A function type. You can pass entire functions as variables!
- `time.Duration`: An integer representing nanoseconds, used for time intervals.
- `0.0`: A `float64` decimal. Here, it represents a 0% jitter (randomness) factor.
- `true`: A boolean. Here, it represents `sliding`, meaning the timer starts *after* `f` executes.

### 6. The "Config Struct" Design Pattern
**Q: What are the trade-offs of passing multiple different data types directly into a public function (e.g., `0.0, true`)?**
**A:** 
- **Readability:** It creates the "Mystery Boolean" problem where it is hard to tell what `0.0` or `true` means without reading the docs.
- **Extensibility:** If you add a new parameter later, you break the signature and break everyone else's code relying on your public library.
- **Solution:** Bundle the parameters into a single custom type (e.g., `WaitOptions struct { Jitter float64, Sliding bool }`).

### 7. Anonymous & Higher-Order Functions
**Q: If `Poll` requires a `ConditionFunc`, how do we declare a condition when calling it?**
**A:** We use an **Anonymous Function** (a function without a name) and declare it on the fly:
```go
wait.Poll(time.Second, time.Minute, func() (bool, error) {
    // Condition logic goes here
    return true, nil 
})
```
Functions that take other functions as arguments are called **Higher-Order Functions**.
