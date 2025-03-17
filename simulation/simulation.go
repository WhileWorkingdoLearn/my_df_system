package sim

import (
	"math/rand"
	"time"
)

type Stream struct {
	Data <-chan byte
}

func NewStream(data []byte) *Stream {
	ch := make(chan byte, 2)
	go func() {
		defer close(ch)
		for _, item := range data {
			min := 50
			max := 100
			randomIntInRange := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)
			ch <- item
		}
	}()
	return &Stream{Data: ch}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func NewReader() Reader {
	return Stream{}
}

func (r Stream) Read(buff []byte) (int, error) {
	res := <-r.Data
	for range len(buff) {
		buff = append(buff, res)
	}
	return len(buff), nil
}
