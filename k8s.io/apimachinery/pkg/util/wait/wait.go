package wait

import (
	"context"
	"time"
)

// Until loops until stop channel is closed, running f every period.

// func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
// 	JitterUntil(f, period, 0.0, true, stopCh)
// }

// UntilWithContext loops until the context is canceled, running f every period.

func UntilWithContext(ctx context.Context, f func(context.Context), period time.Duration) {

	// Notice how we use an anonymous function, and how ctx.Done() bridges the gap!
	JitterUntil(func() { f(ctx) }, period, 0.0, true, ctx.Done())
}

// ConditionFunc returns true if the condition is satisfied, or an error
// if the loop should be aborted.

type ConditionFunc func() (done bool, err error)

// Poll tries a condition func

func Poll(interval, timeout time.Duration, condition ConditionFunc) error {
	return poll(interval, timeout, condition)
}

type Backoff struct {
	Duration time.Duration
	Factor   float64
	Jitter   float64
	Steps    int
	Cap      time.Duration
}

// Questions unanswered so far by the agent.

// Within a func a new function is defined?
// (f(ctx)) -> What is this exactly mean, and how is it related to 'Until' function?

// Can I declare a new function within a  existing func ?
// and if I do so what is the scope of that function ?  -- I guessing it only limited to the function it is declared in.

// Questions answered so far by the agent.

// Question about the code-
// When working with libraries, how do we know what functions are available?
// How to exavtly know what params are required for each function present within a library.

// Syntax shock for me:
// what is happening here "<-chan"
// Why are words like period,StopCh, highlighted differently.
// why are words like Until and JitUntil capitalized.
// is JitUntil a sub-func of the main func Until?

// Connect the dots :
// Poll is a public function and it takes
//1. interval
//2. timeout
//3. ConditionFunc which is a function that takes no arguments and
// returns true if the condition is satisfied, or an error if the loop should be aborted.

// But I have'nt declared a condition yet - for the ConditionFunc.

// 'struct' means a custom defined data type.
// So can it be a strucured data type or an data type  ?
// Just like there is an distintion of b/w private & public functions in go lang.
// is there any such kind for struct too?.

// In the file above, we define a struct 'ConditionFunc' with fields like name, age, city.
// So whenever I create a variable of type 'ConditionFunc', it will have these three fields.
// To an untrained eye, it might appear like an object - so what is the mental model I would need to keep in my mind.
// To identify a Struct. within a function.

// 'Jitter' is the variable within the function or ?
// When creating a Custom data type [struct.]. what are the convention I would need to adhere.
// How does the expotential backoff mechanism works within the given value fields?
