package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
)

type Conn struct {
	conn   chan []byte
	byteCH chan byte
	data   [][]byte
	wg     sync.WaitGroup
}

func NewUDP(data [][]byte) *Conn {
	c := &Conn{data: data, conn: make(chan []byte, 3), byteCH: make(chan byte, 8)}
	c.Dial()
	return c
}

func (udp *Conn) Dial() {
	go func() {

		defer close(udp.byteCH)
		for d := range udp.conn {

			for _, d := range d {
				min := 5
				max := 10
				randomIntInRange := rand.Intn(max-min+1) + min
				time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)
				udp.byteCH <- d
			}
		}
	}()
	go func() {
		defer close(udp.conn)
		for _, d := range udp.data {
			udp.wg.Add(1)

			go func(data []byte) {
				defer udp.wg.Done()
				min := 10
				max := 50
				randomIntInRange := rand.Intn(max-min+1) + min
				time.Sleep(time.Duration(randomIntInRange) * time.Millisecond)
				udp.conn <- data
			}(d)
		}
		udp.wg.Wait()
	}()
}

func (udp *Conn) Read(buff []byte) (int, error) {
	n := 0
	for n < len(buff) {
		d, ok := <-udp.byteCH
		if !ok {
			return n, io.EOF
		}
		buff[n] = d
		n++
	}
	return n, nil

}

const Keylength = 16
const IdLength = 16
const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Packet struct {
	SessionID    []byte
	PrevPacketId []byte
	PacketID     []byte
	NPacketID    []byte
	Data         []byte
}

func (p *Packet) Parse(data []byte) (parsed int, done bool, err error) {
	clrf := bytes.Index(data, []byte("|"))
	if clrf == -1 {
		return 0, false, nil
	}
	if clrf == 0 {
		return 0, true, nil
	}
	seperator := bytes.Index(data, []byte(":"))
	if seperator == -1 {
		return 0, false, nil
	}

	return 0, true, nil
}

const (
	Init uint8 = iota
	Ack
	Stream
	Alive
	Pending
	Closed
)

func main() {
	// msg := "ConnectionKey:PPackedId:PackedId:NPackedId:Data:IsBedirectional:Checksum"
	SessionKey := GenerateKey(Keylength)
	packetKeys := make([]Packet, 0)
	for i := range 3 {
		nKey := GenerateKey(Keylength)
		p := Packet{
			SessionID: SessionKey,
			PacketID:  nKey,
			Data:      []byte(fmt.Sprintf("Hello world %v", i)),
		}
		packetKeys = append(packetKeys, p)
	}
	MSGS := make([][]byte, 0)
	for idx, packet := range packetKeys {
		buff := bytes.NewBuffer(make([]byte, 0))
		buff.Write(packet.SessionID)
		buff.WriteByte(':')
		buff.WriteByte('0')
		buff.WriteByte(':')
		if idx > 0 {
			buff.Write(packetKeys[idx-1].PacketID)
		}
		buff.WriteByte(':')
		buff.Write(packet.PacketID)
		buff.WriteByte(':')
		if idx < len(packetKeys)-1 {
			buff.Write(packetKeys[idx+1].PacketID)
		}
		buff.WriteByte(':')
		buff.Write(packet.Data)
		buff.WriteString(("|"))
		MSGS = append(MSGS, buff.Bytes())
	}

	conn := NewUDP(MSGS)
	buff := make([]byte, 128, 128)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buff[:n])

}

type SessionKey []byte

func GenerateKey(length int) SessionKey {
	b := make([]byte, length)
	for i := range length {
		min := 0
		max := len(Alphabet) - 1
		randomIntInRange := rand.Intn(max-min+1) + min
		b[i] = byte(Alphabet[randomIntInRange])
	}
	return b
}

/*
func (s SessionKey) ToString() string {
	return string(s)
}

type Version int

func (v Version) ToByte() byte {
	fmt.Println("Converts to byte")
	return byte(v)
}

func (v Version) ToString() string {
	return fmt.Sprintf(" %v", v)
}

func (v Version) FromString(num string) (int, error) {
	i, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return int(v), err
	}
	return int(i), nil
}



type Sessions map[string]Connection

type Connection struct {
	id       []byte
	buffer   [][]byte
	received int
	state    uint8
}
*/
