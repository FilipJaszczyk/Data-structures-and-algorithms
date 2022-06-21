package broken

import (
	"fmt"
	"sync"
)

// Deadlock waiting for goroutine that never started
func Deadlock() {
	var data int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()
	fmt.Printf("value of data is: %d", data)
}

// Output of this function results either:
// -value of data is: 0
// -value of data is: 1
// no output
func Undeterministic() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("value of data is: %d", data)
	}
}

// Func literal error
// output would results in 4 4 4
func RangeVariablePassed() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
