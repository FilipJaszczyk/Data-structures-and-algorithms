package samples

import (
	"fmt"
)

func ReceiveValue() int {
	ch := make(chan int)
	go func(a, b int) {
		ch <- a + b
	}(1, 2)
	return <-ch
}

func ProducerConsumer() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Printf("sending: %d \n", i)
			ch <- i
		}
	}()

	go func() {
		for v := range ch {
			fmt.Printf("processing: %d \n", v)
		}
		done <- true
	}()
	<-done
}

// Buffered channel is non blocking one meaning that
// sending data (to the chanel) is not blocked by the consumer.
func BufferedConsumerProducer(size, numbersToProcess int) {
	ch := make(chan int, size)
	done := make(chan bool)
	go func() {
		defer close(ch)
		for i := 1; i <= numbersToProcess; i++ {
			fmt.Printf("sending: %d \n", i)
			ch <- i
		}
	}()

	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Printf("processing: %d \n", v)
			} else {
				done <- true
			}
		}
	}()
	<-done
}
