package sim

import (
	"math/rand"
	"time"
)

type Stream struct {
	Data      <-chan byte
	Closed    chan bool
	hasclosed bool
}

func NewStream(data []byte) Reader {
	ch := make(chan byte, 1)
	Closed := make(chan bool, 1)
	go func() {
		defer close(ch)
		defer close(Closed)
		for _, item := range data {
			min := 50
			max := 100
			randomIntInRange := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)
			ch <- item
		}
		Closed <- true
	}()
	return &Stream{Data: ch, Closed: Closed, hasclosed: false}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (r *Stream) Read(p []byte) (int, error) {

	return 0, nil
}
