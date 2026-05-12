package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Booting up the Distributed System...")

	// 1. Create a modern Context that listens for Ctrl+C (SIGINT) from the keyboard!

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// 2. We will instantiate our Queue and Controller here later
	myQueue := queue.New()
	myController := controller.Controller{queue: myQueue}

	// 3. Boot up the Controller's background workers!
	// We use the 'go' keyword so this doenst block the main thread.
	// We pass in our context, add tell it to create
	go myController.Run(ctx, 5)

	// 4. Start the Informer to watch for events and feed the queue!
	go myController.Informer()

	fmt.Println("System is runnning! Press Ctrl+C to stop.")

	// 3. Block the main thread forever (until the user presses Ctrl+C )
	<-ctx.Done()

	fmt.Println("Gracefully shutting down ...")
}

// Questions unanswered so far by the agent.

// Questions answered so far by the agent.

// ctx is related context.Context ? Like is it an object instance of context.Context
// Why do we need a main.go file, what type of function/objects are stored here.
// Global function/variables/objects are stored here or?
// ctx has delcared twice what is the difference between them.
// to an untrained eye, how can I think about extending a function/varible like ctx.
// I learned about interfaces - how we can call a method using the instance/object of an interface.
