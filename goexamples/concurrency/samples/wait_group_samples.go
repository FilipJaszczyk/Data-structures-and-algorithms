package samples

import (
	"sync"
)

func WgBasicUsage() int {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()
	return data
}
