package sim

import (
	"io"
	"math/rand"
	"sync"
	"time"
)

type Stream struct {
	Data chan byte
	wg   sync.WaitGroup
}

func NewStream() *Stream {
	ch := make(chan byte, 20)
	return &Stream{Data: ch}
}

func (s *Stream) Start(data []byte) {
	s.wg.Wait()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(s.Data)
		for _, item := range data {

			min := 10
			max := 100
			randomIntInRange := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)

			s.Data <- item
		}

	}()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (s *Stream) Read(p []byte) (n int, err error) {
	for i := range p {
		b, ok := <-s.Data
		if !ok {
			if n == 0 {
				return 0, io.EOF
			}
			return n, io.EOF
		}
		p[i] = b
		n++
	}
	return n, nil
}
