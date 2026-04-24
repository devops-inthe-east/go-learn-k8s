package wait

import (
	"time"
)

// Until loops until stop channel is closed, running f every period. 

func Until(f func (),period time.Duration, stopCh <-chan struct{}){
	JitterUntil (f, period, 0.0, true, stopCh)
}

// ConditionFunc returns true if the condition is satisfied, or an error 
// if the loop should be aborted. 

type ConditionFunc func () (done bool, err error)

// Poll tries a condition func 

func Poll(interval, timeout time.Duration, condition ConditionFunc) error {
	return poll(interval, timeout, condition)
}



// Questions unanswered so far by the agent. 

// Connect the dots :
// Poll is a public function and it takes 
//1. interval
//2. timeout
//3. ConditionFunc which is a function that takes no arguments and
// returns true if the condition is satisfied, or an error if the loop should be aborted.

// But I have'nt declared a condition yet - for the ConditionFunc.




// Questions answered so far by the agent. 


// Question about the code- 
// When working with libraries, how do we know what functions are available? 
// How to exavtly know what params are required for each function present within a library.

// Syntax shock for me:
// what is happening here "<-chan"
// Why are words like period,StopCh, highlighted differently.
// why are words like Until and JitUntil capitalized.
// is JitUntil a sub-func of the main func Until?
