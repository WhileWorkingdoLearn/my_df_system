package connection

import "fmt"

type ConnectionState int

const (
	Idle ConnectionState = iota
	ACK
	Open
	Stream
	Reset
	Close
)

type Connection struct {
	state ConnectionState
}

func (c *Connection) State(state ConnectionState) error {
	if state == Idle {

		return nil
	}
	if state == ACK {

		return nil
	}
	if state == Open {
		fmt.Println("is Open")
		return nil
	}

	if state == Stream {

		return nil
	}

	if state == Reset {

		return nil
	}

	if state == Close {
		fmt.Println("Closing")
		c.state = state
		return nil
	}
	return fmt.Errorf("unkown state")
}

func (c *Connection) Listen() {
	for c.state != Close {
		fmt.Println("Listening")
		err := c.State(2)
		if err != nil {
			break
		}
	}
}

func (c *Connection) Close() {
	c.state = Close
}
