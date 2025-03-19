package sim

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stream struct {
	Data      chan byte
	Closed    chan bool
	hasclosed bool
	wg        sync.WaitGroup
}

func NewStream() *Stream {
	ch := make(chan byte, 1)
	return &Stream{Data: ch, hasclosed: false}
}

func (s *Stream) Start(data []byte) {
	s.wg.Wait()
	s.wg.Add(1)
	go func(chOne chan<- byte) {
		defer s.wg.Done()
		defer close(chOne)
		for _, item := range data {

			min := 50
			max := 100
			randomIntInRange := rand.Intn(max-min+1) + min
			time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)
			chOne <- item
		}
		//s.hasclosed = true
	}(s.Data)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func (r *Stream) Read(p []byte) (int, error) {
	idx := 0
	buff := make([]byte, len(p), len(p))

	for idx < len(buff) {
		v, ok := <-r.Data
		if !ok {
			break
		}

		if v > 0 && ok {
			buff[idx] = v
			idx++
		}

	}
	copy(p, buff[:idx])
	fmt.Println(p)
	return idx, nil
}
