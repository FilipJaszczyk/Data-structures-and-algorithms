package main

import "goexamples/concurrency/samples"

type closer1 interface {
	close()
}

type closer2 interface {
	close2()
}

type c struct {
	closer1
	closer2
}

func main() {
	samples.BufferedConsumerProducer(5, 5)

}
