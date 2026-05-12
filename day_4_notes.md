# Day 4: System Architecture and Wiring the Application

## Fundamental Concepts & Q&A

### 1. The Entry Point: `main.go`
**Q: Why do we need a `main.go` file, and what goes inside it?**
**A:** Every executable Go program must have exactly one starting point: a function named `main()` sitting inside `package main`. This file acts as the **"Wiring Room"**. Good Go developers do not put business logic in `main.go`. Instead, it is used to instantiate databases, initialize Queues, boot up Controllers, and wire all the pieces together.

### 2. Graceful Shutdowns and OS Signals
**Q: What does `signal.NotifyContext` do?**
**A:** When you press `Ctrl+C` in your terminal, your Operating System fires a `SIGTERM` (Interrupt) signal to forcefully kill the application. 
`signal.NotifyContext` actively listens to your Operating System for this exact signal. When it hears `Ctrl+C`, instead of letting the application crash violently, it cleanly triggers the `ctx.Done()` channel. This allows your Kubernetes workers to finish their current jobs, save their progress, and shut down gracefully.

### 3. Concurrency Orchestration
**Q: How does the entire distributed system flow together?**
**A:** 
1. `main.go` instantiates the Queue and the Controller.
2. `myController.Run(ctx, 5)` spawns 5 background threads (Goroutines). Because the queue is empty, all 5 threads go to sleep using `q.cond.Wait()`.
3. `myController.Informer()` fakes a Kubernetes event, locks the Mutex, appends an item to the Queue, and rings `q.cond.Signal()`.
4. The signal wakes up exactly *one* of the 5 sleeping threads to process the item.
5. Meanwhile, `<-ctx.Done()` safely blocks the `main.go` thread from exiting until the user stops the program.

### 4. Visibility (Capitalization Revisited)
**Q: Why did we capitalize `Queue` in the `Controller` struct when moving across files?**
**A:** As we learned on Day 1, Go uses capitalization for visibility. Because `main.go` and `controller.go` are different packages, `main.go` cannot inject a queue into a private, lowercase field. The field must be capitalized (`Queue Queue`) so external packages can access it.
