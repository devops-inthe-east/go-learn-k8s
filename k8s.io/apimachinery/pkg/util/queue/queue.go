package workqueue

import "sync"

// Interface represents the operations of a queue

type Interface interface {
	Add(item interface{})
	Len() int
	Get() (item interface{}, shutdown bool)
	Done(item interface{})
}

// Type is the actual implementation of the work queue

type Type struct {
	// queue defines the order in which we will work on item
	queue []interface{}

	// dirty define all of the items that need to be procesed
	// It uses a map to ensure items are unique (acting like a Set).

	dirty map[interface{}]struct{}

	// Cond helps us safely lock thw queue so multiple workds dont crash it

	cond *sync.Cond

	shuttingDown bool
}

// Part 2 of Day 2!

func New() *Type {
	return &Type{
		dirty: make(map[interface{}]struct{}),
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}
func (q *Type) Add(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if q.shuttingDown {
		return
	}

	if _, exists := q.dirty[item]; exists {
		return
	}

	q.dirty[item] = struct{}{}
	q.queue = append(q.queue, item)

	q.cond.Signal()
}


// Get blocks until it can return an item to be processed 

func (q *Type) Get() (item interface {}, shutdown bool) {
		q.cond.L.Lock()
		defer q.cond.L.Unlock()

		// if the queue is empty, go to sleep!
		for len(q.queue) == 0 && !q.shuttingDown {
				q.cond.Wait()
		}

		if len(q.queue) == 0 {
				// We must be shuttung down
				return nil, true
		}
		
		// 1. Grab the very first item in ther slice
		item = q.queue[0]
		
		// 2. Prevent memory leaks!
		
		q.queue[0] = nil

		// 3. Slice off the first item

		q.queue = q.queue[1:]

		return item, false
		


// Questions unanswered so far by the agent.

// q.cond.Wait(): This is the literal function that puts the Go thread to sleep to save CPU.
// (Remember q.cond.Signal() from the Add() function? This is what it wakes up!)
// So essentially  q.cond.Signal() is a wake up call for the waiting threads.
// And q.cond.Wait() is the one that puts the thread to sleep.

// we use the variable ' len(q.queue)' 2 times in the Get function code. 
// Once in line 114 and once in line 116. the operative words 'for' and 'if' 
// are just for logic implementation of what ?
// Why do we delcare the memory leaks line after we get the first item. 
// And before the slice off items. 
// I dont know why memory leaks is a problem when we are working with pointers or concurrency function.


// Slice Manipulation: q.queue = q.queue[1:] is the standard Go syntax for chopping off the first item of a list.
// Memory Management: Why on earth do we set q.queue[0] = nil right before we chop it off?




// Questions answered so far by the agent.

// Batch 1 questions :
//
// Since we are working with multi data types and at time we also create custom data type(struct).

// Is there a convention I would need to follow for assign brackets to these custom types and normal data types
// 'Cond' is locking the queue? How does that work?
// if concurrency can use channels to communicate, What underlying mechanism does cond uses ?
// Are we in the territory of multi threading?

// Batch 2 questions :
// Why do we use make for the dirty map, but not for the queue slice?
// answer: Maybe we are reading the data without copying it.
// the variables 'dirty' & 'cond' are declared 'private' (starts with lowercase letter)
// And then I am using the &Type function to retrive it.
// But why is the sync.NewCond is used with the '&sync.Mutex' variable.
// What does the '&' symbol denotes here.


//  Batch 3 Questions:

// The Receiver: func (q *Type) Add(...) ->
// Notice how there is a set of parentheses before the function name?
// This is Go's version of Object-Oriented Programming!
// The defer keyword -> defer q.cond.L.Unlock() is arguably the most famous keyword in Go.

// I dont why the defer keyword has to be here ? and I am guessing it has an innate meaning
// when we are using the delcared function.
// Is this related to memory management ?

// The "Comma Ok" idiom -> if _, exists := q.dirty[item]; exists is how Go checks if something exists inside a Map.
// I am not aware of the Comma Ok idiom, and hence I am not able to get how it is working here.
// Can you explain this ? What is the correlation between the queue & comma-ok idiom ?
// Why do we using the '_' with the 'exists' variable here?

// Slices and Signals -> append() and Signal().
// What is the need  for the appned () function here? and is it related to the Slices Data Structure ?
// What exactly is the use of signal() function here and is it related to the Concurrency we