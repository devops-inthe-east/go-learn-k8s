package controller

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

// Queue is an interface that describes exactly what  a queue should do!
// Notice there is NO actual logic here, just a blueprintof methids.

type Queue interface {
	Add(item interface{})
	Get() (item interface{}, shutdown bool)
	Done(item interface{})
}

type Controller struct {
	// We now demand that whatever queue is passed to us MUST obey the rules above.
	// We dont care if its a workqueue.Type or a ShoppingCart, as long as it has Add,Get & Done.
	// MUST satisy the Queue interface above. This is the Power of Abstraction

	queue Queue
}

// Run starts the controller's background workers.
func (c *Controller) Run(ctx context.Context, workers int) {
	fmt.Println("Starting Controller ...")

	// Spin up exactly 'workers' numbers of threads!
	for i := 0; i < workers; i++ {
		// The 'go' keyword physically spins off a new background thread (Gorountine)!
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	// This block the main thread from existing until the context is cancelled.
	<-ctx.Done()
	fmt.Println("Shutting Down Controller...")
}

// runWorker is a long-runnning function that continally process the queue.
func (c *Controller) runWorker(ctx context.Context) {
	// An infinite loop! It runs as long as processNextWorkItem returns 'true'.
	for c.processNextWorkItem(ctx) {
	}
}

// processNextWorkItem grabs one item off the queue and processess it.
func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	// In reality: item, shutdown := c.queue.Get()
	fmt.Println("Processing items from queue...")

	// Simulate hards work
	time.Sleep(2 * time.Second)

	// Returing true tells the runWorker loop to keep spinning!
	return true

}

// Informer simulates watching the Kubernestes API server for changes.
func (c *Controller) Informer() {
	// 1. Simulate the Kubernetes API Server by creating a channel that sends data every 2 seconnds.
	fmt.Println("INFEROMER: Alert! Pod 'nginx-web-server' just crashed!")

	// 2. The Informer frabs our queue nd tosses the key into it!
	c.queue.Add("nginx-web-server")

	fmt.Println("INFORMER: I added 'nginx-web-server' to the queue for the background workers to handle.")

}

// Questions unanswered so far by the agent.

// My impresion is that context.Context is also a HOF - that can be passed as an argument to a function.
// But it is not a user defined data type. Nor is it a struct. so then what is it ?
// Also context.Context is the new way of passing values between functions or goroutines.
// If so what was the older way deprecated for?>

// Questions answered so far by the agent.

// What is hard work here ?  -- It can be anything like making an api call, writing to a file, or crunching numbers.
// The runWorker function is an infinite loop , how did should I know that to an untrained eye.
// is wait.UntilWithContext an HOF function??
// The 'go' keyword physically spins off a new background thread (Gorountine)!
